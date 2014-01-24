// A simple low level wrapper around the OpenCL 1.1 C API.
package clw11

/*
#cgo windows linux LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type (
	Bool  C.cl_bool
	Int   C.cl_int
	Uint  C.cl_uint
	Ulong C.cl_ulong
	Size  C.size_t
)

func (b Bool) Bool() bool {
	return b == C.CL_TRUE
}

func ToBool(b bool) Bool {
	if b {
		return C.CL_TRUE
	}
	return C.CL_FALSE
}

func toBytes(p unsafe.Pointer, length int) []byte {
	if p == nil {
		return nil
	}

	var result []byte
	header := (*reflect.SliceHeader)((unsafe.Pointer(&result)))
	header.Cap = length
	header.Len = length
	header.Data = uintptr(p)
	return result
}
