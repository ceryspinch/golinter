package commentlength

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"

	"github.com/fatih/color"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	minCommentLength = 2
	maxCommentLength = 30
)

var Analyzer = &analysis.Analyzer{
	Name:     "commentlength",
	Doc:      "Checks for unnecessarily short or long comments",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CommentGroup)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		commentGroup := node.(*ast.CommentGroup)
		for _, comment := range commentGroup.List {
			words := strings.Split(comment.Text, " ")

			if words[0] == "want" {
				continue
			}
			
			if len(words) > maxCommentLength {
				position := pass.Fset.Position(comment.Pos())
				pass.Reportf(
					comment.Pos(),
					(color.RedString("Comment on line %d contains %d words, ", position.Line, len(words)))+
						color.BlueString("which exceeds the maximum suggested comment length of %d. This could mean that the code is too complex. ", maxCommentLength)+
						color.GreenString("Try to simplify the code so that such a long comment is not needed to understand the code."),
				)
			}

			if len(words) < minCommentLength {
				position := pass.Fset.Position(comment.Pos())
				pass.Reportf(
					comment.Pos(),
					(color.RedString("Comment on line %d contains %d words, ", position.Line, len(words)))+
						color.BlueString("which is shorter than the minimum suggested comment length of %d. This could mean that the comment is unnecessary and does not add any value.", minCommentLength)+
						color.GreenString("Revaluate whether the comment is needed by checking if the code explains itself without it."),
				)
			}
		}
	})

	return nil, nil
}
