package main

import (
	"log"

	"github.com/ceryspinch/golinter/common"
	"github.com/ceryspinch/golinter/rules/commentlength"
	"github.com/ceryspinch/golinter/rules/complexconditional"
	"github.com/ceryspinch/golinter/rules/constantnaming"
	"github.com/ceryspinch/golinter/rules/functionlength"
	"github.com/ceryspinch/golinter/rules/functionnaming"
	"github.com/ceryspinch/golinter/rules/magicnumbers"
	"github.com/ceryspinch/golinter/rules/packagenaming"
	"github.com/ceryspinch/golinter/rules/parameterlist"
	"github.com/ceryspinch/golinter/rules/repeatedstrings"
	"github.com/ceryspinch/golinter/rules/unusedconstant"
	"github.com/ceryspinch/golinter/rules/unusedfunction"
	"github.com/ceryspinch/golinter/rules/unusedvariable"
	"github.com/ceryspinch/golinter/rules/variablenaming"
	"golang.org/x/tools/go/analysis/multichecker"
)

func RunLinter() {
	// Open the JSON file for writing
	file, err := common.OpenJSONFile("output.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	multichecker.Main(
		functionnaming.Analyzer,
		variablenaming.Analyzer,
		packagenaming.Analyzer,
		constantnaming.Analyzer,
		parameterlist.Analyzer,
		functionlength.Analyzer,
		repeatedstrings.Analyzer,
		magicnumbers.Analyzer,
		complexconditional.Analyzer,
		commentlength.Analyzer,
		unusedvariable.Analyzer,
		unusedconstant.Analyzer,
		unusedfunction.Analyzer,
	)

	err = common.CloseJSONFile("output.json")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	RunLinter()
}
