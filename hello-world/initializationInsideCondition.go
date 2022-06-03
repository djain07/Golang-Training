package main

import "fmt"

func initializationInsideCondition() {
	var userName = "Sample Data"
	if userName, userAgeRandom := returnValue(); userName == "Mayank" {
		fmt.Println("This user is not Allowed...")
		fmt.Println(userAgeRandom)
	} else {
		fmt.Println(("User Not Registered"))
	}

	fmt.Println("Sample Degugger")
	fmt.Println(userName)
}

func returnValue() (string, int) {
	userName := ""
	userAge := 0

	fmt.Println("Enter User Name")
	fmt.Scanln(&userName)

	fmt.Println("Enter User Age")
	fmt.Scanln(&userAge)

	return userName, userAge
}
