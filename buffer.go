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
	Memory      C.cl_mem
	MemoryFlags C.cl_mem_flags
	MapFlags    C.cl_map_flags
)

// Bitfield.
const (
	MemoryReadWrite        MemoryFlags = C.CL_MEM_READ_WRITE
	MemoryWriteOnly        MemoryFlags = C.CL_MEM_WRITE_ONLY
	MemoryReadOnly         MemoryFlags = C.CL_MEM_READ_ONLY
	MemoryUseHostPointer   MemoryFlags = C.CL_MEM_USE_HOST_PTR
	MemoryAllocHostPointer MemoryFlags = C.CL_MEM_ALLOC_HOST_PTR
	MemoryCopyHostPointer  MemoryFlags = C.CL_MEM_COPY_HOST_PTR
)

// Bitfield.
const (
	MapRead  MapFlags = C.CL_MAP_READ
	MapWrite MapFlags = C.CL_MAP_WRITE
)

func CreateBuffer(context Context, flags MemoryFlags, size Size, host_ptr unsafe.Pointer) (Memory, error) {

	var err C.cl_int
	memory := C.clCreateBuffer(context, C.cl_mem_flags(flags), C.size_t(size), host_ptr, &err)

	return Memory(memory), toError(err)
}

func ReleaseMemObject(memobject Memory) error {
	return toError(C.clReleaseMemObject(memobject))
}

func EnqueueReadBuffer(command_queue CommandQueue, buffer Memory, blocking_read Bool, offset, cb Size,
	ptr unsafe.Pointer, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueReadBuffer(command_queue, buffer, C.cl_bool(blocking_read), C.size_t(offset),
		C.size_t(cb), ptr, num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

func EnqueueWriteBuffer(command_queue CommandQueue, buffer Memory, blocking_read Bool, offset, cb Size,
	ptr unsafe.Pointer, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueWriteBuffer(command_queue, buffer, C.cl_bool(blocking_read), C.size_t(offset),
		C.size_t(cb), ptr, num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

func EnqueueCopyBuffer(command_queue CommandQueue, src_buffer, dst_buffer Memory, src_offset, dst_offset, cb Size,
	wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueCopyBuffer(command_queue, src_buffer, dst_buffer, C.size_t(src_offset),
		C.size_t(dst_offset), C.size_t(cb), num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

func EnqueueMapBuffer(command_queue CommandQueue, buffer Memory, blocking_map Bool, map_flags MapFlags, offset, cb Size,
	wait_list []Event, event *Event) (unsafe.Pointer, error) {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	var err C.cl_int
	mapped := C.clEnqueueMapBuffer(command_queue, buffer, C.cl_bool(blocking_map), C.cl_map_flags(map_flags),
		C.size_t(offset), C.size_t(cb), num_events_in_wait_list, event_wait_list, (*C.cl_event)(event), &err)

	return mapped, toError(err)
}

func EnqueueUnmapMemObject(command_queue CommandQueue, memobj Memory, mapped_ptr unsafe.Pointer, wait_list []Event,
	event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueUnmapMemObject(command_queue, memobj, mapped_ptr, num_events_in_wait_list,
		event_wait_list, (*C.cl_event)(event)))
}
