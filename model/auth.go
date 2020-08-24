package model

import "time"

type AuthToken struct {
	ID        string     `gorm:"primary_key;uuid" json:"id"`
	Disabled  bool       `gorm:"default:false" json:"disabled"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	User      User       `gorm:"foreignkey:UserRef" json:"user"`
	UserRef   string     `sql:"type:uuid" json:"user_ref"`
}
