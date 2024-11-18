package recvcheck_test

import (
	"testing"

	"github.com/raeperd/recvcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		desc     string
		settings recvcheck.Settings
	}{
		{
			desc:     "basic",
			settings: recvcheck.Settings{},
		},
		{
			desc:     "excluded",
			settings: recvcheck.Settings{},
		},
		{
			desc:     "disablebuiltin",
			settings: recvcheck.Settings{DisableBuiltin: true},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			a := recvcheck.NewAnalyzer(test.settings)

			analysistest.Run(t, analysistest.TestData(), a, test.desc)
		})
	}
}
