package lib

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var signingKey = []byte("AllYourBase")

func generateCookie(name, token, domain string, maxAge int, httpOnly, secure bool) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Domain = domain
	cookie.MaxAge = maxAge
	cookie.HttpOnly = httpOnly
	cookie.Path = "/"
	cookie.Secure = secure
	return cookie
}

func SetCookie(env string, tokens JSON, ctx echo.Context) {
	if tokens == nil {
		return
	}

	switch env {
	case "production":
		ctx.SetCookie(generateCookie("access_token", tokens["accessToken"].(string), "/", 60*60*24, true, true))
		ctx.SetCookie(generateCookie("refresh_token", tokens["refreshToken"].(string), "/", 60*60*24*30, true, true))
		break
	case "development":
		ctx.SetCookie(generateCookie("access_token", tokens["accessToken"].(string), "", 60*60*24, true, false))
		ctx.SetCookie(generateCookie("refresh_token", tokens["refreshToken"].(string), "", 60*60*24*30, true, false))
		break
	default:
		break
	}
}

func generateToken(payload JSON, subject string, expire time.Duration) (string, error) {
	// Create the Claims
	claims := &jwt.MapClaims{
		"exp":     time.Now().Add(expire).Unix(),
		"issuer":  "veloss",
		"subject": subject,
		"payload": payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeToken(decodeToken string) (JSON, error) {
	result, err := jwt.Parse(decodeToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !result.Valid {
		return nil, errors.New("Invalid Token Error")
	}

	return result.Claims.(jwt.MapClaims), nil
}

func GenerateAccessToken(payload JSON, subject string) (string, error) {
	return generateToken(payload, subject, time.Hour*24)
}

func GenerateRefreshToken(payload JSON, subject string) (string, error) {
	return generateToken(payload, subject, time.Hour*24*30)
}
