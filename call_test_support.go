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

uintptr_t call_sum_length_ptr(void) {
    return (uintptr_t)&call_sum_length;
}

uintptr_t call_failure_ptr(void) {
    return (uintptr_t)&call_failure;
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
