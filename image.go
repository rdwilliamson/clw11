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
import "unsafe"

type (
	MemObjectType C.cl_mem_object_type
	ImageFormat   C.cl_image_format
	ChannelOrder  C.cl_channel_order
	ChannelType   C.cl_channel_type
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

func CreateImageFormat(co ChannelOrder, ct ChannelType) ImageFormat {
	return ImageFormat{C.cl_channel_order(co), C.cl_channel_type(ct)}
}

func (f *ImageFormat) Order() ChannelOrder {
	return ChannelOrder(f.image_channel_order)
}

func (f *ImageFormat) Type() ChannelType {
	return ChannelType(f.image_channel_data_type)
}

func GetSupportedImageFormats(context Context, flags MemFlags, image_type MemObjectType, num_entries Uint,
	image_formats *ImageFormat, num_image_formats *Uint) error {

	return toError(C.clGetSupportedImageFormats(context, C.cl_mem_flags(flags), C.cl_mem_object_type(image_type),
		C.cl_uint(num_entries), (*C.cl_image_format)(image_formats), (*C.cl_uint)(num_image_formats)))
}

func CreateImage2D(context Context, flags MemFlags, image_format ImageFormat, image_width, image_height,
	image_row_pitch Size, host_ptr unsafe.Pointer) (Mem, error) {

	var err C.cl_int
	mem := C.clCreateImage2D(context, C.cl_mem_flags(flags), (*C.cl_image_format)(&image_format), C.size_t(image_width),
		C.size_t(image_height), C.size_t(image_row_pitch), host_ptr, &err)

	return Mem(mem), toError(err)
}
