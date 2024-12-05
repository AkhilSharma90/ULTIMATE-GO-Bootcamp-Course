package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	data := "Hello, Go!"
	reader := strings.NewReader(data)

	buf := make([]byte, 5) // Buffer to read 5 bytes at a time
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Print(string(buf[:n])) // Print read bytes
	}
}

//output - Hello, Go!
