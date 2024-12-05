package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if a file exists
	_, err := os.Stat("example.txt")
	if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	} else {
		fmt.Println("File exists.")
	}
}
