package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 8082,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("UDP Server running on port 8082")

	buf := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("Received: %s from %s\n", string(buf[:n]), clientAddr)
	}
}
