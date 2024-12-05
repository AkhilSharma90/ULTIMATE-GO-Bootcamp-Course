package main

import (
	"context"
	"fmt"
)

func printUserID(ctx context.Context) {
	userID := ctx.Value("userID")
	if userID != nil {
		fmt.Println("User ID:", userID)
	} else {
		fmt.Println("No user ID found in context")
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "userID", 42)
	printUserID(ctx)

	// Passing a context without userID
	emptyCtx := context.Background()
	printUserID(emptyCtx)
}
