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
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go task(ctx)

	time.Sleep(3 * time.Second) // Let the program wait
}
