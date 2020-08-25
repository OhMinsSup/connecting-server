package repository

import (
	"connecting-server/model"
	"github.com/jinzhu/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (auth *AuthRepository) FindExists(userId string) (*model.User, bool) {
	var user model.User
	if err := auth.db.Where("userId = ?", userId).First(&user).Error; err != nil {
		return nil, false
	}
	return &user, true
}
