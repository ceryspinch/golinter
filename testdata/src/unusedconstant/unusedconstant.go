package unusedconstant

import "fmt"

const goodbyeMessage = "Bye!" // want "Constant \"goodbyeMessage\" has been declared but is not used, which means that it is redundant. Delete the constant declaration if it is not needed or use it within a function."

const helloMessage = "Hi!"

func helloWorld() {
	fmt.Print(helloMessage)
}