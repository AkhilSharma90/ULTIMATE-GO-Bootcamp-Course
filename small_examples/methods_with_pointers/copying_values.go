package main

import "fmt"

type Point struct {
	X, Y int
}

// Method to copy values from another point
func (p *Point) CopyFrom(other *Point) {
	p.X, p.Y = other.X, other.Y
}
func main() {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{}
	p2.CopyFrom(&p1)
	fmt.Println("Copied point:", p2) // Output: {1 2}
}
