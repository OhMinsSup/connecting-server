package controller

import "github.com/labstack/echo/v4"

func LocalRegister(ctx echo.Context) error {
	ctx.JSON(200, "Hello world")
	return nil
}

