package constantnaming

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "constantnaming",
	Doc:      "Checks for deviations from Go's naming conventions in constant names",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		genDecl := node.(*ast.GenDecl)

		// Skip if not a constant declaration
		if genDecl.Tok != token.CONST {
			return
		}

		for _, spec := range genDecl.Specs {
			valSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			for _, ident := range valSpec.Names {
				constName := ident.Name

				isValid, reason := isValidConstantName(constName)
				if !isValid {
					pass.Reportf(
						ident.Pos(),
						(color.RedString("Constant %q does not follow Go's naming conventions ", constName))+
							color.BlueString("as it contains %s. ", reason)+
							color.GreenString("Instead use CamelCase, for example: exampleConstantName."),
					)
				}
			}
		}
	})

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
