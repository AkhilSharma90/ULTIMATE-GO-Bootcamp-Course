package main

import "fmt"

// Function to calculate average
func calculateAverage(scores []float64) float64 {
	total := 0.0
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))
}

// Function to determine the grade
func assignGrade(avg float64) string {
	switch {
	case avg >= 90:
		return "A"
	case avg >= 80:
		return "B"
	case avg >= 70:
		return "C"
	case avg >= 60:
		return "D"
	default:
		return "F"
	}
}
func main() {
	scores := []float64{85, 92, 78, 90, 88}
	average := calculateAverage(scores)
	grade := assignGrade(average)
	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Average: %.2f\n", average)
	fmt.Printf("Grade: %s\n", grade)
}
