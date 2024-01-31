package unusedfunction

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/fatih/color"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/ceryspinch/golinter/common"
)

var Analyzer = &analysis.Analyzer{
	Name:     "unusedfunction",
	Doc:      "Checks for functions that are declared but never called",
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
		funcDeclPosition := node.Pos()
		fullFuncDeclPosition := pass.Fset.Position(funcDeclPosition)

		// Ignore if main function as this will not need to be called anywhere else
		if funcName != "main" {
			isUsed := checkUnusedFunction(pass, funcName, node.Pos())
			if !isUsed {
				pass.Reportf(
					funcDeclPosition,
					(color.RedString("Function %q has been declared but is not called anywhere, ", funcName))+
						color.BlueString("which means that is is redundant. ")+
						color.GreenString("Delete the function if it is not needed or use call it from within another function."),
				)

				result := common.LintResult{
					FilePath: fullFuncDeclPosition.Filename,
					Line:     fullFuncDeclPosition.Line,
					Message:  fmt.Sprintf("Function %q has been declared but is not called anywhere, which means that is is redundant. Delete the function if it is not needed or use call it from within another function.", funcName),
				}

				common.AppendResultToJSON(result, "output.json")
			}
		}
	})

	return nil, nil
}

func checkUnusedFunction(pass *analysis.Pass, funcName string, declarationPos token.Pos) bool {
	for _, file := range pass.Files {
		found := false
		ast.Inspect(file, func(n ast.Node) bool {
			if callExpr, ok := n.(*ast.CallExpr); ok {
				if ident, ok := callExpr.Fun.(*ast.Ident); ok {
					// Check if the function is called elsewhere and is not the original declaration
					if ident.Name == funcName && ident.Pos() != declarationPos {
						found = true
						return false // Stop inspecting if the function is found
					}
				}
			}
			return true
		})

		if found {
			// Stop searching in other files if the function is found
			return true
		}
	}

	return false
}
