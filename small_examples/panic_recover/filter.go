package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			switch r {
			case "specific error":
				fmt.Println("Recovered from specific error")
			default:
				fmt.Println("Recovered from general panic:", r)
			}
		}
	}()
	panic("specific error")
}
