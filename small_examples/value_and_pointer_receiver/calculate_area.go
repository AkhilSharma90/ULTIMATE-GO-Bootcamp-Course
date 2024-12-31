package main

import "fmt"

type Rectangle struct {
	Width, Height int
}

// Method with value receiver
func (r Rectangle) Area() int {
	return r.Width * r.Height
}
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area:", rect.Area()) // Output: 50
}
