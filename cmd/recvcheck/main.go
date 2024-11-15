package main

import (
	"flag"
	"strings"

	"github.com/raeperd/recvcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	var (
		excludeMethod          string
		noBuiltinExcludeMethod bool
	)
	flag.StringVar(&excludeMethod, "exclude-method", "",
		"exclude the method signature from the check, seperated by '/'")
	flag.BoolVar(&noBuiltinExcludeMethod, "no-builtin-exclude-method", false,
		`disables the default exclude methods such as "MarshalText() ([]byte, error)"`)
	flag.Parse()

	setting := recvcheck.Setting{
		NoBuiltinExcludeMethod: noBuiltinExcludeMethod,
		ExcludeMethod:          strings.Split(excludeMethod, "/"),
	}
	singlechecker.Main(recvcheck.NewAnalyzer(setting))
}
