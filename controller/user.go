package controller

import (
	"connecting-server/lib"
	"github.com/labstack/echo/v4"
)

func UserProfileInfo(ctx echo.Context) error {
	return ctx.JSON(200, lib.JSON{
		"id": ctx.Get("id"),
	})
}
