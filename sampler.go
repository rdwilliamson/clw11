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
	"unsafe"
)

type (
	Sampler        C.cl_sampler
	AddressingMode C.cl_addressing_mode
	FilterMode     C.cl_filter_mode
	SamplerInfo    C.cl_sampler_info
)

// Creates a sampler object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateSampler.html
func CreateSampler(context Context, normalized_coords Bool, addressing_mode AddressingMode,
	filter_mode FilterMode) (Sampler, error) {

	var err C.cl_int
	sampler := C.clCreateSampler(context, C.cl_bool(normalized_coords), C.cl_addressing_mode(addressing_mode),
		C.cl_filter_mode(filter_mode), &err)

	return Sampler(sampler), toError(err)
}

// Increments the sampler reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainSampler.html
func RetainSampler(sampler Sampler) error {
	return toError(C.clRetainSampler(sampler))
}

// Decrements the sampler reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseSampler.html
func ReleaseSampler(sampler Sampler) error {
	return toError(C.clReleaseSampler(sampler))
}

// Returns information about the sampler object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetSamplerInfo.html
func GetSamplerInfo(sampler Sampler, param_name SamplerInfo, param_value_size Size, param_value unsafe.Pointer,
	param_value_size_ret *Size) error {

	return toError(C.clGetSamplerInfo(sampler, C.cl_sampler_info(param_name), C.size_t(param_value_size),
		param_value, (*C.size_t)(param_value_size_ret)))
}
