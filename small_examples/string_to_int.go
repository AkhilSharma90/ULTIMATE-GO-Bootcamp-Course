package main
import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str) // Convert string to integer
    if err == nil {
        fmt.Println("Converted Number:", num) // Output: Converted Number: 123
    }
    strFloat := "3.14"
    f, err := strconv.ParseFloat(strFloat, 64) // Convert string to float
    if err == nil {
        fmt.Println("Converted Float:", f) // Output: Converted Float: 3.14
    }
}