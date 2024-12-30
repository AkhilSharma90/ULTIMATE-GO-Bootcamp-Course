package main

type Shape struct{}

// Only one method named Area is allowed in the same type
func (s Shape) Area() int { return 0 }

// func (s Shape) Area(a int) int { return a } // This will cause a compilation error
