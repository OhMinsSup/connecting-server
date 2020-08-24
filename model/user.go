package model

import "time"

type User struct {
	ID          string      `gorm:"primary_key;uuid" json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at"`
	UserProfile UserProfile `gorm:"foreignkey:UserRef" json:"user_profile"`
	UserMeta    UserMeta    `gorm:"foreignkey:UserRef" json:"user_meta"`
}

type UserProfile struct {
	ID         string     `gorm:"primary_key;uuid" json:"id"`
	UserId     string     `sql:"index" json:"user_id"`
	Phone      string     `sql:"index" json:"phone"`
	Username   string     `json:"username"`
	Birthday   string     `json:"birthday"`
	Thumbnail  string     `json:"thumbnail"`
	CoverImage string     `json:"cover_image"`
	SttsMsg    string     `gorm:"type:text" json:"stts_msg"` // status message
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	UserRef    string     `sql:"type:uuid" json:"user_ref"`
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
