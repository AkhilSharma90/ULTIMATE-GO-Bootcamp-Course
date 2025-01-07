package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello from Anonymous Goroutine!")
	}()
	time.Sleep(1 * time.Second)
}
