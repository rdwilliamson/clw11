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
	Program C.cl_program
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
	return nil
}
