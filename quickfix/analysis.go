package quickfix

import (
	"honnef.co/go/tools/analysis/lint"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzers = lint.InitializeAnalyzers(Docs, map[string]*analysis.Analyzer{
	"QF1000": {
		Run:      CheckStringsIndexByte,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	"QF1001": {
		Run:      CheckDeMorgan,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
})
