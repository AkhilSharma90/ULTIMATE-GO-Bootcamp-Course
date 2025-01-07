package main

import (
	"fmt"
	"time"
)

func trackExecutionTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}
func main() {
	defer trackExecutionTime(time.Now())
	// Simulate some work
	time.Sleep(2 * time.Second)
	fmt.Println("Work completed")
}
