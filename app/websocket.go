package app

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

const (
	SOCKET_MAX_MESSAGE_SIZE_KB  = 8 * 1024 // 8KB
)

func ConnectWebSocket(ctx echo.Context) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  SOCKET_MAX_MESSAGE_SIZE_KB,
		WriteBufferSize: SOCKET_MAX_MESSAGE_SIZE_KB,
		CheckOrigin:     nil,
	}

	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			panic(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
