package types

import (
	"github.com/gofiber/websocket/v2"
)

type WsPayload struct {
	Action   string          `json:"action"`
	Username string          `json:"username"`
	Message  string          `json:"message"`
	Conn     *websocket.Conn `json:"-"`
}
type Test struct {
	Message     string          `json:"message"`
	MessageType int             `json:"message_type"`
	Conn        *websocket.Conn `json:"-"`
}

type WebSocketConnection struct {
	*websocket.Conn
}
type WsJsonRresponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}
type ErrorPage struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
}
