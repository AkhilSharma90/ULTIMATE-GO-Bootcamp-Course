package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

// Method to scale the dimensions
func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}
func main() {
	rect := Rectangle{Width: 5, Height: 10}
	rect.Scale(3)
	fmt.Println("Scaled rectangle:", rect) // Output: {15 30}
}
