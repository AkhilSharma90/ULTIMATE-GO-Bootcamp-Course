package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url_list_file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%s: Error - %v\n", url, err)
			continue
		}
		fmt.Printf("%s: %s\n", url, resp.Status)
		resp.Body.Close()
	}
}
