package main

// #include "v8golang.h"
// #cgo CXXFLAGS: -fno-rtti -fpic -std=c++14 -DV8_COMPRESS_POINTERS -DV8_31BIT_SMIS_ON_64BIT_ARCH
// #cgo darwin linux CXXFLAGS: -I${SRCDIR} -I${SRCDIR}/include
// #cgo LDFLAGS: -pthread -lv8
// #cgo darwin LDFLAGS: -L${SRCDIR}/libv8/darwin_x86_64
// #cgo linux LDFLAGS: -L${SRCDIR}/libv8/linux_x86_64
import "C"
import (
	"sync"
	"fmt"
)

var v8init sync.Once

func main() {
	Init()

	fmt.Println(Version())

	ptr := C.v8_Isolate_New(C.StartupData{data: nil, length: 0})

	C.v8_Isolate_Terminate(ptr)
}

func Init() {
	v8init.Do(func() {
		C.v8_Initialize()
	})
}

func Version() string {
	return C.GoString(C.Version())
}