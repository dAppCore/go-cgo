package cgo

/*
#include <stdint.h>

static uintptr_t cgo_test_sum = 0;
static int cgo_test_rc = 0;

static void cgo_test_reset(void) {
	cgo_test_sum = 0;
	cgo_test_rc = 0;
}

static uintptr_t cgo_test_sum_value(void) {
	return cgo_test_sum;
}

static int cgo_test_fn_0(void) {
	cgo_test_sum += 1;
	return cgo_test_rc;
}

static int cgo_test_fn_1(uintptr_t a1) {
	cgo_test_sum = a1;
	return cgo_test_rc;
}

static int cgo_test_fn_18(
	uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5,
	uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10,
	uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15,
	uintptr_t a16, uintptr_t a17, uintptr_t a18
) {
	cgo_test_sum =
		a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 +
		a10 + a11 + a12 + a13 + a14 + a15 + a16 + a17 + a18;
	return cgo_test_rc;
}

static void* cgo_test_ptr_0(void) { return (void*)cgo_test_fn_0; }
static void* cgo_test_ptr_1(void) { return (void*)cgo_test_fn_1; }
static void* cgo_test_ptr_18(void) { return (void*)cgo_test_fn_18; }
*/
import "C"

import "unsafe"

func testCallReset() {
	C.cgo_test_reset()
}

func testCallPtr0() unsafe.Pointer {
	return C.cgo_test_ptr_0()
}

func testCallPtr1() unsafe.Pointer {
	return C.cgo_test_ptr_1()
}

func testCallPtr18() unsafe.Pointer {
	return C.cgo_test_ptr_18()
}

func testCallSum() uintptr {
	return uintptr(C.cgo_test_sum_value())
}

func testWithErrno(rc C.int) (int, error) {
	return WithErrno(func() C.int {
		return rc
	})
}
