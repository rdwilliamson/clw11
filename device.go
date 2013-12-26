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
	"unsafe"
)

type (
	DeviceID               C.cl_device_id
	DeviceInfo             C.cl_device_info
	DeviceExecCapabilities C.cl_device_exec_capabilities
	DeviceFPConfig         C.cl_device_fp_config
	DeviceLocalMemType     C.cl_device_local_mem_type
	DeviceMemCacheType     C.cl_device_mem_cache_type
	DeviceType             C.cl_device_type
)

// Bitfield.
const (
	DeviceTypeDefault     DeviceType = C.CL_DEVICE_TYPE_DEFAULT
	DeviceTypeCpu         DeviceType = C.CL_DEVICE_TYPE_CPU
	DeviceTypeGpu         DeviceType = C.CL_DEVICE_TYPE_GPU
	DeviceTypeAccelerator DeviceType = C.CL_DEVICE_TYPE_ACCELERATOR
	DeviceTypeAll         DeviceType = C.CL_DEVICE_TYPE_ALL
)

const (
	DeviceTypeInfo                   DeviceInfo = C.CL_DEVICE_TYPE // Appended "Info" due to conflict with type.
	DeviceVendorID                   DeviceInfo = C.CL_DEVICE_VENDOR_ID
	DeviceMaxComputeUnits            DeviceInfo = C.CL_DEVICE_MAX_COMPUTE_UNITS
	DeviceMaxWorkItemDimensions      DeviceInfo = C.CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS
	DeviceMaxWorkGroupSize           DeviceInfo = C.CL_DEVICE_MAX_WORK_GROUP_SIZE
	DeviceMaxWorkItemSizes           DeviceInfo = C.CL_DEVICE_MAX_WORK_ITEM_SIZES
	DevicePreferredVectorWidthChar   DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_CHAR
	DevicePreferredVectorWidthShort  DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_SHORT
	DevicePreferredVectorWidthInt    DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_INT
	DevicePreferredVectorWidthLong   DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_LONG
	DevicePreferredVectorWidthFloat  DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_FLOAT
	DevicePreferredVectorWidthDouble DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_DOUBLE
	DeviceMaxClockFrequency          DeviceInfo = C.CL_DEVICE_MAX_CLOCK_FREQUENCY
	DeviceAddressBits                DeviceInfo = C.CL_DEVICE_ADDRESS_BITS
	DeviceMaxReadImageArgs           DeviceInfo = C.CL_DEVICE_MAX_READ_IMAGE_ARGS
	DeviceMaxWriteImageArgs          DeviceInfo = C.CL_DEVICE_MAX_WRITE_IMAGE_ARGS
	DeviceMaxMemAllocSize            DeviceInfo = C.CL_DEVICE_MAX_MEM_ALLOC_SIZE
	DeviceImage2dMaxWidth            DeviceInfo = C.CL_DEVICE_IMAGE2D_MAX_WIDTH
	DeviceImage2dMaxHeight           DeviceInfo = C.CL_DEVICE_IMAGE2D_MAX_HEIGHT
	DeviceImage3dMaxWidth            DeviceInfo = C.CL_DEVICE_IMAGE3D_MAX_WIDTH
	DeviceImage3dMaxHeight           DeviceInfo = C.CL_DEVICE_IMAGE3D_MAX_HEIGHT
	DeviceImage3dMaxDepth            DeviceInfo = C.CL_DEVICE_IMAGE3D_MAX_DEPTH
	DeviceImageSupport               DeviceInfo = C.CL_DEVICE_IMAGE_SUPPORT
	DeviceMaxParameterSize           DeviceInfo = C.CL_DEVICE_MAX_PARAMETER_SIZE
	DeviceMaxSamplers                DeviceInfo = C.CL_DEVICE_MAX_SAMPLERS
	DeviceMemBaseAddrAlign           DeviceInfo = C.CL_DEVICE_MEM_BASE_ADDR_ALIGN
	DeviceMinDataTypeAlignSize       DeviceInfo = C.CL_DEVICE_MIN_DATA_TYPE_ALIGN_SIZE
	DeviceSingleFpConfig             DeviceInfo = C.CL_DEVICE_SINGLE_FP_CONFIG
	DeviceDoubleFpConfig             DeviceInfo = C.CL_DEVICE_DOUBLE_FP_CONFIG
	DeviceGlobalMemCacheType         DeviceInfo = C.CL_DEVICE_GLOBAL_MEM_CACHE_TYPE
	DeviceGlobalMemCachelineSize     DeviceInfo = C.CL_DEVICE_GLOBAL_MEM_CACHELINE_SIZE
	DeviceGlobalMemCacheSize         DeviceInfo = C.CL_DEVICE_GLOBAL_MEM_CACHE_SIZE
	DeviceGlobalMemSize              DeviceInfo = C.CL_DEVICE_GLOBAL_MEM_SIZE
	DeviceMaxConstantBufferSize      DeviceInfo = C.CL_DEVICE_MAX_CONSTANT_BUFFER_SIZE
	DeviceMaxConstantArgs            DeviceInfo = C.CL_DEVICE_MAX_CONSTANT_ARGS
	DeviceLocalMemTypeInfo           DeviceInfo = C.CL_DEVICE_LOCAL_MEM_TYPE // Appended "Info" due to conflict with type.
	DeviceLocalMemSize               DeviceInfo = C.CL_DEVICE_LOCAL_MEM_SIZE
	DeviceErrorCorrectionSupport     DeviceInfo = C.CL_DEVICE_ERROR_CORRECTION_SUPPORT
	DeviceProfilingTimerResolution   DeviceInfo = C.CL_DEVICE_PROFILING_TIMER_RESOLUTION
	DeviceEndianLittle               DeviceInfo = C.CL_DEVICE_ENDIAN_LITTLE
	DeviceAvailable                  DeviceInfo = C.CL_DEVICE_AVAILABLE
	DeviceCompilerAvailable          DeviceInfo = C.CL_DEVICE_COMPILER_AVAILABLE
	DeviceExecutionCapabilities      DeviceInfo = C.CL_DEVICE_EXECUTION_CAPABILITIES
	DeviceQueueProperties            DeviceInfo = C.CL_DEVICE_QUEUE_PROPERTIES
	DeviceName                       DeviceInfo = C.CL_DEVICE_NAME
	DeviceVendor                     DeviceInfo = C.CL_DEVICE_VENDOR
	DriverVersion                    DeviceInfo = C.CL_DRIVER_VERSION
	DeviceProfile                    DeviceInfo = C.CL_DEVICE_PROFILE
	DeviceVersion                    DeviceInfo = C.CL_DEVICE_VERSION
	DeviceExtensions                 DeviceInfo = C.CL_DEVICE_EXTENSIONS
	DevicePlatform                   DeviceInfo = C.CL_DEVICE_PLATFORM
	DevicePreferredVectorWidthHalf   DeviceInfo = C.CL_DEVICE_PREFERRED_VECTOR_WIDTH_HALF
	DeviceHostUnifiedMemory          DeviceInfo = C.CL_DEVICE_HOST_UNIFIED_MEMORY
	DeviceNativeVectorWidthChar      DeviceInfo = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_CHAR
	DeviceNativeVectorWidthShort     DeviceInfo = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_SHORT
	DeviceNativeVectorWidthInt       DeviceInfo = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_INT
	DeviceNativeVectorWidthLong      DeviceInfo = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_LONG
	DeviceNativeVectorWidthFloat     DeviceInfo = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_FLOAT
	DeviceNativeVectorWidthDouble    DeviceInfo = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_DOUBLE
	DeviceNativeVectorWidthHalf      DeviceInfo = C.CL_DEVICE_NATIVE_VECTOR_WIDTH_HALF
	DeviceOpenclCVersion             DeviceInfo = C.CL_DEVICE_OPENCL_C_VERSION
)

// Bitfield.
const (
	FPDenorm         DeviceFPConfig = C.CL_FP_DENORM
	FPFma            DeviceFPConfig = C.CL_FP_FMA
	FPInfNan         DeviceFPConfig = C.CL_FP_INF_NAN
	FPRoundToInf     DeviceFPConfig = C.CL_FP_ROUND_TO_INF
	FPRoundToNearest DeviceFPConfig = C.CL_FP_ROUND_TO_NEAREST
	FPRoundToZero    DeviceFPConfig = C.CL_FP_ROUND_TO_ZERO
)

const (
	None           DeviceMemCacheType = C.CL_NONE
	ReadOnlyCache  DeviceMemCacheType = C.CL_READ_ONLY_CACHE
	ReadWriteCache DeviceMemCacheType = C.CL_READ_WRITE_CACHE
)

const (
	Global DeviceLocalMemType = C.CL_GLOBAL
	Local  DeviceLocalMemType = C.CL_LOCAL
)

// Bitfield
const (
	ExecKernel       DeviceExecCapabilities = C.CL_EXEC_KERNEL
	ExecNativeKernel DeviceExecCapabilities = C.CL_EXEC_NATIVE_KERNEL
)

func GetDeviceIDs(platform PlatformID, deviceType DeviceType, numEntries Uint, devices *DeviceID,
	numDevices *Uint) error {

	return NewError(C.clGetDeviceIDs(C.cl_platform_id(platform), C.cl_device_type(deviceType), C.cl_uint(numEntries),
		(*C.cl_device_id)(devices), (*C.cl_uint)(numDevices)))
}

func GetDeviceInfo(device DeviceID, paramName DeviceInfo, paramValueSize Size, paramValue unsafe.Pointer,
	paramValueSizeRet *Size) error {

	return NewError(C.clGetDeviceInfo(C.cl_device_id(device), C.cl_device_info(paramName), C.size_t(paramValueSize),
		paramValue, (*C.size_t)(paramValueSizeRet)))
}
