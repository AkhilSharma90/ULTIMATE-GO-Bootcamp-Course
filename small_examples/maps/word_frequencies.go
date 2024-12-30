package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is awesome go is fun"
	words := strings.Fields(text)
	wordCount := make(map[string]int)

	// //map is called wordCount
	// "go": 2
	// "is": 2
	// "awesome": 1
	// "fun": 1

	// Count word frequencies
	for i := 0; i < len(words); i++ {
		word := words[i]
		wordCount[word]++
	}
	// Collect keys in a slice
	keys := []string{}
	for key := range wordCount {
		keys = append(keys, key)
	}
	// Iterate over keys and print word frequencies
	fmt.Println("Word frequencies:")
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		fmt.Printf("%s: %d\n", key, wordCount[key])
	}
}
