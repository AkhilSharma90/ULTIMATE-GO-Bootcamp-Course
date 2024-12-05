package main

import (
	"fmt"
	"os"
)

func main() {
	// Creating a directory
	err := os.Mkdir("exampleDir", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Listing directory contents
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		fmt.Println("File/Dir:", file.Name())
	}

	// Removing the directory
	err = os.Remove("exampleDir")
	if err != nil {
		fmt.Println("Error removing directory:", err)
		return
	}
	fmt.Println("Directory removed successfully.")
}
