package parameterlist

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "longparameterlist",
	Doc:  "Checks for unusually long parameter lists that may suggest a function is doing too many things",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		paramList := funcDecl.Type.Params
		if paramList.NumFields() >= 5 {
			pass.Reportf(
				node.Pos(),
				"Function %q has 5 or more parameters, which may suggest that the function is doing more than one thing. Try to split the function up into smaller ones that do one thing each, this makes the functions more readable, maintainable and testable.",
				funcDecl.Name.String(),
			)
		}

		return true
	}

	for _, fileAST := range pass.Files {
		ast.Inspect(fileAST, inspect)
	}

	return nil, nil
}
