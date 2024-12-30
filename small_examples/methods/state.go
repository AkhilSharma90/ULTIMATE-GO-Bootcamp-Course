package main

import "fmt"

type LightSwitch struct {
	IsOn bool
}

// Method to check if the switch is on
func (ls LightSwitch) IsSwitchedOn() bool {
	return ls.IsOn
}
func main() {
	switch1 := LightSwitch{IsOn: true}
	switch2 := LightSwitch{IsOn: false}
	fmt.Println("Switch1 is on:", switch1.IsSwitchedOn()) // Output: true
	fmt.Println("Switch2 is on:", switch2.IsSwitchedOn()) // Output: false
}
