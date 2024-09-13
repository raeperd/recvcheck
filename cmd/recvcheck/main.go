package main

import (
	"github.com/raeperd/recvcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(recvcheck.Analyzer)
}
