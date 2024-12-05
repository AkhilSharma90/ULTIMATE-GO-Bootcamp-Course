package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	name := flag.String("name", "Guest", "Your name")
	age := flag.Int("age", 18, "Your age")

	// Customize usage
	flag.Usage = func() {
		fmt.Println("Usage of this program:")
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
}

//go run main.go -help

/*
Usage of this program:
  -age int
        Your age (default 18)
  -name string
        Your name (default "Guest")
*/
