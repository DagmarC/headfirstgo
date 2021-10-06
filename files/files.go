package main

import (
	"fmt"
	"github.com/headfirstgo/files/recursion"
)

// handlePanic Pattern: will be called from main everytime at the end (defer). If no panic has occurred then the p will be nil,
// so we can return from the function call.
// If the panic was raised via Error then we use type assertion and if yes we can recover from it.
// If some other panic has occurred e.g. panic("some string"), then "some string" is not error type, so we will re-raise
// the panic again.
func handlePanic() {
	// Recover needs to be called somewhere else not in the place where panic occurred, because everything after panic()
	// will be discarded. If some panic() was called,
	// recover() has the some value (it accepts interface{}), otherwise it is <nil>.
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func main() {
	defer handlePanic()
	recursion.ScanFiles("C:\\Users")
}
