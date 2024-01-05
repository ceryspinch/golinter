package repeatedstrings // want "String literal \"hello\" is repeated 3 times, which may cause problems during maintenance. Consider defining it as a constant instead so that if you need to update the value, you do not have to do it for every single instance."

import "fmt"

func helloWorld() {
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println(3)
}
