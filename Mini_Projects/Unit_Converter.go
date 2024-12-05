package main

import "fmt"

func main() {
	fmt.Println("Unit Converter")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Println("3. Kilometers to Miles")
	fmt.Println("4. Miles to Kilometers")
	fmt.Println("5. Kilograms to Pounds")
	fmt.Println("6. Pounds to Kilograms")
	fmt.Print("Choose an option: ")

	var choice int
	fmt.Scan(&choice)

	var value float64
	fmt.Print("Enter value to convert: ")
	fmt.Scan(&value)

	switch choice {
	case 1:
		fmt.Printf("Result: %.2f°F\n", value*9/5+32)
	case 2:
		fmt.Printf("Result: %.2f°C\n", (value-32)*5/9)
	case 3:
		fmt.Printf("Result: %.2f miles\n", value*0.621371)
	case 4:
		fmt.Printf("Result: %.2f km\n", value/0.621371)
	case 5:
		fmt.Printf("Result: %.2f lbs\n", value*2.20462)
	case 6:
		fmt.Printf("Result: %.2f kg\n", value/2.20462)
	default:
		fmt.Println("Invalid option.")
	}
}
