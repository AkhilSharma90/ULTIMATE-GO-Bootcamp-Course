package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}
}