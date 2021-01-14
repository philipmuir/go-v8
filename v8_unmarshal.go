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
	"reflect"
)

func (v *Value) Unmarshal(t reflect.Type) (*reflect.Value, error) {
	if t == valueType {
		v := reflect.ValueOf(v)
		return &v, nil
	}

	switch t.Kind() {
	case reflect.Bool:
		if value, err := v.Bool(); err != nil {
			return nil, err
		} else {
			v := reflect.ValueOf(value).Convert(t)
			return &v, nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if value, err := v.Int64(); err != nil {
			return nil, err
		} else {
			v := reflect.ValueOf(value).Convert(t)
			return &v, nil
		}
	case reflect.Float32, reflect.Float64:
		if value, err := v.Float64(); err != nil {
			return nil, err
		} else {
			v := reflect.ValueOf(value).Convert(t)
			return &v, nil
		}
	case reflect.Array, reflect.Slice:
	case reflect.Func:
	case reflect.Ptr, reflect.Interface:
	case reflect.Map:
	case reflect.String:
		v := reflect.ValueOf(v.String()).Convert(t)
		return &v, nil
	case reflect.Struct:
	}

	panic(fmt.Sprintf("unsupported kind: %v", t.Kind()))
}
