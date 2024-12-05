package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("example.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 5) // Read 5 bytes at a time
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}
