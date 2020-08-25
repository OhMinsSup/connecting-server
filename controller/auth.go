package controller

import (
	"connecting-server/app"
	"connecting-server/dto"
	"connecting-server/lib"
	"connecting-server/service"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LocalRegister(ctx echo.Context) error {
	body := new(dto.LocalRegisterBody)
	if err := ctx.Bind(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, app.BadRequestErrorResponse(err))
	}

	db := ctx.Get("db").(*gorm.DB)
	authService := service.NewAuthService(db)

	result, err := authService.LocalRegisterService(*body)
	if err != nil {
		return ctx.JSON(err.Code, err)
	}

	tokens := result.GenerateUserToken(db)
	lib.SetCookie(lib.GetEnvWithKey("APP_ENV"), tokens, ctx)
	return ctx.JSON(http.StatusOK, lib.JSON{
		"user":   result,
		"tokens": tokens,
	})
}
