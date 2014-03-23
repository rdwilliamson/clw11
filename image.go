package clw11

/*
#define CL_USE_DEPRECATED_OPENCL_1_1_APIS
#ifdef __APPLE__
#include "OpenCL/opencl.h"
#else
#include "CL/opencl.h"
#endif
*/
import "C"

type (
	ChannelOrder C.cl_channel_order
	ChannelType  C.cl_channel_type
)

const (
	R         ChannelOrder = C.CL_R
	A         ChannelOrder = C.CL_A
	RG        ChannelOrder = C.CL_RG
	RA        ChannelOrder = C.CL_RA
	RGB       ChannelOrder = C.CL_RGB
	RGBA      ChannelOrder = C.CL_RGBA
	BGRA      ChannelOrder = C.CL_BGRA
	ARGB      ChannelOrder = C.CL_ARGB
	Intensity ChannelOrder = C.CL_INTENSITY
	Luminance ChannelOrder = C.CL_LUMINANCE
	Rx        ChannelOrder = C.CL_Rx
	RGx       ChannelOrder = C.CL_RGx
	RGBx      ChannelOrder = C.CL_RGBx
)

const (
	SnormInt8      ChannelType = C.CL_SNORM_INT8
	SnormInt16     ChannelType = C.CL_SNORM_INT16
	UnormInt8      ChannelType = C.CL_UNORM_INT8
	UnormInt16     ChannelType = C.CL_UNORM_INT16
	UnormShort565  ChannelType = C.CL_UNORM_SHORT_565
	UnormShort555  ChannelType = C.CL_UNORM_SHORT_555
	UnormInt101010 ChannelType = C.CL_UNORM_INT_101010
	SignedInt8     ChannelType = C.CL_SIGNED_INT8
	SignedInt16    ChannelType = C.CL_SIGNED_INT16
	SignedInt32    ChannelType = C.CL_SIGNED_INT32
	UnsignedInt8   ChannelType = C.CL_UNSIGNED_INT8
	UnsignedInt16  ChannelType = C.CL_UNSIGNED_INT16
	UnsignedInt32  ChannelType = C.CL_UNSIGNED_INT32
	HalfFloat      ChannelType = C.CL_HALF_FLOAT
	Float          ChannelType = C.CL_FLOAT
)
