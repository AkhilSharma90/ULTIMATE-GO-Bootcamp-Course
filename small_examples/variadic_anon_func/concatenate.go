package main

import "fmt"

// Variadic function to concatenate strings
func concatenate(words ...string) string {
	result := ""
	for _, word := range words {
		result += word + " "
	}
	return result
}
func main() {
	fmt.Println(concatenate("Go", "is", "awesome!")) // Output: "Go is awesome! "
	fmt.Println(concatenate("Hello"))                // Output: "Hello "
}
