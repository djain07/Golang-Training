package main

import (
	"fmt"
)

var userInfoDataRandom string = "saydfasiydfi"
var userAge int
var userDesignation string = "Developer"

func creatingVariable() {
	var variableOne int = 100000000000000000
	var variableTwo = 3.14
	var variableThree int64 = 10
	var userName string = "Mayank Gupta"
	var userAge int = 100
	fmt.Println(userAge)
	fmt.Println(userName)
	fmt.Println(float64(variableOne) + variableTwo)
	fmt.Println(variableOne + int(variableThree))

	otherFunction()
}

func otherFunction() {
	fmt.Println(userAge)
}
