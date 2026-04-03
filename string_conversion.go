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
*/
import "C"

import (
	"fmt"
	"strconv"
	"syscall"
	"unsafe"
)

// SizeT converts a Go int into a C size_t for cgo calls.
//
//	size := cgo.SizeT(len(data))
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

// Int converts a Go int into a C int for cgo calls.
//
//	rc := cgo.Int(returnCode)
func Int(value int) C.int {
	if value < -2147483648 || value > 2147483647 {
		panic("cgo.Int: value exceeds C.int range")
	}
	return C.int(value)
}

// Call invokes a C function pointer and maps a non-zero return into a Go error.
//
//	err := cgo.Call(unsafe.Pointer(C.some_function), cgo.SizeT(len(data)))
//	err == nil indicates success.
func Call(function unsafe.Pointer, args ...interface{}) error {
	if function == nil {
		panic("cgo.Call: function pointer is nil")
	}

	var result uintptr
	target := uintptr(function)

	switch len(args) {
	case 0:
		result = uintptr(C.cgo_call_0(C.uintptr_t(target)))
	case 1:
		a0, ok := toSyscallArg(args[0])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = uintptr(C.cgo_call_1(C.uintptr_t(target), C.uintptr_t(a0)))
	case 2:
		a0, ok := toSyscallArg(args[0])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a1, ok := toSyscallArg(args[1])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = uintptr(C.cgo_call_2(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1)))
	case 3:
		a0, ok := toSyscallArg(args[0])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a1, ok := toSyscallArg(args[1])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a2, ok := toSyscallArg(args[2])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = uintptr(C.cgo_call_3(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2)))
	case 4:
		a0, ok := toSyscallArg(args[0])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a1, ok := toSyscallArg(args[1])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a2, ok := toSyscallArg(args[2])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a3, ok := toSyscallArg(args[3])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = uintptr(C.cgo_call_4(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3)))
	case 5:
		a0, ok := toSyscallArg(args[0])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a1, ok := toSyscallArg(args[1])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a2, ok := toSyscallArg(args[2])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a3, ok := toSyscallArg(args[3])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a4, ok := toSyscallArg(args[4])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = uintptr(C.cgo_call_5(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4)))
	case 6:
		a0, ok := toSyscallArg(args[0])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a1, ok := toSyscallArg(args[1])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a2, ok := toSyscallArg(args[2])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a3, ok := toSyscallArg(args[3])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a4, ok := toSyscallArg(args[4])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a5, ok := toSyscallArg(args[5])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = uintptr(C.cgo_call_6(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5)))
	default:
		panic("cgo.Call: unsupported argument count: max 6")
	}

	if result != 0 {
		return fmt.Errorf("cgo.Call: return code %d", int(result))
	}
	return nil
}

func toSyscallArg(value interface{}) (uintptr, bool) {
	switch typed := value.(type) {
	case nil:
		return 0, true
	case uintptr:
		return typed, true
	case unsafe.Pointer:
		return uintptr(typed), true
	case C.char:
		return uintptr(typed), true
	case C.int:
		return uintptr(typed), true
	case C.size_t:
		return uintptr(typed), true
	case *C.char:
		return uintptr(unsafe.Pointer(typed)), true
	case *byte:
		return uintptr(unsafe.Pointer(typed)), true
	case bool:
		if typed {
			return 1, true
		}
		return 0, true
	case int:
		return uintptr(typed), true
	case int8:
		return uintptr(typed), true
	case int16:
		return uintptr(typed), true
	case int32:
		return uintptr(typed), true
	case int64:
		return uintptr(typed), true
	case uint:
		return uintptr(typed), true
	case uint8:
		return uintptr(typed), true
	case uint16:
		return uintptr(typed), true
	case uint32:
		return uintptr(typed), true
	case uint64:
		return uintptr(typed), true
	default:
		return 0, false
	}
}

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

// Errno converts a C error number to a Go error.
//
//	rc := cgo.Errno(-2)
func Errno(rc C.int) error {
	if rc == 0 {
		return nil
	}
	return syscall.Errno(rc)
}

// WithErrno runs a function that returns C.int and maps the result to Go error.
//
//	result, err := cgo.WithErrno(func() C.int {
//		return C.my_function()
//	})
func WithErrno(fn func() C.int) (int, error) {
	result := fn()
	if err := Errno(result); err != nil {
		return int(result), err
	}
	return int(result), nil
}
