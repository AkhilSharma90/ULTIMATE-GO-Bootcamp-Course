package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a new random source
	fmt.Println("Random number between 0 and 99:", r.Intn(100))
	fmt.Println("Random number between 1 and 10:", r.Intn(10)+1)

	f := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("Random float between 0.0 and 1.0:", f.Float64())
	fmt.Println("Random float between 5.0 and 10.0:", 5.0+(f.Float64()*5.0))

	p := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := p.Perm(10) // Generate a random permutation of numbers 0 to 9
	fmt.Println("Random permutation:", numbers)
}
