package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2025, 01, 1, 14, 0, 0, 0, time.UTC)
	end := time.Now()

	duration := end.Sub(start)

	fmt.Println("Duration:", duration)
	fmt.Println("Hours:", duration.Hours())
	fmt.Println("Minutes:", duration.Minutes())
}

/*
Duration: 2h32m45.123456s
Hours: 2.5469787394444444
Minutes: 152.81872436666668
*/
