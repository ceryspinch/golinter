package golinter

import (
	"github.com/ceryspinch/golinter/rules/constantnaming"
	"github.com/ceryspinch/golinter/rules/functionlength"
	"github.com/ceryspinch/golinter/rules/functionnaming"
	"github.com/ceryspinch/golinter/rules/magicnumbers"
	"github.com/ceryspinch/golinter/rules/packagenaming"
	"github.com/ceryspinch/golinter/rules/parameterlist"
	"github.com/ceryspinch/golinter/rules/repeatedstrings"
	"github.com/ceryspinch/golinter/rules/variablenaming"
	"golang.org/x/tools/go/analysis/multichecker"
)

func RunLinter() {
	multichecker.Main(
		functionnaming.Analyzer,
		variablenaming.Analyzer,
		packagenaming.Analyzer,
		constantnaming.Analyzer,
		parameterlist.Analyzer,
		functionlength.Analyzer,
		repeatedstrings.Analyzer,
		magicnumbers.Analyzer,
	)
}
