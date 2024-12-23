package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <name>")
		return
	}
	name := os.Args[1]
	fmt.Printf("Hello, %s!\n", name)
}
