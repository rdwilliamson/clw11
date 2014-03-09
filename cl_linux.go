package clw11

/*
#cgo LDFLAGS: -lOpenCL

#include "CL/opencl.h"
*/
import "C"
import (
	"errors"
)

const (
	GLContext  ContextProperties = C.CL_GL_CONTEXT_KHR
	EGLDisplay ContextProperties = C.CL_EGL_DISPLAY_KHR
	GLXDisplay ContextProperties = C.CL_GLX_DISPLAY_KHR
)

var InvalidGLSharegroupReference = errors.New("invalid GL sharegroup reference")

func init() {
	errorMap[C.CL_INVALID_GL_SHAREGROUP_REFERENCE_KHR] = InvalidGLSharegroupReference
}