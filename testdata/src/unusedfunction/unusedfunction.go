package unusedfunction

import "fmt"


func helloWorld() { // want "Function \"helloWorld\" has been declared but is not called anywhere, which means that is is redundant. Delete the function if it is not needed or use call it from within another function."
	fmt.Println("Hello World")
	goodbyeWorld()
}

func goodbyeWorld() {
	fmt.Println("goodbyeWorld")
}

