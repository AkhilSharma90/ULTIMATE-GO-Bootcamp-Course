package main

import "fmt"

func main() {
	x := 10
	defer fmt.Println("Deferred x:", x)
	x = 20
	fmt.Println("Changed x:", x)
}
