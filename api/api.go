package api

import (
	v1 "connecting-server/api/v1"
	"github.com/labstack/echo/v4"
)

func ApplyRoutes(e *echo.Echo)  {
	api := e.Group("/api")
	v1.ApplyRoutes(api)
}
