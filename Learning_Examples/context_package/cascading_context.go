package main

import (
	"context"
	"fmt"
	"time"
)

func childTask(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Child task stopped due to:", ctx.Err())
	}
}

func main() {
	parentCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	childCtx, _ := context.WithCancel(parentCtx)
	go childTask(childCtx)

	time.Sleep(3 * time.Second)
}
