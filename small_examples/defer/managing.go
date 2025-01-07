package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // Ensure file is closed
	// Perform operations on the file
	fmt.Println("File opened successfully")
}
