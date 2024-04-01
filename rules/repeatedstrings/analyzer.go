package repeatedstrings

import (
	"go/ast"
	"go/token"

	"github.com/fatih/color"
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
	// Map to store string literals and their number of occurrences
	stringLiteralsCount := make(map[string]int)
	// Map to store string literals and the position of their first use
	stringLiteralsFirstUse := make(map[string]token.Pos)

	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BasicLit)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		literal := node.(*ast.BasicLit)

		// Skip if literal is not a string
		if literal.Kind != token.STRING {
			return
		}

		// Get string value and add to the map as a key
		stringValue := literal.Value
		stringLiteralsCount[stringValue]++

		// Store position of string's first use
		stringLiteralsFirstUse[stringValue] = literal.Pos()
	})

	for str, count := range stringLiteralsCount {
		if count > 1 {
			pass.Reportf(
				stringLiteralsFirstUse[str],
				(color.RedString("String literal %s is repeated %d times, ", str, count))+
					color.CyanString("which may cause problems during maintenance. ")+
					color.YellowString("Consider defining it as a constant instead so that if you need to update the value, you do not have to do it for every single instance.\n"),
			)
		}
	}

	return nil, nil
}
