// func main() {
//     // Anonymous function assigned to a variable
//     closure := func() {
//         // Accessing variables from the surrounding scope
//     }
//     closure()
// }

//example -

package main

import "fmt"

func main() {
	x := 10
	closure := func() {
		fmt.Println("x =", x)
	}
	closure() // Output: x = 10
}
