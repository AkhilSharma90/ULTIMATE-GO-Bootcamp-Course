package main

import "fmt"

func main() {
	ch := make(chan int)
	// Send data into the channel
	go func() {
		ch <- 42
	}()
	// Receive data from the channel
	value := <-ch
	fmt.Println("Received:", value)
}
