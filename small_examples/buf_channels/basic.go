package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Buffered channel with capacity 2
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // Receives 1
	fmt.Println(<-ch) // Receives 2
}
