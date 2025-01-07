package main
import "fmt"
func IdentityT any T {
    return value
}
func main() {
    fmt.Println(Identity(42))           // Output: 42
    fmt.Println(Identity("Hello, Go")) // Output: Hello, Go
}