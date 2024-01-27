package unusedvariable

import (
	"go/ast"
	"go/token"

	"github.com/fatih/color"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "unusedvariable",
	Doc:      "Checks for variables that are declared but never used",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// Go ode will not compile if variables are declared but not used within functions,
	// so this code only checks for package level declarations in the form: 'var example string' or 'var example = "hello"'
	// that pass compilation but are not then used within any functions throughout the code.
	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		decl := node.(*ast.GenDecl)
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
				isUsed := checkUnusedVariable(pass, varName, spec.Pos())
				if !isUsed {
					pass.Reportf(
						ident.Pos(),
						(color.RedString("Variable %q has been declared but is not used, ", varName))+
							color.BlueString("which means that is is redundant. ")+
							color.GreenString("Delete the variable declaration if it is not needed or use it within a function."),
					)
				}
			}
		}
	})

	return nil, nil
}

func checkUnusedVariable(pass *analysis.Pass, varName string, declarationPos token.Pos) bool {
	for _, file := range pass.Files {
		found := false
		ast.Inspect(file, func(n ast.Node) bool {
			if ident, ok := n.(*ast.Ident); ok {
				// Check if the identifier has the same name as the variable we're looking for and is not the original declaration
				if ident.Name == varName && ident.Pos() != declarationPos {
					found = true
					return false // Stop inspecting if the variable is found
				}
			}
			return true
		})

		if found {
			// Stop searching in other files if the variable is found
			return true
		}
	}

	return false
}
