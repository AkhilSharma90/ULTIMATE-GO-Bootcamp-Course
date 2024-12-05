package main

import (
	"flag"
	"fmt"
	"strings"
)

type csvFlag []string

func (c *csvFlag) String() string {
	return fmt.Sprint(*c)
}

func (c *csvFlag) Set(value string) error {
	*c = strings.Split(value, ",")
	return nil
}

func main() {
	var items csvFlag

	// Define a custom flag
	flag.Var(&items, "items", "Comma-separated list of items")

	flag.Parse()

	fmt.Println("Items:", items)
}

//go run main.go -items=apple,banana,orange

//Items: [apple banana orange]
