package main

import "fmt"

func printData(data interface{}) {
	switch v := data.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case []int:
		fmt.Println("Slice of integers:", v)
	default:
		fmt.Println("Unknown type")
	}
}
func main() {
	printData(42)
	printData("Hello, Go!")
	printData([]int{1, 2, 3})
}
