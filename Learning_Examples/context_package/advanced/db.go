package main

import (
	"context"
	"fmt"
	"time"
)

func queryDatabase(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Query successful")
	case <-ctx.Done():
		fmt.Println("Query canceled")
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go queryDatabase(ctx)
	time.Sleep(4 * time.Second)
}
