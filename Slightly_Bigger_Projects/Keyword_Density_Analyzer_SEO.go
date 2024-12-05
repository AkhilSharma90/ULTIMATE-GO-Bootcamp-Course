package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	wordCount := make(map[string]int)

	for _, word := range words {
		word = strings.Trim(word, ".,!?\"'") // Remove punctuation
		wordCount[word]++
	}

	return wordCount
}

func calculateDensity(wordCount map[string]int, totalWords int) {
	fmt.Println("Keyword Density Analysis:")
	for word, count := range wordCount {
		density := (float64(count) / float64(totalWords)) * 100
		fmt.Printf("%s: %.2f%% (%d occurrences)\n", word, density, count)
	}
}

func main() {
	fmt.Println("Enter the path to the text file:")
	var filePath string
	fmt.Scanln(&filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + " ")
	}

	text := content.String()
	wordCount := countWords(text)
	totalWords := len(strings.Fields(text))
	calculateDensity(wordCount, totalWords)
}
