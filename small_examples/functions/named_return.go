package main

import "fmt"

// Function with named return values
func rectangleProperties(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Automatically returns area and perimeter
}
func main() {
	area, perimeter := rectangleProperties(5, 3)
	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)
}
