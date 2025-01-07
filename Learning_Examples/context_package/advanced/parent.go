package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s stopped\n", name)
			return
		default:
			fmt.Printf("%s running\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	parentCtx, cancel := context.WithCancel(context.Background())
	childCtx, _ := context.WithTimeout(parentCtx, 2*time.Second)
	go worker(childCtx, "Child Worker")
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
