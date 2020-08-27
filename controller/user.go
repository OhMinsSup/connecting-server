package controller

import (
	"connecting-server/service"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func UserProfileInfo(ctx echo.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	id := ctx.Get("id").(string)
	userService := service.NewUserService(db, id)

	result, err := userService.Profile()
	if err != nil {
		return ctx.JSON(err.Code, err)
	}
	return ctx.JSON(200, result)
}
