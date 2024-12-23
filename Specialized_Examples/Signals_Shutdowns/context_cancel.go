package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker shutting down...")
			return
		default:
			fmt.Println("Worker is processing...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Run a worker Goroutine
	go worker(ctx)

	// Set up signal handling
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan // Block until a signal is received
	fmt.Println("Signal received. Initiating shutdown...")
	cancel()                    // Cancel the context
	time.Sleep(2 * time.Second) // Wait for worker to shut down
	fmt.Println("Application exiting.")
}
