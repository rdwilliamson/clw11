package clw11

/*
#cgo windows linux LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif

extern void contextCallback(char *errinfo, void *private_info, size_t cb, void *user_data);

void callContextCallback(const char *errinfo, const void *private_info, size_t cb, void *user_data)
{
	contextCallback((char*)errinfo, (void*)private_info, cb, user_data);
}
*/
import "C"
import "unsafe"

type (
	Context           C.cl_context
	ContextProperties C.cl_context_properties
)

const (
	ContextPlatform ContextProperties = C.CL_CONTEXT_PLATFORM
)

var contextCallbackCounter int // FIXME broken, copy event's implementation

func CreateContext(properties []ContextProperties, devices []DeviceID,
	callback func(err string, data []byte)) (Context, error) {

	var propertiesValue *C.cl_context_properties
	if properties != nil {
		properties = append(properties, 0)
		propertiesValue = (*C.cl_context_properties)(unsafe.Pointer(&properties[0]))
	}

	// FIXME broken, copy event's implementation
	var cCallbackFunction *[0]byte
	if callback != nil {
		contextCallbackMap[contextCallbackCounter] = callback
		contextCallbackCounter++
		cCallbackFunction = (*[0]byte)(C.callContextCallback)
	}

	var clErr C.cl_int
	context := Context(C.clCreateContext(propertiesValue, C.cl_uint(len(devices)),
		(*C.cl_device_id)(unsafe.Pointer(&devices[0])), cCallbackFunction,
		unsafe.Pointer(uintptr(contextCallbackCounter)), &clErr))

	// FIXME broken, copy event's implementation
	if err := toError(clErr); err != nil {
		if callback != nil {
			contextCallbackCounter--
			delete(contextCallbackMap, contextCallbackCounter)
		}
		return context, err
	}

	return context, nil
}
