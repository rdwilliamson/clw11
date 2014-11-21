package clw11

/*
#cgo LDFLAGS: -lOpenCL

#define CL_USE_DEPRECATED_OPENCL_1_1_APIS
#include "CL/opencl.h"
*/
import "C"
import (
	"errors"
)

const (
	GLContext  ContextProperties = C.CL_GL_CONTEXT_KHR
	EGLDisplay ContextProperties = C.CL_EGL_DISPLAY_KHR
	WGLHDC     ContextProperties = C.CL_WGL_HDC_KHR
)

var InvalidGLSharegroupReference = errors.New("cl: invalid GL sharegroup reference")

func init() {
	errorMap[C.CL_INVALID_GL_SHAREGROUP_REFERENCE_KHR] = InvalidGLSharegroupReference
}
