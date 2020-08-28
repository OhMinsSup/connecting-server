package user

import (
	"connecting-server/controller"
	"connecting-server/middlewares"
	"github.com/labstack/echo/v4"
)

func ApplyRoutes(e *echo.Group) {
	user := e.Group("/user")

	user.GET("/profile", controller.UserProfileInfo, middlewares.Authorized)
}
