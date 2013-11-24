package clw11

/*
#cgo LDFLAGS: -lOpenCL
#include "CL/opencl.h"
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

var InvalidGlSharegroupReferenceKHR = errors.New("invalid gl sharegroup reference khr")

var InvalidD3d10DeviceKHR = errors.New("invalid d3d10 device khr")
var InvalidD3d10ResourceKHR = errors.New("invalid d3d10 resource khr")
var D3d10ResourceAlreadyAcquiredKHR = errors.New("d3d10 resource already acquired khr")
var D3d10ResourceNotAcquiredKHR = errors.New("d3d10 resource not acquired khr")

var errorMap = map[C.cl_int]error{
	-1:  DeviceNotFound,
	-2:  DeviceNotAvailable,
	-3:  CompilerNotAvailable,
	-4:  MemObjectAllocationFailure,
	-5:  OutOfResources,
	-6:  OutOfHostMemory,
	-7:  ProfilingInfoNotAvailable,
	-8:  MemCopyOverlap,
	-9:  ImageFormatMismatch,
	-10: ImageFormatNotSupported,
	-11: BuildProgramFailure,
	-12: MapFailure,
	-13: MisalignedSubBufferOffset,
	-14: ExecStatusErrorForEventsInWaitList,

	-30: InvalidValue,
	-31: InvalidDeviceType,
	-32: InvalidPlatform,
	-33: InvalidDevice,
	-34: InvalidContext,
	-35: InvalidQueueProperties,
	-36: InvalidCommandQueue,
	-37: InvalidHostPtr,
	-38: InvalidMemObject,
	-39: InvalidImageFormatDescriptor,
	-40: InvalidImageSize,
	-41: InvalidSampler,
	-42: InvalidBinary,
	-43: InvalidBuildOptions,
	-44: InvalidProgram,
	-45: InvalidProgramExecutable,
	-46: InvalidKernelName,
	-47: InvalidKernelDefinition,
	-48: InvalidKernel,
	-49: InvalidArgIndex,
	-50: InvalidArgValue,
	-51: InvalidArgSize,
	-52: InvalidKernelArgs,
	-53: InvalidWorkDimension,
	-54: InvalidWorkGroupSize,
	-55: InvalidWorkItemSize,
	-56: InvalidGlobalOffset,
	-57: InvalidEventWaitList,
	-58: InvalidEvent,
	-59: InvalidOperation,
	-60: InvalidGlObject,
	-61: InvalidBufferSize,
	-62: InvalidMipLevel,
	-63: InvalidGlobalWorkSize,
	-64: InvalidProperty,

	-1000: InvalidGlSharegroupReferenceKHR,

	-1002: InvalidD3d10DeviceKHR,
	-1003: InvalidD3d10ResourceKHR,
	-1004: D3d10ResourceAlreadyAcquiredKHR,
	-1005: D3d10ResourceNotAcquiredKHR,
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
	if code == C.CL_SUCCESS /* 0 */ {
		return nil
	}

	if err := errorMap[code]; err != nil {
		return wrapError(err)
	}

	panic("unknown OpenCL error")
}
