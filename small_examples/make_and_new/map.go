package main

import "fmt"

func main() {
	scores := make(map[string]int) // Map with string keys and int values
	scores["Alice"] = 95
	scores["Bob"] = 85
	fmt.Println(scores) // Output: map[Alice:95 Bob:85]
}
