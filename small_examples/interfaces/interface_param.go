package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width, Height float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	shapes := []Shape{rect, circle}
	fmt.Println("Total Area:", TotalArea(shapes))
}
