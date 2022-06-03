package main

import "fmt"

func calculatingDataByValue() {
	calculatedData := 100

	calculatedData = addByValue(calculatedData)
	calculatedData = divideByValue(calculatedData)
	calculatedData = minusByValue(calculatedData)

	fmt.Println(calculatedData)
}

func addByValue(calculatedData int) int {
	return calculatedData + 100
}

func divideByValue(calculatedData int) int {
	return calculatedData / 10
}

func minusByValue(calculatedData int) int {
	return calculatedData - 5
}
