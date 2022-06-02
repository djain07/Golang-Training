package main

import (
	"fmt"
)

func rawStringLiteral() {
	userDescription := `
		This is the User named TechnoFunnel \\ terrtret
		Aged as 20
		This user is working righ now in Golang.
		
		Thanks for the session
	`
	fmt.Println(userDescription)
}
