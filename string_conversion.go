package cgo

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"math"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

type cStringAllocation struct {
	base  unsafe.Pointer
	freed atomic.Bool
}

var cStringAllocations sync.Map
var freedPointers sync.Map

// SizeT converts a Go int into a C size_t for C APIs.
//
//	bufferSize := SizeT(len(payload))
func SizeT(value int) C.size_t {
	if value < 0 {
		panic("cgo.SizeT: negative values are not representable as C.size_t")
	}

	var max uint64
	switch unsafe.Sizeof(C.size_t(0)) {
	case 4:
		max = math.MaxUint32
	case 8:
		max = math.MaxUint64
	default:
		panic("cgo.SizeT: unsupported C.size_t size")
	}

	if uint64(value) > max {
		panic("cgo.SizeT: value exceeds C.size_t range")
	}

	return C.size_t(value)
}

// Int converts a Go int into a C int for C APIs.
//
//	remaining := Int(2)
func Int(value int) C.int {
	if value < 0 {
		panic("cgo.Int: value exceeds C.int range")
	}

	var max int64
	switch unsafe.Sizeof(C.int(0)) {
	case 4:
		max = math.MaxInt32
	case 8:
		max = math.MaxInt64
	default:
		panic("cgo.Int: unsupported C.int size")
	}

	if int64(value) > max {
		panic("cgo.Int: value exceeds C.int range")
	}

	return C.int(value)
}

// Errno converts a C errno value to a Go error.
//
//	rc := C.some_function()
//	err := Errno(rc)
func Errno(rc C.int) error {
	if rc == 0 {
		return nil
	}
	return syscall.Errno(rc)
}

// WithErrno calls a C function and returns the errno as a Go error.
//
//	result, err := WithErrno(func() C.int {
//	    return C.my_function(args...)
//	})
func WithErrno(fn func() C.int) (int, error) {
	rc := fn()
	return int(rc), Errno(rc)
}

// GoString converts a C string to Go string safely.
//
//	goStr := GoString(cStr)
func GoString(cs *C.char) string {
	if cs == nil {
		return ""
	}
	return C.GoString(cs)
}

// CString converts a Go string to C string. Caller must free via cgo.Free().
//
//	cStr := CString(goStr)
//	defer cgo.Free(unsafe.Pointer(cStr))
func CString(s string) *C.char {
	size := len(s) + 1
	raw := C.malloc(C.size_t(size))
	if raw == nil {
		panic("cgo.CString: C allocation failed")
	}

	buf := unsafe.Slice((*byte)(raw), size)
	copy(buf, s)
	buf[len(s)] = 0

	cString := (*C.char)(raw)
	addr := uintptr(unsafe.Pointer(cString))
	cStringAllocations.Store(addr, &cStringAllocation{
		base: raw,
	})
	freedPointers.Delete(addr)
	return cString
}

// Free releases C-allocated memory. Idempotent.
//
//	cgo.Free(ptr)
func Free(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}

	addr := uintptr(ptr)
	if alloc, ok := cStringAllocations.Load(addr); ok {
		allocation := alloc.(*cStringAllocation)
		if allocation.freed.CompareAndSwap(false, true) {
			C.free(allocation.base)
			cStringAllocations.Delete(addr)
			freedPointers.Store(addr, struct{}{})
		}
		return
	}

	if _, loaded := freedPointers.LoadOrStore(addr, struct{}{}); loaded {
		return
	}

	C.free(ptr)
	freedPointers.Store(addr, struct{}{})
}
