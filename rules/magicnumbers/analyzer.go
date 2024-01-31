package magicnumbers

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ceryspinch/golinter/common"
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
		literal := node.(*ast.BasicLit)
		// Skip any non-integer literals
		if literal.Kind != token.INT {
			return
		}

		value := literal.Value
		position := pass.Fset.Position(literal.Pos())

		pass.Reportf(
			literal.Pos(),
			(color.RedString("Possible magic number detected: %s. ", value))+
				color.BlueString("This can make the code difficult to understand and could lead to errors during maintenance. ")+
				color.GreenString("Consider defining it as a constant instead."),
		)

		result := common.LintResult{
			FilePath: position.Filename,
			Line:     position.Line,
			Message:  fmt.Sprintf("Possible magic number detected: %s. This can make the code difficult to understand and could lead to errors during maintenance. Consider defining it as a constant instead.", value),
		}

		common.AppendResultToJSON(result, "output.json")
	})

	return nil, nil
}
