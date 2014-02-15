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
import "unsafe"

type (
	Kernel              C.cl_kernel
	KernelInfo          C.cl_kernel_info
	KernelWorkGroupInfo C.cl_kernel_work_group_info
)

const (
	KernelFunctionName    = KernelInfo(C.CL_KERNEL_FUNCTION_NAME)
	KernelNumArgs         = KernelInfo(C.CL_KERNEL_NUM_ARGS)
	KernelReference_count = KernelInfo(C.CL_KERNEL_REFERENCE_COUNT)
	KernelContext         = KernelInfo(C.CL_KERNEL_CONTEXT)
	KernelProgram         = KernelInfo(C.CL_KERNEL_PROGRAM)
)

const (
	KernelWorkGroupSize                  = KernelWorkGroupInfo(C.CL_KERNEL_WORK_GROUP_SIZE)
	KernelCompileWorkGroupSize           = KernelWorkGroupInfo(C.CL_KERNEL_COMPILE_WORK_GROUP_SIZE)
	KernelLocalMemSize                   = KernelWorkGroupInfo(C.CL_KERNEL_LOCAL_MEM_SIZE)
	KernelPreferredWorkGroupSizeMultiple = KernelWorkGroupInfo(C.CL_KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE)
	KernelPrivateMemSize                 = KernelWorkGroupInfo(C.CL_KERNEL_PRIVATE_MEM_SIZE)
)

func CreateKernel(program Program, kernel_name string) (Kernel, error) {

	name := C.CString(kernel_name)
	defer C.free(unsafe.Pointer(name))

	var err C.cl_int
	kernel := C.clCreateKernel(program, name, &err)

	return Kernel(kernel), toError(err)
}

func GetKernelInfo(kernel Kernel, param_name KernelInfo, param_value_size Size, param_value unsafe.Pointer,
	param_value_size_ret *Size) error {

	return toError(C.clGetKernelInfo(kernel, C.cl_kernel_info(param_name), C.size_t(param_value_size),
		param_value, (*C.size_t)(param_value_size_ret)))
}

func GetKernelWorkGroupInfo(kernel Kernel, device DeviceID, param_name KernelWorkGroupInfo, param_value_size Size,
	param_value unsafe.Pointer, param_value_size_ret *Size) error {

	return toError(C.clGetKernelWorkGroupInfo(kernel, device, C.cl_kernel_work_group_info(param_name),
		C.size_t(param_value_size), param_value, (*C.size_t)(param_value_size_ret)))
}
