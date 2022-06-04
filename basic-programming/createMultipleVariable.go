package main

import (
	"fmt"
	"reflect"
)

var (
	userSalary0  byte  = 50
	userSalary9  byte  = 255
	userSalary19 int16 = 3000
	userSalary29 byte  = 23
	userSalary39 byte  = 23
	userSalary49       = 23.56
	userInfo           = "TechnoFunnel"
	isEmployed         = true
)

var (
	userSalary011, userSalary911, userAddress byte    = 50, 30, 8
	userSalary4911                            float32 = 23.56
	userInfo11                                string  = "TechnoFunnel"
	isEmployed11, isEmployable                bool    = true, false
	studentAge, studentSalary, studentAddress byte    = 50, 30, 8
)

func createMultipleVariable() {
	var userInfoData string = "Data"
	userAge := 66
	fmt.Println(reflect.TypeOf((userSalary49)))
	fmt.Println(userInfoData)
	fmt.Println(userAge)
}
