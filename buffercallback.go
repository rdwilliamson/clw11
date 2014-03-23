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

type BufferCallbackFunc func(memobj Mem, userData interface{})

type bufferCallbackData struct {
	function BufferCallbackFunc
	userData interface{}
}

type bufferCallbackCollection struct {
	sync.Mutex
	callbackMap map[uintptr]bufferCallbackData
	counter     uintptr
}

func (bcc *bufferCallbackCollection) add(function BufferCallbackFunc, userData interface{}) uintptr {

	bcc.Lock()
	key := bcc.counter
	bcc.counter++
	bcc.callbackMap[key] = bufferCallbackData{function, userData}
	bcc.Unlock()

	return key
}

func (bcc *bufferCallbackCollection) get(key uintptr) (BufferCallbackFunc, interface{}) {

	bcc.Lock()
	data := bcc.callbackMap[key]
	delete(bcc.callbackMap, key)
	bcc.Unlock()

	return data.function, data.userData
}

var (
	BufferCallbackFunction = bufferCallback
	bufferCallbacks        = bufferCallbackCollection{callbackMap: map[uintptr]bufferCallbackData{}}
)

//export bufferCallback
func bufferCallback(memobj C.cl_mem, user_data unsafe.Pointer) {

	callback, userData := bufferCallbacks.get(uintptr(user_data))
	callback(Mem(memobj), userData)
}
