package model

import "time"

type Message struct {
	ID        string     `gorm:"primary_key;uuid;" json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Room      Room       `gorm:"foreignkey:RoomRef" json:"room"`
	RoomRef   string     `sql:"type:uuid" json:"room_ref"`
	User      User       `gorm:"foreignkey:UserRef" json:"user"`
	UserRef   string     `sql:"type:uuid" json:"user_ref"`
}
