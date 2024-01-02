package packagenaming

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "packagenaming",
	Doc:      "Checks for deviations from Go's naming conventions in package names",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		file := node.(*ast.File)
		packageName := file.Name.String()

		isValid, reason := isValidPackageName(packageName)
		if !isValid {
			pass.Reportf(
				file.Package,
				"Package name %q does not follow Go's naming conventions as it contains an %s. Package names should be short and only contain lowercase letters.",
				packageName, reason,
			)
		}

	})

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
