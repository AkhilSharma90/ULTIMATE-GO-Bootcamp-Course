package main

import "fmt"

func causePanic() {
	panic("Panic from causePanic")
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()
	fmt.Println("Calling causePanic")
	causePanic()
	fmt.Println("This will not be printed")
}
