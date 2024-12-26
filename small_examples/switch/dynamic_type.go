package main
import "fmt"
func printType(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Printf("Type: int, Value: %d\n", v)
    case string:
        fmt.Printf("Type: string, Value: %s\n", v)
    case bool:
        fmt.Printf("Type: bool, Value: %v\n", v)
    default:
        fmt.Printf("Unknown type\n")
    }
}

func main() {
    printType(42)
    printType("Hello")
    printType(true)
    printType(3.14)
}