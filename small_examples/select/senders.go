package main

import "fmt"

func sendData(ch chan string, msg string) {
	ch <- msg
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go sendData(ch1, "Message 1")
	go sendData(ch2, "Message 2")
	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case msg := <-ch2:
		fmt.Println("Received:", msg)
	}
}
