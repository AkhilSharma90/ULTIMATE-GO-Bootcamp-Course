package main

import "fmt"

func main() {
	var data interface{}
	data = 42
	fmt.Println("Integer:", data)
	data = "Hello, Go!"
	fmt.Println("String:", data)
	data = []int{1, 2, 3}
	fmt.Println("Slice:", data)
}
