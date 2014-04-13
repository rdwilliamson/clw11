package clw11

/*
#define CL_USE_DEPRECATED_OPENCL_1_1_APIS
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
	ContextInfo       C.cl_context_info
)

const (
	ContextPlatform ContextProperties = C.CL_CONTEXT_PLATFORM
)

const (
	ContextReferenceCount ContextInfo = C.CL_CONTEXT_REFERENCE_COUNT
	ContextDevices        ContextInfo = C.CL_CONTEXT_DEVICES
	ContextPropertiesInfo ContextInfo = C.CL_CONTEXT_PROPERTIES // Appended "Info" due to conflict with type.
	ContextNumDevices     ContextInfo = C.CL_CONTEXT_NUM_DEVICES
)

// Creates an OpenCL context.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateContext.html
func CreateContext(properties []ContextProperties, devices []DeviceID, callback ContextCallbackFunc,
	user_data interface{}) (Context, error) {

	var propertiesPtr *C.cl_context_properties
	if properties != nil {
		propertiesPtr = (*C.cl_context_properties)(&properties[0])
	}

	key := contextCallbacks.add(callback, user_data)

	var err C.cl_int
	context := C.clCreateContext(propertiesPtr, C.cl_uint(len(devices)), (*C.cl_device_id)(unsafe.Pointer(&devices[0])),
		(*[0]byte)(C.callContextCallback), unsafe.Pointer(key), &err)

	if err != C.CL_SUCCESS {
		// If the C side setting of the callback failed the get callback will
		// remove the callback from the map.
		contextCallbacks.get(key)
	}

	return Context(context), toError(err)
}

// Create an OpenCL context from a device type that identifies the specific
// device(s) to use.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateContextFromType.html
func CreateContextFromType(properties []ContextProperties, device_type DeviceType, callback ContextCallbackFunc,
	user_data interface{}) (Context, error) {

	var propertiesPtr *C.cl_context_properties
	if properties != nil {
		propertiesPtr = (*C.cl_context_properties)(&properties[0])
	}

	key := contextCallbacks.add(callback, user_data)

	var err C.cl_int
	context := C.clCreateContextFromType(propertiesPtr, C.cl_device_type(device_type),
		(*[0]byte)(C.callContextCallback), unsafe.Pointer(key), &err)

	if err != C.CL_SUCCESS {
		// If the C side setting of the callback failed GetCallback will remove
		// the callback from the map.
		contextCallbacks.get(key)
	}

	return Context(context), toError(err)
}

// Increment the context reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainContext.html
func RetainContext(context Context) error {
	return toError(C.clRetainContext(context))
}

// Decrement the context reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseContext.html
func ReleaseContext(context Context) error {
	return toError(C.clReleaseContext(context))
}

// Query information about a context.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetContextInfo.html
func GetContextInfo(context Context, param_name ContextInfo, param_value_size Size, param_value unsafe.Pointer,
	param_value_size_ret *Size) error {

	return toError(C.clGetContextInfo(context, C.cl_context_info(param_name), C.size_t(param_value_size), param_value,
		(*C.size_t)(param_value_size_ret)))
}
