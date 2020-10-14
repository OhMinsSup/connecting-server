package repository

import (
	"connecting-server/app"
	"connecting-server/model"
	"github.com/jinzhu/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (msg *MessageRepository) CreateMessage() (*model.Message, *app.ErrorException) {
	tx := msg.db.Begin()

	msgModel := model.Message{}

	if err := tx.Create(&msgModel).Error; err != nil {
		tx.Rollback()
		return nil, app.InteralServerErrorResponse(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, app.InteralServerErrorResponse(err)
	}

	return &msgModel, nil
}
