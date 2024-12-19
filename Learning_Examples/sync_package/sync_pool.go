package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new instance...")
			return "New Object"
		},
	}

	// Get an object from the pool
	obj := pool.Get()
	fmt.Println(obj)

	// Put an object back into the pool
	pool.Put("Reusable Object")

	// Get an object again
	obj2 := pool.Get()
	fmt.Println(obj2)
}

//sync.Pool is ideal for managing temporary objects to reduce memory allocations in performance-critical applications.