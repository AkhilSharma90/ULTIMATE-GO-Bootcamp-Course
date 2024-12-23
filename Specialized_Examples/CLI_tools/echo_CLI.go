package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // Get arguments excluding the program name
	if len(args) == 0 {
		fmt.Println("Please provide some input.")
		return
	}
	fmt.Println("Echo:", args)
}
