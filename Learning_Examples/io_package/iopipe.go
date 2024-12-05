package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r, w := io.Pipe()

	// Writer Goroutine
	go func() {
		defer w.Close()
		data := "Hello, Pipe!"
		w.Write([]byte(data))
	}()

	// Reader
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		fmt.Println("Error copying:", err)
		return
	}

	fmt.Println(buf.String())
}

//output - Hello, Pipe!
