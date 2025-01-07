package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Println("Deferred Execution")
	fmt.Println("End")
}
