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

type (
	Bool  C.cl_bool
	Int   C.cl_int
	Uint  C.cl_uint
	Ulong C.cl_ulong
	Size  C.size_t
)

func ToGoBool(b Bool) bool {
	return b != C.CL_FALSE
}

func Flush(cq CommandQueue) error {
	return toError(C.clFlush(cq))
}

func Finish(cq CommandQueue) error {
	return toError(C.clFinish(cq))
}
