package socket

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

const wsURL = "wss://ws.blockchain.info/inv"

func Init() {
	log.Println("socket init")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	log.Printf("Connecting to %v...", wsURL)
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	// READ
	go func() {
		defer conn.Close()
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				return
			}
			log.Printf("Received message: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte("{\"op\": \"ping\"}"))
			if err != nil {
				log.Println("write: ", err)
				return
			}
			log.Println("Ping")
		case <-interrupt:
			log.Println("interrupt caught")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			conn.Close()
			return
		}
	}

}
