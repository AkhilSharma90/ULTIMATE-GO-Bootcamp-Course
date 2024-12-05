package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task stopped due to:", ctx.Err())
	case <-time.After(2 * time.Second):
		fmt.Println("Task completed successfully")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Ensure proper cleanup of context resources

	go task(ctx)

	time.Sleep(3 * time.Second) // Let the program wait
}
