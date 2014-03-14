// A simple low level wrapper around the OpenCL 1.1 C API.
package clw11

/*
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

const (
	True  = Bool(C.CL_TRUE)
	False = Bool(C.CL_FALSE)
)
