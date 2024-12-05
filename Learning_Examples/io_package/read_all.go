package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, Go!")

	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println(string(data)) // Convert bytes to string and print
}
