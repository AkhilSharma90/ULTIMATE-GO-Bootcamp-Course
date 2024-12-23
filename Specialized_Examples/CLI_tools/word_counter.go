package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter a sentence:")
	var input string
	fmt.Scanln(&input)

	wordCount := len(strings.Fields(input))
	fmt.Printf("Word Count: %d\n", wordCount)
}
