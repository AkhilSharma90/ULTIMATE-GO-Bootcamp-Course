package main

import "fmt"

func batchProcessor(ch chan int) {
	sum := 0
	for value := range ch {
		sum += value
	}
	fmt.Println("Total Sum:", sum)
}
func main() {
	ch := make(chan int, 5)
	go batchProcessor(ch)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}
