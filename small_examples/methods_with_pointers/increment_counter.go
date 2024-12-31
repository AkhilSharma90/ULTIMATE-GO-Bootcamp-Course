package main

import "fmt"

type Counter struct {
	Value int
}

// Method to increment the counter
func (c *Counter) Increment() {
	c.Value++
}
func main() {
	counter := Counter{}
	counter.Increment()
	fmt.Println("Counter value:", counter.Value) // Output: 1
}
