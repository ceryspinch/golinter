package repeatedstrings

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "repeatedstrings",
	Doc:      "Checks for repeated strings that could be constants",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// Map to store string literals and their occurrences
	stringLiterals := make(map[string]int)

	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BasicLit)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		basicLit := node.(*ast.BasicLit)
		if basicLit.Kind != token.STRING {
			return
		}

		// Get string value and add to the map as a key
		stringValue := basicLit.Value
		stringLiterals[stringValue]++
	})

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
