package main

import (
	"flag"

	"github.com/raeperd/recvcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	var (
		noBuiltinExcludeMethod bool
	)
	flag.BoolVar(&noBuiltinExcludeMethod, "no-builtin-exclude-method", false,
		`disables the default exclude methods such as "MarshalText"`)

	setting := recvcheck.Setting{
		NoBuiltinExcludeMethod: noBuiltinExcludeMethod,
	}
	singlechecker.Main(recvcheck.NewAnalyzer(setting))
}
