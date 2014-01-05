package clw11

/*
#cgo LDFLAGS: -framework OpenCL

#include "OpenCL/opencl.h"
#include "OpenCL/cl_gl.h"
*/
import "C"
import (
	"errors"
)

const (
	GLShareGroupApple ContextProperties = C.CL_CONTEXT_PROPERTY_USE_CGL_SHAREGROUP_APPLE
)

var InvalidGLContextApple = errors.New("invalid GL context apple")

func init() {
	errorMap[C.CL_INVALID_GL_CONTEXT_APPLE] = InvalidGLContextApple
}
