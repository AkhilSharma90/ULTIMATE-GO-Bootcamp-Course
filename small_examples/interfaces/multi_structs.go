package main
import (
    "fmt"
    "math"
)
type Shape interface {
    Area() float64
    Perimeter() float64
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
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
func (c Circle) Area() float64 {
    return math.Pi  c.Radius  c.Radius
}
func (c Circle) Perimeter() float64 {
    return 2  math.Pi  c.Radius
}
func main() {
    rect := Rectangle{Width: 10, Height: 5}
    circle := Circle{Radius: 7}
    shapes := []Shape{rect, circle}
    for _, shape := range shapes {
        fmt.Println("Shape Details:")
        fmt.Println("Area:", shape.Area())
        fmt.Println("Perimeter:", shape.Perimeter())
    }
}