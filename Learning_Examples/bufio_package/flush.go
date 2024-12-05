package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("output.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 1; i <= 5; i++ {
		writer.WriteString(fmt.Sprintf("Line %d\n", i))
	}
	writer.Flush()
}
