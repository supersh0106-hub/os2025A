package main

import (
	"fmt"
	"reflect"
)

func main() {
	// casting
	var length float64 = 3.2
	var width int = 2
	fmt.Println("면적은", int(length)*width)
	fmt.Println("길이 > 너비?", int(length) > width)
	fmt.Println("형변환", reflect.TypeOf(int(length)))
	fmt.Println("원본", reflect.TypeOf(length))
}
