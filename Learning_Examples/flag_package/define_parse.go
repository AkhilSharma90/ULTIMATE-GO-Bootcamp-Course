package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	name := flag.String("name", "Guest", "Your name")
	age := flag.Int("age", 18, "Your age")

	// Parse command-line arguments
	flag.Parse()

	// Print flag values
	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
}

//go run main.go -name=John -age=30

//Name: John Age: 30
