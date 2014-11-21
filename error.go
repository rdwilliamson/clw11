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
import (
	"errors"
)

var (
	DeviceNotFound                     = errors.New("cl: device not found")
	DeviceNotAvailable                 = errors.New("cl: device not available")
	CompilerNotAvailable               = errors.New("cl: compiler not available")
	MemObjectAllocationFailure         = errors.New("cl: mem object allocation failure")
	OutOfResources                     = errors.New("cl: out of resources")
	OutOfHostMemory                    = errors.New("cl: out of host memory")
	ProfilingInfoNotAvailable          = errors.New("cl: profiling info not available")
	MemCopyOverlap                     = errors.New("cl: mem copy overlap")
	ImageFormatMismatch                = errors.New("cl: image format mismatch")
	ImageFormatNotSupported            = errors.New("cl: image format not supported")
	BuildProgramFailure                = errors.New("cl: build program failure")
	MapFailure                         = errors.New("cl: map failure")
	MisalignedSubBufferOffset          = errors.New("cl: misaligned sub buffer offset")
	ExecStatusErrorForEventsInWaitList = errors.New("cl: exec status error for events in wait list")
)

var (
	InvalidValue                 = errors.New("cl: invalid value")
	InvalidDeviceType            = errors.New("cl: invalid device type")
	InvalidPlatform              = errors.New("cl: invalid platform")
	InvalidDevice                = errors.New("cl: invalid device")
	InvalidContext               = errors.New("cl: invalid context")
	InvalidQueueProperties       = errors.New("cl: invalid queue properties")
	InvalidCommandQueue          = errors.New("cl: invalid command queue")
	InvalidHostPtr               = errors.New("cl: invalid host ptr")
	InvalidMemObject             = errors.New("cl: invalid mem object")
	InvalidImageFormatDescriptor = errors.New("cl: invalid image format descriptor")
	InvalidImageSize             = errors.New("cl: invalid image size")
	InvalidSampler               = errors.New("cl: invalid sampler")
	InvalidBinary                = errors.New("cl: invalid binary")
	InvalidBuildOptions          = errors.New("cl: invalid build options")
	InvalidProgram               = errors.New("cl: invalid program")
	InvalidProgramExecutable     = errors.New("cl: invalid program executable")
	InvalidKernelName            = errors.New("cl: invalid kernel name")
	InvalidKernelDefinition      = errors.New("cl: invalid kernel definition")
	InvalidKernel                = errors.New("cl: invalid kernel")
	InvalidArgIndex              = errors.New("cl: invalid arg index")
	InvalidArgValue              = errors.New("cl: invalid arg value")
	InvalidArgSize               = errors.New("cl: invalid arg size")
	InvalidKernelArgs            = errors.New("cl: invalid kernel args")
	InvalidWorkDimension         = errors.New("cl: invalid work dimension")
	InvalidWorkGroupSize         = errors.New("cl: invalid work group size")
	InvalidWorkItemSize          = errors.New("cl: invalid work item size")
	InvalidGlobalOffset          = errors.New("cl: invalid global offset")
	InvalidEventWaitList         = errors.New("cl: invalid event wait list")
	InvalidEvent                 = errors.New("cl: invalid event")
	InvalidOperation             = errors.New("cl: invalid operation")
	InvalidGlObject              = errors.New("cl: invalid gl object")
	InvalidBufferSize            = errors.New("cl: invalid buffer size")
	InvalidMipLevel              = errors.New("cl: invalid mip level")
	InvalidGlobalWorkSize        = errors.New("cl: invalid global work size")
	InvalidProperty              = errors.New("cl: invalid property")
)

var errorMap = map[C.cl_int]error{

	C.CL_DEVICE_NOT_FOUND:                          DeviceNotFound,
	C.CL_DEVICE_NOT_AVAILABLE:                      DeviceNotAvailable,
	C.CL_COMPILER_NOT_AVAILABLE:                    CompilerNotAvailable,
	C.CL_MEM_OBJECT_ALLOCATION_FAILURE:             MemObjectAllocationFailure,
	C.CL_OUT_OF_RESOURCES:                          OutOfResources,
	C.CL_OUT_OF_HOST_MEMORY:                        OutOfHostMemory,
	C.CL_PROFILING_INFO_NOT_AVAILABLE:              ProfilingInfoNotAvailable,
	C.CL_MEM_COPY_OVERLAP:                          MemCopyOverlap,
	C.CL_IMAGE_FORMAT_MISMATCH:                     ImageFormatMismatch,
	C.CL_IMAGE_FORMAT_NOT_SUPPORTED:                ImageFormatNotSupported,
	C.CL_BUILD_PROGRAM_FAILURE:                     BuildProgramFailure,
	C.CL_MAP_FAILURE:                               MapFailure,
	C.CL_MISALIGNED_SUB_BUFFER_OFFSET:              MisalignedSubBufferOffset,
	C.CL_EXEC_STATUS_ERROR_FOR_EVENTS_IN_WAIT_LIST: ExecStatusErrorForEventsInWaitList,

	C.CL_INVALID_VALUE:                   InvalidValue,
	C.CL_INVALID_DEVICE_TYPE:             InvalidDeviceType,
	C.CL_INVALID_PLATFORM:                InvalidPlatform,
	C.CL_INVALID_DEVICE:                  InvalidDevice,
	C.CL_INVALID_CONTEXT:                 InvalidContext,
	C.CL_INVALID_QUEUE_PROPERTIES:        InvalidQueueProperties,
	C.CL_INVALID_COMMAND_QUEUE:           InvalidCommandQueue,
	C.CL_INVALID_HOST_PTR:                InvalidHostPtr,
	C.CL_INVALID_MEM_OBJECT:              InvalidMemObject,
	C.CL_INVALID_IMAGE_FORMAT_DESCRIPTOR: InvalidImageFormatDescriptor,
	C.CL_INVALID_IMAGE_SIZE:              InvalidImageSize,
	C.CL_INVALID_SAMPLER:                 InvalidSampler,
	C.CL_INVALID_BINARY:                  InvalidBinary,
	C.CL_INVALID_BUILD_OPTIONS:           InvalidBuildOptions,
	C.CL_INVALID_PROGRAM:                 InvalidProgram,
	C.CL_INVALID_PROGRAM_EXECUTABLE:      InvalidProgramExecutable,
	C.CL_INVALID_KERNEL_NAME:             InvalidKernelName,
	C.CL_INVALID_KERNEL_DEFINITION:       InvalidKernelDefinition,
	C.CL_INVALID_KERNEL:                  InvalidKernel,
	C.CL_INVALID_ARG_INDEX:               InvalidArgIndex,
	C.CL_INVALID_ARG_VALUE:               InvalidArgValue,
	C.CL_INVALID_ARG_SIZE:                InvalidArgSize,
	C.CL_INVALID_KERNEL_ARGS:             InvalidKernelArgs,
	C.CL_INVALID_WORK_DIMENSION:          InvalidWorkDimension,
	C.CL_INVALID_WORK_GROUP_SIZE:         InvalidWorkGroupSize,
	C.CL_INVALID_WORK_ITEM_SIZE:          InvalidWorkItemSize,
	C.CL_INVALID_GLOBAL_OFFSET:           InvalidGlobalOffset,
	C.CL_INVALID_EVENT_WAIT_LIST:         InvalidEventWaitList,
	C.CL_INVALID_EVENT:                   InvalidEvent,
	C.CL_INVALID_OPERATION:               InvalidOperation,
	C.CL_INVALID_GL_OBJECT:               InvalidGlObject,
	C.CL_INVALID_BUFFER_SIZE:             InvalidBufferSize,
	C.CL_INVALID_MIP_LEVEL:               InvalidMipLevel,
	C.CL_INVALID_GLOBAL_WORK_SIZE:        InvalidGlobalWorkSize,
	C.CL_INVALID_PROPERTY:                InvalidProperty,
}

// CodeToError converts an OpenCL int to an OpenCL error. It panics if there is
// no corresponding error.
func CodeToError(code Int) error {
	return toError(C.cl_int(code))
}

func toError(code C.cl_int) error {
	if code == C.CL_SUCCESS {
		return nil
	}
	if err := errorMap[code]; err != nil {
		return err
	}
	panic("unknown OpenCL error")
}
