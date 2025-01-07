import "golang.org/x/exp/constraints"
// Generic function that works with ordered types
func MaxT constraints.Ordered T {
    if a > b {
        return a
    }
    return b
}
func main() {
    fmt.Println(Max(10, 20))          // Output: 20
    fmt.Println(Max("apple", "zoo")) // Output: zoo
}