package controller

import (
	"connecting-server/service"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func UserProfileInfo(ctx echo.Context) error {
	id := ctx.Get("id").(string)
	db := ctx.Get("db").(*gorm.DB)
	userService := service.NewUserService(db, id)

	result, err := userService.Profile()
	if err != nil {
		return ctx.JSON(err.Code, err)
	}
	return ctx.JSON(200, result)
}
