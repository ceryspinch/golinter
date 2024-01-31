package variablenaming

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/ceryspinch/golinter/common"
	"github.com/fatih/color"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	exampleVarName = "exampleVariableName"
)

var Analyzer = &analysis.Analyzer{
	Name:     "variablenaming",
	Doc:      "Checks for deviations from Go's naming conventions in variable names",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
		(*ast.AssignStmt)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		switch decl := node.(type) {
		// Check for variable declarations (e.g. var exampleVariable string)
		case *ast.GenDecl:
			// Skip if not a variable declaration
			if decl.Tok != token.VAR {
				return
			}

			for _, spec := range decl.Specs {
				valSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}

				for _, ident := range valSpec.Names {
					varName := ident.Name
					varPosition := ident.Pos()
					fullVarPosition := pass.Fset.Position(varPosition)

					isValid, reason := isValidVariableName(varName)
					if !isValid {
						pass.Reportf(
							varPosition,
							(color.RedString("Variable %q in variable declaration does not follow Go's naming conventions ", varName))+
								color.BlueString("as it contains %s. ", reason)+
								color.GreenString("Instead use CamelCase, for example: %q.", exampleVarName),
						)

						result := common.LintResult{
							FilePath: fullVarPosition.Filename,
							Line:     fullVarPosition.Line,
							Message:  fmt.Sprintf("Variable %q in variable declaration not follow Go's naming conventions, as it contains %s. Instead use CamelCase, for example: %q.", varName, reason, exampleVarName),
						}

						common.AppendResultToJSON(result, "output.json")
					}
				}
			}

		case *ast.AssignStmt:
			// Check for variable assignments (e.g. exampleVariable := "hello" or exampleVariable = "hello")
			for _, varDecl := range decl.Lhs {
				if ident, ok := varDecl.(*ast.Ident); ok {
					varName := ident.Name
					varPosition := ident.Pos()
					fullVarPosition := pass.Fset.Position(varPosition)

					isValid, reason := isValidVariableName(varName)
					if !isValid {
						pass.Reportf(
							varPosition,
							(color.RedString("Variable %q in variable assignment does not follow Go's naming conventions ", varName))+
								color.BlueString("as it contains %s. ", reason)+
								color.GreenString("Instead use CamelCase, for example: %q.", exampleVarName),
						)

						result := common.LintResult{
							FilePath: fullVarPosition.Filename,
							Line:     fullVarPosition.Line,
							Message:  fmt.Sprintf("Variable %q in variable assignment not follow Go's naming conventions, as it contains %s. Instead use CamelCase, for example: %q.", varName, reason, exampleVarName),
						}

						common.AppendResultToJSON(result, "output.json")
					}
				}
			}
		}
	})

	return nil, nil
}

func isValidVariableName(varName string) (bool, string) {
	if strings.Contains(varName, "_") {
		return false, "an underscore"
	}

	if strings.ToUpper(varName) == varName {
		return false, "only uppercase letters"
	}

	return true, ""
}
