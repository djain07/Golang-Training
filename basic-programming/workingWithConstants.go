package main

import (
	"fmt"
)

func workingWithConstants() {
	const data int = 10
	// fmt.Println(&data)
	// const name string = ""
	data2 := testfn(data)
	fmt.Println(data2)

}
func testfn(data1 int) int {
	data1 = 100
	return (data1 + 5)
}
