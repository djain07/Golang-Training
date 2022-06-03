package main

import "fmt"

func conditionalOperator() {
	var userName string
	var userAge int
	fmt.Scanln(&userName)
	fmt.Scanln(&userAge)

	if userName == "Mayank" && userAge == 30 {
		fmt.Println("Welcome Guest")
	} else {
		fmt.Println("Welcome User")
	}
}
