package main

import "fmt"

type MyInt int

// Method on the custom type
func (m MyInt) Double() MyInt {
	return m * 2
}
func main() {
	num := MyInt(5)
	fmt.Println("Double:", num.Double()) // Output: 10
}
