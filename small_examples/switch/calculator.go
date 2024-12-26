package main
import "fmt"
func main() {
    a, b := 10, 5
    operation := "*"
    switch operation {
    case "+":
        fmt.Println("Addition:", a+b)
    case "-":
        fmt.Println("Subtraction:", a-b)
    case "*":
        fmt.Println("Multiplication:", a*b)
    case "/":
        if b != 0 {
            fmt.Println("Division:", a/b)
        } else {
            fmt.Println("Cannot divide by zero")
        }
    default:
        fmt.Println("Unknown operation")
    }
}

