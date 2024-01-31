package unusedconstant

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
	Name:     "unusedconstant",
	Doc:      "Checks for constants that are declared but never used",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		decl := node.(*ast.GenDecl)
		if decl.Tok != token.CONST {
			return
		}

		for _, spec := range decl.Specs {
			valSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			for _, ident := range valSpec.Names {
				constName := ident.Name
				constPosition := spec.Pos()
				fullConstPosition := pass.Fset.Position(constPosition)

				isUsed := checkUnusedConstant(pass, constName, constPosition)

				if !isUsed {
					pass.Reportf(
						constPosition,
						(color.RedString("Constant %q has been declared but is not used, ", constName))+
							color.BlueString("which means that is is redundant. ")+
							color.GreenString("Delete the constant declaration if it is not needed or use it within a function."),
					)

					result := common.LintResult{
						FilePath: fullConstPosition.Filename,
						Line:     fullConstPosition.Line,
						Message:  fmt.Sprintf("Constant %q has been declared but is not used, which means that is is redundant. Delete the constant declaration if it is not needed or use it within a function.", constName),
					}

					common.AppendResultToJSON(result, "output.json")
				}
			}
		}
	})

	return nil, nil
}

func checkUnusedConstant(pass *analysis.Pass, constName string, declarationPos token.Pos) bool {
	for _, file := range pass.Files {
		found := false
		ast.Inspect(file, func(n ast.Node) bool {
			if ident, ok := n.(*ast.Ident); ok {
				// Check if the identifier has the same name as the constant we're looking for and is not the original declaration
				if ident.Name == constName && ident.Pos() != declarationPos {
					found = true
					return false // Stop inspecting if the constant is found
				}
			}
			return true
		})

		if found {
			// Stop searching in other files if the constant is found
			return true
		}
	}
	return false
}
