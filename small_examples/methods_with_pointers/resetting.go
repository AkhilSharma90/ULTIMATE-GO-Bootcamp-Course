package main

import "fmt"

type Counter struct {
	Value int
}

// Method with a pointer receiver to reset the counter
func (c *Counter) Reset() {
	c.Value = 0
}
func main() {
	counter := Counter{Value: 42}
	counter.Reset()                                    // Resets the counter
	fmt.Println("Counter after reset:", counter.Value) // Output: 0
}
