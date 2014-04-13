package clw11

/*
#define CL_USE_DEPRECATED_OPENCL_1_1_APIS
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

type ProgramCallbackFunc func(program Program, userData interface{})

type programCallbackData struct {
	function ProgramCallbackFunc
	userData interface{}
}

type programCallbackCollection struct {
	sync.Mutex
	callbackMap map[uintptr]programCallbackData
	counter     uintptr
}

func (ccc *programCallbackCollection) add(function ProgramCallbackFunc, userData interface{}) uintptr {

	ccc.Lock()
	key := ccc.counter
	ccc.counter++
	ccc.callbackMap[key] = programCallbackData{function, userData}
	ccc.Unlock()

	return key
}

func (ccc *programCallbackCollection) get(key uintptr) (ProgramCallbackFunc, interface{}) {

	ccc.Lock()
	data := ccc.callbackMap[key]
	delete(ccc.callbackMap, key)
	ccc.Unlock()

	return data.function, data.userData
}

var (
	programCallbackFunction = programCallback
	programCallbacks        = programCallbackCollection{callbackMap: map[uintptr]programCallbackData{}}
)

//export programCallback
func programCallback(program C.cl_program, user_data unsafe.Pointer) {

	callback, userData := programCallbacks.get(uintptr(user_data))
	callback(Program(program), userData)
}
