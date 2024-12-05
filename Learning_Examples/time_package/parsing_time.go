package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02"
	timeStr := "2024-11-19"

	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	fmt.Println("Parsed Time:", parsedTime)
}

//Parsed Time: 2024-11-19 00:00:00 +0000 UTC
