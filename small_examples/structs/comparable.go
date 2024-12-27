package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 2}
	fmt.Println("Are points equal?", p1 == p2) // Output: true
}
