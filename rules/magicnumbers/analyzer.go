package magicnumbers

import (
	"go/ast"
	"go/token"

	"github.com/fatih/color"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "magicnumbers",
	Doc:      "Checks for integer literals that could be defined as constants instead",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BasicLit)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		expr := node.(*ast.BasicLit)
		if expr.Kind != token.INT {
			return
		}

		value := expr.Value
		pass.Reportf(
			node.Pos(),
			(color.RedString("Possible magic number detected: %s. ", value))+
				color.BlueString("This can make the code difficult to understand and could lead to errors during maintenance. ")+
				color.GreenString("Consider defining it as a constant instead."),
		)
	})

	return nil, nil
}
