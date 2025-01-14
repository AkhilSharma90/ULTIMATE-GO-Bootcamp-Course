package main

import (
	"fmt"
	"time"
)

func worker(ch chan int, quit chan bool) {
	for {
		select {
		case task := <-ch:
			fmt.Println("Processing task:", task)
		case <-quit:
			fmt.Println("Worker shutting down")
			return
		}
	}
}
func main() {
	ch := make(chan int, 3)
	quit := make(chan bool)
	go worker(ch, quit)
	for i := 1; i <= 3; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	quit <- true
	time.Sleep(1 * time.Second)
}
