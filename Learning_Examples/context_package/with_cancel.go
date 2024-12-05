package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker received cancellation signal:", ctx.Err())
			return
		default:
			fmt.Println("Worker is working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second) // Let the worker run for 2 seconds
	cancel()                    // Cancel the context
	time.Sleep(1 * time.Second) // Allow some time for cleanup
}
