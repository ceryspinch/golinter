package main

import (
	"go-linter/pkg/functionnaming"
	"go-linter/pkg/packagenaming"
	"go-linter/pkg/variablenaming"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		functionnaming.Analyzer,
		variablenaming.Analyzer,
		packagenaming.Analyzer,
	)
}
