package service

import (
	"connecting-server/app"
	"connecting-server/dto"
	"connecting-server/model"
	"connecting-server/repository"
	"errors"
	"github.com/jinzhu/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (auth *AuthService) LocalRegisterService(body dto.LocalRegisterBody) (*model.User, *app.ErrorException) {
	userRepository := repository.NewUserRepository(auth.db)

	if exists := userRepository.FindExists(body.UserId); !exists {
		return nil, app.AlreadyExistsErrorResponse(errors.New("이미 존재하는 유저 아이디 입니다."))
	}

	user, err := userRepository.CreateUser(body)
	if err != nil {
		return nil, err
	}

	return user, nil
}
