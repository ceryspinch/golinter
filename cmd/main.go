package main

import (
	"go-linter/pkg/namingconventions"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(namingconventions.Analyzer)
}
