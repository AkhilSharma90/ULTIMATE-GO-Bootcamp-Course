package main
import "fmt"

func main() {
    var num float64 = 3.141592653589793
    // Convert to float32
    num32 := float32(num)
    // Convert back to float64
    num64 := float64(num32)
    fmt.Println("Original float64:", num)   // 3.141592653589793
    fmt.Println("Converted to float32:", num32) // Precision may be lost
    fmt.Println("Converted back to float64:", num64) // Precision still limited by float32
}