package main

import "fmt"

func worker(id int, ch chan int) {
	ch <- id * 10
}
func main() {
	ch := make(chan int)
	go worker(1, ch)
	go worker(2, ch)
	go worker(3, ch)
	// Receiving data from the channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
