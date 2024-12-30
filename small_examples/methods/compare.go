package main

import "fmt"

type Box struct {
	Width, Height int
}

// Method to check if two boxes are equal
func (b Box) Equals(other Box) bool {
	return b.Width == other.Width && b.Height == other.Height
}
func main() {
	b1 := Box{Width: 5, Height: 5}
	b2 := Box{Width: 5, Height: 5}
	b3 := Box{Width: 3, Height: 4}
	fmt.Println("b1 equals b2:", b1.Equals(b2)) // Output: true
	fmt.Println("b1 equals b3:", b1.Equals(b3)) // Output: false
}
