package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <logfile> <keyword>")
		return
	}

	logfile := os.Args[1]
	keyword := os.Args[2]

	file, err := os.Open(logfile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), keyword) {
			count++
		}
	}

	fmt.Printf("The keyword '%s' appeared %d times\n", keyword, count)
}
