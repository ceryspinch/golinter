package parameterlist

import (
	"go/ast"

	"github.com/fatih/color"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "longparameterlist",
	Doc:      "Checks for long parameter lists that may suggest a function is doing too many things",
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

		paramList := funcDecl.Type.Params
		if paramList.NumFields() >= 5 {
			pass.Reportf(
				funcPosition,
				(color.RedString("Function %q has five or more parameters, ", funcName))+
					color.CyanString("which may suggest that the function is doing more than one thing or is too complex and could be difficult to read, maintain and test. ")+
					color.YellowString("Try to split the function up into smaller ones that do one thing each.\n"),
			)
		}

	})

	return nil, nil
}
