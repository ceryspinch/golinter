package namingconventions

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "namingconvention",
	Doc:  "Checks for functions, variables, constants and file names that do not follow Go's naming conventions",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		// Function name checking
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		funcName := funcDecl.Name.String()

		// Check if function name contains underscores
		if strings.Contains(funcName, "_") {
			pass.Reportf(
				node.Pos(),
				"Function %s does not follow Go naming conventions as it contains an underscore. Instead use camelCase, for example %q",
				funcName, "exampleFunctionNameWithGoodNaming")
		}

		// Check if function name uses camelCase

		// TODO: Flesh out logic here

		// TODO: Constant name checking

		// TODO: Variable name checking

		// TODO: File name checking

		return true

	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}

	return nil, nil
}
