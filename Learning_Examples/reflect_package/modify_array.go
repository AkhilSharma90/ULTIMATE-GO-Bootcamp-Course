package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [3]int{1, 2, 3}
	v := reflect.ValueOf(&arr).Elem()

	for i := 0; i < v.Len(); i++ {
		v.Index(i).SetInt(int64(v.Index(i).Int() * 2))
	}

	fmt.Println("Modified Array:", arr)
}
