package main

import "fmt"

func logStep(step string) {
	fmt.Println("Executing step:", step)
}
func main() {
	defer logStep("1")
	defer logStep("2")
	defer logStep("3")
	fmt.Println("Main function running")
}
