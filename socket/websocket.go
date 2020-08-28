package socket

import (
	"connecting-server/lib"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	SOCKET_MAX_MESSAGE_SIZE_KB  = 8 * 1024 // 8KB
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

func ConnectWebSocket(ctx echo.Context) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  SOCKET_MAX_MESSAGE_SIZE_KB,
		WriteBufferSize: SOCKET_MAX_MESSAGE_SIZE_KB,
		CheckOrigin:     OriginChecker("*"),
	}

	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		lib.DefaultErrorLog("webSocket connecting err.", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	defer ws.Close()

	//wc := NewWebConn(ws, ctx)

	return nil
	//for {
	//	// Write
	//	err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
	//	if err != nil {
	//		lib.DefaultErrorLog("webSocket connecting err.", err)
	//		return echo.NewHTTPError(http.StatusInternalServerError, "connect web_socket.connect.upgrade.app_error")
	//	}
	//
	//	// Read
	//	_, msg, err := ws.ReadMessage()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("%s\n", msg)
	//}
}
