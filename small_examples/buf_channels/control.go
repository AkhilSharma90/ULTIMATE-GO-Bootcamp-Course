package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan bool) {
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d finished\n", id)
	<-ch // Release a slot
}
func main() {
	ch := make(chan bool, 2) // Limit concurrency to 2
	for i := 1; i <= 5; i++ {
		ch <- true
		go worker(i, ch)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("All workers completed")
}
