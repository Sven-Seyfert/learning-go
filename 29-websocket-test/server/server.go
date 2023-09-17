package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

var upgrader = websocket.Upgrader{}

func main() {
	// Register endpoint and listen for connections.
	http.HandleFunc("/ws", register)
	if err := http.ListenAndServe("localhost:80", nil); err != nil {
		panic(err)
	}
}

func register(writer http.ResponseWriter, req *http.Request) {
	// Upgrade to websocket
	conn, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Read within go routine
	go consume(conn)

	// For demo purposes
	sendSomeMessages(conn)
}

func consume(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Println("Client says:", string(message))
	}
}

func sendSomeMessages(conn *websocket.Conn) {
	time.Sleep(time.Second * 5) //nolint:gomnd
	send(conn, "Hello from the server")

	time.Sleep(time.Second)
	send(conn, "It's lovely out today, isn't it")

	time.Sleep(time.Second * 2) //nolint:gomnd
	send(conn, "OK, See you.")
}

func send(conn *websocket.Conn, message string) {
	fmt.Println("Server send:", message)

	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		fmt.Println(err)
	}
}
