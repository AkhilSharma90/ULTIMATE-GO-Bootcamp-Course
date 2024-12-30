package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func main() {
	products := []Product{
		{"Laptop", 1200.99},
		{"Mouse", 25.50},
		{"Keyboard", 45.00},
	}
	var total float64
	for _, product := range products {
		total += product.Price
	}
	fmt.Printf("Total cost: $%.2f\n", total)
}
