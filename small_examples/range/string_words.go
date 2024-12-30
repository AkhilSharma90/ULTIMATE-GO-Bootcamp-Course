package main

import (
	"fmt"
	"unicode"
)

func main() {
	text := "Go is fast and fun!"
	words := []string{}
	word := ""
	for _, r := range text {
		if unicode.IsSpace(r) {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	fmt.Println("Words:", words)
}
