package main

import (
	"github.com/ceryspinch/go-linter/pkg/constantnaming"
	"github.com/ceryspinch/go-linter/pkg/functionlength"
	"github.com/ceryspinch/go-linter/pkg/functionnaming"
	"github.com/ceryspinch/go-linter/pkg/magicnumbers"
	"github.com/ceryspinch/go-linter/pkg/packagenaming"
	"github.com/ceryspinch/go-linter/pkg/parameterlist"
	"github.com/ceryspinch/go-linter/pkg/repeatedstrings"
	"github.com/ceryspinch/go-linter/pkg/variablenaming"
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
