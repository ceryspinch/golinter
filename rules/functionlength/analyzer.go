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
		funcName := funcDecl.Name.String()
		funcPosition := funcDecl.Pos()

		// Count the number of lines in the function.
		startLine := pass.Fset.Position(funcDecl.Pos()).Line
		endLine := pass.Fset.Position(funcDecl.End()).Line
		numLines := endLine - startLine + 1

		if numLines >= 10 {
			pass.Reportf(
				funcPosition,
				(color.RedString("Function %q is %d lines long, ", funcName, numLines))+
					color.CyanString("which may suggest that the function is doing more than one thing or is too complex and could be difficult to read, maintain and test. ")+
					color.YellowString("Try to split the function up into smaller ones that do one thing each.\n"),
			)
		}
	})

	return nil, nil
}
