package complexconditional

func checkX() bool {
	x := 0
	if x == 1 || x == 2 || x == 3 || x == 4 { // want "Complex if statement condition detected. Consider refactoring for better readability."
		return true
	}
	return false
}

func checkY() bool {
	x := 0
	if x%2 == 0 { // want "Multiple nested if statements detected. Consider refactoring to improve readability."
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
