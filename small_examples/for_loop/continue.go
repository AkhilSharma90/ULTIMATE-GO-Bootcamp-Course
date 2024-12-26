package main
import "fmt"
func main() {
    num := 1
    for num <= 10 {
        if num%3 == 0 {
            num++
            continue // Skip multiples of 3
        }
        fmt.Println(num)
        num++
    }
}