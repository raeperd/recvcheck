package recvcheck_test

import (
	"testing"

	"github.com/raeperd/recvcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), recvcheck.NewAnalyzer(recvcheck.Setting{}), "test")
}
