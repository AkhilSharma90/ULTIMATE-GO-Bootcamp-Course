package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	mapping := map[string]int{"Alice": 25, "Bob": 30}

	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Map: %v\n", mapping)

	// Pretty-print with %+v
	fmt.Printf("Map (detailed): %+v\n", mapping)
}

/*Sample Output
Slice: [1 2 3]
Map: map[Alice:25 Bob:30]
Map (detailed): map[Alice:25 Bob:30]
*/
