package main
import "fmt"
func main() {
    day := "Friday"
    switch day {
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }
}
