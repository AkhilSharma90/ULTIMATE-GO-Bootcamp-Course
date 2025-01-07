package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task canceled")
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go task(ctx)
	time.Sleep(4 * time.Second)
}
