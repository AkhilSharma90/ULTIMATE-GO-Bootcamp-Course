package main

import "fmt"

func sendData(id int, ch chan string) {
	ch <- fmt.Sprintf("Data from worker %d", id)
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go sendData(1, ch1)
	go sendData(2, ch2)
	// Receiving data sequentially
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}
