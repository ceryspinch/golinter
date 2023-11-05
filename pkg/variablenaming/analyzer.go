package variablenaming

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "variablenaming",
	Doc:  "Checks for deviations from Go's naming conventions in variable names",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		switch decl := node.(type) {
		// Check for variable declarations (e.g. var exampleVariable string)
		case *ast.GenDecl:
			// If declaration is not a variable then skip checking logic
			if decl.Tok != token.VAR {
				return true
			}

			for _, spec := range decl.Specs {
				valSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				for _, ident := range valSpec.Names {
					varName := ident.Name

					result, reason := isValidVariableName(varName)
					if !result {
						pass.Reportf(
							ident.Pos(),
							"Variable %q in variable declaration does not follow Go's naming conventions as it %s. Instead use CamelCase, for example %q.",
							varName, reason, "exampleVariableName")
					}
				}
			}

		case *ast.AssignStmt:
			// Check for variable assignments (e.g. exampleVariable := "hello" or exampleVariable = "hello")
			for _, varDecl := range decl.Lhs {
				if ident, ok := varDecl.(*ast.Ident); ok {
					varName := ident.Name

					result, reason := isValidVariableName(varName)
					if !result {
						pass.Reportf(
							ident.Pos(),
							"Variable %q in variable assignment does not follow Go's naming conventions as it %s. Instead use CamelCase, for example %q.",
							varName, reason, "exampleVariableName")
					}
				}
			}
		}

		return true
	}

	for _, fileAST := range pass.Files {
		ast.Inspect(fileAST, inspect)
	}

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
