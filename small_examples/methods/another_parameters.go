package main

import "fmt"

type Circle struct {
	Radius float64
}

// Method to update the radius
func (c *Circle) UpdateRadius(newRadius float64) {
	c.Radius = newRadius
}
func main() {
	c := Circle{Radius: 5.0}
	fmt.Println("Original Radius:", c.Radius)
	c.UpdateRadius(10.0)
	fmt.Println("Updated Radius:", c.Radius)
}
