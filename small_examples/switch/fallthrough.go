package main
import "fmt"

func main() {
    grade := "A"
    switch grade {
    case "A":
        fmt.Println("Excellent!")
        fallthrough
    case "B":
        fmt.Println("Well done!")
        fallthrough
    case "C":
        fmt.Println("Good effort!")
    default:
        fmt.Println("Keep trying!")
    }
}
