package test

import (
	"time"
)

type RPC struct { // want `the methods of "RPC" use pointer receiver and non pointer receiver.`
	result int
	done   chan struct{}
}

func (rpc *RPC) compute() {
	time.Sleep(time.Second) // strenuous computation intensifies
	rpc.result = 42
	close(rpc.done)
}

func (RPC) version() int {
	return 1 // never going to need to change this
}
