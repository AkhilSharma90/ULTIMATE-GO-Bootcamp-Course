package main

import (
	"fmt"
	"time"
)

func printNumbers(id int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	for i := 1; i <= 3; i++ {
		go printNumbers(i)
	}
	time.Sleep(1 * time.Second)
}
