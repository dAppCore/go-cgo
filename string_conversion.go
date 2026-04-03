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

*/
import "C"

import (
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

// Call invokes a C function pointer and maps a non-zero return into a Go error.
//
//	err := Call(unsafe.Pointer(C.some_function), buffer.Ptr(), SizeT(len(data)))
//	if err != nil {
//		return err
//	}
func Call(function unsafe.Pointer, args ...interface{}) error {
	if function == nil {
		panic("cgo.Call: function pointer is nil")
	}

	var result C.int
	target := uintptr(function)

	switch len(args) {
	case 0:
		result = C.cgo_call_0(C.uintptr_t(target))
	case 1:
		a0, ok := toSyscallArg(args[0])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_1(C.uintptr_t(target), C.uintptr_t(a0))
	case 2:
		a0, ok := toSyscallArg(args[0])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a1, ok := toSyscallArg(args[1])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_2(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1))
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
		result = C.cgo_call_3(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2))
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
		result = C.cgo_call_4(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3))
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
		result = C.cgo_call_5(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4))
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
		result = C.cgo_call_6(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5))
	case 7:
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
		a6, ok := toSyscallArg(args[6])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_7(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5), C.uintptr_t(a6))
	case 8:
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
		a6, ok := toSyscallArg(args[6])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a7, ok := toSyscallArg(args[7])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_8(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5), C.uintptr_t(a6), C.uintptr_t(a7))
	case 9:
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
		a6, ok := toSyscallArg(args[6])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a7, ok := toSyscallArg(args[7])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a8, ok := toSyscallArg(args[8])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_9(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5), C.uintptr_t(a6), C.uintptr_t(a7), C.uintptr_t(a8))
	case 10:
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
		a6, ok := toSyscallArg(args[6])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a7, ok := toSyscallArg(args[7])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a8, ok := toSyscallArg(args[8])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a9, ok := toSyscallArg(args[9])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_10(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5), C.uintptr_t(a6), C.uintptr_t(a7), C.uintptr_t(a8), C.uintptr_t(a9))
	case 11:
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
		a6, ok := toSyscallArg(args[6])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a7, ok := toSyscallArg(args[7])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a8, ok := toSyscallArg(args[8])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a9, ok := toSyscallArg(args[9])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a10, ok := toSyscallArg(args[10])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_11(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5), C.uintptr_t(a6), C.uintptr_t(a7), C.uintptr_t(a8), C.uintptr_t(a9), C.uintptr_t(a10))
	case 12:
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
		a6, ok := toSyscallArg(args[6])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a7, ok := toSyscallArg(args[7])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a8, ok := toSyscallArg(args[8])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a9, ok := toSyscallArg(args[9])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a10, ok := toSyscallArg(args[10])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a11, ok := toSyscallArg(args[11])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_12(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5), C.uintptr_t(a6), C.uintptr_t(a7), C.uintptr_t(a8), C.uintptr_t(a9), C.uintptr_t(a10), C.uintptr_t(a11))
	case 13:
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
		a6, ok := toSyscallArg(args[6])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a7, ok := toSyscallArg(args[7])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a8, ok := toSyscallArg(args[8])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a9, ok := toSyscallArg(args[9])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a10, ok := toSyscallArg(args[10])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a11, ok := toSyscallArg(args[11])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a12, ok := toSyscallArg(args[12])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_13(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5), C.uintptr_t(a6), C.uintptr_t(a7), C.uintptr_t(a8), C.uintptr_t(a9), C.uintptr_t(a10), C.uintptr_t(a11), C.uintptr_t(a12))
	case 14:
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
		a6, ok := toSyscallArg(args[6])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a7, ok := toSyscallArg(args[7])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a8, ok := toSyscallArg(args[8])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a9, ok := toSyscallArg(args[9])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a10, ok := toSyscallArg(args[10])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a11, ok := toSyscallArg(args[11])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a12, ok := toSyscallArg(args[12])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a13, ok := toSyscallArg(args[13])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_14(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5), C.uintptr_t(a6), C.uintptr_t(a7), C.uintptr_t(a8), C.uintptr_t(a9), C.uintptr_t(a10), C.uintptr_t(a11), C.uintptr_t(a12), C.uintptr_t(a13))
	case 15:
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
		a6, ok := toSyscallArg(args[6])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a7, ok := toSyscallArg(args[7])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a8, ok := toSyscallArg(args[8])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a9, ok := toSyscallArg(args[9])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a10, ok := toSyscallArg(args[10])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a11, ok := toSyscallArg(args[11])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a12, ok := toSyscallArg(args[12])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a13, ok := toSyscallArg(args[13])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a14, ok := toSyscallArg(args[14])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_15(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5), C.uintptr_t(a6), C.uintptr_t(a7), C.uintptr_t(a8), C.uintptr_t(a9), C.uintptr_t(a10), C.uintptr_t(a11), C.uintptr_t(a12), C.uintptr_t(a13), C.uintptr_t(a14))
	case 16:
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
		a6, ok := toSyscallArg(args[6])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a7, ok := toSyscallArg(args[7])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a8, ok := toSyscallArg(args[8])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a9, ok := toSyscallArg(args[9])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a10, ok := toSyscallArg(args[10])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a11, ok := toSyscallArg(args[11])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a12, ok := toSyscallArg(args[12])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a13, ok := toSyscallArg(args[13])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a14, ok := toSyscallArg(args[14])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		a15, ok := toSyscallArg(args[15])
		if !ok {
			panic("cgo.Call: unsupported argument type")
		}
		result = C.cgo_call_16(C.uintptr_t(target), C.uintptr_t(a0), C.uintptr_t(a1), C.uintptr_t(a2), C.uintptr_t(a3), C.uintptr_t(a4), C.uintptr_t(a5), C.uintptr_t(a6), C.uintptr_t(a7), C.uintptr_t(a8), C.uintptr_t(a9), C.uintptr_t(a10), C.uintptr_t(a11), C.uintptr_t(a12), C.uintptr_t(a13), C.uintptr_t(a14), C.uintptr_t(a15))
	default:
		panic("cgo.Call: unsupported argument count: max 16")
	}

	if result != 0 {
		return Errno(result)
	}
	return nil
}

func toSyscallArg(value interface{}) (uintptr, bool) {
	switch typed := value.(type) {
	case nil:
		return 0, true
	case uintptr:
		return typed, true
	case *Buffer:
		if typed == nil {
			return 0, true
		}
		return uintptr(typed.Ptr()), true
	case unsafe.Pointer:
		return uintptr(typed), true
	case C.char:
		return uintptr(typed), true
	case C.schar:
		return uintptr(typed), true
	case C.uchar:
		return uintptr(typed), true
	case C.short:
		return uintptr(typed), true
	case C.ushort:
		return uintptr(typed), true
	case C.int:
		return uintptr(typed), true
	case C.long:
		return uintptr(typed), true
	case C.longlong:
		return uintptr(typed), true
	case C.ulonglong:
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
//	cString := CString("example")
//	result := GoString(cString)
//	Free(unsafe.Pointer(cString))
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
	return C.CString(value)
}

// Free releases memory previously returned by CString.
//
//	cString := CString("hello")
//	Free(unsafe.Pointer(cString))
func Free(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}
	C.free(ptr)
}

// Errno converts a C error number to a Go error.
//
//	resultCode := Errno(-2)
func Errno(resultCode C.int) error {
	if resultCode == 0 {
		return nil
	}
	return syscall.Errno(resultCode)
}

// WithErrno runs a function that returns C.int and maps the result to Go error.
//
//	resultCode, err := WithErrno(func() C.int {
//		return C.my_function()
//	})
func WithErrno(fn func() C.int) (int, error) {
	result := fn()
	if err := Errno(result); err != nil {
		return int(result), err
	}
	return int(result), nil
}
