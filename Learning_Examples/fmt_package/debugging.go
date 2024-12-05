package main

import "fmt"

func main() {
	data := struct {
		Name  string
		Age   int
		Admin bool
	}{"Alice", 30, true}

	fmt.Printf("Detailed struct: %#v\n", data)
}

/*Sample Output
Detailed struct: struct { Name string; Age int; Admin bool }{Name:"Alice", Age:30, Admin:true}
*/
