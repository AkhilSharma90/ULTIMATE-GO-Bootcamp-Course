package main

import "fmt"

func main() {
	// Print without newline
	fmt.Print("Hello, ")

	// Print with newline
	fmt.Println("World!")

	// Print with formatting
	name := "Alice"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

/*sample output
Hello, World!
Name: Alice, Age: 30
*/
