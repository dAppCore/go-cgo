package cgo

/*
#include <stdint.h>
#include <stdlib.h>

typedef int (*cgo_call_int_fn0_t)(void);
typedef int (*cgo_call_int_fn1_t)(uintptr_t);
typedef int (*cgo_call_int_fn2_t)(uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn3_t)(uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn4_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn5_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn6_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn7_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn8_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn9_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn10_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn11_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn12_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn13_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn14_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn15_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn16_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn17_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
typedef int (*cgo_call_int_fn18_t)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);

int cgo_call_0(uintptr_t fn) {
	return ((cgo_call_int_fn0_t)fn)();
}

int cgo_call_1(uintptr_t fn, uintptr_t a0) {
	return ((cgo_call_int_fn1_t)fn)(a0);
}

int cgo_call_2(uintptr_t fn, uintptr_t a0, uintptr_t a1) {
	return ((cgo_call_int_fn2_t)fn)(a0, a1);
}

int cgo_call_3(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2) {
	return ((cgo_call_int_fn3_t)fn)(a0, a1, a2);
}

int cgo_call_4(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3) {
	return ((cgo_call_int_fn4_t)fn)(a0, a1, a2, a3);
}

int cgo_call_5(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4) {
	return ((cgo_call_int_fn5_t)fn)(a0, a1, a2, a3, a4);
}

int cgo_call_6(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5) {
	return ((cgo_call_int_fn6_t)fn)(a0, a1, a2, a3, a4, a5);
}

int cgo_call_7(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6) {
	return ((cgo_call_int_fn7_t)fn)(a0, a1, a2, a3, a4, a5, a6);
}

int cgo_call_8(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7) {
	return ((cgo_call_int_fn8_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7);
}

int cgo_call_9(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8) {
	return ((cgo_call_int_fn9_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

int cgo_call_10(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9) {
	return ((cgo_call_int_fn10_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

int cgo_call_11(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10) {
	return ((cgo_call_int_fn11_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

int cgo_call_12(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11) {
	return ((cgo_call_int_fn12_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

int cgo_call_13(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12) {
	return ((cgo_call_int_fn13_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

int cgo_call_14(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13) {
	return ((cgo_call_int_fn14_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

int cgo_call_15(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14) {
	return ((cgo_call_int_fn15_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14);
}

int cgo_call_16(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15) {
	return ((cgo_call_int_fn16_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15);
}

int cgo_call_17(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15, uintptr_t a16) {
	return ((cgo_call_int_fn17_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16);
}

int cgo_call_18(uintptr_t fn, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15, uintptr_t a16, uintptr_t a17) {
	return ((cgo_call_int_fn18_t)fn)(a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17);
}

*/
import "C"

import (
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

var freedCStringPointers sync.Map

type cStringFreeState struct {
	freed atomic.Bool
}

// SizeT converts a Go int into a C size_t for C APIs.
//
//	bufferSize := SizeT(len(payload))
func SizeT(value int) C.size_t {
	if value < 0 {
		panic("cgo.SizeT: negative values are not representable as C.size_t")
	}

	if value > 0 {
		sizeBits := int(unsafe.Sizeof(C.size_t(0)) * 8)
		if sizeBits < strconv.IntSize {
			maxSize := (uint64(1) << sizeBits) - 1
			if uint64(value) > maxSize {
				panic("cgo.SizeT: value exceeds C.size_t range")
			}
		}
	}
	return C.size_t(value)
}

// Int converts a Go int into a C int for C APIs.
//
//	remaining := Int(2)
func Int(value int) C.int {
	cIntBits := cIntBitSize()
	if cIntBits < strconv.IntSize {
		maxValue := (int64(1) << (cIntBits - 1)) - 1
		minValue := -maxValue - 1
		casted := int64(value)
		if casted < minValue || casted > maxValue {
			panic("cgo.Int: value exceeds C.int range")
		}
	}
	return C.int(value)
}

func cIntBitSize() int {
	return int(unsafe.Sizeof(C.int(0)) * 8)
}

// Call invokes a C function pointer and maps a non-zero return code to an error.
//
// Supported argument types are unsafe.Pointer, *Buffer, *C.char, []byte, int,
// int32, int64, uint, uint32, uint64, C.size_t, C.int, and C.uintptr_t.
//
//	err := Call(unsafe.Pointer(C.some_function), buffer.Ptr(), SizeT(len(payload)))
//	if err != nil { return err }
func Call(function unsafe.Pointer, args ...interface{}) error {
	if function == nil {
		panic("cgo.Call: function pointer is nil")
	}

	var result C.int
	target := uintptr(function)
	unsupportedArgument := func(index int) {
		panic("cgo.Call: unsupported argument type at argument " + strconv.Itoa(index))
	}
	toCallArg := func(index int) C.uintptr_t {
		converted, ok := toCallArgValue(args[index])
		if !ok {
			unsupportedArgument(index + 1)
		}
		return C.uintptr_t(converted)
	}

	switch len(args) {
	case 0:
		result = C.cgo_call_0(C.uintptr_t(target))
	case 1:
		result = C.cgo_call_1(C.uintptr_t(target), toCallArg(0))
	case 2:
		result = C.cgo_call_2(C.uintptr_t(target), toCallArg(0), toCallArg(1))
	case 3:
		result = C.cgo_call_3(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2))
	case 4:
		result = C.cgo_call_4(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3))
	case 5:
		result = C.cgo_call_5(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4))
	case 6:
		result = C.cgo_call_6(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5))
	case 7:
		result = C.cgo_call_7(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6))
	case 8:
		result = C.cgo_call_8(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7))
	case 9:
		result = C.cgo_call_9(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7), toCallArg(8))
	case 10:
		result = C.cgo_call_10(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7), toCallArg(8), toCallArg(9))
	case 11:
		result = C.cgo_call_11(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7), toCallArg(8), toCallArg(9), toCallArg(10))
	case 12:
		result = C.cgo_call_12(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7), toCallArg(8), toCallArg(9), toCallArg(10), toCallArg(11))
	case 13:
		result = C.cgo_call_13(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7), toCallArg(8), toCallArg(9), toCallArg(10), toCallArg(11), toCallArg(12))
	case 14:
		result = C.cgo_call_14(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7), toCallArg(8), toCallArg(9), toCallArg(10), toCallArg(11), toCallArg(12), toCallArg(13))
	case 15:
		result = C.cgo_call_15(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7), toCallArg(8), toCallArg(9), toCallArg(10), toCallArg(11), toCallArg(12), toCallArg(13), toCallArg(14))
	case 16:
		result = C.cgo_call_16(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7), toCallArg(8), toCallArg(9), toCallArg(10), toCallArg(11), toCallArg(12), toCallArg(13), toCallArg(14), toCallArg(15))
	case 17:
		result = C.cgo_call_17(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7), toCallArg(8), toCallArg(9), toCallArg(10), toCallArg(11), toCallArg(12), toCallArg(13), toCallArg(14), toCallArg(15), toCallArg(16))
	case 18:
		result = C.cgo_call_18(C.uintptr_t(target), toCallArg(0), toCallArg(1), toCallArg(2), toCallArg(3), toCallArg(4), toCallArg(5), toCallArg(6), toCallArg(7), toCallArg(8), toCallArg(9), toCallArg(10), toCallArg(11), toCallArg(12), toCallArg(13), toCallArg(14), toCallArg(15), toCallArg(16), toCallArg(17))
	default:
		panic("cgo.Call: unsupported argument count: max 18")
	}

	runtime.KeepAlive(args)

	if result != 0 {
		return Errno(result)
	}
	return nil
}

func toCallArgValue(value interface{}) (uintptr, bool) {
	switch typed := value.(type) {
	case nil:
		return 0, true
	case unsafe.Pointer:
		return uintptr(typed), true
	case *Buffer:
		if typed == nil {
			return 0, true
		}
		return uintptr(typed.Ptr()), true
	case *C.char:
		return uintptr(unsafe.Pointer(typed)), true
	case int:
		return uintptr(typed), true
	case int32:
		return uintptr(typed), true
	case int64:
		return uintptr(typed), true
	case uint:
		return uintptr(typed), true
	case uint32:
		return uintptr(typed), true
	case uint64:
		return uintptr(typed), true
	case uintptr:
		return typed, true
	case C.size_t:
		return uintptr(typed), true
	case C.int:
		return uintptr(typed), true
	case []byte:
		if len(typed) == 0 {
			return 0, true
		}
		return uintptr(unsafe.Pointer(&typed[0])), true
	}
	return 0, false
}

// GoString converts a null-terminated C string to a Go string.
//
//	native := GoString(CString("hello"))
func GoString(cs *C.char) string {
	if cs == nil {
		return ""
	}
	return C.GoString(cs)
}

// CString converts a Go string to a C string.
//
//	cString := CString("hello")
//	defer Free(unsafe.Pointer(cString))
func CString(value string) *C.char {
	size := len(value) + 1
	memory := C.malloc(C.size_t(size))
	if memory == nil {
		panic("cgo.CString: C allocation failed")
	}

	bytes := unsafe.Slice((*byte)(memory), size)
	copy(bytes, value)
	bytes[len(value)] = 0

	freedCStringPointers.Delete(memory)
	return (*C.char)(memory)
}

// Free releases memory previously allocated by this package.
// It is safe to call more than once on the same pointer.
//
//	cString := CString("hello")
//	Free(unsafe.Pointer(cString))
func Free(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}

	stateValue, _ := freedCStringPointers.LoadOrStore(ptr, &cStringFreeState{})
	state := stateValue.(*cStringFreeState)
	if !state.freed.CompareAndSwap(false, true) {
		return
	}

	C.free(ptr)
}

// Errno converts a C errno-like return value to a Go error.
//
//	err := Errno(C.int(13))
func Errno(resultCode C.int) error {
	if resultCode == 0 {
		return nil
	}

	return syscall.Errno(resultCode)
}

// WithErrno runs a C-style function and converts the C return value to (result, error).
//
//	result, err := WithErrno(func() C.int {
//		return C.my_function()
//	})
func WithErrno(fn func() C.int) (int, error) {
	result := fn()
	if err := Errno(result); err != nil {
		return int(result), err
	}
	return int(result), nil
}
