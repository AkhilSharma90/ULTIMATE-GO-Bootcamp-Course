package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Data ready"
	}()
	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("No data yet, moving on...")
	}
}
