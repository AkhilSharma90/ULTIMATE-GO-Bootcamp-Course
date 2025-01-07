package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fetchAPI(name string) {
	start := time.Now()
	// Simulating a random API response time
	responseTime := time.Duration(rand.Intn(2000)+500) * time.Millisecond
	time.Sleep(responseTime)
	fmt.Printf("API %s responded in %v\n", name, time.Since(start))
}
func main() {
	rand.Seed(time.Now().UnixNano()) // Seed for random delay
	apis := []string{"UserService", "OrderService", "PaymentService", "InventoryService"}
	fmt.Println("Starting API requests...")
	start := time.Now()
	for _, api := range apis {
		go fetchAPI(api) // Run each API call in a goroutine
	}
	// Wait for all goroutines to finish
	time.Sleep(3 * time.Second)
	fmt.Printf("All API requests completed in %v\n", time.Since(start))
}
