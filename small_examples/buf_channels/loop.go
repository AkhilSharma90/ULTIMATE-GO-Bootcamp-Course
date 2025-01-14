package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	//close(ch) // Closing the channel allows iteration over it
	for value := range ch {
		fmt.Println(value)
	}
}
