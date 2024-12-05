package main

import (
	"fmt"
	"math"
)

func main() {
	num := 7.65
	fmt.Println("Ceiling of 7.65:", math.Ceil(num)) // Round up
	fmt.Println("Floor of 7.65:", math.Floor(num))  // Round down
	fmt.Println("Round of 7.65:", math.Round(num))  // Round to nearest
}

/*
Ceiling of 7.65: 8
Floor of 7.65: 7
Round of 7.65: 8
*/
