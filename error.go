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
import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

var DeviceNotFound = errors.New("device not found")
var DeviceNotAvailable = errors.New("device not available")
var CompilerNotAvailable = errors.New("compiler not available")
var MemObjectAllocationFailure = errors.New("mem object allocation failure")
var OutOfResources = errors.New("out of resources")
var OutOfHostMemory = errors.New("out of host memory")
var ProfilingInfoNotAvailable = errors.New("profiling info not available")
var MemCopyOverlap = errors.New("mem copy overlap")
var ImageFormatMismatch = errors.New("image format mismatch")
var ImageFormatNotSupported = errors.New("image format not supported")
var BuildProgramFailure = errors.New("build program failure")
var MapFailure = errors.New("map failure")
var MisalignedSubBufferOffset = errors.New("misaligned sub buffer offset")
var ExecStatusErrorForEventsInWaitList = errors.New("exec status error for events in wait list")

var InvalidValue = errors.New("invalid value")
var InvalidDeviceType = errors.New("invalid device type")
var InvalidPlatform = errors.New("invalid platform")
var InvalidDevice = errors.New("invalid device")
var InvalidContext = errors.New("invalid context")
var InvalidQueueProperties = errors.New("invalid queue properties")
var InvalidCommandQueue = errors.New("invalid command queue")
var InvalidHostPtr = errors.New("invalid host ptr")
var InvalidMemObject = errors.New("invalid mem object")
var InvalidImageFormatDescriptor = errors.New("invalid image format descriptor")
var InvalidImageSize = errors.New("invalid image size")
var InvalidSampler = errors.New("invalid sampler")
var InvalidBinary = errors.New("invalid binary")
var InvalidBuildOptions = errors.New("invalid build options")
var InvalidProgram = errors.New("invalid program")
var InvalidProgramExecutable = errors.New("invalid program executable")
var InvalidKernelName = errors.New("invalid kernel name")
var InvalidKernelDefinition = errors.New("invalid kernel definition")
var InvalidKernel = errors.New("invalid kernel")
var InvalidArgIndex = errors.New("invalid arg index")
var InvalidArgValue = errors.New("invalid arg value")
var InvalidArgSize = errors.New("invalid arg size")
var InvalidKernelArgs = errors.New("invalid kernel args")
var InvalidWorkDimension = errors.New("invalid work dimension")
var InvalidWorkGroupSize = errors.New("invalid work group size")
var InvalidWorkItemSize = errors.New("invalid work item size")
var InvalidGlobalOffset = errors.New("invalid global offset")
var InvalidEventWaitList = errors.New("invalid event wait list")
var InvalidEvent = errors.New("invalid event")
var InvalidOperation = errors.New("invalid operation")
var InvalidGlObject = errors.New("invalid gl object")
var InvalidBufferSize = errors.New("invalid buffer size")
var InvalidMipLevel = errors.New("invalid mip level")
var InvalidGlobalWorkSize = errors.New("invalid global work size")
var InvalidProperty = errors.New("invalid property")

// var InvalidGLSharegroupReferenceKHR = errors.New("invalid GL sharegroup reference KHR")

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

	// -1000: InvalidGLSharegroupReferenceKHR,

	// -1002: InvalidD3d10DeviceKHR,
	// -1003: InvalidD3d10ResourceKHR,
	// -1004: D3d10ResourceAlreadyAcquiredKHR,
	// -1005: D3d10ResourceNotAcquiredKHR,
}

// Error with calling function.
type Error struct {
	Function string
	Err      error
}

func (err Error) Error() string {
	return fmt.Sprint(err.Function, ": ", err.Err.Error())
}

// Gets "package.function" from call stack for error.
func wrapError(err error) error {
	pc, _, _, _ := runtime.Caller(2)
	name := runtime.FuncForPC(pc).Name()
	last := strings.LastIndex(name, "/")
	if last == -1 {
		last = 0
	} else {
		last++
	}
	return &Error{name[last:], err}
}

func NewError(code C.cl_int) error {
	if code == C.CL_SUCCESS {
		return nil
	}

	if err := errorMap[code]; err != nil {
		return wrapError(err)
	}

	panic("unknown OpenCL error")
}
