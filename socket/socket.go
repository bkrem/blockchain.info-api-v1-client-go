package socket

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var ops = map[string]string{
	"ping":           "ping",
	"unconfirmedTxs": "unconfirmed_sub",
	"newBlocks":      "blocks_sub",
	"txsOnAddr":      "addr_sub",
	"pingBlock":      "ping_block", // debug
	"pingTx":         "ping_tx",    // debug
}

const wsURL = "wss://ws.blockchain.info/inv"

func dial() (*websocket.Conn, error) {
	log.Printf("Connecting to %v...", wsURL)
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	return conn, err
}

func read(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			return
		}
		log.Printf("Received:\n%s", message)
	}
}

func monitor(conn *websocket.Conn) {
	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-interrupt:
			log.Println("interrupt caught, closing connection...")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
				log.Println("done triggered")
			case <-time.After(time.Second):
			}
			conn.Close()
			return
		}
	}
}

func send(conn *websocket.Conn, op string) {
	jsonMsg, err := json.Marshal(map[string]string{"op": op})
	if err != nil {
		log.Println("send .Marshal error:", err)
		return
	}
	msg := string(jsonMsg)
	log.Printf("Sending msg: %v", msg)
	err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println("send error:", err)
		return
	}
}

func sendWithArgs(conn *websocket.Conn, op string, args map[string]string) {
	args["op"] = op
	jsonArgs, err := json.Marshal(args)
	if err != nil {
		log.Println("sendWithArgs .Marshal error:", err)
		return
	}
	msg := string(jsonArgs)
	log.Printf("Sending msg (with args): %v", msg)
	err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println("sendWithArgs error:", err)
		return
	}
}

// ========================================================================

func keepAlive(conn *websocket.Conn) {
	const keepAliveInterval = 30
	ticker := time.NewTicker(time.Second * keepAliveInterval)
	for {
		select {
		case <-ticker.C:
			log.Println("keepAlive op: Ping")
			send(conn, ops["ping"])
		}
	}
}

// SubUnconfirmedTxs opens a subscription for all new bitcoin transactions.
func SubUnconfirmedTxs(conn *websocket.Conn) {
	send(conn, ops["unconfirmedTxs"])
}

// SubNewBlocks opens a subscription for any new blocks that are minted
// on the bitcoin blockchain.
func SubNewBlocks(conn *websocket.Conn) {
	send(conn, ops["newBlocks"])
}

// SubAddress opens a subscription for any new transactions for
// the passed bitcoin address `addr`.
func SubAddress(conn *websocket.Conn, addr string) {
	sendWithArgs(conn, ops["txsOnAddr"],
		map[string]string{"addr": addr})
}

// Connect initialises the secure websocket connection
// to wss://ws.blockchain.info/inv .
func Connect() {
	conn, err := dial()
	if err != nil {
		log.Fatal("dial error: ", err)
	}
	defer conn.Close()

	go keepAlive(conn)
	SubNewBlocks(conn)
	// SubAddress(conn, "1DTh7XPb42PgCFnuMHSitMPWxCfNNFej8n")

	go read(conn)
	monitor(conn)
}
