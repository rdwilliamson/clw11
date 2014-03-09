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

type EventCallbackFunc func(e Event, ces CommandExecutionStatus, userData interface{})

type eventCallbackData struct {
	function EventCallbackFunc
	userData interface{}
}

type eventCallbackCollection struct {
	sync.Mutex
	callbackMap map[uintptr]eventCallbackData
	counter     uintptr
}

func (ecc *eventCallbackCollection) add(function EventCallbackFunc, userData interface{}) uintptr {

	ecc.Lock()
	key := ecc.counter
	ecc.counter++
	ecc.callbackMap[key] = eventCallbackData{function, userData}
	ecc.Unlock()

	return key
}

func (ecc *eventCallbackCollection) get(key uintptr) (EventCallbackFunc, interface{}) {

	ecc.Lock()
	data := ecc.callbackMap[key]
	delete(ecc.callbackMap, key)
	ecc.Unlock()

	return data.function, data.userData
}

var (
	EventCallbackFunction = eventCallback
	eventCallbacks        = eventCallbackCollection{callbackMap: map[uintptr]eventCallbackData{}}
)

//export eventCallback
func eventCallback(event C.cl_event, event_command_exec_status C.cl_int, user_data unsafe.Pointer) {

	callback, userData := eventCallbacks.get(uintptr(user_data))
	callback(Event(event), CommandExecutionStatus(event_command_exec_status), userData)
}
