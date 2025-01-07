package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	for task := range ch {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go worker(1, ch1)
	go worker(2, ch2)
	for i := 1; i <= 5; i++ {
		select {
		case ch1 <- i:
		case ch2 <- i:
		}
	}
	close(ch1)
	close(ch2)
}
