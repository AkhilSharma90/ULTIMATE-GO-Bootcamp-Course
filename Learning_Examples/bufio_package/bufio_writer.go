package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Hello, World!\n")
	writer.Flush() // Ensure all data is written to the file
}
