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
	endpointUrl := "ws://localhost:80/ws"

	conn, _, err := websocket.DefaultDialer.Dial(endpointUrl, nil)
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
	time.Sleep(2 * time.Second)

	send(conn, "What you up to?")
	time.Sleep(3 * time.Second)

	send(conn, "OK, bye now")
	time.Sleep(10 * time.Second)
}

func send(conn *websocket.Conn, message string) {
	fmt.Println("Client send:", message)

	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		fmt.Println(err)
	}
}
