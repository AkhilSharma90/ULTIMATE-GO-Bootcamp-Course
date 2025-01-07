package main

import (
	"fmt"
	"time"
)

func worker(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Task completed"
}
func main() {
	ch := make(chan string, 1)
	go worker(ch)
	fmt.Println("Waiting for task...")
	fmt.Println(<-ch)
}
