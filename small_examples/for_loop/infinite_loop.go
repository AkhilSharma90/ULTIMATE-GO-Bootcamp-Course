package main
import "fmt"
func main() {
    counter := 1
    // Infinite loop
    for {
        fmt.Println("Counter:", counter)
        counter++
        if counter > 5 {
            break // Exit the loop
        }
    }
    fmt.Println("Exited the infinite loop!")
}