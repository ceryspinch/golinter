package functionnaming

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "functionnaming",
	Doc:      "Checks for deviations from Go's naming conventions in function names",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		funcDecl := node.(*ast.FuncDecl)
		funcName := funcDecl.Name.String()

		isValid, reason := isValidFunctionName(funcName)
		if !isValid {
			pass.Reportf(
				node.Pos(),
				"Function %q does not follow Go's naming conventions as it contains %s. Instead use Camel Case, for example %q for private functions or %q for public functions.",
				funcName, reason, "examplePrivateFunctionName", "ExamplePrivateFunctionName")
		}
	})

	return nil, nil
}

func isValidFunctionName(funcName string) (bool, string) {
	if strings.Contains(funcName, "_") {
		return false, "an underscore"
	}

	if strings.ToUpper(funcName) == funcName {
		return false, "only uppercase letters"
	}

	return true, ""
}
