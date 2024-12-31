package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Before scaling:", rect)
	rect.Scale(2) // Modify the original instance
	fmt.Println("After scaling:", rect)
}
