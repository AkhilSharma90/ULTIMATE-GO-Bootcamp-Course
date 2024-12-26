package main

import (
    "fmt"
    "strconv"
)
func main() {
    num := 123
    str := strconv.Itoa(num) // Convert integer to string
    fmt.Println("String:", str) // Output: String: 123
	
    f := 3.14
    strFloat := strconv.FormatFloat(f, 'f', 2, 64) // Format float to string
    fmt.Println("String Float:", strFloat) // Output: String Float: 3.14
}