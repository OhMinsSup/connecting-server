package middlewares

import (
	"connecting-server/lib"
	"connecting-server/model"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
	"time"
)

func refresh(ctx echo.Context, refreshToken string) (string, error) {
	decodeTokenData, errDecode := lib.DecodeToken(refreshToken)
	if errDecode != nil {
		return "", errors.New("INVALID_TOKEN")
	}

	db := ctx.Get("db").(*gorm.DB)
	payload := decodeTokenData["payload"].(map[string]interface{})

	var userModel model.User
	if err := db.Where("id = ?", payload["user_id"].(string)).First(&userModel).Error; err != nil {
		return "", err
	}

	tokenId := payload["token_id"].(string)
	exp := int64(decodeTokenData["exp"].(float64))

	tokens := userModel.RefreshUserToken(tokenId, exp, refreshToken)
	lib.SetCookie(lib.GetEnvWithKey("APP_ENV"), tokens, ctx)
	return payload["token_id"].(string), nil
}

func ConsumeUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Path() == "/auth/logout" {
			return next(ctx)
		}

		var accessToken string
		var refreshToken string

		accessCookie, errAccess := ctx.Cookie("access_token")
		if errAccess != nil {
			authorization := ctx.Request().Header.Get("Authorization")
			if authorization != "" {
				sp := strings.Split(authorization, "Bearer ")
				// invalid token
				if len(sp) < 1 {
					return next(ctx)
				}
				accessToken = sp[1]
			}
		} else {
			accessToken = accessCookie.Value
		}

		refreshCookie, errRefresh := ctx.Cookie("refresh_token")
		if errRefresh != nil {
			return next(ctx)
		} else {
			refreshToken = refreshCookie.Value
		}

		decodeTokenData, errDecode := lib.DecodeToken(accessToken)
		if errDecode != nil {
			id, err := refresh(ctx, refreshToken)
			if err != nil {
				ctx.Set("id", nil)
				return next(ctx)
			}
			ctx.Set("id", id)
			return next(ctx)
		}

		payload := decodeTokenData["payload"].(map[string]interface{})
		tokenExpire := int64(decodeTokenData["exp"].(float64))
		now := time.Now().Unix()
		diff := tokenExpire - now

		if diff < 60*60 && refreshToken != "" {
			log.Println("refreshToken")
			id, err := refresh(ctx, refreshToken)
			if err != nil {
				ctx.Set("id", nil)
				return next(ctx)
			}

			ctx.Set("id", id)
			return next(ctx)
		}

		ctx.Set("id", payload["user_id"])
		return next(ctx)
	}
}
