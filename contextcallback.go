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

type ContextCallbackFunc func(err string, data []byte, userData interface{})

type contextCallbackData struct {
	function ContextCallbackFunc
	userData interface{}
}

type contextCallbackCollection struct {
	sync.Mutex
	callbackMap map[uintptr]contextCallbackData
	counter     uintptr
}

func (ccc *contextCallbackCollection) add(function ContextCallbackFunc, userData interface{}) uintptr {

	ccc.Lock()
	key := ccc.counter
	ccc.counter++
	ccc.callbackMap[key] = contextCallbackData{function, userData}
	ccc.Unlock()

	return key
}

func (ccc *contextCallbackCollection) get(key uintptr) (ContextCallbackFunc, interface{}) {

	ccc.Lock()
	data := ccc.callbackMap[key]
	ccc.Unlock()

	return data.function, data.userData
}

func (ccc *contextCallbackCollection) delete(key uintptr) {

	ccc.Lock()
	delete(ccc.callbackMap, key)
	ccc.Unlock()
}

var (
	contextCallbackFunction = contextCallback
	contextCallbacks        = contextCallbackCollection{callbackMap: map[uintptr]contextCallbackData{}}
)

//export contextCallback
func contextCallback(errinfo *C.char, private_info unsafe.Pointer, cb C.size_t, user_data unsafe.Pointer) {

	callback, userData := contextCallbacks.get(uintptr(user_data))
	callback(C.GoString(errinfo), C.GoBytes(private_info, C.int(cb)), userData)
}
