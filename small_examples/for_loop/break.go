package main
import "fmt"

func main() {
    num := 1
    for  {
        if num == 5 {
            fmt.Println("Breaking the loop at 5")
            break
        }
        fmt.Println(num)
        num++
    }
    fmt.Println("Loop exited.")
}