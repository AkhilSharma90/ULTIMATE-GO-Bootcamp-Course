package main

import "fmt"

func cleanup(resource string) {
	fmt.Println("Cleaning up:", resource)
}
func main() {
	defer cleanup("Database Connection")
	defer cleanup("File Handle")
	defer cleanup("Network Connection")
	fmt.Println("Performing operations...")
}
