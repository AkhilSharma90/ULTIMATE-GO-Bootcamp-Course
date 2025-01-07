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
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Worker running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(2 * time.Second)
	cancel() // Stop the worker
	time.Sleep(1 * time.Second)
}
