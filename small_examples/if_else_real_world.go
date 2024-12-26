package main

import "fmt"

func main() {
    var temperature int
    fmt.Print("Enter the temperature: ")
    fmt.Scan(&temperature)
    if temperature < 0 {
        fmt.Println("It's freezing!")
    } else if temperature < 20 {
        fmt.Println("It's cold")
    } else if temperature < 30 {
        fmt.Println("It's warm")
    } else {
        fmt.Println("It's hot!")
    }
}