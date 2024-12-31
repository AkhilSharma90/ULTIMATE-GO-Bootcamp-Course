package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct
type Rectangle struct {
	Width, Height float64
}

// Implement the methods of the interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	var shape Shape = rect // Rectangle satisfies the Shape interface
	fmt.Println("Area:", shape.Area())
	fmt.Println("Perimeter:", shape.Perimeter())
}
