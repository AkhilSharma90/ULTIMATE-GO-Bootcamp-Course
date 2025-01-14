package main

import (
	"fmt"
	"time"
)

func worker(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Task done"
}
func main() {
	ch := make(chan string, 1)
	go worker(ch)
	time.Sleep(2 * time.Second)
	if len(ch) == 0 {
		fmt.Println("Task timeout")
	} else {
		fmt.Println(<-ch)
	}
}
