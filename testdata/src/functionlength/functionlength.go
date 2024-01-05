package functionlength

import "fmt"

func longFunction() { // want "Function \"longFunction\" is 14 lines long, which may suggest that the function is doing more than one thing or is too complex and could be difficult to read, maintain and test. Try to split the function up into smaller ones that do one thing each."
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)
	fmt.Println(11)
	fmt.Println(12)
}

func shortFunction() { 
	i := 0
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}
}