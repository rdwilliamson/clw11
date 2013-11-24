package clw11

/*
#cgo LDFLAGS: -lOpenCL
#include "CL/opencl.h"
*/
import "C"

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

func GetPlatformInfo(platform PlatformID, paramName PlatformInfo, paramValueSize Size, paramValue []byte,
	paramValueSizeRet *Size) error {

	return NewError(C.clGetPlatformInfo(C.cl_platform_id(platform), C.cl_platform_info(paramName),
		C.size_t(paramValueSize), voidPointer(paramValue), (*C.size_t)(paramValueSizeRet)))
}
