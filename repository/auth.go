package repository

import "github.com/jinzhu/gorm"

type AuthRepository struct {
	db *gorm.DB
}
