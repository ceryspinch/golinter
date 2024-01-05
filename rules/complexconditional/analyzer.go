package complexconditional

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "complexconditional",
	Doc:      "Checks for complex conditional if statements",
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
				"Complex if statement condition detected with %d boolean expressions. This can make the code difficult to read and maintain. Consider refactoring by moving these long conditional checks into separate functions to be called.", complexity)
		}
	})

	// Check for nested if statements
	inspector.Preorder([]ast.Node{(*ast.IfStmt)(nil)}, func(node ast.Node) {
		ifStmt := node.(*ast.IfStmt)
		nestedIfCount := countNestedIfs(ifStmt)

		if nestedIfCount > 1 {
			pass.Reportf(
				ifStmt.Pos(),
				"Multiple, %d, nested if statements detected. This can make the code difficult to read, maintain and test. Consider refactoring by checking for invalid conditions first, simplifying conditions or using a switch statement instead.", nestedIfCount)
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
