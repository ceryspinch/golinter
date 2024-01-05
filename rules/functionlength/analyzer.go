package functionlength

import (
	"go/ast"

	"github.com/fatih/color"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "longfunction",
	Doc:      "Checks for long functions that may suggest they are too complex",
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
				(color.RedString("Function %q is %d lines long, ", funcDecl.Name.String(), numLines))+
					color.BlueString("which may suggest that the function is doing more than one thing or is too complex and could be difficult to read, maintain and test. ")+
					color.GreenString("Try to split the function up into smaller ones that do one thing each."),
			)
		}
	})

	return nil, nil
}
