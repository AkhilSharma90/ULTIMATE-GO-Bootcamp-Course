package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	text := "Go is an open-source programming language."
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
