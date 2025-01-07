package main

import (
	"errors"
	"fmt"
)

func riskyOperation(flag bool) error {
	if !flag {
		return errors.New("regular error occurred")
	}
	panic("critical error occurred")
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	err := riskyOperation(false)
	if err != nil {
		fmt.Println("Handled error:", err)
	}
	riskyOperation(true)
}
