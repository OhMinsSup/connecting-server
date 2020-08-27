package repository

import (
	"connecting-server/app"
	"connecting-server/dto"
	"connecting-server/model"
	"github.com/jinzhu/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRawQuery struct {
	ID                 string    `json:"id"`
	CountryCode        string    `json:"country_code"`
	Phone              string    `json:"phone"`
	Birthday           string    `json:"birthday"`
	ThumbnailUrl       string    `json:"thumbnail_url"`
	CoverImageUrl      string    `json:"cover_image_url"`
	SttsMsg            string    `json:"stts_msg"`
	ShareProfileChange bool      `json:"share_profile_change"`
	AllowAddingFriends bool      `json:"allow_adding_friends"`
	FollowView         bool      `json:"follow_view"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (user *UserRepository) Profile(id string) *UserRawQuery {
	var userModel UserRawQuery
	sql := `
	SELECT u.id, u.created_at, u.updated_at, up.country_code, up.phone, up.birthday,
    up.thumbnail_url, up.cover_image_url, up.stts_msg,
    um.share_profile_change, um.allow_adding_friends, um.follow_view
	FROM "users" AS u
    INNER JOIN "user_profiles" AS up ON up.user_ref = (u.id)::uuid
    INNER JOIN "user_meta" As um ON um.user_ref = (u.id)::uuid
    WHERE u.id = ?`

	if err := user.db.Raw(sql, id).Scan(&userModel).Error; err != nil {
		return nil
	}
	return &userModel
}

func (user *UserRepository) FindExists(userId string) bool {
	var userModel model.User
	if err := user.db.Where("user_id = ?", userId).First(&userModel).Error; gorm.IsRecordNotFoundError(err) {
		return true
	}
	return false
}

func (user *UserRepository) FindUser(userId string) *model.User {
	var userModel model.User
	if err := user.db.Where("user_id = ?", userId).First(&userModel).Error; err != nil {
		return nil
	}
	return &userModel
}

func (user *UserRepository) CreateUser(body dto.LocalRegisterBody) (*model.User, *app.ErrorException) {
	tx := user.db.Begin()

	hash, err := model.HashPassword(body.Password)
	if err != nil {
		return nil, app.InteralServerErrorResponse(err)
	}

	userModel := model.User{
		UserId:   body.UserId,
		Password: hash,
	}

	if err := tx.Create(&userModel).Error; err != nil {
		tx.Rollback()
		return nil, app.InteralServerErrorResponse(err)
	}

	userProfileModal := model.UserProfile{
		Username:     body.Username,
		CountryCode:  body.CountryCode,
		Phone:        body.Phone,
		Birthday:     body.Birthday,
		ThumbnailUrl: body.ThumbnailUrl,
		SttsMsg:      body.SttsMsg,
		UserRef:      userModel.ID,
	}

	if err := tx.Create(&userProfileModal).Error; err != nil {
		tx.Rollback()
		return nil, app.InteralServerErrorResponse(err)
	}

	userMetaModel := model.UserMeta{
		UserRef: userModel.ID,
	}

	if err := tx.Create(&userMetaModel).Error; err != nil {
		tx.Rollback()
		return nil, app.InteralServerErrorResponse(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, app.InteralServerErrorResponse(err)
	}

	return &userModel, nil
}
