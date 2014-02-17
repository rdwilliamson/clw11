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

func SetKernelArg(kernel Kernel, arg_index Uint, arg_size Size, arg_value unsafe.Pointer) error {
	return toError(C.clSetKernelArg(kernel, C.cl_uint(arg_index), C.size_t(arg_size), arg_value))
}

func EnqueueNDRangeKernel(command_queue CommandQueue, kernel Kernel, global_work_offset, global_work_size,
	local_work_size []Size, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)
	return toError(C.clEnqueueNDRangeKernel(command_queue, kernel, C.cl_uint(len(global_work_offset)),
		(*C.size_t)(&global_work_offset[0]), (*C.size_t)(&global_work_size[0]), (*C.size_t)(&local_work_size[0]),
		num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}
