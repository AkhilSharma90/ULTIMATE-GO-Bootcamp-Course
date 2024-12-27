package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}
type Order struct {
	OrderID  string
	Products []Product
	Total    float64
}

func main() {
	order := Order{
		OrderID: "ORD123",
		Products: []Product{
			{Name: "Laptop", Price: 1200.50},
			{Name: "Mouse", Price: 25.99},
		},
		Total: 1226.49,
	}
	fmt.Println("Order ID:", order.OrderID)
	fmt.Println("Total:", order.Total)
	for _, product := range order.Products {
		fmt.Printf("Product: %s, Price: %.2f\n", product.Name, product.Price)
	}
}
