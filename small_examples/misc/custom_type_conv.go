package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
    var c Celsius = 100
    f := Fahrenheit(c*9/5 + 32) // Convert Celsius to Fahrenheit
    fmt.Println("Fahrenheit:", f)
}