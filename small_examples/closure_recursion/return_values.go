package main

import "fmt"

func main() {
	multiplier := func(factor int) func(int) int {
		return func(input int) int {
			return input * factor
		}
	}
	double := multiplier(2)
	triple := multiplier(3)
	fmt.Println("Double 5:", double(5)) // Output: Double 5: 10
	fmt.Println("Triple 4:", triple(4)) // Output: Triple 4: 12
}
