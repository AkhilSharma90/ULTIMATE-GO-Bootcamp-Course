package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	future := now.Add(2 * time.Hour)

	fmt.Println("Now:", now)
	fmt.Println("2 Hours Later:", future)
}

//Now: 2024-11-19 14:32:45.123456 -0700 MST
//2 Hours Later: 2024-11-19 16:32:45.123456 -0700 MST
