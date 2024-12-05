package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define a flag
	mode := flag.String("mode", "default", "Mode of operation")

	flag.Parse()

	// Access non-flag arguments
	args := flag.Args()

	fmt.Println("Mode:", *mode)
	fmt.Println("Additional arguments:", args)
}

//go run main.go -mode=fast file1.txt file2.txt

/*
Mode: fast
Additional arguments: [file1.txt file2.txt]
*/
