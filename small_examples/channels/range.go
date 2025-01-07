package main

import "fmt"

func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // Close the channel to stop range iteration
}
func main() {
	ch := make(chan int)
	go sendNumbers(ch)
	for num := range ch {
		fmt.Println("Received:", num)
	}
}
