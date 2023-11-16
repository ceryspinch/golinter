package constantnaming

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "constantnaming",
	Doc:  "Checks for deviations from Go's naming conventions in constant names",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		decl, ok := node.(*ast.GenDecl)
		if !ok || decl.Tok != token.CONST {
			return true
		}

		for _, spec := range decl.Specs {
			valSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			for _, ident := range valSpec.Names {
				constName := ident.Name

				result, reason := isValidConstantName(constName)
				if !result {
					pass.Reportf(
						ident.Pos(),
						"Constant %q does not follow Go's naming conventions as it %s. Instead use CamelCase, for example %q.",
						constName, reason, "exampleConstantName")
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

func isValidConstantName(varName string) (bool, string) {
	if strings.Contains(varName, "_") {
		return false, "an underscore"
	}

	if strings.ToUpper(varName) == varName {
		return false, "only uppercase letters"
	}

	return true, ""
}
