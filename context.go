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

func CreateContext(properties []ContextProperties, devices []DeviceID, callback ContextCallbackFunc,
	user_data interface{}) (Context, error) {

	var propertiesValue *C.cl_context_properties
	if properties != nil {
		properties = append(properties, 0)
		propertiesValue = (*C.cl_context_properties)(unsafe.Pointer(&properties[0]))
	}

	key := contextCallbackMap.setCallback(callback, user_data)

	var clErr C.cl_int
	context := Context(C.clCreateContext(propertiesValue, C.cl_uint(len(devices)),
		(*C.cl_device_id)(unsafe.Pointer(&devices[0])), (*[0]byte)(C.callContextCallback), unsafe.Pointer(key), &clErr))
	err := toError(clErr)

	if err != nil {
		// If the C side setting of the callback failed GetCallback will remove
		// the callback from the map.
		contextCallbackMap.getCallback(key)
	}

	return context, err
}

func CreateContextFromType() {

}

func RetainContext() {

}

func ReleaseContext() {

}

func GetContextInfo() {

}
