package main

import "fmt"

type Counter struct {
	Count int
}

// Method to increment the counter
func (c *Counter) Increment() {
	c.Count++
}
func main() {
	c := Counter{}
	fmt.Println("Initial count:", c.Count)
	c.Increment()
	fmt.Println("After increment:", c.Count) // Output: 1
}
