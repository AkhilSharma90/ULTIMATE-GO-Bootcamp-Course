package main

import "fmt"

func main() {
    outer:
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            if i == 2 && j == 2 {
                fmt.Println("Breaking outer loop")
                break outer
            }
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
}