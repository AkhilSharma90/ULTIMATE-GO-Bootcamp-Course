package main

import (
    "fmt"
    "strconv"
	"reflect"
)

func main() {
    boolTrue := true
    boolFalse := false
    strTrue := strconv.FormatBool(boolTrue)
    strFalse := strconv.FormatBool(boolFalse)
    fmt.Println("Boolean true as string:", strTrue)   // "true"
    fmt.Println("Boolean false as string:", strFalse) // "false"

	fmt.Println(reflect.TypeOf(strTrue))
	fmt.Println(reflect.TypeOf(strFalse))
	
}