package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	conn := createConnection()

	// Read within go routine
	go consume(conn)

	// For demo purposes
	sendSomeMessages(conn)
}

func createConnection() *websocket.Conn {
	endpointURL := "ws://localhost:80/ws"

	conn, _, err := websocket.DefaultDialer.Dial(endpointURL, nil)
	if err != nil {
		panic(err)
	}

	return conn
}

func consume(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Println("Server says:", string(message))
	}
}

func sendSomeMessages(conn *websocket.Conn) {
	send(conn, "Hello server, this is client")
	time.Sleep(time.Second * 2) //nolint:gomnd

	send(conn, "What you up to?")
	time.Sleep(time.Second * 3) //nolint:gomnd

	send(conn, "OK, bye now")
	time.Sleep(time.Second * 10) //nolint:gomnd
}

func send(conn *websocket.Conn, message string) {
	fmt.Println("Client send:", message)

	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		fmt.Println(err)
	}
}
