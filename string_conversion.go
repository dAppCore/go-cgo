package cgo

/*
#include <stdlib.h>
*/
import "C"

import "unsafe"

// GoString converts a null-terminated C string to a Go string.
//
//	cStr := C.CString("example")
//	result := cgo.GoString(cStr)
//	cgo.Free(unsafe.Pointer(cStr))
func GoString(cs *C.char) string {
	if cs == nil {
		return ""
	}
	return C.GoString(cs)
}

// CString converts a Go string to a C string.
//
//	cStr := cgo.CString("hello")
//	defer cgo.Free(unsafe.Pointer(cStr))
func CString(value string) *C.char {
	return C.CString(value)
}

// Free releases memory previously returned by CString.
//
//	cStr := cgo.CString("hello")
//	cgo.Free(unsafe.Pointer(cStr))
func Free(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}
	C.free(ptr)
}
