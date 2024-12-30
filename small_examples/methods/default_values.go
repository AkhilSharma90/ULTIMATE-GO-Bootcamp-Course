package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

// Method to provide a default score if not set
func (s Student) GetScore() int {
	if s.Score == 0 {
		return 50 // Default score
	}
	return s.Score
}
func main() {
	student1 := Student{Name: "Alice"}
	student2 := Student{Name: "Bob", Score: 80}
	fmt.Println(student1.Name, "Score:", student1.GetScore()) // Output: 50
	fmt.Println(student2.Name, "Score:", student2.GetScore()) // Output: 80
}
