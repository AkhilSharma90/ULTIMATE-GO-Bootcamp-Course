// func(parameters) returnType {
//     // Function body
// }

package main

import "fmt"

func main() {
	// Define and call an anonymous function
	func() {
		fmt.Println("Hello from an anonymous function!")
	}()
}
