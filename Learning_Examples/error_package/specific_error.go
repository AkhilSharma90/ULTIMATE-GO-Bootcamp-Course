package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func getResource() error {
	return fmt.Errorf("resource error: %w", ErrNotFound)
}

func main() {
	err := getResource()
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Specific error detected: Not Found")
	} else {
		fmt.Println("Different error occurred")
	}
}

//output - Specific error detected: Not Found
