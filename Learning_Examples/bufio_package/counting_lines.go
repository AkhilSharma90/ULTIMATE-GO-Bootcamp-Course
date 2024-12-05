package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("example.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	fmt.Printf("Total lines: %d\n", lineCount)
}
