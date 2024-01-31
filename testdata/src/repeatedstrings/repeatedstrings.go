package repeatedstrings

import "fmt"

func helloWorld() {
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello") // want "String literal \"hello\" is repeated 3 times, which may cause problems during maintenance. Consider defining it as a constant instead so that if you need to update the value, you do not have to do it for every single instance."
	fmt.Println(3)
}
