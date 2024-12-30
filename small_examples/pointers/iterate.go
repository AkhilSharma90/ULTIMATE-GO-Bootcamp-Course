package main

import "fmt"

func main() {
	arr := [3]int{10, 20, 30}
	ptr := &arr
	for i := 0; i < len(*ptr); i++ {
		fmt.Printf("Element %d: %d\n", i, (*ptr)[i])
	}
}
