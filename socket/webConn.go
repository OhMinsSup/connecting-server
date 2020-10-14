package socket

import (
	"connecting-server/model"
	"connecting-server/repository"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
)

// WebConn은 사용자에 대한 단일 웹 소켓 연결을 나타냅니다.
// 데이터 송수신을 관리하기위한 모든 necesarry 상태를 포함합니다.
// 웹 소켓입니다.
type WebConn struct {
	db        *gorm.DB
	webSocket *websocket.Conn
	context   echo.Context
	send      chan *model.Message
	userRef   string
	roomRef   string
}

// 현재 접속 중인 전체 클라이언트 리스트
var clients []*WebConn

const messageBufferSize = 256

func NewWebConn(ws *websocket.Conn, db *gorm.DB, context echo.Context, roomRef, userRef string) {
	wc := &WebConn{
		db:        db,
		webSocket: ws,
		context:   context,
		send:      make(chan *model.Message, messageBufferSize),
		userRef:   userRef,
		roomRef:   roomRef,
	}

	clients = append(clients, wc)

	go wc.readLoop()
	go wc.writeLoop()
}

func (ws *WebConn) Close() {
	for i, client := range clients {
		if client == ws {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}

	close(ws.send)
	ws.webSocket.Close()
	log.Printf("close connection. addr: %s", ws.webSocket.RemoteAddr())
}

func (wc *WebConn) readLoop() {
	msgRepository := repository.NewMessageRepository(wc.db)
	// 메세지 수신 대기
	for {
		msg, err := wc.read()
		if err != nil {
			// 오류가 발생하면 메세지 수신 루프 종료
			log.Println("read message error:", err)
			break
		}

		msgRepository.CreateMessage()
		wc.broadcast(msg)
	}
	wc.Close()
}

func (wc *WebConn) writeLoop() {
	//	클라이언트의 send 채널 메세지 수신 대기
	for msg := range wc.send {
		if wc.roomRef == msg.RoomRef {
			wc.write(msg)
		}
	}
}

func (wc *WebConn) broadcast(m *model.Message) {
	for _, client := range clients {
		client.send <- m
	}
}

func (wc *WebConn) read() (*model.Message, error) {
	var msg *model.Message

	// 웹소켓 커넥션에 JSON 형태의 메세지가 전달되면 Message 타입으로 메세지를 읽음
	if err := wc.webSocket.ReadJSON(&msg); err != nil {
		return nil, err
	}

	tx := wc.db.Begin()
	if err := tx.Update(model.Message{
		UserRef: wc.userRef,
	}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Message 에 현재 시간과 사용자 정보 세팅
	log.Println("read from websocket:", msg)
	return msg, tx.Commit().Error
}

func (wc *WebConn) write(m *model.Message) error {
	log.Println("write to websocket:", m)
	// 웹 소켓 커넥션에 JSON 형태로 메세지 세팅
	return wc.webSocket.WriteJSON(m)
}
