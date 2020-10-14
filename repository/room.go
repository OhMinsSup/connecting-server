package repository

import (
	"connecting-server/app"
	"connecting-server/model"
	"github.com/jinzhu/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (room *RoomRepository) createRoom(name string) (*model.Room, *app.ErrorException) {
	tx := room.db.Begin()

	roomModel := model.Room{
		Name: name,
	}

	if err := tx.Create(&roomModel).Error; err != nil {
		tx.Rollback()
		return nil, app.InteralServerErrorResponse(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, app.InteralServerErrorResponse(err)
	}

	return &roomModel, nil
}
