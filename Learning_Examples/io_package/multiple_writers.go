package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf1, buf2 bytes.Buffer

	writer := io.MultiWriter(&buf1, &buf2)
	data := "Hello, Go!"

	_, err := writer.Write([]byte(data))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	fmt.Println("Buffer 1:", buf1.String()) // Output data in buf1
	fmt.Println("Buffer 2:", buf2.String()) // Output data in buf2
}
