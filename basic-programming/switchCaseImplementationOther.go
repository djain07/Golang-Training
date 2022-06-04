package main

import (
	"fmt"
)

// Director
// VP
// Manager
// Employee

func main() {
	userdata := "CEO"

	switch userdata {
	case "Director":
		fmt.Println("User was a Director")
		fallthrough
	case "VP":
		fmt.Println("User was a VP")
		fallthrough
	case "Manager":
	case "CEO":
		fmt.Println("User was a Manager")
		fallthrough
	case "Employee":
		fmt.Println("User was a Employee")
	}

}
