package main

import "fmt"

func task(done chan bool) {
	fmt.Println("Task started")
	// Simulate some work
	done <- true
}
func main() {
	done := make(chan bool)
	go task(done)
	// Wait for task to complete
	<-done
	fmt.Println("Task completed")
}
