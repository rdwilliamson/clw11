// A simple low level wrapper around the OpenCL 1.1 C API.
package clw11

/*
#cgo LDFLAGS: -lOpenCL
#include "CL/opencl.h"
*/
import "C"
import (
	"unsafe"
)

type (
	Bool  C.cl_bool
	Int   C.cl_int
	Uint  C.cl_uint
	Ulong C.cl_ulong
	Size  C.size_t
)

func Pointer(buffer []byte) unsafe.Pointer {
	if buffer != nil {
		return unsafe.Pointer(&buffer[0])
	}
	return nil
}

func ToGoBool(b Bool) bool {
	return b != C.CL_FALSE
}
