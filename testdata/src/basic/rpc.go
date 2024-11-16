package basic

import (
	"time"
)

type RPC struct { // want `the methods of "RPC" use pointer receiver and non-pointer receiver.`
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

// Following main function cause data race error
// reference: https://dave.cheney.net/2015/11/18/wednesday-pop-quiz-spot-the-race
// func main() {
// 	rpc := &RPC{done: make(chan struct{})}
//
// 	go rpc.compute()         // kick off computation in the background
// 	version := rpc.version() // grab some other information while we're waiting
// 	<-rpc.done               // wait for computation to finish
// 	result := rpc.result
//
// 	fmt.Printf("RPC computation complete, result: %d, version: %d\n", result, version)
// }
