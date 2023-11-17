package repeatedstrings

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "repeatedstrings",
	Doc:  "Checks for repeated strings that could be constants",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	// Map to store string literals and their occurrences
	stringLiterals := make(map[string]int)

	inspect := func(node ast.Node) bool {
		// Check for basic literals of type string
		if basicLit, ok := node.(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
			stringValue := basicLit.Value
			stringLiterals[stringValue]++
		}
		return true
	}

	for _, fileAST := range pass.Files {
		ast.Inspect(fileAST, inspect)
	}

	// Check for repeated string literals
	for str, count := range stringLiterals {
		if count > 1 {
			pass.Reportf(
				pass.Files[0].Pos(),
				"String literal %s is repeated %d times. Consider defining it as a constant instead so that if you need to update the value, you do not have to do it for every single instance.",
				str,
				count,
			)
		}
	}

	return nil, nil
}
