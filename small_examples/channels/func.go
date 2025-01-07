package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}
func main() {
	ch := make(chan int)
	go square(5, ch)
	fmt.Println("Square is:", <-ch)
}
