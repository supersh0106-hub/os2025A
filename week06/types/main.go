package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var name string
	// name = "Kim seong hyun"

	// var name = "Kim seong hyun"

	// name := "Kim seong hyun"

	// name := 2.71

	var name float64
	name = 2.71
	fmt.Println(name, reflect.TypeOf(name))
}
