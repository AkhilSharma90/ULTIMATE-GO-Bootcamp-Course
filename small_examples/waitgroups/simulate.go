package main

import (
	"fmt"
	"sync"
	"time"
)

func startService(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting %s...\n", name)
	time.Sleep(2 * time.Second)
	fmt.Printf("%s started\n", name)
}
func main() {
	var wg sync.WaitGroup
	services := []string{"Auth", "Database", "Cache"}
	for _, service := range services {
		wg.Add(1)
		go startService(service, &wg)
	}
	wg.Wait()
	fmt.Println("All services started")
}
