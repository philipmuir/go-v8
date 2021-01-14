package v8
// #include "v8_c_bridge.h"
// #cgo CXXFLAGS: -fno-rtti -fpic -std=c++14 -DV8_COMPRESS_POINTERS -DV8_31BIT_SMIS_ON_64BIT_ARCH
// #cgo darwin linux CXXFLAGS: -I${SRCDIR} -I${SRCDIR}/include
// #cgo LDFLAGS: -pthread -lv8
// #cgo windows LDFLAGS: -lv8_libplatform
// #cgo darwin LDFLAGS: -L${SRCDIR}/libv8/darwin_x86_64
// #cgo linux LDFLAGS: -L${SRCDIR}/libv8/linux_x86_64
import "C"

import (
	"fmt"
	"sync"
)

// Version exposes the compiled-in version of the linked V8 library.  This can
// be used to test for specific javascript functionality support (e.g. ES6
// destructuring isn't supported before major version 5.).
var Version = struct{ Major, Minor, Build, Patch int }{
	Major: int(C.version.major),
	Minor: int(C.version.minor),
	Build: int(C.version.build),
	Patch: int(C.version.patch),
}

// PromiseState defines the state of a promise: either pending, resolved, or
// rejected. Promises that are pending have no result value yet. A promise that
// is resolved has a result value, and a promise that is rejected has a result
// value that is usually the error.
type PromiseState uint8

const (
	PromiseStatePending PromiseState = iota
	PromiseStateResolved
	PromiseStateRejected
	kNumPromiseStates
)

var promiseStateStrings = [kNumPromiseStates]string{"Pending", "Resolved", "Rejected"}

func (s PromiseState) String() string {
	if s < 0 || s >= kNumPromiseStates {
		return fmt.Sprintf("InvalidPromiseState:%d", int(s))
	}
	return promiseStateStrings[s]
}

var initOnce sync.Once

func Initialize() {
	initOnce.Do(func() {
		C.v8_Initialize()
	})
}
