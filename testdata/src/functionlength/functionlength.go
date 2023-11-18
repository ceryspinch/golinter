package functionlength

import "fmt"

func longFunction() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)
	fmt.Println(11)
	fmt.Println(12) // want "Function \"longFunction\" is 13 lines long, which may suggest that the function is doing more than one thing or is too complex. Consider refactoring it to improve readability and maintainability."
}
