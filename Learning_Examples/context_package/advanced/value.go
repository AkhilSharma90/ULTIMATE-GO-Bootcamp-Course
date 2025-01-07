package main

import (
	"context"
	"fmt"
)

func worker(ctx context.Context) {
	userID := ctx.Value("userID").(string)
	fmt.Println("Worker running for user:", userID)
}
func main() {
	ctx := context.WithValue(context.Background(), "userID", "12345")
	worker(ctx)
}
