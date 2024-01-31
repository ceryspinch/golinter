package parameterlist

import (
	"fmt"
	"go/ast"

	"github.com/ceryspinch/golinter/common"
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
		fullFuncPosition := pass.Fset.Position(funcPosition)

		paramList := funcDecl.Type.Params
		if paramList.NumFields() >= 5 {
			pass.Reportf(
				funcPosition,
				(color.RedString("Function %q has five or more parameters, ", funcName))+
					color.BlueString("which may suggest that the function is doing more than one thing or is too complex and could be difficult to read, maintain and test. ")+
					color.GreenString("Try to split the function up into smaller ones that do one thing each."),
			)

			result := common.LintResult{
				FilePath: fullFuncPosition.Filename,
				Line:     fullFuncPosition.Line,
				Message:  fmt.Sprintf("Function %q has five or more parameters, which may suggest that the function is doing more than one thing or is too complex and could be difficult to read, maintain and test. Try to split the function up into smaller ones that do one thing each.", funcName),
			}

			common.AppendResultToJSON(result, "output.json")
		}

	})

	return nil, nil
}
