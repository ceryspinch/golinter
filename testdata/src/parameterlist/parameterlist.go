package parameterlist

func oneParameterFunction(x int)

func fiveParametersFunction(a, b, c, d, e int) {} // want "Function \"fiveParametersFunction\" has five or more parameters, which may suggest that the function is doing more than one thing or is too complex and could be difficult to read, maintain and test. Try to split the function up into smaller ones that do one thing each."