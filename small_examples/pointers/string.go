package main

import "fmt"

func main() {
	str := "Hello, Go!"
	ptr := &str                        // Pointer to string
	fmt.Println("Value of str:", *ptr) // Output: Hello, Go!
	*ptr = "Hello, Pointers!"          // Modify string through pointer
	fmt.Println("Updated value of str:", str)
}
