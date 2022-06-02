package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func workingWithStringConversion() {
	userAge := "345345"
	userSalary := 6345624

	//Integer Value to String
	salaryString := strconv.Itoa(userSalary)
	fmt.Println(salaryString)

	// String Value to Integer
	convertedInteger, err := strconv.Atoi(userAge)

	if err == nil {
		fmt.Println("Conversion Succesfully Done")
		fmt.Println(convertedInteger)
		fmt.Println(reflect.TypeOf(convertedInteger))
	} else {
		fmt.Println("Conversion Failed")
	}
}
