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
	Event                  C.cl_event
	EventInfo              C.cl_event_info
	ProfilingInfo          C.cl_profiling_info
	CommandType            C.cl_command_type
	CommandExecutionStatus C.cl_int
)

const (
	EventCommandQueue           EventInfo = C.CL_EVENT_COMMAND_QUEUE
	EventCommandType            EventInfo = C.CL_EVENT_COMMAND_TYPE
	EventReferenceCount         EventInfo = C.CL_EVENT_REFERENCE_COUNT
	EventCommandExecutionStatus EventInfo = C.CL_EVENT_COMMAND_EXECUTION_STATUS
	EventContext                EventInfo = C.CL_EVENT_CONTEXT
)

const (
	CommandNdrangeKernel        CommandType = C.CL_COMMAND_NDRANGE_KERNEL
	CommandTask                 CommandType = C.CL_COMMAND_TASK
	CommandNativeKernel         CommandType = C.CL_COMMAND_NATIVE_KERNEL
	CommandReadBuffer           CommandType = C.CL_COMMAND_READ_BUFFER
	CommandWriteBuffer          CommandType = C.CL_COMMAND_WRITE_BUFFER
	CommandCopyBuffer           CommandType = C.CL_COMMAND_COPY_BUFFER
	CommandReadImage            CommandType = C.CL_COMMAND_READ_IMAGE
	CommandWriteImage           CommandType = C.CL_COMMAND_WRITE_IMAGE
	CommandCopyImage            CommandType = C.CL_COMMAND_COPY_IMAGE
	CommandCopyImageToBuffer    CommandType = C.CL_COMMAND_COPY_IMAGE_TO_BUFFER
	CommandCopyBufferToImage    CommandType = C.CL_COMMAND_COPY_BUFFER_TO_IMAGE
	CommandMapBuffer            CommandType = C.CL_COMMAND_MAP_BUFFER
	CommandMapImage             CommandType = C.CL_COMMAND_MAP_IMAGE
	CommandUnmapMemoryObject    CommandType = C.CL_COMMAND_UNMAP_MEM_OBJECT
	CommandMarker               CommandType = C.CL_COMMAND_MARKER
	CommandAcquireGlObjects     CommandType = C.CL_COMMAND_ACQUIRE_GL_OBJECTS
	CommandReleaseGlObjects     CommandType = C.CL_COMMAND_RELEASE_GL_OBJECTS
	CommandReadBufferRectangle  CommandType = C.CL_COMMAND_READ_BUFFER_RECT
	CommandWriteBufferRectangle CommandType = C.CL_COMMAND_WRITE_BUFFER_RECT
	CommandCopyBufferRectangle  CommandType = C.CL_COMMAND_COPY_BUFFER_RECT
	CommandUser                 CommandType = C.CL_COMMAND_USER
)

const (
	Complete  CommandExecutionStatus = C.CL_COMPLETE
	Running   CommandExecutionStatus = C.CL_RUNNING
	Submitted CommandExecutionStatus = C.CL_SUBMITTED
	Queued    CommandExecutionStatus = C.CL_QUEUED
)

const (
	ProfilingCommandQueued ProfilingInfo = C.CL_PROFILING_COMMAND_QUEUED
	ProfilingCommandSubmit ProfilingInfo = C.CL_PROFILING_COMMAND_SUBMIT
	ProfilingCommandStart  ProfilingInfo = C.CL_PROFILING_COMMAND_START
	ProfilingCommandEnd    ProfilingInfo = C.CL_PROFILING_COMMAND_END
)

func Flush(cq CommandQueue) error {
	return toError(C.clFlush(cq))
}

func Finish(cq CommandQueue) error {
	return toError(C.clFinish(cq))
}

func toCWaitList(wait_list []Event) (event_wait_list *C.cl_event, num_events_in_wait_list C.cl_uint) {
	if wait_list != nil {
		event_wait_list = (*C.cl_event)(&wait_list[0])
		num_events_in_wait_list = C.cl_uint(len(wait_list))
	}
	return
}

func CreateUserEvent(context Context) (Event, error) {
	var err C.cl_int
	result := C.clCreateUserEvent(context, &err)
	return Event(result), toError(err)
}

func GetEventInfo(event Event, paramName EventInfo, paramValueSize Size, paramValue unsafe.Pointer,
	paramValueSizeRet *Size) error {

	return toError(C.clGetEventInfo(event, C.cl_event_info(paramName), C.size_t(paramValueSize), paramValue,
		(*C.size_t)(paramValueSizeRet)))
}

func GetEventProfilingInfo(event Event, paramName ProfilingInfo, paramValueSize Size, paramValue unsafe.Pointer,
	paramValueSizeRet *Size) error {

	return toError(C.clGetEventProfilingInfo(event, C.cl_profiling_info(paramName), C.size_t(paramValueSize),
		paramValue, (*C.size_t)(paramValueSizeRet)))
}

func SetEventCallback(event Event, command_exec_callback_type CommandExecutionStatus,
	callback func(event Event, event_command_exec_status CommandExecutionStatus, user_data interface{}),
	user_data interface{}) error {

	return nil
}