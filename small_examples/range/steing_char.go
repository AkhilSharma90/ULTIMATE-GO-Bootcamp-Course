package main

import "fmt"

func main() {
	text := "Hello, 世界!"
	count := 0
	for range text {
		count++
	}
	fmt.Printf("The string '%s' has %d characters.\n", text, count)
}
