package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

// Method with a pointer receiver to modify the original instance
func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	rect.Scale(2)                               // Modify the original instance
	fmt.Printf("Scaled Rectangle: %+v\n", rect) // Output: {Width: 20 Height: 10}
}
