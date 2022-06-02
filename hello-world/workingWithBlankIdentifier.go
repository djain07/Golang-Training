package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func workingWithBlankIdentifier() {
	userAge := "345345"
	userSalary := 6345624

	var iRavi_5555K string = "Ravi"

	//Integer Value to String
	salaryString := strconv.Itoa(userSalary)
	fmt.Println(salaryString)

	// String Value to Integer
	convertedInteger, _ := strconv.Atoi(userAge)

	fmt.Println(convertedInteger)
	fmt.Println(reflect.TypeOf(convertedInteger))

	fmt.Println(iRavi_5555K)

	fmt.Println(userDesignation)
	fmt.Println(userInfoDataRandom)
	fmt.Println(userDesignation)
}
