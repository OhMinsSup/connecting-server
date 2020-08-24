package auth

import (
	"connecting-server/controller"
	"github.com/labstack/echo/v4"
)

func ApplyRoutes(e *echo.Group) {
	auth := e.Group("/auth")

	auth.POST("/register/local", controller.LocalRegister)
}
