package main

import "fmt"

type Engine struct{}

func (e Engine) Start() {
	fmt.Println("Engine started!")
}

type Car struct {
	Engine // Embedded struct
}

func main() {
	c := Car{}
	c.Start() // Access method of the embedded struct
}
