package main

import "fmt"

func main() {
	rows := 2
	cols := 3
	matrix := new([2][3]int) // Allocate a 2x3 matrix
	fmt.Println("Default matrix:", *matrix)
	// Assign values
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = (i + 1) * (j + 1)
		}
	}
	fmt.Println("Updated matrix:", *matrix)
}
