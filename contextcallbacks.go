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
	"unsafe"
)

//export callback
func callback(errinfo *C.char, private_info unsafe.Pointer, cb C.size_t, user_data unsafe.Pointer) {
}
