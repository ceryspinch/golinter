package packagenaming

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "packagenaming",
	Doc:  "Checks for deviations from Go's naming conventions in package names",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.File:
			packageName := node.Name.Name

			result, reason := isValidPackageName(packageName)
			if !result {
				pass.Reportf(
					node.Package,
					"Package name %q does not follow Go's naming conventions as it contains an %s. Package names should be short and only contain lowercase letters.",
					packageName, reason,
				)
			}
		}

		return true
	}

	for _, fileAST := range pass.Files {
		ast.Inspect(fileAST, inspect)
	}

	return nil, nil
}

func isValidPackageName(packageName string) (bool, string) {
	if strings.Contains(packageName, "_") {
		return false, "underscore"
	}

	for _, char := range packageName {
		if unicode.IsUpper(char) {
			return false, "uppercase letter"
		}
	}
	
	return true, ""
}
