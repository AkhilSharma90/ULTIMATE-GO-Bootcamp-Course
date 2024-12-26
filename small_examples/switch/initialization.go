package main
import "fmt"
func main() {
    
    switch num := 6; {
    case num%2 == 0:
        fmt.Println("Even")
    case num%2 != 0:
        fmt.Println("Odd")
    }
    
}
