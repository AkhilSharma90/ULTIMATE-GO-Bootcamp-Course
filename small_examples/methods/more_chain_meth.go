package main

import "fmt"

type Point struct {
	X, Y int
}

// Method to move the point along X-axis
func (p *Point) MoveX(dx int) *Point {
	p.X += dx
	return p
}

// Method to move the point along Y-axis
func (p *Point) MoveY(dy int) *Point {
	p.Y += dy
	return p
}
func main() {
	p := &Point{X: 0, Y: 0}
	p.MoveX(5).MoveY(10)
	fmt.Printf("Moved Point: %+v\n", p) // Output: Moved Point: &{X:5 Y:10}
}
