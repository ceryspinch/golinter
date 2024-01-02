package functionlength_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ceryspinch/golinter/rules/functionlength"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, functionlength.Analyzer, "functionlength")
}
