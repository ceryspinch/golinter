package complexconditional

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "complexconditional",
	Doc:      "Checks for complex conditional if or switch statements",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.IfStmt)(nil),
	}

	// Check for complex conditional statements in if conditions
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		ifStmt := node.(*ast.IfStmt)
		complexity := calculateComplexity(ifStmt.Cond)

		if complexity > 3 {
			pass.Reportf(
				ifStmt.Pos(),
				"Complex if statement condition detected. Consider refactoring for better readability.")
		}
	})

	// Check for nested if statements
	inspector.Preorder([]ast.Node{(*ast.IfStmt)(nil)}, func(node ast.Node) {
		ifStmt := node.(*ast.IfStmt)
		nestedIfCount := countNestedIfs(ifStmt)

		if nestedIfCount > 1 {
			pass.Reportf(
				ifStmt.Pos(),
				"Multiple nested if statements detected. Consider refactoring to improve readability.")
		}
	})

	return nil, nil
}

func calculateComplexity(expr ast.Expr) int {
	complexity := 0
	ast.Inspect(expr, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.BinaryExpr:
			complexity++
		}
		return true
	})
	return complexity
}

func countNestedIfs(ifStmt *ast.IfStmt) int {
	count := 0
	ast.Inspect(ifStmt.Body, func(node ast.Node) bool {
		if innerIfStmt, ok := node.(*ast.IfStmt); ok {
			count++
			count += countNestedIfs(innerIfStmt)
		}
		return true
	})
	return count
}
