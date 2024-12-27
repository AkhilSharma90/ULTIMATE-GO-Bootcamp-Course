package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{4, 5, 6}

	fmt.Println("arr1 == arr2:", arr1 == arr2) // True
	fmt.Println("arr1 == arr3:", arr1 == arr3) // False
}
