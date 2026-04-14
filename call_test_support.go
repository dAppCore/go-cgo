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

int call_seven_args(uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6) {
	if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13 && a4 == 14 && a5 == 15 && a6 == 16) {
		return 0;
	}
	return 13;
}

int call_eight_args(uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7) {
	if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13 && a4 == 14 && a5 == 15 && a6 == 16 && a7 == 17) {
		return 0;
	}
	return 13;
}

int call_nine_args(uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8) {
	if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13 && a4 == 14 && a5 == 15 && a6 == 16 && a7 == 17 && a8 == 18) {
		return 0;
	}
	return 13;
}

int call_thirteen_args(uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12) {
	if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13 && a4 == 14 && a5 == 15 && a6 == 16 && a7 == 17 && a8 == 18 && a9 == 19 && a10 == 20 && a11 == 21 && a12 == 22) {
		return 0;
	}
	return 13;
}

int call_fourteen_args(uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13) {
	if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13 && a4 == 14 && a5 == 15 && a6 == 16 && a7 == 17 && a8 == 18 && a9 == 19 && a10 == 20 && a11 == 21 && a12 == 22 && a13 == 23) {
		return 0;
	}
	return 13;
}

int call_fifteen_args(uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14) {
	if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13 && a4 == 14 && a5 == 15 && a6 == 16 && a7 == 17 && a8 == 18 && a9 == 19 && a10 == 20 && a11 == 21 && a12 == 22 && a13 == 23 && a14 == 24) {
		return 0;
	}
	return 13;
}

int call_sixteen_args(uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15) {
	if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13 && a4 == 14 && a5 == 15 && a6 == 16 && a7 == 17 && a8 == 18 && a9 == 19 && a10 == 20 && a11 == 21 && a12 == 22 && a13 == 23 && a14 == 24 && a15 == 25) {
		return 0;
	}
	return 13;
}

int call_seventeen_args(uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15, uintptr_t a16) {
	if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13 && a4 == 14 && a5 == 15 && a6 == 16 && a7 == 17 && a8 == 18 && a9 == 19 && a10 == 20 && a11 == 21 && a12 == 22 && a13 == 23 && a14 == 24 && a15 == 25 && a16 == 26) {
		return 0;
	}
	return 13;
}

int call_eighteen_args(uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15, uintptr_t a16, uintptr_t a17) {
	if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13 && a4 == 14 && a5 == 15 && a6 == 16 && a7 == 17 && a8 == 18 && a9 == 19 && a10 == 20 && a11 == 21 && a12 == 22 && a13 == 23 && a14 == 24 && a15 == 25 && a16 == 26 && a17 == 27) {
		return 0;
	}
	return 13;
}

int call_buffer_argument(uintptr_t value) {
	return value == 0 ? 13 : 0;
}

int call_zero_byte_slice_argument(uintptr_t value) {
	return value == 0 ? 0 : 13;
}

int call_c_types_arguments(long a0, unsigned long a1, long long a2, unsigned long long a3) {
	if (a0 == 10 && a1 == 11 && a2 == 12 && a3 == 13) {
		return 0;
	}
	return 13;
}

int call_uintptr_argument(uintptr_t value) {
	return value == 42 ? 0 : 13;
}

int call_any_pointer_argument(uintptr_t value) {
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

uintptr_t call_seven_args_ptr(void) {
	return (uintptr_t)&call_seven_args;
}

uintptr_t call_eight_args_ptr(void) {
	return (uintptr_t)&call_eight_args;
}

uintptr_t call_nine_args_ptr(void) {
	return (uintptr_t)&call_nine_args;
}

uintptr_t call_thirteen_args_ptr(void) {
	return (uintptr_t)&call_thirteen_args;
}

uintptr_t call_fourteen_args_ptr(void) {
	return (uintptr_t)&call_fourteen_args;
}

uintptr_t call_fifteen_args_ptr(void) {
	return (uintptr_t)&call_fifteen_args;
}

uintptr_t call_sixteen_args_ptr(void) {
	return (uintptr_t)&call_sixteen_args;
}

uintptr_t call_seventeen_args_ptr(void) {
	return (uintptr_t)&call_seventeen_args;
}

uintptr_t call_eighteen_args_ptr(void) {
	return (uintptr_t)&call_eighteen_args;
}

uintptr_t call_buffer_argument_ptr(void) {
	return (uintptr_t)&call_buffer_argument;
}

uintptr_t call_zero_byte_slice_argument_ptr(void) {
	return (uintptr_t)&call_zero_byte_slice_argument;
}

uintptr_t call_c_types_arguments_ptr(void) {
	return (uintptr_t)&call_c_types_arguments;
}

uintptr_t call_uintptr_argument_ptr(void) {
	return (uintptr_t)&call_uintptr_argument;
}

uintptr_t call_any_pointer_argument_ptr(void) {
	return (uintptr_t)&call_any_pointer_argument;
}

*/
import "C"

import "unsafe"

func cFunctionPointer(value C.uintptr_t) unsafe.Pointer {
	return unsafe.Pointer(uintptr(value))
}

func callSumLengthFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_sum_length_ptr())
}

func callFailureFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_failure_ptr())
}

func callSixArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_six_args_ptr())
}

func callSevenArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_seven_args_ptr())
}

func callEightArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_eight_args_ptr())
}

func callNineArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_nine_args_ptr())
}

func callThirteenArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_thirteen_args_ptr())
}

func callFourteenArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_fourteen_args_ptr())
}

func callFifteenArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_fifteen_args_ptr())
}

func callSixteenArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_sixteen_args_ptr())
}

func callSeventeenArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_seventeen_args_ptr())
}

func callEighteenArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_eighteen_args_ptr())
}

func callBufferArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_buffer_argument_ptr())
}

func callZeroByteSliceArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_zero_byte_slice_argument_ptr())
}

func callCTypeArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_c_types_arguments_ptr())
}

func callCUintptrArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_uintptr_argument_ptr())
}

func callAnyPointerArgumentFunction() unsafe.Pointer {
	return cFunctionPointer(C.call_any_pointer_argument_ptr())
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

func callSupportsCIntegerArguments() error {
	return Call(callCTypeArgumentFunction(), C.long(10), C.ulong(11), C.longlong(12), C.ulonglong(13))
}

func callSupportsCUintptrArgument() error {
	return Call(callCUintptrArgumentFunction(), C.uintptr_t(42))
}
