package variablenaming

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	reportMsg = "Variable %q in variable declaration does not follow Go's naming conventions as it contains %s. Instead use CamelCase, for example %q."
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
		(*ast.FuncDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		switch decl := node.(type) {
		// Check for variable declarations (e.g. var exampleVariable string)
		case *ast.GenDecl:
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

					isValid, reason := isValidVariableName(varName)
					if !isValid {
						pass.Reportf(
							ident.Pos(),
							reportMsg,
							varName, reason, "exampleVariableName")
					}
				}
			}

		case *ast.AssignStmt:
			// Check for variable assignments (e.g. exampleVariable := "hello" or exampleVariable = "hello")
			for _, varDecl := range decl.Lhs {
				if ident, ok := varDecl.(*ast.Ident); ok {
					varName := ident.Name

					isValid, reason := isValidVariableName(varName)
					if !isValid {
						pass.Reportf(
							ident.Pos(),
							reportMsg,
							varName, reason, "exampleVariableName")
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
