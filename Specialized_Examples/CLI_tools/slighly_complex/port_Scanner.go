package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	host := "127.0.0.1"
	for port := 1; port <= 1024; port++ {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			continue
		}
		fmt.Printf("Port %d is open\n", port)
		conn.Close()
	}
}
