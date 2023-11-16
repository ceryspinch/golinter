package functionlength

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "longfunction",
	Doc:  "Checks for unusually long functions that may suggest it is doing too many things",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		// Count the number of lines in the function.
		startLine := pass.Fset.Position(funcDecl.Pos()).Line
		endLine := pass.Fset.Position(funcDecl.End()).Line
		numLines := endLine - startLine + 1

		// TODO: Decide on number to use here that is research/evidence backed
		if numLines >= 20 {
			pass.Reportf(
				node.Pos(),
				"Function %q is %d lines long, which may suggest that the function is doing more than one thing or is too complex. Consider refactoring it to improve readability and maintainability.",
				funcDecl.Name.String(),
				numLines,
			)
		}
		return true
	}

	for _, fileAST := range pass.Files {
		ast.Inspect(fileAST, inspect)
	}

	return nil, nil
}
