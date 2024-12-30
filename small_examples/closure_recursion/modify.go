package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		counter++ // Modify the outer variable
	}
	increment()
	increment()
	fmt.Println("Counter:", counter) // Output: Counter: 2
}
