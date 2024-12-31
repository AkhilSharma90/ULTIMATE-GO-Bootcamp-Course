package main

import "fmt"

type Pair struct {
	A, B int
}

// Method to swap the values
func (p *Pair) Swap() {
	p.A, p.B = p.B, p.A
}
func main() {
	pair := Pair{A: 10, B: 20}
	pair.Swap()
	fmt.Println("After swap:", pair) // Output: {20 10}
}
