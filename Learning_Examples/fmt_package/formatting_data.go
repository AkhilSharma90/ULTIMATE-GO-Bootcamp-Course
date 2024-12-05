package main

import "fmt"

func main() {
	integer := 42
	float := 3.14159
	boolean := true
	str := "Gopher"

	fmt.Printf("Integer: %d\n", integer)
	fmt.Printf("Float: %.2f\n", float) // Two decimal places
	fmt.Printf("Boolean: %t\n", boolean)
	fmt.Printf("String: %s\n", str)
}

/*sample output
Integer: 42
Float: 3.14
Boolean: true
String: Gopher
*/
