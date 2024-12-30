package main

import "fmt"

// Conversion functions
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}
func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}
func main() {
	c := 25.0
	fmt.Printf("%.2f°C = %.2f°F\n", c, celsiusToFahrenheit(c))
	fmt.Printf("%.2f°C = %.2fK\n", c, celsiusToKelvin(c))
	f := 77.0
	fmt.Printf("%.2f°F = %.2f°C\n", f, fahrenheitToCelsius(f))
	k := 298.15
	fmt.Printf("%.2fK = %.2f°C\n", k, kelvinToCelsius(k))
}
