package model

import (
	"connecting-server/lib"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	ID          string      `gorm:"primary_key;uuid" json:"id"`
	UserId      string      `sql:"index" json:"user_id"`
	Password    string      `json:"password"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at"`
	UserProfile UserProfile `gorm:"foreignkey:UserRef" json:"user_profile"`
	UserMeta    UserMeta    `gorm:"foreignkey:UserRef" json:"user_meta"`
}

func HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(hash), err
}

func (user *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plain))
	return err == nil
}

func (user *User) GenerateUserToken(db *gorm.DB) lib.JSON {
	tx := db.Begin()

	authTokenModel := AuthToken{
		UserRef: user.ID,
	}

	if err := tx.Create(&authTokenModel).Error; err != nil {
		tx.Rollback()
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil
	}

	accessSubject := "access_token"
	accessPayload := lib.JSON{
		"user_id": user.ID,
	}
	accessToken, _ := lib.GenerateAccessToken(accessPayload, accessSubject)

	refreshSubject := "refresh_token"
	refreshPayload := lib.JSON{
		"user_id":  user.ID,
		"token_id": authTokenModel.ID,
	}
	refreshToken, _ := lib.GenerateRefreshToken(refreshPayload, refreshSubject)

	return lib.JSON{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
}

func (user *User) RefreshUserToken(tokenId string, refreshTokenExp int64, originalRefreshToken string) lib.JSON {
	now := time.Now().Unix()
	diff := refreshTokenExp - now

	log.Println("refreshing..")
	refreshToken := originalRefreshToken

	if diff < 60*60*24*15 {
		log.Println("refreshing refreshToken")
		accessSubject := "access_token"
		accessPayload := lib.JSON{
			"user_id": user.ID,
		}
		accessToken, _ := lib.GenerateAccessToken(accessPayload, accessSubject)
		log.Println("accessToken", accessToken)
		refreshSubject := "refresh_token"
		refreshPayload := lib.JSON{
			"user_id":  user.ID,
			"token_id": tokenId,
		}
		refreshToken, _ = lib.GenerateRefreshToken(refreshPayload, refreshSubject)
		log.Println("refreshToken", refreshToken)
		return lib.JSON{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		}
	}

	return nil
}

type UserProfile struct {
	ID            string     `gorm:"primary_key;uuid;" json:"id"`
	Username      string     `sql:"index" json:"username"`
	CountryCode   string     `json:"country_code"`
	Phone         string     `sql:"index" json:"phone"`
	Birthday      string     `json:"birthday"`
	ThumbnailUrl  string     `json:"thumbnail_url"`
	CoverImageUrl string     `json:"cover_image_url"`
	SttsMsg       string     `gorm:"type:text" json:"stts_msg"` // status message
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	UserRef       string     `sql:"type:uuid" json:"user_ref"`
}

type UserMeta struct {
	ID                 string     `gorm:"primary_key;uuid" json:"id"`
	ShareProfileChange bool       `gorm:"default:true" json:"share_profile_change"`
	AllowAddingFriends bool       `gorm:"default:true" json:"allow_adding_friends"`
	FollowView         bool       `gorm:"default:false" json:"follow_view"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
	UserRef            string     `sql:"type:uuid" json:"user_ref"`
}
