package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- int) { // Send-only channel
	ch <- 100
}
func receiveData(ch <-chan int) { // Receive-only channel
	fmt.Println("Received:", <-ch)
}
func main() {
	ch := make(chan int)
	go sendData(ch)
	go receiveData(ch)
	// Allow goroutines to complete
	time.Sleep(1 * time.Second)
}
