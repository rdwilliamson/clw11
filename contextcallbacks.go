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

type ContextCallbackFunc func(err string, data []byte, userData interface{})

type contextCallbackData struct {
	function ContextCallbackFunc
	userData interface{}
}

type contextCallbackMapStruct struct {
	sync.Mutex
	callbackMap map[uintptr]contextCallbackData
	counter     uintptr
}

func (ccm *contextCallbackMapStruct) setCallback(function ContextCallbackFunc, userData interface{}) uintptr {

	ccm.Lock()
	key := ccm.counter
	ccm.counter++
	ccm.callbackMap[key] = contextCallbackData{function, userData}
	ccm.Unlock()

	return key
}

func (ccm *contextCallbackMapStruct) getCallback(key uintptr) (ContextCallbackFunc, interface{}) {

	ccm.Lock()
	data := ccm.callbackMap[key]
	delete(ccm.callbackMap, key)
	ccm.Unlock()

	return data.function, data.userData
}

var (
	contextCallbackFunction = contextCallback
	contextCallbackMap      = contextCallbackMapStruct{callbackMap: map[uintptr]contextCallbackData{}}
)

//export contextCallback
func contextCallback(errinfo *C.char, private_info unsafe.Pointer, cb C.size_t, user_data unsafe.Pointer) {

}
