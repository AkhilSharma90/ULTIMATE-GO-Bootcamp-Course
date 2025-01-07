package main

import (
	"context"
	"fmt"
	"time"
)

func apiCall(ctx context.Context) {
	select {
	case <-time.After(4 * time.Second):
		fmt.Println("API call succeeded")
	case <-ctx.Done():
		fmt.Println("API call timed out")
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go apiCall(ctx)
	time.Sleep(5 * time.Second)
}
