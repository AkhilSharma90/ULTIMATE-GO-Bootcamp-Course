package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, quit chan bool) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			fmt.Printf("Worker %d shutting down\n", id)
			return
		default:
			fmt.Printf("Worker %d working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	var wg sync.WaitGroup
	quit := make(chan bool)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg, quit)
	}
	time.Sleep(2 * time.Second)
	close(quit)
	wg.Wait()
	fmt.Println("All workers shut down gracefully")
}
