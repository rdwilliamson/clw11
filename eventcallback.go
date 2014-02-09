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
	"sync"
	"unsafe"
)

type eventCallbackData struct {
	function func(e Event, ces CommandExecutionStatus, userData interface{})
	userData interface{}
}

var (
	eventCallbackMapLock sync.RWMutex
	eventCallbackMap     = make(map[uintptr]eventCallbackData)
	eventCallbackCounter uintptr
)

var eventCallbackFunc = eventCallback

//export eventCallback
func eventCallback(event C.cl_event, event_command_exec_status C.cl_int, user_data unsafe.Pointer) {

	eventCallbackMapLock.RLock()
	callback := eventCallbackMap[uintptr(user_data)]
	eventCallbackMapLock.RUnlock()

	callback.function(Event(event), CommandExecutionStatus(event_command_exec_status), callback.userData)
}
