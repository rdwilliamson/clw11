package clw11

/*
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
