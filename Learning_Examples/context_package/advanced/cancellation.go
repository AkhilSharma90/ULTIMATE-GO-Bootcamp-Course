package main

import (
	"context"
	"fmt"
	"time"
)

func stage(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s stopped\n", name)
			return
		default:
			fmt.Printf("%s processing\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go stage(ctx, "Stage 1")
	go stage(ctx, "Stage 2")
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
