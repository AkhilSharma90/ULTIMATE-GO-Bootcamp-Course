package main
import "fmt"

func main() {
    num := 28
    switch {
    case num < 10:
        fmt.Println("Less than 10")
    case num >= 10 && num <= 20:
        fmt.Println("Between 10 and 20")
    default:
        fmt.Println("Greater than 20")
    }
}