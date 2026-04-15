package cgo

/*
#include <stdlib.h>
#include <stdint.h>

static inline int cgo_call_0(void *fn) {
	return ((int (*)(void))fn)();
}

static inline int cgo_call_1(void *fn, uintptr_t a1) {
	return ((int (*)(uintptr_t))fn)(a1);
}

static inline int cgo_call_2(void *fn, uintptr_t a1, uintptr_t a2) {
	return ((int (*)(uintptr_t, uintptr_t))fn)(a1, a2);
}

static inline int cgo_call_3(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3);
}

static inline int cgo_call_4(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4);
}

static inline int cgo_call_5(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5);
}

static inline int cgo_call_6(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6);
}

static inline int cgo_call_7(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7);
}

static inline int cgo_call_8(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8);
}

static inline int cgo_call_9(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static inline int cgo_call_10(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static inline int cgo_call_11(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static inline int cgo_call_12(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static inline int cgo_call_13(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static inline int cgo_call_14(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14);
}

static inline int cgo_call_15(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15);
}

static inline int cgo_call_16(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15, uintptr_t a16) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16);
}

static inline int cgo_call_17(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15, uintptr_t a16, uintptr_t a17) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17);
}

static inline int cgo_call_18(void *fn, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15, uintptr_t a16, uintptr_t a17, uintptr_t a18) {
	return ((int (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18);
}
*/
import "C"

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

// Call invokes a C function pointer and maps a non-zero return code to an error.
//
//	err := Call(unsafe.Pointer(C.some_function), buffer.Ptr(), SizeT(len(payload)))
func Call(function unsafe.Pointer, args ...interface{}) error {
	if function == nil {
		panic("cgo.Call: function pointer is nil")
	}
	if len(args) > 18 {
		panic(fmt.Sprintf("cgo.Call: unsupported arity: %d", len(args)))
	}

	encoded := make([]uintptr, len(args))
	for i, arg := range args {
		encoded[i] = encodeCallArg(i+1, arg)
	}

	var rc C.int
	switch len(encoded) {
	case 0:
		rc = C.cgo_call_0(function)
	case 1:
		rc = C.cgo_call_1(function, C.uintptr_t(encoded[0]))
	case 2:
		rc = C.cgo_call_2(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]))
	case 3:
		rc = C.cgo_call_3(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]))
	case 4:
		rc = C.cgo_call_4(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]))
	case 5:
		rc = C.cgo_call_5(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]))
	case 6:
		rc = C.cgo_call_6(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]))
	case 7:
		rc = C.cgo_call_7(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]))
	case 8:
		rc = C.cgo_call_8(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]))
	case 9:
		rc = C.cgo_call_9(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]), C.uintptr_t(encoded[8]))
	case 10:
		rc = C.cgo_call_10(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]), C.uintptr_t(encoded[8]), C.uintptr_t(encoded[9]))
	case 11:
		rc = C.cgo_call_11(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]), C.uintptr_t(encoded[8]), C.uintptr_t(encoded[9]), C.uintptr_t(encoded[10]))
	case 12:
		rc = C.cgo_call_12(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]), C.uintptr_t(encoded[8]), C.uintptr_t(encoded[9]), C.uintptr_t(encoded[10]), C.uintptr_t(encoded[11]))
	case 13:
		rc = C.cgo_call_13(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]), C.uintptr_t(encoded[8]), C.uintptr_t(encoded[9]), C.uintptr_t(encoded[10]), C.uintptr_t(encoded[11]), C.uintptr_t(encoded[12]))
	case 14:
		rc = C.cgo_call_14(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]), C.uintptr_t(encoded[8]), C.uintptr_t(encoded[9]), C.uintptr_t(encoded[10]), C.uintptr_t(encoded[11]), C.uintptr_t(encoded[12]), C.uintptr_t(encoded[13]))
	case 15:
		rc = C.cgo_call_15(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]), C.uintptr_t(encoded[8]), C.uintptr_t(encoded[9]), C.uintptr_t(encoded[10]), C.uintptr_t(encoded[11]), C.uintptr_t(encoded[12]), C.uintptr_t(encoded[13]), C.uintptr_t(encoded[14]))
	case 16:
		rc = C.cgo_call_16(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]), C.uintptr_t(encoded[8]), C.uintptr_t(encoded[9]), C.uintptr_t(encoded[10]), C.uintptr_t(encoded[11]), C.uintptr_t(encoded[12]), C.uintptr_t(encoded[13]), C.uintptr_t(encoded[14]), C.uintptr_t(encoded[15]))
	case 17:
		rc = C.cgo_call_17(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]), C.uintptr_t(encoded[8]), C.uintptr_t(encoded[9]), C.uintptr_t(encoded[10]), C.uintptr_t(encoded[11]), C.uintptr_t(encoded[12]), C.uintptr_t(encoded[13]), C.uintptr_t(encoded[14]), C.uintptr_t(encoded[15]), C.uintptr_t(encoded[16]))
	case 18:
		rc = C.cgo_call_18(function, C.uintptr_t(encoded[0]), C.uintptr_t(encoded[1]), C.uintptr_t(encoded[2]), C.uintptr_t(encoded[3]), C.uintptr_t(encoded[4]), C.uintptr_t(encoded[5]), C.uintptr_t(encoded[6]), C.uintptr_t(encoded[7]), C.uintptr_t(encoded[8]), C.uintptr_t(encoded[9]), C.uintptr_t(encoded[10]), C.uintptr_t(encoded[11]), C.uintptr_t(encoded[12]), C.uintptr_t(encoded[13]), C.uintptr_t(encoded[14]), C.uintptr_t(encoded[15]), C.uintptr_t(encoded[16]), C.uintptr_t(encoded[17]))
	default:
		panic("cgo.Call: unsupported arity")
	}

	runtime.KeepAlive(args)
	return Errno(rc)
}

func encodeCallArg(position int, arg interface{}) uintptr {
	if arg == nil {
		panic(fmt.Sprintf("cgo.Call: unsupported argument type at argument %d", position))
	}

	switch v := arg.(type) {
	case unsafe.Pointer:
		return uintptr(v)
	case []byte:
		if len(v) == 0 {
			return 0
		}
		return uintptr(unsafe.Pointer(&v[0]))
	case *Buffer:
		if v == nil {
			return 0
		}
		return uintptr(v.Ptr())
	case *C.char:
		if v == nil {
			return 0
		}
		return uintptr(unsafe.Pointer(v))
	case C.size_t:
		return uintptr(v)
	case C.int:
		return uintptr(v)
	case int:
		return uintptr(v)
	case int32:
		return uintptr(v)
	case int64:
		return uintptr(v)
	case uint:
		return uintptr(v)
	case uint32:
		return uintptr(v)
	case uint64:
		return uintptr(v)
	default:
		if reflect.TypeOf(arg).Kind() == reflect.Uintptr {
			return uintptr(reflect.ValueOf(arg).Uint())
		}
		panic(fmt.Sprintf("cgo.Call: unsupported argument type at argument %d", position))
	}
}
