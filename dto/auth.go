package dto

type LocalRegisterBody struct {
	UserId       string `json:"user_id" binding:"exists,alphanum,min=4,max=255"`
	Password     string `json:"password" binding:"exists,min=8,max=255"`
	Username     string `json:"username"`
	Birthday     string `json:"birthday"`
	CountryCode  string `json:"country_code"`
	Phone        string `json:"phone"`
	ThumbnailUrl string `json:"thumbnail_url" binding:"omitempty,url"`
	SttsMsg      string `json:"stts_msg" binding:"max=1024"`
}
