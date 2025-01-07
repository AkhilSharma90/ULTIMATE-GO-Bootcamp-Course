package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		for {
			ch1 <- "Ping"
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			ch2 <- "Pong"
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
