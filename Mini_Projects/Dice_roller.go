package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice(sides, rolls int) []int {
	rand.Seed(time.Now().UnixNano())
	results := make([]int, rolls)
	for i := 0; i < rolls; i++ {
		results[i] = rand.Intn(sides) + 1
	}
	return results
}

func main() {
	var sides, rolls int
	fmt.Print("Enter the number of sides on the dice: ")
	fmt.Scan(&sides)
	fmt.Print("Enter the number of dice to roll: ")
	fmt.Scan(&rolls)

	results := rollDice(sides, rolls)
	fmt.Println("Rolls:", results)
	sum := 0
	for _, roll := range results {
		sum += roll
	}
	fmt.Println("Total:", sum)
}
