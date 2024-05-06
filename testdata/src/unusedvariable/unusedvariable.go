package unusedvariable

import "fmt"

var greetingMessage string // want "Variable \"greetingMessage\" has been declared but is not used, which means that it is redundant. Delete the variable declaration if it is not needed or use it within a function."

var goodbyeMessage = "Bye!" // want "Variable \"goodbyeMessage\" has been declared but is not used, which means that it is redundant. Delete the variable declaration if it is not needed or use it within a function."

var usedVariable = "Hi!"

func helloWorld() {
	fmt.Print(usedVariable)
}
