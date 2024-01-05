package parameterlist

import (
	"go/ast"

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

		paramList := funcDecl.Type.Params
		if paramList.NumFields() >= 5 {
			pass.Reportf(
				node.Pos(),
				"Function %q has five or more parameters, which may suggest that the function is doing more than one thing and could be difficult to read, maintain and test. Try to split the function up into smaller ones that do one thing each.",
				funcDecl.Name.String(),
			)
		}

	})

	return nil, nil
}
