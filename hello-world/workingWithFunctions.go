package main

import (
	"fmt"
)

var empName string = "Mayank"
var empAge string = "20"

func workingWithFunctions() {
	empAge := "1000"
	fmt.Printf("Hello World from Main Function: %v", &empAge)
	displayData(&empAge)
	fmt.Println(empAge)
}

func displayData(empAgePointer *string) {
	*empAgePointer = "40"
	fmt.Printf("Hello World from Main Function: %v", *empAgePointer)
}
