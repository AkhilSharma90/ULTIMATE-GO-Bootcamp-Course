package main
import "fmt"
func main() {
    row := 1
    for row <= 3 {
        col := 1
        for col <= 3 {
            fmt.Printf("%d  %d \n", row, col)
            col++
        }
        row++
    }
}
