package main

import (
	"fmt"
)

const (
	my_age = 22
)

var (
	greeting_message = "hello"
)

func hello_world() {
	fmt.Println("Hello World")
	fmt.Println("Hello World")
}

func badVariableNaming() {
	var var_decl_bad_naming string
	var_decl_bad_naming = "hello"
	fmt.Println(var_decl_bad_naming)

	var_assignment_bad_naming := "test"
	fmt.Print(var_assignment_bad_naming)
}

func tooManyParams(a, b, c, d, e string) {
	fmt.Print(a, b, c, d, e)
}

func countToTen() {
	fmt.Print(1)
	fmt.Print(2)
	fmt.Print(3)
	fmt.Print(4)
	fmt.Print(5)
	fmt.Print(6)
	fmt.Print(7)
	fmt.Print(8)
	fmt.Print(9)
	fmt.Print(10)
}

func complexConditional() {
	x := 5
	if x == 0 || x == 1 || x == 2 || x == 3 {
		fmt.Println("yes")
	}
	fmt.Println("no")
}
