package main

import (
	"fmt"
	"sync"
)

var (
	data  int
	mutex sync.Mutex
)

func readData() int {
	mutex.Lock()
	defer mutex.Unlock()
	return data
}
func writeData(value int) {
	mutex.Lock()
	data = value
	mutex.Unlock()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		writeData(42)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Read Data:", readData())
	}()
	wg.Wait()
}
