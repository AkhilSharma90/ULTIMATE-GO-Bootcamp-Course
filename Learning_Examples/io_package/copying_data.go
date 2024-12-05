package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	src := strings.NewReader("Hello, Go!")
	dst := &strings.Builder{}

	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying:", err)
		return
	}

	fmt.Println(dst.String()) // Output the copied data
}
