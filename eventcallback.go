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

type eventCallbackMapStruct struct {
	sync.Mutex
	callbackMap map[uintptr]eventCallbackData
	counter     uintptr
}

func (ecm *eventCallbackMapStruct) setCallback(function EventCallbackFunc, userData interface{}) uintptr {

	ecm.Lock()
	key := ecm.counter
	ecm.counter++
	ecm.callbackMap[key] = eventCallbackData{function, userData}
	ecm.Unlock()

	return key
}

func (ecm *eventCallbackMapStruct) getCallback(key uintptr) (EventCallbackFunc, interface{}) {

	ecm.Lock()
	data := ecm.callbackMap[key]
	delete(ecm.callbackMap, key)
	ecm.Unlock()

	return data.function, data.userData
}

var (
	EventCallbackFunction = eventCallback
	eventCallbackMap      = eventCallbackMapStruct{callbackMap: map[uintptr]eventCallbackData{}}
)

//export eventCallback
func eventCallback(event C.cl_event, event_command_exec_status C.cl_int, user_data unsafe.Pointer) {

	callback, userData := eventCallbackMap.getCallback(uintptr(user_data))
	callback(Event(event), CommandExecutionStatus(event_command_exec_status), userData)
}
