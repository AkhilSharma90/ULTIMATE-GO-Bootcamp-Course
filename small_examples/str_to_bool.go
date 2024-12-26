package main

import (
    "fmt"
    "strconv"
)

func main() {
    strTrue := "true"
    strFalse := "false"
	
    boolValue, err := strconv.ParseBool(strTrue)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Converted to boolean:", boolValue) // true
    }
    boolValue, err = strconv.ParseBool(strFalse)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Converted to boolean:", boolValue) // false
    }
    // Invalid input
    invalidStr := "not_a_boolean"
    _, err = strconv.ParseBool(invalidStr)
    if err != nil {
        fmt.Println("Error:", err) // Error: strconv.ParseBool: parsing "not_a_boolean": invalid syntax
    }
}