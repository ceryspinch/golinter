package constantnaming_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ceryspinch/golinter/rules/constantnaming"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, constantnaming.Analyzer, "constantnaming")
}
