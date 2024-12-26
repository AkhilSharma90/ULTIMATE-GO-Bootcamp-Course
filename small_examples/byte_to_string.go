package main
import "fmt"
func main() {
    // Byte slice to string
    byteSlice := []byte{'H', 'e', 'l', 'l', 'o'}
    str := string(byteSlice)
    fmt.Println("String:", str) // Output: Hello
    // String to byte slice
    s := "GoLang"
    bytes := []byte(s)
    fmt.Println("Byte Slice:", bytes) // Output: [71 111 76 97 110 103]
}