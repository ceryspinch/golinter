package main

import (
	"fmt"
)

const (
	my_constant = "hello"
)

var (
	bad_var_name = "hello"
)

func bad_function_name() {
	fmt.Println("Hello World")
}

func badVariableNaming() {
	const BADCONST = "hello"
	var var_decl_bad_naming string
	var_decl_bad_naming = "hello"
	fmt.Println(var_decl_bad_naming)

	var_assignment_bad_naming := "test"
	fmt.Print(var_assignment_bad_naming)
}

func tooManyParams(a, b, c, d, e string) {
	fmt.Print(a, b, c, d, e)
}
