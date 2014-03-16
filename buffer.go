package clw11

/*
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
	Mem      C.cl_mem
	MemFlags C.cl_mem_flags
	MapFlags C.cl_map_flags
)

// Bitfield.
const (
	MemReadWrite        MemFlags = C.CL_MEM_READ_WRITE
	MemWriteOnly        MemFlags = C.CL_MEM_WRITE_ONLY
	MemReadOnly         MemFlags = C.CL_MEM_READ_ONLY
	MemUseHostPointer   MemFlags = C.CL_MEM_USE_HOST_PTR
	MemAllocHostPointer MemFlags = C.CL_MEM_ALLOC_HOST_PTR
	MemCopyHostPointer  MemFlags = C.CL_MEM_COPY_HOST_PTR
)

// Bitfield.
const (
	MapRead  MapFlags = C.CL_MAP_READ
	MapWrite MapFlags = C.CL_MAP_WRITE
)

// Creates a buffer object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateBuffer.html
func CreateBuffer(context Context, flags MemFlags, size Size, host_ptr unsafe.Pointer) (Mem, error) {

	var err C.cl_int
	memory := C.clCreateBuffer(context, C.cl_mem_flags(flags), C.size_t(size), host_ptr, &err)

	return Mem(memory), toError(err)
}

// Increments the memory object reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainMemObject.html
func RetainMemObject(memobj Mem) error {
	return toError(C.clRetainMemObject(memobj))
}

// Decrements the memory object reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseMemObject.html
func ReleaseMemObject(memobj Mem) error {
	return toError(C.clReleaseMemObject(memobj))
}

func EnqueueReadBuffer(command_queue CommandQueue, buffer Mem, blocking_read Bool, offset, cb Size,
	ptr unsafe.Pointer, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueReadBuffer(command_queue, buffer, C.cl_bool(blocking_read), C.size_t(offset),
		C.size_t(cb), ptr, num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

func EnqueueWriteBuffer(command_queue CommandQueue, buffer Mem, blocking_read Bool, offset, cb Size,
	ptr unsafe.Pointer, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueWriteBuffer(command_queue, buffer, C.cl_bool(blocking_read), C.size_t(offset),
		C.size_t(cb), ptr, num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

func EnqueueCopyBuffer(command_queue CommandQueue, src_buffer, dst_buffer Mem, src_offset, dst_offset, cb Size,
	wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueCopyBuffer(command_queue, src_buffer, dst_buffer, C.size_t(src_offset),
		C.size_t(dst_offset), C.size_t(cb), num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

func EnqueueMapBuffer(command_queue CommandQueue, buffer Mem, blocking_map Bool, map_flags MapFlags, offset, cb Size,
	wait_list []Event, event *Event) (unsafe.Pointer, error) {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	var err C.cl_int
	mapped := C.clEnqueueMapBuffer(command_queue, buffer, C.cl_bool(blocking_map), C.cl_map_flags(map_flags),
		C.size_t(offset), C.size_t(cb), num_events_in_wait_list, event_wait_list, (*C.cl_event)(event), &err)

	return mapped, toError(err)
}

func EnqueueUnmapMemObject(command_queue CommandQueue, memobj Mem, mapped_ptr unsafe.Pointer, wait_list []Event,
	event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueUnmapMemObject(command_queue, memobj, mapped_ptr, num_events_in_wait_list,
		event_wait_list, (*C.cl_event)(event)))
}
