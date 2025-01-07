package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

// Method with a value receiver
func (r Rectangle) ModifyWidth(newWidth int) {
	r.Width = newWidth // This changes only the copy, not the original
}

// Method with a pointer receiver
func (r *Rectangle) ModifyHeight(newHeight int) {
	r.Height = newHeight // This changes the original
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	// Call value receiver method
	rect.ModifyWidth(20)
	fmt.Println("After ModifyWidth:", rect.Width) // Output: 10 (unchanged)

	// Call pointer receiver method
	rect.ModifyHeight(15)
	fmt.Println("After ModifyHeight:", rect.Height) // Output: 15 (changed)
}
