package main

import "fmt"

func main() {
	scores := make(map[string]int, 10) // Map with initial capacity of 10
	scores["Alice"] = 95
	scores["Bob"] = 85
	fmt.Println(scores) // Output: map[Alice:95 Bob:85]
}
