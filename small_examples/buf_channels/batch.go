package main

import "fmt"

func batchProcessor(ch chan int, done chan bool) {
	sum := 0
	for value := range ch {
		sum += value
	}
	fmt.Println("Total Sum:", sum)
	// Notify the main function that processing is done
	done <- true
}

func main() {
	ch := make(chan int, 5)
	done := make(chan bool) // Channel for signaling completion

	for i := 1; i <= 5; i++ {
		ch <- i // Send data to the buffered channel
	}

	close(ch) // Close the channel to signal no more data will be sent

	<-done // Wait for the batchProcessor to signal completion
	go batchProcessor(ch, done)
}
