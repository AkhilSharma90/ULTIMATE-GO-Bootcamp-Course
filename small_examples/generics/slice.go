package main
import "fmt"
func FilterT any bool) []T {
    var result []T
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}
func main() {
    numbers := []int{1, 2, 3, 4, 5}
    evenNumbers := Filter(numbers, func(n int) bool {
        return n%2 == 0
    })
    fmt.Println(evenNumbers) // Output: [2 4]
}