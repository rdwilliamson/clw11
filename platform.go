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
	"unsafe"
)

type (
	PlatformID   C.cl_platform_id
	PlatformInfo C.cl_platform_info
)

const (
	PlatformProfile    PlatformInfo = C.CL_PLATFORM_PROFILE
	PlatformVersion    PlatformInfo = C.CL_PLATFORM_VERSION
	PlatformName       PlatformInfo = C.CL_PLATFORM_NAME
	PlatformVendor     PlatformInfo = C.CL_PLATFORM_VENDOR
	PlatformExtensions PlatformInfo = C.CL_PLATFORM_EXTENSIONS
)

func GetPlatformIDs(numEntries Uint, platforms *PlatformID, numPlatforms *Uint) error {
	return NewError(C.clGetPlatformIDs(C.cl_uint(numEntries), (*C.cl_platform_id)(platforms),
		(*C.cl_uint)(numPlatforms)))
}

func GetPlatformInfo(platform PlatformID, paramName PlatformInfo, paramValueSize Size, paramValue unsafe.Pointer,
	paramValueSizeRet *Size) error {

	return NewError(C.clGetPlatformInfo(C.cl_platform_id(platform), C.cl_platform_info(paramName),
		C.size_t(paramValueSize), paramValue, (*C.size_t)(paramValueSizeRet)))
}
