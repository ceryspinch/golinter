package magicnumbers

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "magicnumbers",
	Doc:  "Checks for integer literals that could be defined as constants instead",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		// Check only expressions
		expr, ok := node.(*ast.BasicLit)
		if !ok || expr.Kind != token.INT {
			return true
		}
		value := expr.Value
		pass.Reportf(
			node.Pos(),
			"Possible magic number detected: %s. Consider defining it as a constant instead to make the code more readable and maintainable.",
			value,
		)

		return true
	}

	for _, fileAST := range pass.Files {
		ast.Inspect(fileAST, inspect)
	}

	return nil, nil
}
