package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch)
}
func consumer(ch chan int) {
	for value := range ch {
		fmt.Println("Consumed:", value)
	}
}
func main() {
	ch := make(chan int, 3)
	go producer(ch)
	consumer(ch)
}
