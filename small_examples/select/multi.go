package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "Message from channel 1"
	}()
	go func() {
		ch2 <- "Message from channel 2"
	}()
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}
