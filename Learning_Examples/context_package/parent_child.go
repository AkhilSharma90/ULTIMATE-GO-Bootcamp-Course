package main

//Here’s an example of a parent function passing a context to a child function.

import (
	"context"
	"fmt"
)

func printMessage(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled:", ctx.Err())
		return
	default:
		fmt.Println("Hello from the function!")
	}
}

func main() {
	ctx := context.Background() // Create a base context
	printMessage(ctx)
}
