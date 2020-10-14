package socket

import (
	"connecting-server/lib"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	SOCKET_MAX_MESSAGE_SIZE_KB = 8 * 1024 // 8KB
)

func CheckOrigin(r *http.Request, allowedOrigins string) bool {
	origin := r.Header.Get("Origin")
	if origin == "" {
		return true
	}

	if allowedOrigins == "*" {
		return true
	}

	for _, allowed := range strings.Split(allowedOrigins, " ") {
		if allowed == origin {
			return true
		}
	}
	return false
}

func OriginChecker(allowedOrigins string) func(*http.Request) bool {
	return func(r *http.Request) bool {
		return CheckOrigin(r, allowedOrigins)
	}
}

func ConnectWebSocket(ctx echo.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  SOCKET_MAX_MESSAGE_SIZE_KB,
		WriteBufferSize: SOCKET_MAX_MESSAGE_SIZE_KB,
		CheckOrigin:     OriginChecker("*"),
	}

	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		lib.DefaultErrorLog("webSocket connecting err:", err)
		panic(err)
	}

	db := ctx.Get("db").(*gorm.DB)
	userId := ctx.Get("id").(string)

	NewWebConn(ws, db, ctx, "", userId)
}
