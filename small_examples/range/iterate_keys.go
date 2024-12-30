package main

import "fmt"

func main() {
	capitals := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}
	for country := range capitals {
		fmt.Println("Country:", country)
	}
}
