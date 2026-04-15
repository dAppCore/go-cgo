package cgo

/*
#include <stdint.h>
#include <stdlib.h>

static uintptr_t cgo_test_sum = 0;
static int cgo_test_rc = 0;

static void cgo_test_reset(void) {
	cgo_test_sum = 0;
	cgo_test_rc = 0;
}

static void cgo_test_set_rc(int rc) {
	cgo_test_rc = rc;
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

static int cgo_test_fn_2(uintptr_t a1, uintptr_t a2) {
	cgo_test_sum = a1 + a2;
	return cgo_test_rc;
}

static int cgo_test_fn_3(uintptr_t a1, uintptr_t a2, uintptr_t a3) {
	cgo_test_sum = a1 + a2 + a3;
	return cgo_test_rc;
}

static int cgo_test_fn_4(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4) {
	cgo_test_sum = a1 + a2 + a3 + a4;
	return cgo_test_rc;
}

static int cgo_test_fn_5(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5;
	return cgo_test_rc;
}

static int cgo_test_fn_6(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6;
	return cgo_test_rc;
}

static int cgo_test_fn_7(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7;
	return cgo_test_rc;
}

static int cgo_test_fn_8(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8;
	return cgo_test_rc;
}

static int cgo_test_fn_9(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9;
	return cgo_test_rc;
}

static int cgo_test_fn_10(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + a10;
	return cgo_test_rc;
}

static int cgo_test_fn_11(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + a10 + a11;
	return cgo_test_rc;
}

static int cgo_test_fn_12(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + a10 + a11 + a12;
	return cgo_test_rc;
}

static int cgo_test_fn_13(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + a10 + a11 + a12 + a13;
	return cgo_test_rc;
}

static int cgo_test_fn_14(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + a10 + a11 + a12 + a13 + a14;
	return cgo_test_rc;
}

static int cgo_test_fn_15(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + a10 + a11 + a12 + a13 + a14 + a15;
	return cgo_test_rc;
}

static int cgo_test_fn_16(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15, uintptr_t a16) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + a10 + a11 + a12 + a13 + a14 + a15 + a16;
	return cgo_test_rc;
}

static int cgo_test_fn_17(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9, uintptr_t a10, uintptr_t a11, uintptr_t a12, uintptr_t a13, uintptr_t a14, uintptr_t a15, uintptr_t a16, uintptr_t a17) {
	cgo_test_sum = a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + a10 + a11 + a12 + a13 + a14 + a15 + a16 + a17;
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
static void* cgo_test_ptr_2(void) { return (void*)cgo_test_fn_2; }
static void* cgo_test_ptr_3(void) { return (void*)cgo_test_fn_3; }
static void* cgo_test_ptr_4(void) { return (void*)cgo_test_fn_4; }
static void* cgo_test_ptr_5(void) { return (void*)cgo_test_fn_5; }
static void* cgo_test_ptr_6(void) { return (void*)cgo_test_fn_6; }
static void* cgo_test_ptr_7(void) { return (void*)cgo_test_fn_7; }
static void* cgo_test_ptr_8(void) { return (void*)cgo_test_fn_8; }
static void* cgo_test_ptr_9(void) { return (void*)cgo_test_fn_9; }
static void* cgo_test_ptr_10(void) { return (void*)cgo_test_fn_10; }
static void* cgo_test_ptr_11(void) { return (void*)cgo_test_fn_11; }
static void* cgo_test_ptr_12(void) { return (void*)cgo_test_fn_12; }
static void* cgo_test_ptr_13(void) { return (void*)cgo_test_fn_13; }
static void* cgo_test_ptr_14(void) { return (void*)cgo_test_fn_14; }
static void* cgo_test_ptr_15(void) { return (void*)cgo_test_fn_15; }
static void* cgo_test_ptr_16(void) { return (void*)cgo_test_fn_16; }
static void* cgo_test_ptr_17(void) { return (void*)cgo_test_fn_17; }
static void* cgo_test_ptr_18(void) { return (void*)cgo_test_fn_18; }
*/
import "C"

import "unsafe"

func testCallReset() {
	C.cgo_test_reset()
}

func testCallSetRC(rc C.int) {
	C.cgo_test_set_rc(rc)
}

func testCallPtr0() unsafe.Pointer {
	return C.cgo_test_ptr_0()
}

func testCallPtr1() unsafe.Pointer {
	return C.cgo_test_ptr_1()
}

func testCallPtr2() unsafe.Pointer  { return C.cgo_test_ptr_2() }
func testCallPtr3() unsafe.Pointer  { return C.cgo_test_ptr_3() }
func testCallPtr4() unsafe.Pointer  { return C.cgo_test_ptr_4() }
func testCallPtr5() unsafe.Pointer  { return C.cgo_test_ptr_5() }
func testCallPtr6() unsafe.Pointer  { return C.cgo_test_ptr_6() }
func testCallPtr7() unsafe.Pointer  { return C.cgo_test_ptr_7() }
func testCallPtr8() unsafe.Pointer  { return C.cgo_test_ptr_8() }
func testCallPtr9() unsafe.Pointer  { return C.cgo_test_ptr_9() }
func testCallPtr10() unsafe.Pointer { return C.cgo_test_ptr_10() }
func testCallPtr11() unsafe.Pointer { return C.cgo_test_ptr_11() }
func testCallPtr12() unsafe.Pointer { return C.cgo_test_ptr_12() }
func testCallPtr13() unsafe.Pointer { return C.cgo_test_ptr_13() }
func testCallPtr14() unsafe.Pointer { return C.cgo_test_ptr_14() }
func testCallPtr15() unsafe.Pointer { return C.cgo_test_ptr_15() }
func testCallPtr16() unsafe.Pointer { return C.cgo_test_ptr_16() }
func testCallPtr17() unsafe.Pointer { return C.cgo_test_ptr_17() }
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

func testCallUintptr(value C.uintptr_t) error {
	return Call(testCallPtr1(), value)
}

func testCallSizeT(value C.size_t) error {
	return Call(testCallPtr1(), value)
}

func testCallCInt(value C.int) error {
	return Call(testCallPtr1(), value)
}

func testMalloc(size C.size_t) unsafe.Pointer {
	return C.malloc(size)
}
