package main
import "fmt"
func main() {
    f := 3.14
    i := int(f) // Fractional part is truncated
    fmt.Println("Float:", f, "Integer:", i) // Output: Float: 3.14 Integer: 3
}