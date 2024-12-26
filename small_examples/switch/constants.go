package main

import "fmt"

const (
    Red    = "red"
    Blue   = "blue"
    Green  = "green"
    Yellow = "yellow"
)

func main() {
    color := Green

    switch color {
    case Red:
        fmt.Println("Stop")
    case Green:
        fmt.Println("Go")
    case Yellow:
        fmt.Println("Slow down")
    default:
        fmt.Println("Unknown color")
    }
}