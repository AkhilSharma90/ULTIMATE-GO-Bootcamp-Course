package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a sentence:")

	// Create a scanner to read input from the user
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Count the words using strings.Fields
		wordCount := len(strings.Fields(input))
		fmt.Printf("Word Count: %d\n", wordCount)
	}

	// Handle any potential errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
