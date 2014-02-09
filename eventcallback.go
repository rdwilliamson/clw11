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

type eventCallbackGoFunction func(e Event, ces CommandExecutionStatus, userData interface{})

type eventCallbackData struct {
	function eventCallbackGoFunction
	userData interface{}
}

type eventCallbackMapStruct struct {
	sync.Mutex
	callbackMap map[uintptr]eventCallbackData
	counter     uintptr
}

func (ecm *eventCallbackMapStruct) SetCallback(function eventCallbackGoFunction, userData interface{}) uintptr {
	ecm.Lock()
	defer ecm.Unlock()

	key := ecm.counter
	ecm.callbackMap[key] = eventCallbackData{function, userData}
	ecm.counter++

	return key
}

func (ecm *eventCallbackMapStruct) GetCallback(key uintptr) (eventCallbackGoFunction, interface{}) {
	ecm.Lock()
	defer ecm.Unlock()

	data := ecm.callbackMap[key]
	delete(ecm.callbackMap, key)

	return data.function, data.userData
}

var (
	eventCallbackFunc = eventCallback
	eventCallbackMap  = eventCallbackMapStruct{callbackMap: map[uintptr]eventCallbackData{}}
)

//export eventCallback
func eventCallback(event C.cl_event, event_command_exec_status C.cl_int, user_data unsafe.Pointer) {

	callback, userData := eventCallbackMap.GetCallback(uintptr(user_data))
	callback(Event(event), CommandExecutionStatus(event_command_exec_status), userData)
}
