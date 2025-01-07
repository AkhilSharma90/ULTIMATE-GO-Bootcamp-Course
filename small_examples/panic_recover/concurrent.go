package main

import (
	"fmt"
	"time"
)

func safeGoroutine() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in goroutine:", r)
		}
	}()
	fmt.Println("Goroutine started")
	panic("Goroutine panic")
}
func main() {
	go safeGoroutine()
	time.Sleep(1 * time.Second) // Allow goroutine to finish
	fmt.Println("Main program continues")
}
