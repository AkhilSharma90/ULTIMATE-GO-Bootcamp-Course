package main

import "fmt"

func main() {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Critical error")
	fmt.Println("This will not be executed")
}
