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

var eventCallbackMap = make(map[int]func(err string, data []byte))

var eventCallbackFunc = eventCallback

//export eventCallback
func eventCallback(event *C.cl_event, event_command_exec_status C.cl_int, user_data unsafe.Pointer) {
	fmt.Println(event, event_command_exec_status, user_data)
}
