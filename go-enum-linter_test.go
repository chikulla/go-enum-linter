package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestEnumRestrictionAnalyzer(t *testing.T) {
	testdata := analysistest.TestData() // points "testdata" dir

	analysistest.Run(t, testdata, EnumRestrictionAnalyzer, "valid", "invalid")
}
