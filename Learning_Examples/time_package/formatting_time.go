package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Formatted Time:", currentTime.Format("2006-01-02 15:04:05"))
}

//Formatted Time: 2024-11-19 14:32:45
