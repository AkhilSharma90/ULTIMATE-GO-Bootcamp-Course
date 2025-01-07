package main

import (
	"fmt"
	"strings"
)

func toUpper(input string, ch chan string) {
	ch <- strings.ToUpper(input)
}
func main() {
	words := []string{"go", "is", "awesome"}
	ch := make(chan string)
	for _, word := range words {
		go toUpper(word, ch)
	}
	for range words {
		fmt.Println("Processed:", <-ch)
	}
}
