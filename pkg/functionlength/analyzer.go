package functionlength

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "longfunction",
	Doc:      "Checks for unusually long functions that may suggest it is too complex",
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

		// Count the number of lines in the function.
		startLine := pass.Fset.Position(funcDecl.Pos()).Line
		endLine := pass.Fset.Position(funcDecl.End()).Line
		numLines := endLine - startLine + 1

		// TODO: Decide on number to use here that is research/evidence backed
		if numLines >= 10 {
			pass.Reportf(
				node.Pos(),
				"Function %q is %d lines long, which may suggest that the function is doing more than one thing or is too complex. Consider refactoring it to improve readability and maintainability.",
				funcDecl.Name.String(),
				numLines,
			)
		}
	})

	return nil, nil
}
