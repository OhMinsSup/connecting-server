package v1

import (
	"connecting-server/api/v1/auth"
	"connecting-server/api/v1/user"
	"github.com/labstack/echo/v4"
)

func ApplyRoutes(e *echo.Group) {
	v1 := e.Group("/v1.0")

	auth.ApplyRoutes(v1)
	user.ApplyRoutes(v1)
}
