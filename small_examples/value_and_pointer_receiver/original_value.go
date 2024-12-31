package main

import "fmt"

type Circle struct {
	Radius float64
}

// Method with a value receiver
func (c Circle) DoubleRadius() {
	c.Radius *= 2 // Modify the copy
}
func main() {
	circle := Circle{Radius: 5.0}
	circle.DoubleRadius()
	fmt.Println("Circle radius:", circle.Radius) // Output: 5.0
}
