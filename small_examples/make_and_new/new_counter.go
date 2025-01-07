package main

import "fmt"

func incrementCounter(counter *int) {
	*counter++
}
func main() {
	counter := new(int) // Allocate memory for a counter
	fmt.Println("Initial counter:", *counter)
	incrementCounter(counter)
	incrementCounter(counter)
	fmt.Println("Final counter:", *counter)
}
