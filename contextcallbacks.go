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
	"fmt"
	"unsafe"
)

//export callback
func callback(errinfo *C.char, private_info unsafe.Pointer, cb C.size_t, user_data unsafe.Pointer) {
	errString := C.GoString(errinfo)
	private := C.GoBytes(private_info, C.int(cb))
	functionID := *(*int)(user_data)
	fmt.Println(errString)
	fmt.Println(private)
	fmt.Println(functionID)
	// TODO use user data to lookup the callback function. User can use closures
	// to pass private data.
}
