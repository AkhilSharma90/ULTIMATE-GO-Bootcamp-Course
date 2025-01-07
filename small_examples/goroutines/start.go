package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Goroutine!")
}
func main() {
	go sayHello() // Run as a goroutine
	fmt.Println("Hello from Main!")
	time.Sleep(1 * time.Second) // Wait to see goroutine output
}
