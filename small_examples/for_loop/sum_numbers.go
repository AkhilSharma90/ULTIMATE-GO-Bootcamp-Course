package main
import "fmt"
func main() {
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum of numbers from 1 to 100 is:", sum)
}

