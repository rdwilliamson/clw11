package clw11

/*
#define CL_USE_DEPRECATED_OPENCL_1_1_APIS
#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif

extern void nativeKernel(cl_mem memobj, void *user_data);

void callNativeKernel(cl_mem memobj, void *user_data)
{
	nativeKernel(memobj, user_data);
}
*/
import "C"
import "unsafe"

type (
	Kernel              C.cl_kernel
	KernelInfo          C.cl_kernel_info
	KernelWorkGroupInfo C.cl_kernel_work_group_info
)

const (
	KernelFunctionName   = KernelInfo(C.CL_KERNEL_FUNCTION_NAME)
	KernelNumArgs        = KernelInfo(C.CL_KERNEL_NUM_ARGS)
	KernelReferenceCount = KernelInfo(C.CL_KERNEL_REFERENCE_COUNT)
	KernelContext        = KernelInfo(C.CL_KERNEL_CONTEXT)
	KernelProgram        = KernelInfo(C.CL_KERNEL_PROGRAM)
)

const (
	KernelWorkGroupSize                  = KernelWorkGroupInfo(C.CL_KERNEL_WORK_GROUP_SIZE)
	KernelCompileWorkGroupSize           = KernelWorkGroupInfo(C.CL_KERNEL_COMPILE_WORK_GROUP_SIZE)
	KernelLocalMemSize                   = KernelWorkGroupInfo(C.CL_KERNEL_LOCAL_MEM_SIZE)
	KernelPreferredWorkGroupSizeMultiple = KernelWorkGroupInfo(C.CL_KERNEL_PREFERRED_WORK_GROUP_SIZE_MULTIPLE)
	KernelPrivateMemSize                 = KernelWorkGroupInfo(C.CL_KERNEL_PRIVATE_MEM_SIZE)
)

// Creates a kernal object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateKernel.html
func CreateKernel(program Program, kernel_name string) (Kernel, error) {

	name := C.CString(kernel_name)
	defer C.free(unsafe.Pointer(name))

	var err C.cl_int
	kernel := C.clCreateKernel(program, name, &err)

	return Kernel(kernel), toError(err)
}

// Creates kernel objects for all kernel functions in a program object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateKernelsInProgram.html
func CreateKernelsInProgram(program Program, kernels []Kernel, num_kernels_ret *Uint) error {

	var num_kernels C.cl_uint
	var cKernels *C.cl_kernel
	if kernels != nil {
		num_kernels = C.cl_uint(len(kernels))
		cKernels = (*C.cl_kernel)(&kernels[0])
	}

	return toError(C.clCreateKernelsInProgram(program, num_kernels, cKernels, (*C.cl_uint)(num_kernels_ret)))
}

// Increments the kernel object reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainKernel.html
func RetainKernel(kernel Kernel) error {
	return toError(C.clRetainKernel(kernel))
}

// Decrements the kernel reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseKernel.html
func ReleaseKernel(kernel Kernel) error {
	return toError(C.clReleaseKernel(kernel))
}

// Used to set the argument value for a specific argument of a kernel.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clSetKernelArg.html
func SetKernelArg(kernel Kernel, arg_index Uint, arg_size Size, arg_value unsafe.Pointer) error {
	return toError(C.clSetKernelArg(kernel, C.cl_uint(arg_index), C.size_t(arg_size), arg_value))
}

// Returns information about the kernel object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetKernelInfo.html
func GetKernelInfo(kernel Kernel, param_name KernelInfo, param_value_size Size, param_value unsafe.Pointer,
	param_value_size_ret *Size) error {

	return toError(C.clGetKernelInfo(kernel, C.cl_kernel_info(param_name), C.size_t(param_value_size),
		param_value, (*C.size_t)(param_value_size_ret)))
}

// Returns information about the kernel object that may be specific to a device.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetKernelWorkGroupInfo.html
func GetKernelWorkGroupInfo(kernel Kernel, device DeviceID, param_name KernelWorkGroupInfo, param_value_size Size,
	param_value unsafe.Pointer, param_value_size_ret *Size) error {

	return toError(C.clGetKernelWorkGroupInfo(kernel, device, C.cl_kernel_work_group_info(param_name),
		C.size_t(param_value_size), param_value, (*C.size_t)(param_value_size_ret)))
}

// Enqueues a command to execute a kernel on a device.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueNDRangeKernel.html
func EnqueueNDRangeKernel(command_queue CommandQueue, kernel Kernel, global_work_offset, global_work_size,
	local_work_size []Size, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)
	return toError(C.clEnqueueNDRangeKernel(command_queue, kernel, C.cl_uint(len(global_work_offset)),
		(*C.size_t)(&global_work_offset[0]), (*C.size_t)(&global_work_size[0]), (*C.size_t)(&local_work_size[0]),
		num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

// Enqueues a command to execute a kernel on a device.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueTask.html
func EnqueueTask(command_queue CommandQueue, kernel Kernel, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)
	return toError(C.clEnqueueTask(command_queue, kernel, num_events_in_wait_list, event_wait_list,
		(*C.cl_event)(event)))
}

// Enqueues a command to execute a native C/C++ function not compiled using the
// OpenCL compiler.
// UNTESTED
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueNativeKernel.html
func EnqueueNativeKernel(command_queue CommandQueue, user_func NativeKernelFunc, userData interface{},
	mem_object_list []Mem, args_mem_loc *unsafe.Pointer, wait_list []Event, event *Event) error {

	var num_mem_object Uint
	var mem_list *Mem
	if mem_object_list != nil && len(mem_object_list) > 0 {
		num_mem_object = Uint(len(mem_object_list))
		mem_list = &mem_object_list[0]
	}
	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	key := nativeKernelCollection.add(user_func, userData)

	err := toError(C.clEnqueueNativeKernel(command_queue, (*[0]byte)(C.callNativeKernel), unsafe.Pointer(&key),
		C.size_t(unsafe.Sizeof(key)), C.cl_uint(num_mem_object), (*C.cl_mem)(mem_list), args_mem_loc,
		num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))

	if err != nil {
		nativeKernelCollection.get(key)
	}

	return err
}
