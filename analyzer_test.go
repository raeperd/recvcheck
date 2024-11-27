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
			desc:     "builtinmethods",
			settings: recvcheck.Settings{},
		},
		{
			desc:     "disablebuiltin",
			settings: recvcheck.Settings{DisableBuiltin: true},
		},
		{
			desc:     "exclusions",
			settings: recvcheck.Settings{Exclusions: []string{"SQL.Value"}},
		},
		{
			desc:     "exclusionswildcard",
			settings: recvcheck.Settings{Exclusions: []string{"*.Value"}},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			a := recvcheck.NewAnalyzer(test.settings)

			analysistest.Run(t, analysistest.TestData(), a, test.desc)
		})
	}
}
