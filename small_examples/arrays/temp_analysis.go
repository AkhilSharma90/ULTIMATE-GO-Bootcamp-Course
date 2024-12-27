package main

import (
	"fmt"
)

func main() {
	// Daily temperatures recorded over a week (in Celsius)
	temperatures := [7]float64{23.5, 25.0, 22.8, 26.1, 24.3, 23.0, 27.5}
	// Calculate the average temperature
	var total float64
	for _, temp := range temperatures {
		total += temp
	}
	average := total / float64(len(temperatures))
	// Find the hottest day
	maxTemp := temperatures[0]
	day := 0
	for i, temp := range temperatures {
		if temp > maxTemp {
			maxTemp = temp
			day = i
		}
	}
	// Print the results
	fmt.Printf("Average Temperature: %.2f°C\n", average)
	fmt.Printf("Hottest Day: Day %d with %.2f°C\n", day+1, maxTemp)
}
