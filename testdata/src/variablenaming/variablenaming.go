package variablenaming

import "fmt"

const notAVar = "test"

var my_variable_name string // want "Variable \"my_variable_name\" in variable declaration does not follow Go's naming conventions as it contains an underscore. Instead use CamelCase, for example \"exampleVariableName\"."

var MYVARIABLENAME string // want "Variable \"MYVARIABLENAME\" in variable declaration does not follow Go's naming conventions as it contains only uppercase letters. Instead use CamelCase, for example \"exampleVariableName\"."

func variableAssignment() {
	my_variable_name := "test" // want "Variable \"my_variable_name\" in variable assignment does not follow Go's naming conventions as it contains an underscore. Instead use CamelCase, for example \"exampleVariableName\"."
	fmt.Print(my_variable_name)
}
