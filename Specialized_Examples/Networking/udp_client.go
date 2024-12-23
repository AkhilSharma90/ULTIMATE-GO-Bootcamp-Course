package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr := net.UDPAddr{
		Port: 8082,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.DialUDP("udp", nil, &serverAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	message := []byte("Hello, UDP Server!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
