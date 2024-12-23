package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, wg *sync.WaitGroup, stopChan chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			fmt.Printf("Worker %d shutting down...\n", id)
			return
		default:
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stopChan := make(chan struct{})

	// Start multiple workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg, stopChan)
	}

	// Set up signal handling
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan // Block until a signal is received
	fmt.Println("\nShutdown signal received. Stopping workers...")

	close(stopChan) // Signal all workers to stop
	wg.Wait()       // Wait for all workers to finish
	fmt.Println("All workers stopped. Exiting.")
}
