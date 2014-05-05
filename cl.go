// A simple wrapper around the OpenCL 1.1 C API.
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
	Bool   C.cl_bool
	Char   C.cl_char
	Uchar  C.cl_uchar
	Short  C.cl_short
	Ushort C.cl_ushort
	Int    C.cl_int
	Uint   C.cl_uint
	Long   C.cl_long
	Ulong  C.cl_ulong

	Half   C.cl_half
	Float  C.cl_float
	Double C.cl_double

	Size C.size_t

	Char2  C.cl_char2
	Char4  C.cl_char4
	Char8  C.cl_char8
	Char16 C.cl_char16

	Uchar2  C.cl_uchar2
	Uchar4  C.cl_uchar4
	Uchar8  C.cl_uchar8
	Uchar16 C.cl_uchar16

	Short2  C.cl_short2
	Short4  C.cl_short4
	Short8  C.cl_short8
	Short16 C.cl_short16

	Ushort2  C.cl_ushort2
	Ushort4  C.cl_ushort4
	Ushort8  C.cl_ushort8
	Ushort16 C.cl_ushort16

	Int2  C.cl_int2
	Int4  C.cl_int4
	Int8  C.cl_int8
	Int16 C.cl_int16

	Uint2  C.cl_uint2
	Uint4  C.cl_uint4
	Uint8  C.cl_uint8
	Uint16 C.cl_uint16

	Long2  C.cl_long2
	Long4  C.cl_long4
	Long8  C.cl_long8
	Long16 C.cl_long16

	Ulong2  C.cl_ulong2
	Ulong4  C.cl_ulong4
	Ulong8  C.cl_ulong8
	Ulong16 C.cl_ulong16

	Float2  C.cl_float2
	Float4  C.cl_float4
	Float8  C.cl_float8
	Float16 C.cl_float16

	Double2  C.cl_double2
	Double4  C.cl_double4
	Double8  C.cl_double8
	Double16 C.cl_double16
)

const (
	True  = Bool(C.CL_TRUE)
	False = Bool(C.CL_FALSE)
)
