package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {

		ch <- "Data ready"
	}()
	time.Sleep(1 * time.Second)
	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("No data yet, moving on...")
	}
}
