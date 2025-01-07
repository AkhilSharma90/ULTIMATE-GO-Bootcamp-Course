//make(Type, length, capacity)

package main

import "fmt"

func main() {
	nums := make([]int, 5) // Slice of length 5
	fmt.Println(nums)      // Output: [0 0 0 0 0]
}
