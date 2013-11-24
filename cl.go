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
	Int  C.cl_int
	Uint C.cl_uint
	Size C.size_t
)

func voidPointer(buffer []byte) unsafe.Pointer {
	if buffer != nil {
		return unsafe.Pointer(&buffer[0])
	}
	return nil
}
