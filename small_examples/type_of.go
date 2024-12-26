package main

import (
    "fmt"
    "reflect"
)
func main() {
    var x = 42
    // Get the type and value of the variable
    t := reflect.TypeOf(x)
    v := reflect.ValueOf(x)
    fmt.Println("Type:", t)            // Type: int
    fmt.Println("Value:", v)           // Value: 42
    fmt.Println("Kind:", t.Kind())     // Kind: int
    fmt.Println("Value as int:", v.Int()) // Value as int: 42
}