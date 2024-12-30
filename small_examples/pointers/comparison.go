package main

import "fmt"

func main() {
	a, b := 10, 20
	ptr1 := &a
	ptr2 := &b
	ptr3 := &a
	fmt.Println(ptr1 == ptr2) // Output: false
	fmt.Println(ptr1 == ptr3) // Output: true
}
