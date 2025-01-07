package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		defer fmt.Println("Deferred in loop:", i)
	}
}
