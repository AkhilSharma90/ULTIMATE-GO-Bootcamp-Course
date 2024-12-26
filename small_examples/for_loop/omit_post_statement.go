package main
import "fmt"
func main() {
    for i := 1; i <= 5;{
        fmt.Println(i)
        i++ // Update inside the loop
    }
}