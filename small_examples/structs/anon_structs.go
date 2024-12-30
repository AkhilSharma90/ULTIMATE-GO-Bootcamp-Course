package main

import "fmt"

func main() {
	car := struct {
		Brand string
		Year  int
	}{Brand: "Tesla", Year: 2022}

	fmt.Println(car)

}
