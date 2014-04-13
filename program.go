package clw11

/*
#define CL_USE_DEPRECATED_OPENCL_1_1_APIS
#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif

extern void programCallback(cl_program program, void *user_data);

void callProgramCallback(cl_program program, void *user_data)
{
	programCallback(program, user_data);
}
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

// Creates a program object for a context, and loads the source code specified
// by the text strings in the strings array into the program object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateProgramWithSource.html
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

// Creates a program object for a context, and loads the binary bits specified
// by binary into the program object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateProgramWithBinary.html
func CreateProgramWithBinary(context Context, devices []DeviceID, binaries [][]byte,
	binary_status []error) (Program, error) {

	num_devices := len(devices)
	lengths := make([]C.size_t, num_devices)
	cBinaries := make([]*C.uchar, num_devices)
	errors := make([]C.cl_int, num_devices)
	for i := range devices {
		lengths[i] = C.size_t(len(binaries[i]))
		cBinaries[i] = (*C.uchar)(&binaries[i][0])
	}

	var err C.cl_int
	program := C.clCreateProgramWithBinary(context, C.cl_uint(num_devices), (*C.cl_device_id)(&devices[0]),
		(*C.size_t)(&lengths[0]), (**C.uchar)(&cBinaries[0]), (*C.cl_int)(&errors[0]), &err)

	for i := range binary_status {
		binary_status[i] = toError(errors[i])
	}

	return Program(program), toError(err)
}

// Increments the program reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clRetainProgram.html
func RetainProgram(program Program) error {
	return toError(C.clRetainProgram(program))
}

// Decrements the program reference count.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clReleaseProgram.html
func ReleaseProgram(program Program) error {
	return toError(C.clReleaseProgram(program))
}

// Allows the implementation to release the resources allocated by the OpenCL
// compiler.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clUnloadCompiler.html
func UnloadCompiler() error {
	return toError(C.clUnloadCompiler())
}

// Builds (compiles and links) a program executable from the program source or
// binary.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clBuildProgram.html
func BuildProgram(program Program, devices []DeviceID, options string, callback ProgramCallbackFunc,
	userData interface{}) error {

	cOptions := C.CString(options)
	defer C.free(unsafe.Pointer(cOptions))

	key := programCallbacks.add(callback, userData)

	err := toError(C.clBuildProgram(program, C.cl_uint(len(devices)), (*C.cl_device_id)(&devices[0]), cOptions,
		(*[0]byte)(C.callProgramCallback), unsafe.Pointer(key)))

	if err != nil {
		// If the C side setting of the callback failed the get callback will
		// remove the callback from the map.
		programCallbacks.get(key)
	}

	return err
}

// Returns information about the program object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetProgramInfo.html
func GetProgramInfo(program Program, param_name ProgramInfo, param_value_size Size, param_value unsafe.Pointer,
	param_value_size_ret *Size) error {

	return toError(C.clGetProgramInfo(program, C.cl_program_info(param_name), C.size_t(param_value_size), param_value,
		(*C.size_t)(param_value_size_ret)))
}

// Returns build information for each device in the program object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetProgramBuildInfo.html
func GetProgramBuildInfo(program Program, device DeviceID, param_name ProgramBuildInfo, param_value_size Size,
	param_value unsafe.Pointer, param_value_size_ret *Size) error {

	return toError(C.clGetProgramBuildInfo(program, device, C.cl_program_build_info(param_name),
		C.size_t(param_value_size), param_value, (*C.size_t)(param_value_size_ret)))
}
