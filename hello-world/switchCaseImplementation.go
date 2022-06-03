package main

import (
	"fmt"
	"reflect"
)

func switchCaseImplementation() {
	userdata := false

	switch reflect.TypeOf(userdata).String() {
	case "int":
		fmt.Println("This is of Type Integer")
	case "float32":
		fmt.Println("This is of Type Float32")
	case "float64":
		fmt.Println("This is of Type Float64")
		fallthrough
	case "complex64":
		fmt.Println("This is of Type Complex64")
		fallthrough
	case "complex128":
		fmt.Println("This is of Type Complex128")
	default:
		fmt.Println("None of them Matched")
	}

}
