package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var userSalary byte = 255
var userSalary1 int16 = 3000
var userSalary2 byte = 30
var userSalary3 byte = 30
var userSalary4 byte = 30
var userSalary5 byte = 30

func shortHandNotation() {
	variableOne := "Some String"
	variableTwo := 2
	variableThree := 3
	someFloat := 3.14
	userSalary := 20
	stringSalaryValue := strconv.Itoa(userSalary)
	userName := "Mayank Gupta"
	fmt.Println("User Name is: " + userName)
	fmt.Println("User Name is: " + stringSalaryValue)
	fmt.Println(reflect.TypeOf(userName))
	fmt.Printf("First Value: %v Second Value: %v ThirdValue: %v %v", variableOne, variableTwo, variableThree, someFloat)

}
