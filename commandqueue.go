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

type (
	CommandQueue           C.cl_command_queue
	CommandQueueProperties C.cl_command_queue_properties
)

// Bitfield.
const (
	QueueOutOfOrderExecModeEnable CommandQueueProperties = C.CL_QUEUE_OUT_OF_ORDER_EXEC_MODE_ENABLE
	QueueProfilingEnable          CommandQueueProperties = C.CL_QUEUE_PROFILING_ENABLE
)

func CreateCommandQueue(c Context, d DeviceID, p CommandQueueProperties) (CommandQueue, error) {
	var err C.cl_int
	result := C.clCreateCommandQueue(c, d, C.cl_command_queue_properties(p), &err)
	return CommandQueue(result), NewError(err)
}
