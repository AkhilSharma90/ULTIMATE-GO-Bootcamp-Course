package main

import (
	"errors"
	"fmt"
)

func performTask() error {
	return errors.New("failed to perform task")
}

func main() {
	err := performTask()
	if err != nil {
		wrappedErr := fmt.Errorf("error in main function: %w", err)
		fmt.Println("Wrapped Error:", wrappedErr)
	}
}

//Output - Wrapped Error: error in main function: failed to perform task
