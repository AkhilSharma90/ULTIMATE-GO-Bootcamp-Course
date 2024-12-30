package main

import "fmt"

func main() {
	var ptr *int // Declare a pointer without initialization
	if ptr == nil {
		fmt.Println("ptr is nil") // Output: ptr is nil
	}
}
