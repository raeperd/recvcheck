package test

import (
	"time"
)

type RPC struct {
	result int
	done   chan struct{}
}

func (rpc *RPC) compute() {
	time.Sleep(time.Second) // strenuous computation intensifies
	rpc.result = 42
	close(rpc.done)
}

func (RPC) version() int { // want "receiver type should be a pointer"
	return 1 // never going to need to change this
}
