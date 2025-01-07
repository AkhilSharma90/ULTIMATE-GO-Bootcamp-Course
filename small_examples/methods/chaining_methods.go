package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

// Method to set width
func (r *Rectangle) SetWidth(width int) *Rectangle {
	r.Width = width
	return r
}

// Method to set height
func (r *Rectangle) SetHeight(height int) *Rectangle {
	r.Height = height
	return r
}
func main() {
	rect := Rectangle{}
	rect.SetWidth(10).SetHeight(20)      // Chain methods
	fmt.Printf("Rectangle: %+v\n", rect) // Output: {Width: 10 Height: 20}
}
