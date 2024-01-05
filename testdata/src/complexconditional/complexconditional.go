package complexconditional

func checkX() bool {
	x := 0
	if x == 1 || x == 2 || x == 3 || x == 4 { // want "Complex if statement condition detected with 7 boolean expressions. This can make the code difficult to read and maintain. Consider refactoring by moving these long conditional checks into separate functions to be called."
		return true
	}
	return false
}

func checkY() bool {
	x := 0
	if x%2 == 0 { // want "Multiple, 3, nested if statements detected. This can make the code difficult to read, maintain and test. Consider refactoring by checking for invalid conditions first, simplifying conditions or using a switch statement instead."
		if x == 2 || x == 4 {
			if x == 2 {
				return true
			}
			return true
		}
		return false
	}
	return false
}

func validCheckX() bool {
	x := 0
	if x <= 4 {
		return true
	}
	return false
}
