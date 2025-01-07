package main

import (
	"fmt"
	"sync"
)

var (
	total int
	mutex sync.Mutex
)

func addToTotal(value int) {
	mutex.Lock()
	defer mutex.Unlock()
	total += value
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			addToTotal(val)
		}(i)
	}
	wg.Wait()
	fmt.Println("Total Sum:", total)
}
