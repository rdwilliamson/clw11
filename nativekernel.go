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

type NativeKernelFunc func(args interface{})

type nativeKernelFunctionData struct {
	function NativeKernelFunc
	userData interface{}
}

type nativeKernelFunctionCollection struct {
	sync.Mutex
	functionMap map[uintptr]nativeKernelFunctionData
	counter     uintptr
}

func (nkfc *nativeKernelFunctionCollection) add(function NativeKernelFunc, userData interface{}) uintptr {

	nkfc.Lock()
	key := nkfc.counter
	nkfc.counter++
	nkfc.functionMap[key] = nativeKernelFunctionData{function, userData}
	nkfc.Unlock()

	return key
}

func (nkfc *nativeKernelFunctionCollection) get(key uintptr) (NativeKernelFunc, interface{}) {

	nkfc.Lock()
	data := nkfc.functionMap[key]
	delete(nkfc.functionMap, key)
	nkfc.Unlock()

	return data.function, data.userData
}

var (
	nativeKernelFunction   = nativeKernel
	nativeKernelCollection = nativeKernelFunctionCollection{functionMap: map[uintptr]nativeKernelFunctionData{}}
)

//export nativeKernel
func nativeKernel(key unsafe.Pointer) {

	callback, userData := nativeKernelCollection.get(uintptr(key))
	callback(userData)
}
