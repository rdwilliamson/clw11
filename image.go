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
	ImageInfo     C.cl_image_info
	ChannelOrder  C.cl_channel_order
	ChannelType   C.cl_channel_type
)

const (
	MemObjectBuffer  MemObjectType = C.CL_MEM_OBJECT_BUFFER
	MemObjectImage2D MemObjectType = C.CL_MEM_OBJECT_IMAGE2D
	MemObjectImage3D MemObjectType = C.CL_MEM_OBJECT_IMAGE3D
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
	Float32        ChannelType = C.CL_FLOAT // Appended "32" due to conflict with type.
)

const (
	ImageFormatInfo  ImageInfo = C.CL_IMAGE_FORMAT // Appended "Info" due to conflict with type.
	ImageElementSize ImageInfo = C.CL_IMAGE_ELEMENT_SIZE
	ImageRowPitch    ImageInfo = C.CL_IMAGE_ROW_PITCH
	ImageSlicePitch  ImageInfo = C.CL_IMAGE_SLICE_PITCH
	ImageWidth       ImageInfo = C.CL_IMAGE_WIDTH
	ImageHeight      ImageInfo = C.CL_IMAGE_HEIGHT
	ImageDepth       ImageInfo = C.CL_IMAGE_DEPTH
)

// Create an image format descriptor.
func CreateImageFormat(co ChannelOrder, ct ChannelType) ImageFormat {
	return ImageFormat{C.cl_channel_order(co), C.cl_channel_type(ct)}
}

// Retruns the number of channels and the channel layout.
func (f *ImageFormat) ChannelOrder() ChannelOrder {
	return ChannelOrder(f.image_channel_order)
}

// Returns the size of the channel data type.
func (f *ImageFormat) ChannelType() ChannelType {
	return ChannelType(f.image_channel_data_type)
}

// Creates a 2D image object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateImage2D.html
func CreateImage2D(context Context, flags MemFlags, image_format ImageFormat, image_width, image_height,
	image_row_pitch Size, host_ptr unsafe.Pointer) (Mem, error) {

	var err C.cl_int
	mem := C.clCreateImage2D(context, C.cl_mem_flags(flags), (*C.cl_image_format)(&image_format), C.size_t(image_width),
		C.size_t(image_height), C.size_t(image_row_pitch), host_ptr, &err)

	return Mem(mem), toError(err)
}

// Creates a 3D image object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clCreateImage3D.html
func CreateImage3D(context Context, flags MemFlags, image_format ImageFormat, image_width, image_height, image_depth,
	image_row_pitch, image_slice_pitch Size, host_ptr unsafe.Pointer) (Mem, error) {

	var err C.cl_int
	mem := C.clCreateImage3D(context, C.cl_mem_flags(flags), (*C.cl_image_format)(&image_format), C.size_t(image_width),
		C.size_t(image_height), C.size_t(image_depth), C.size_t(image_row_pitch), C.size_t(image_slice_pitch), host_ptr,
		&err)

	return Mem(mem), toError(err)
}

// Get the list of image formats supported by an OpenCL implementation.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetSupportedImageFormats.html
func GetSupportedImageFormats(context Context, flags MemFlags, image_type MemObjectType, num_entries Uint,
	image_formats *ImageFormat, num_image_formats *Uint) error {

	return toError(C.clGetSupportedImageFormats(context, C.cl_mem_flags(flags), C.cl_mem_object_type(image_type),
		C.cl_uint(num_entries), (*C.cl_image_format)(image_formats), (*C.cl_uint)(num_image_formats)))
}

// Enqueues a command to read from a 2D or 3D image object to host memory.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueMapImage.html
func EnqueueReadImage(command_queue CommandQueue, image Mem, blocking_read Bool, origin, region [3]Size, row_pitch,
	slice_pitch Size, ptr unsafe.Pointer, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueReadImage(command_queue, image, C.cl_bool(blocking_read), (*C.size_t)(&origin[0]),
		(*C.size_t)(&region[0]), C.size_t(row_pitch), C.size_t(slice_pitch), ptr, num_events_in_wait_list,
		event_wait_list, (*C.cl_event)(event)))
}

// Enqueues a command to write to a 2D or 3D image object from host memory.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueMapImage.html
func EnqueueWriteImage(command_queue CommandQueue, image Mem, blocking_read Bool, origin, region [3]Size, row_pitch,
	slice_pitch Size, ptr unsafe.Pointer, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueWriteImage(command_queue, image, C.cl_bool(blocking_read), (*C.size_t)(&origin[0]),
		(*C.size_t)(&region[0]), C.size_t(row_pitch), C.size_t(slice_pitch), ptr, num_events_in_wait_list,
		event_wait_list, (*C.cl_event)(event)))
}

// Enqueues a command to copy image objects.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueCopyImage.html
func EnqueueCopyImage(command_queue CommandQueue, src_image, dst_image Mem, src_origin, dst_origin, region [3]Size,
	wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueCopyImage(command_queue, src_image, dst_image, (*C.size_t)(&src_origin[0]),
		(*C.size_t)(&dst_origin[0]), (*C.size_t)(&region[0]), num_events_in_wait_list, event_wait_list,
		(*C.cl_event)(event)))
}

// Enqueues a command to copy an image object to a buffer object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueCopyImageToBuffer.html
func EnqueueCopyImageToBuffer(command_queue CommandQueue, src_image, dst_buffer Mem, src_origin, region [3]Size,
	dst_offset Size, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueCopyImageToBuffer(command_queue, src_image, dst_buffer, (*C.size_t)(&src_origin[0]),
		(*C.size_t)(&region[0]), C.size_t(dst_offset), num_events_in_wait_list, event_wait_list,
		(*C.cl_event)(event)))
}

// Enqueues a command to copy a buffer object to an image object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueCopyBufferToImage.html
func EnqueueCopyBufferToImage(command_queue CommandQueue, src_buffer, dst_image Mem, src_offset Size, dst_origin,
	region [3]Size, wait_list []Event, event *Event) error {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	return toError(C.clEnqueueCopyBufferToImage(command_queue, src_buffer, dst_image, C.size_t(src_offset),
		(*C.size_t)(&dst_origin[0]), (*C.size_t)(&region[0]), num_events_in_wait_list, event_wait_list,
		(*C.cl_event)(event)))
}

// Enqueues a command to map a region of an image object into the host address
// space and returns a pointer to this mapped region.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clEnqueueMapImage.html
func EnqueueMapImage(command_queue CommandQueue, image Mem, blocking_map Bool, map_flags MapFlags, origin,
	region [3]Size, image_row_pitch, image_slice_pitch *Size, wait_list []Event, event *Event) (unsafe.Pointer, error) {

	event_wait_list, num_events_in_wait_list := toEventList(wait_list)

	var err C.cl_int
	mapped := C.clEnqueueMapImage(command_queue, image, C.cl_bool(blocking_map), C.cl_map_flags(map_flags),
		(*C.size_t)(&origin[0]), (*C.size_t)(&region[0]), (*C.size_t)(image_row_pitch), (*C.size_t)(image_slice_pitch),
		num_events_in_wait_list, event_wait_list, (*C.cl_event)(event), &err)

	return mapped, toError(err)
}

// Get information specific to an image object.
// http://www.khronos.org/registry/cl/sdk/1.1/docs/man/xhtml/clGetImageInfo.html
func GetImageInfo(image Mem, param_name ImageInfo, param_value_size Size, param_value unsafe.Pointer,
	param_value_size_ret *Size) error {

	return toError(C.clGetImageInfo(image, C.cl_image_info(param_name), C.size_t(param_value_size), param_value,
		(*C.size_t)(param_value_size_ret)))
}
