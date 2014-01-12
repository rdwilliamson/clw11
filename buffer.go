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
	Memory      C.cl_mem
	MemoryFlags C.cl_mem_flags
)

const (
	MemoryReadWrite        MemoryFlags = C.CL_MEM_READ_WRITE
	MemoryWriteOnly        MemoryFlags = C.CL_MEM_WRITE_ONLY
	MemoryReadOnly         MemoryFlags = C.CL_MEM_READ_ONLY
	MemoryUseHostPointer   MemoryFlags = C.CL_MEM_USE_HOST_PTR
	MemoryAllocHostPointer MemoryFlags = C.CL_MEM_ALLOC_HOST_PTR
	MemoryCopyHostPointer  MemoryFlags = C.CL_MEM_COPY_HOST_PTR
)

func CreateBuffer(context Context, flags MemoryFlags, size Size, host_ptr []byte) (Memory, error) {
	var host unsafe.Pointer
	if host_ptr != nil {
		host = unsafe.Pointer(&host_ptr[0])
	}
	var err C.cl_int
	memory := C.clCreateBuffer(context, C.cl_mem_flags(flags), C.size_t(size), host, &err)
	return Memory(memory), NewError(err)
}

func ReleaseMemObject(memobject Memory) error {
	return NewError(C.clReleaseMemObject(memobject))
}
