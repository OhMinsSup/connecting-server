package socket

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// WebConn은 사용자에 대한 단일 웹 소켓 연결을 나타냅니다.
// 데이터 송수신을 관리하기위한 모든 necesarry 상태를 포함합니다.
// 웹 소켓입니다.
type WebConn struct {
	WebSocket *websocket.Conn
	context   echo.Context
	UserId    string

	endWritePump chan struct{}
	pumpFinished chan struct{}
}

func NewWebConn(ws *websocket.Conn, ctx echo.Context) *WebConn {
	//userId := ctx.Get("id")
	//log.Println("userId", userId)
	//if userId != nil {
	//	// TODO userID가 없는 경우 인증을 하지 않는 경우
	//	log.Println("인증해줘~~~~~")
	//	return nil
	//}

	wc := &WebConn{
		WebSocket:    ws,
		context:      ctx,
		//UserId:       userId.(string),
		UserId: "1d09d054-e6e1-11ea-8a95-acde48001122",
		endWritePump: make(chan struct{}),
		pumpFinished: make(chan struct{}),
	}

	return wc
}
