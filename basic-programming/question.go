package main

import "fmt"

func question() {
	calculatedData := 10
	updateString(&calculatedData)

	calculatedData = calculatedData + 10
	updateString(&calculatedData)

	fmt.Println(calculatedData)
}

// Input: "Memory Location to calculatedData Value"
func updateString(calculatedDataPointer *int) {
	fmt.Println(calculatedDataPointer)
	fmt.Println(*calculatedDataPointer)
}
