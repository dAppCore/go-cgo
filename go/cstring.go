package cgo

/*
#include <stdlib.h>
*/
import "C"

import "unsafe"

// AdoptCString copies a C-allocated null-terminated string into a Go
// string and frees the C side. Use for C APIs that return malloc'd
// strings the caller must release (mlx error messages, sqlite error
// rows, libcurl headers, etc.) — the common shape is:
//
//	msg := C.get_last_error()         // malloc'd char*
//	goMsg := cgo.AdoptCString(msg)    // copy + free in one call
//
// Replaces the two-step:
//
//	goMsg := C.GoString(msg)
//	C.free(unsafe.Pointer(msg))
//
// with a single intent-named call. Returns "" for nil input. The
// caller must not reference the C string after Adopt — the underlying
// memory is freed before the function returns.
func AdoptCString(cStr *C.char) string {
	if cStr == nil {
		return ""
	}
	s := C.GoString(cStr)
	C.free(unsafe.Pointer(cStr))
	return s
}

// AdoptCStringN copies n bytes from a C-allocated buffer into a Go
// string and frees the C side. Use when the C buffer is not null-
// terminated (mlx_string_data when a length is known up front,
// fixed-width C struct fields written as C strings).
//
//	buf := C.mlx_string_data(str)         // not null-terminated
//	n   := C.size_t(C.mlx_string_size(str))
//	goStr := cgo.AdoptCStringN(buf, int(n))
//
// Returns "" for nil input or non-positive n. Frees buf regardless
// of length — the C side relinquishes ownership.
func AdoptCStringN(cStr *C.char, n int) string {
	if cStr == nil {
		return ""
	}
	if n <= 0 {
		C.free(unsafe.Pointer(cStr))
		return ""
	}
	s := C.GoStringN(cStr, C.int(n))
	C.free(unsafe.Pointer(cStr))
	return s
}
