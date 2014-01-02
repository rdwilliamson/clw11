package clw11

/*
#cgo windows linux LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif

extern void callback(char *errinfo, void *private_info, size_t cb, void *user_data);

void callCallback(const char *errinfo, const void *private_info, size_t cb, void *user_data)
{
	callback((char*)errinfo, (void*)private_info, cb, user_data);
}
*/
import "C"
import (
	"unsafe"
)

type (
	Context           C.cl_context
	ContextProperties C.cl_context_properties
)

const (
	ContextPlatform ContextProperties = C.CL_CONTEXT_PLATFORM

	// ContextD3D10Device ContextProperties = C.CL_CONTEXT_D3D10_DEVICE_KHR

	// GLContext          ContextProperties = C.CL_GL_CONTEXT_KHR
	// EGLDisplay         ContextProperties = C.CL_EGL_DISPLAY_KHR
	// GLXDisplay         ContextProperties = C.CL_GLX_DISPLAY_KHR
	// WGLHDC             ContextProperties = C.CL_WGL_HDC_KHR
	// CGLSharegroup      ContextProperties = C.CL_CGL_SHAREGROUP_KHR
)

func CreateContext(properties []ContextProperties, devices []DeviceID) (Context, error) {

	var propertiesValue *C.cl_context_properties
	if properties != nil {
		properties = append(properties, 0)
		propertiesValue = (*C.cl_context_properties)(unsafe.Pointer(&properties[0]))
	}

	var err C.cl_int
	result := Context(C.clCreateContext(propertiesValue,
		C.cl_uint(len(devices)), (*C.cl_device_id)(unsafe.Pointer(&devices[0])), nil, nil, &err))

	return result, NewError(err)
}
