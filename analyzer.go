package recvcheck

import (
	"errors"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "recvcheck",
	Doc:      "checks for receiver type consistency",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	// _ = pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	return nil, errors.New("not implemented")
}
