package service

import (
	"connecting-server/app"
	"connecting-server/repository"
	"errors"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
	id string
}

func NewUserService(db *gorm.DB, id string) *UserService {
	return &UserService{
		db: db,
		id: id,
	}
}

func (user *UserService) Profile() (*repository.UserRawQuery, *app.ErrorException) {
	userRepository := repository.NewUserRepository(user.db)

	result := userRepository.Profile(user.id)
	if result == nil {
		return nil, app.NotFoundErrorResponse(errors.New("유저가 존재하지 않습니다."))
	}
	
	return result, nil
}
