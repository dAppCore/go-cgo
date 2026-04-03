package cgo

/*
#include <stddef.h>
#include <stdint.h>

int call_sum_length(char* data, size_t length) {
    if (data == NULL || length == 0) {
        return 1;
    }
    return 0;
}

int call_failure(void) {
    return 13;
}

int call_six_args(uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5) {
    if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13 && a4 == 14 && a5 == 15) {
        return 0;
    }
    return 13;
}

int call_buffer_argument(uintptr_t value) {
    return value == 0 ? 13 : 0;
}

uintptr_t call_sum_length_ptr(void) {
    return (uintptr_t)&call_sum_length;
}

uintptr_t call_failure_ptr(void) {
    return (uintptr_t)&call_failure;
}

uintptr_t call_six_args_ptr(void) {
    return (uintptr_t)&call_six_args;
}

uintptr_t call_buffer_argument_ptr(void) {
    return (uintptr_t)&call_buffer_argument;
}

*/
import "C"

import "unsafe"

func callSumLengthFunction() unsafe.Pointer {
	function := C.call_sum_length_ptr()
	return *(*unsafe.Pointer)(unsafe.Pointer(&function))
}

func callFailureFunction() unsafe.Pointer {
	function := C.call_failure_ptr()
	return *(*unsafe.Pointer)(unsafe.Pointer(&function))
}

func callSixArgumentFunction() unsafe.Pointer {
	function := C.call_six_args_ptr()
	return *(*unsafe.Pointer)(unsafe.Pointer(&function))
}

func callBufferArgumentFunction() unsafe.Pointer {
	function := C.call_buffer_argument_ptr()
	return *(*unsafe.Pointer)(unsafe.Pointer(&function))
}

func callWithErrnoZero() (int, error) {
	return WithErrno(func() C.int {
		return 0
	})
}

func callWithErrnoFailure() (int, error) {
	return WithErrno(func() C.int {
		return 2
	})
}
