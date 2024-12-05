package main

import "fmt"

func main() {
	num := 42
	fmt.Printf("|%6d|\n", num)  // Right-align, 6 spaces wide
	fmt.Printf("|%-6d|\n", num) // Left-align, 6 spaces wide

	str := "Go"
	fmt.Printf("|%10s|\n", str)  // Right-align, 10 spaces wide
	fmt.Printf("|%-10s|\n", str) // Left-align, 10 spaces wide
}

/*sample output
|    42|
|42    |
|        Go|
|Go        |
*/
