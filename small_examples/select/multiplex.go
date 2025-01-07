package main

import "fmt"

func sendData(ch chan string, msg string) {
	ch <- msg
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go sendData(ch1, "Data from Channel 1")
	go sendData(ch2, "Data from Channel 2")
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
