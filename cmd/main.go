package main

import (
	"go-linter/pkg/constantnaming"
	"go-linter/pkg/functionlength"
	"go-linter/pkg/functionnaming"
	"go-linter/pkg/magicnumbers"
	"go-linter/pkg/packagenaming"
	"go-linter/pkg/parameterlist"
	"go-linter/pkg/repeatedstrings"
	"go-linter/pkg/variablenaming"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
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
