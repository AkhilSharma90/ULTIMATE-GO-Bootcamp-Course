package main

import "fmt"

func sendData(id int, ch chan int, done chan bool) {
	ch <- id * 10 // Send data to the channel
	done <- true  // Notify that the goroutine is done
}

func main() {
	ch := make(chan int, 5)
	done := make(chan bool) // Channel to signal when goroutines are done

	for i := 1; i <= 3; i++ {
		go sendData(i, ch, done)
	}

	// Wait for all 3 goroutines to finish
	for i := 0; i < 3; i++ {
		<-done
	}

	close(ch) // Close the channel once all goroutines are done

	// Read and print values from the channel
	for value := range ch {
		fmt.Println("Received:", value)
	}
}
