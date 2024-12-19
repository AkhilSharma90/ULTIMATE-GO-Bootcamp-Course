package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initializing...")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initialize)
		}()
	}

	wg.Wait()
}

//sync.Once guarantees that the initialize function runs only once, even when called from multiple goroutines.