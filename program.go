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
	Program          C.cl_program
	ProgramInfo      C.cl_program_info
	ProgramBuildInfo C.cl_program_build_info
)

const (
	ProgramReferenceCount = ProgramInfo(C.CL_PROGRAM_REFERENCE_COUNT)
	ProgramContext        = ProgramInfo(C.CL_PROGRAM_CONTEXT)
	ProgramNumDevices     = ProgramInfo(C.CL_PROGRAM_NUM_DEVICES)
	ProgramDevices        = ProgramInfo(C.CL_PROGRAM_DEVICES)
	ProgramSource         = ProgramInfo(C.CL_PROGRAM_SOURCE)
	ProgramBinarySizes    = ProgramInfo(C.CL_PROGRAM_BINARY_SIZES)
	ProgramBinaries       = ProgramInfo(C.CL_PROGRAM_BINARIES)
)

const (
	ProgramBuildStatus  = ProgramBuildInfo(C.CL_PROGRAM_BUILD_STATUS)
	ProgramBuildOptions = ProgramBuildInfo(C.CL_PROGRAM_BUILD_OPTIONS)
	ProgramBuildLog     = ProgramBuildInfo(C.CL_PROGRAM_BUILD_LOG)
)

func CreateProgramWithSource(context Context, sources [][]byte) (Program, error) {

	count := len(sources)
	strings := make([]unsafe.Pointer, count)
	lengths := make([]C.size_t, count)
	for i := range sources {
		strings[i] = unsafe.Pointer(&sources[i][0])
		lengths[i] = C.size_t(len(sources[i]))
	}

	var err C.cl_int
	program := C.clCreateProgramWithSource(context, C.cl_uint(count), (**C.char)(unsafe.Pointer(&strings[0])),
		&lengths[0], &err)

	return Program(program), toError(err)
}

func BuildProgram(program Program, devices []DeviceID, options string, callback func()) error {
	// TODO handle callbacks

	cOptions := C.CString(options)
	defer C.free(unsafe.Pointer(cOptions))

	err := toError(C.clBuildProgram(program, C.cl_uint(len(devices)), (*C.cl_device_id)(&devices[0]), cOptions,
		nil, nil))

	return err
}

func GetProgramInfo(program Program, param_name ProgramInfo, param_value_size Size, param_value unsafe.Pointer,
	param_value_size_ret *Size) error {

	return toError(C.clGetProgramInfo(program, C.cl_program_info(param_name), C.size_t(param_value_size), param_value,
		(*C.size_t)(param_value_size_ret)))
}

func GetProgramBuildInfo(program Program, device DeviceID, param_name ProgramBuildInfo, param_value_size Size,
	param_value unsafe.Pointer, param_value_size_ret *Size) error {

	return toError(C.clGetProgramBuildInfo(program, device, C.cl_program_build_info(param_name),
		C.size_t(param_value_size), param_value, (*C.size_t)(param_value_size_ret)))
}
