package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}

	dir := os.Args[1]
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Files in", dir, ":")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
