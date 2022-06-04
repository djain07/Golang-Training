package main

import "fmt"

// Can we declare a constant without initializing a value ?
// Can I send Contants to Function ?
// Whether we can send constant to a function and update ?

func calculatingDataByReference() {
	calculatedData := 100
	const data byte = 10
	add(&calculatedData)
	add(&calculatedData)
	divide(&calculatedData)
	minus(&calculatedData)
	fmt.Println(data)
	fmt.Println(calculatedData)
}

func add(calculatedDataPointer *int) {
	*calculatedDataPointer = *calculatedDataPointer + 100
}

func divide(calculatedDataPointer *int) {
	*calculatedDataPointer = *calculatedDataPointer / 10
}

func minus(calculatedDataPointer *int) {
	*calculatedDataPointer = *calculatedDataPointer - 5
}
