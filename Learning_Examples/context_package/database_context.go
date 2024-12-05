package main

import (
	"context"
	"fmt"
	"time"
)

func queryDatabase(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second): // Simulate database delay
		fmt.Println("Query successful")
	case <-ctx.Done():
		fmt.Println("Query canceled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	queryDatabase(ctx)
}
