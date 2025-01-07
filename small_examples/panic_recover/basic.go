package main

import "fmt"

func main() {
	fmt.Println("Starting the program")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Before panic")
	panic("Something went wrong")
	fmt.Println("This will not be printed")
}
