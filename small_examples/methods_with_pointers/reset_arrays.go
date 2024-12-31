package main

import "fmt"

type Scores struct {
	Values [3]int
}

// Method to reset scores
func (s *Scores) Reset() {
	for i := range s.Values {
		s.Values[i] = 0
	}
}
func main() {
	scores := Scores{Values: [3]int{10, 20, 30}}
	scores.Reset()
	fmt.Println("Scores after reset:", scores.Values) // Output: [0 0 0]
}
