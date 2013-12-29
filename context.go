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

type (
	Context           C.cl_context
	ContextProperties C.cl_context_properties
)

func clCreateContext(properties []ContextProperties, devices []DeviceID) (Context, error) {
	var result Context
	return result, nil
}
