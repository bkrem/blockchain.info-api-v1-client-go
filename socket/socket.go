package socket

import "github.com/gorilla/websocket"

const wsURL = "wss://ws.blockchain.info/inv"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
