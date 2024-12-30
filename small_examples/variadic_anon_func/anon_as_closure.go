package main

import "fmt"

func main() {
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Println("Counter:", increment()) // Output: 1
	fmt.Println("Counter:", increment()) // Output: 2
}
