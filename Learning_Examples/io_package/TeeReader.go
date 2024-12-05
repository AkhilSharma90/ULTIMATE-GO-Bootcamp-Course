package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	src := bytes.NewReader([]byte("Hello, Go!"))
	var log bytes.Buffer

	reader := io.TeeReader(src, &log)
	buf := make([]byte, 8)
	_, err := reader.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Read data:", string(buf))
	fmt.Println("Logged data:", log.String())
}

/*
Read data: Hello, G
Logged data: Hello, G
*/
