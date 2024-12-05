package main

import (
	"errors"
	"fmt"
)

func main() {
	originalErr := errors.New("original error")
	wrappedErr := fmt.Errorf("wrapping error: %w", originalErr)

	fmt.Println("Wrapped Error:", wrappedErr)
	fmt.Println("Unwrapped Error:", errors.Unwrap(wrappedErr))
}

/*
Wrapped Error: wrapping error: original error
Unwrapped Error: original error
*/
