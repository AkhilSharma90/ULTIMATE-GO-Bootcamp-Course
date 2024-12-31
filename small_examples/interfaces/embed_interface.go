package main

import "fmt"

type Printer interface {
	Print()
}
type Scanner interface {
	Scan()
}
type MultiFunctionDevice interface {
	Printer
	Scanner
}
type OfficeMachine struct{}

func (o OfficeMachine) Print() {
	fmt.Println("Printing...")
}
func (o OfficeMachine) Scan() {
	fmt.Println("Scanning...")
}
func main() {
	var device MultiFunctionDevice = OfficeMachine{}
	device.Print()
	device.Scan()
}
