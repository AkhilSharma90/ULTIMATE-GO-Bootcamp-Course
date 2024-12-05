package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 5; i++ {
		fmt.Println("Tick:", <-ticker.C)
	}
}

/*
Tick: 2024-11-19 14:32:45.123456 -0700 MST
Tick: 2024-11-19 14:32:46.123456 -0700 MST
Tick: 2024-11-19 14:32:47.123456 -0700 MST
Tick: 2024-11-19 14:32:48.123456 -0700 MST
Tick: 2024-11-19 14:32:49.123456 -0700 MST
*/
