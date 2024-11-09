package connection

import (
	"context"
	"github.com/gorilla/websocket"
	"sync"
)

type WebSocketConnection struct {
	Conn     *websocket.Conn
	CtxClose context.CancelFunc
	mutex    sync.Mutex
}

func (ws *WebSocketConnection) GetID() string {
	return ws.Conn.RemoteAddr().String()
}

func (ws *WebSocketConnection) Write(msg []byte) error {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	err := ws.Conn.WriteMessage(websocket.BinaryMessage, msg)

	return err
}

func (ws *WebSocketConnection) Close() {
	if ws.CtxClose != nil {
		ws.CtxClose()
	}
}
