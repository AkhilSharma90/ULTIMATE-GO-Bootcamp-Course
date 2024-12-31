package main

import "fmt"

type Counter struct {
	Value int
}

// Value receiver: does not modify the receiver
func (c Counter) GetValue() int {
	return c.Value
}

// Pointer receiver: modifies the receiver
func (c *Counter) Increment() {
	c.Value++
}
func main() {
	counter := Counter{Value: 10}
	fmt.Println("Initial Value:", counter.GetValue())   // Output: 10
	counter.Increment()                                 // Modifies original
	fmt.Println("After Increment:", counter.GetValue()) // Output: 11
}
