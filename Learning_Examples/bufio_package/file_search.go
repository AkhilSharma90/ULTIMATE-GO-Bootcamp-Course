package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("example.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	searchWord := "Go"
	found := false

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchWord) {
			found = true
			break
		}
	}

	if found {
		fmt.Printf("The word '%s' was found in the file.\n", searchWord)
	} else {
		fmt.Printf("The word '%s' was not found in the file.\n", searchWord)
	}
}
