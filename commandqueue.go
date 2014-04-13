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

type (
	CommandQueue           C.cl_command_queue
	CommandQueueProperties C.cl_command_queue_properties
)

// Bitfield.
const (
	QueueOutOfOrderExecModeEnable CommandQueueProperties = C.CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE
	QueueProfilingEnable          CommandQueueProperties = C.CL_QUEUE_PROFILING_ENABLE
)

// Create a command-queue on a specific device.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateCommandQueue.html
func CreateCommandQueue(context Context, device DeviceID, properties CommandQueueProperties) (CommandQueue, error) {
	var err C.cl_int
	result := C.clCreateCommandQueue(context, device, C.cl_command_queue_properties(properties), &err)
	return CommandQueue(result), toError(err)
}

// Increments the command_queue reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainCommandQueue.html
func RetainCommandQueue(command_queue CommandQueue) error {
	return toError(C.clRetainCommandQueue(command_queue))
}

// Decrements the command_queue reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseCommandQueue.html
func ReleaseCommandQueue(command_queue CommandQueue) error {
	return toError(C.clReleaseCommandQueue(command_queue))
}

// Enqueues a marker command.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueMarker.html
func EnqueueMarker(command_queue CommandQueue, event *Event) error {
	return toError(C.clEnqueueMarker(command_queue, (*C.cl_event)(event)))
}

// Enqueues a wait for a specific event or a list of events to complete before
// any future commands queued in the command-queue are executed.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueWaitForEvents.html
func EnqueueWaitForEvents(command_queue CommandQueue, wait_list []Event) error {
	event_list, num_events := toEventList(wait_list)
	return toError(C.clEnqueueWaitForEvents(command_queue, num_events, event_list))
}

// A synchronization point that enqueues a barrier operation.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueBarrier.html
func EnqueueBarrier(command_queue CommandQueue) error {
	return toError(C.clEnqueueBarrier(command_queue))
}

// Issues all previously queued OpenCL commands in a command-queue to the device
// associated with the command-queue.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clFlush.html
func Flush(cq CommandQueue) error {
	return toError(C.clFlush(cq))
}

// Blocks until all previously queued OpenCL commands in a command-queue are
// issued to the associated device and have completed.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clFinish.html
func Finish(cq CommandQueue) error {
	return toError(C.clFinish(cq))
}
