package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Authorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		exists := ctx.Get("id")
		log.Println(ctx.Get("id"))
		if exists != nil {
			return next(ctx)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "로그인후 이용해주세요")
	}
}
