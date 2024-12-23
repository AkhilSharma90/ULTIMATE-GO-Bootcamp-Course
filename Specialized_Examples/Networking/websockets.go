package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)

		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println("Error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("WebSocket server running on http://localhost:8080/ws")
	http.ListenAndServe(":8080", nil)
}
