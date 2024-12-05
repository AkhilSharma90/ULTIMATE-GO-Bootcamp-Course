package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf bytes.Buffer
	writer := io.Writer(&buf)

	data := "Hello, Go!"
	_, err := writer.Write([]byte(data))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	fmt.Println(buf.String()) // Output the written data
}

//output - Hello, Go!
