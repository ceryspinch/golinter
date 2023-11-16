package functionnaming

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "functionnaming",
	Doc:  "Checks for deviations from Go's naming conventions in function names",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		funcName := funcDecl.Name.String()

		result, reason := isValidFunctionName(funcName)
		// Check if function name contains underscores
		if !result {
			pass.Reportf(
				node.Pos(),
				"Function %q does not follow Go's naming conventions as it contains %s. Instead use Camel Case, for example %q for private functions or %q for public functions.",
				funcName, reason, "examplePrivateFunctionName", "ExamplePrivateFunctionName")
		}

		return true
	}

	for _, fileAST := range pass.Files {
		ast.Inspect(fileAST, inspect)
	}

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
