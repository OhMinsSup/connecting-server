package user

import (
	"connecting-server/controller"
	"github.com/labstack/echo/v4"
)

func ApplyRoutes(e *echo.Group) {
	user := e.Group("/user")

	user.GET("/profile", controller.UserProfileInfo)
}
