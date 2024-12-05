package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter the file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Print("Enter the keyword to search: ")
	var keyword string
	fmt.Scanln(&keyword)

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Printf("Line %d: %s\n", lineNumber, line)
			found = true
		}
		lineNumber++
	}

	if !found {
		fmt.Println("No matches found.")
	}
}
