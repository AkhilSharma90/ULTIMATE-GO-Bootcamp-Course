package main
import "fmt"
func main() {
    i := 1 // Declare outside the loop
    for ; i <= 5; i++ {
        fmt.Println(i)
    }
}