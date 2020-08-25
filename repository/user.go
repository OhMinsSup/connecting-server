package repository

import (
	"connecting-server/app"
	"connecting-server/dto"
	"connecting-server/model"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (user *UserRepository) FindExists(userId string) bool {
	var userModel model.User
	if err := user.db.Where("user_id = ?", userId).First(&userModel).Error; gorm.IsRecordNotFoundError(err) {
		return true
	}
	return false
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
