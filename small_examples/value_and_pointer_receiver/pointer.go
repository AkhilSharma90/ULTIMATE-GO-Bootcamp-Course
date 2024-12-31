package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}
func main() {
	rect := &Rectangle{Width: 10, Height: 5} // Pointer to struct
	rect.Scale(2)                            // Call directly on pointer
	fmt.Println("Scaled rectangle:", rect)   // Output: &{20 10}
}
