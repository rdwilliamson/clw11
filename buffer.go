package clw11

/*
#define CL_USE_DEPRECATED_OPENCL_1_1_APIS
#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif

extern void bufferCallback(cl_mem memobj, void *user_data);

void callBufferCallback(cl_mem memobj, void *user_data)
{
	bufferCallback(memobj, user_data);
}
*/
import "C"
import (
	"unsafe"
)

type (
	Mem              C.cl_mem
	MemFlags         C.cl_mem_flags
	MapFlags         C.cl_map_flags
	BufferRegion     C.cl_buffer_region
	BufferCreateType C.cl_buffer_create_type
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

const (
	BufferCreateTypeRegion BufferCreateType = C.CL_BUFFER_CREATE_TYPE_REGION
)

// Creates a buffer object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateBuffer.html
func CreateBuffer(context Context, flags MemFlags, size Size, host_ptr unsafe.Pointer) (Mem, error) {

	var err C.cl_int
	memory := C.clCreateBuffer(context, C.cl_mem_flags(flags), C.size_t(size), host_ptr, &err)

	return Mem(memory), toError(err)
}

// Creates a buffer object (referred to as a sub-buffer object) from an existing
// buffer object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateSubBuffer.html
func CreateSubBuffer(buffer Mem, flags MemFlags, buffer_create_type BufferCreateType, buffer_create_info unsafe.Pointer,
	errcode_ret *Int) (Mem, error) {

	var err C.cl_int
	memory := C.clCreateSubBuffer(buffer, C.cl_mem_flags(flags), C.cl_buffer_create_type(buffer_create_type),
		buffer_create_info, &err)

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

// Enqueue commands to read from a buffer object to host memory.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueReadBuffer.html
func EnqueueReadBuffer(command_queue CommandQueue, buffer Mem, blocking_read Bool, offset, cb Size,
	ptr unsafe.Pointer, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueReadBuffer(command_queue, buffer, C.cl_bool(blocking_read), C.size_t(offset),
		C.size_t(cb), ptr, num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

// Enqueue commands to write to a buffer object from host memory.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueWriteBuffer.html
func EnqueueWriteBuffer(command_queue CommandQueue, buffer Mem, blocking_read Bool, offset, cb Size,
	ptr unsafe.Pointer, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueWriteBuffer(command_queue, buffer, C.cl_bool(blocking_read), C.size_t(offset),
		C.size_t(cb), ptr, num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

// Enqueue commands to read from a rectangular region from a buffer object to
// host memory.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueReadBufferRect.html
func EnqueueReadBufferRect(command_queue CommandQueue, buffer Mem, blocking_read Bool, buffer_origin, host_origin,
	region [3]Size, buffer_row_pitch, buffer_slice_pitch, host_row_pitch, host_slice_pitch Size, ptr unsafe.Pointer,
	wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueReadBufferRect(command_queue, buffer, C.cl_bool(blocking_read),
		(*C.size_t)(&buffer_origin[0]), (*C.size_t)(&host_origin[0]), (*C.size_t)(&region[0]),
		C.size_t(buffer_row_pitch), C.size_t(buffer_slice_pitch), C.size_t(host_row_pitch), C.size_t(host_slice_pitch),
		ptr, num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

// Enqueue commands to write a rectangular region to a buffer object from host
// memory.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueWriteBufferRect.html
func EnqueueWriteBufferRect(command_queue CommandQueue, buffer Mem, blocking_read Bool, buffer_origin, host_origin,
	region [3]Size, buffer_row_pitch, buffer_slice_pitch, host_row_pitch, host_slice_pitch Size, ptr unsafe.Pointer,
	wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueWriteBufferRect(command_queue, buffer, C.cl_bool(blocking_read),
		(*C.size_t)(&buffer_origin[0]), (*C.size_t)(&host_origin[0]), (*C.size_t)(&region[0]),
		C.size_t(buffer_row_pitch), C.size_t(buffer_slice_pitch), C.size_t(host_row_pitch), C.size_t(host_slice_pitch),
		ptr, num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

// Enqueues a command to copy from one buffer object to another.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueCopyBuffer.html
func EnqueueCopyBuffer(command_queue CommandQueue, src_buffer, dst_buffer Mem, src_offset, dst_offset, cb Size,
	wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueCopyBuffer(command_queue, src_buffer, dst_buffer, C.size_t(src_offset),
		C.size_t(dst_offset), C.size_t(cb), num_events_in_wait_list, event_wait_list, (*C.cl_event)(event)))
}

// Enqueues a command to copy a rectangular region from the buffer object to
// another buffer object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueCopyBufferRect.html
func EnqueueCopyBufferRect(command_queue CommandQueue, src_buffer, dst_buffer Mem, src_origin, dst_origin,
	region [3]Size, src_row_pitch, src_slice_pitch, dst_row_pitch, dst_slice_pitch Size, wait_list []Event,
	event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueCopyBufferRect(command_queue, src_buffer, dst_buffer, (*C.size_t)(&src_origin[0]),
		(*C.size_t)(&dst_origin[0]), (*C.size_t)(&region[0]), C.size_t(src_row_pitch), C.size_t(src_slice_pitch),
		C.size_t(dst_row_pitch), C.size_t(dst_slice_pitch), num_events_in_wait_list, event_wait_list,
		(*C.cl_event)(event)))
}

// Enqueues a command to map a region of the buffer object given by buffer into
// the host address space and returns a pointer to this mapped region.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueMapBuffer.html
func EnqueueMapBuffer(command_queue CommandQueue, buffer Mem, blocking_map Bool, map_flags MapFlags, offset, cb Size,
	wait_list []Event, event *Event) (unsafe.Pointer, error) {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	var err C.cl_int
	mapped := C.clEnqueueMapBuffer(command_queue, buffer, C.cl_bool(blocking_map), C.cl_map_flags(map_flags),
		C.size_t(offset), C.size_t(cb), num_events_in_wait_list, event_wait_list, (*C.cl_event)(event), &err)

	return mapped, toError(err)
}

// Enqueues a command to unmap a previously mapped region of a memory object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueUnmapMemObject.html
func EnqueueUnmapMemObject(command_queue CommandQueue, memobj Mem, mapped_ptr unsafe.Pointer, wait_list []Event,
	event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueUnmapMemObject(command_queue, memobj, mapped_ptr, num_events_in_wait_list,
		event_wait_list, (*C.cl_event)(event)))
}

// Registers a user callback function that will be called when the memory object
// is deleted and its resources freed.
func SetMemObjectDestructorCallback(memobj Mem, callback BufferCallbackFunc, user_data interface{}) error {

	key := bufferCallbacks.add(callback, user_data)

	err := toError(C.clSetMemObjectDestructorCallback(C.cl_mem(memobj), (*[0]byte)(C.callBufferCallback),
		unsafe.Pointer(key)))

	if err != nil {
		// If the C side setting of the callback failed GetCallback will remove
		// the callback from the map.
		bufferCallbacks.get(key)
	}

	return err
}
