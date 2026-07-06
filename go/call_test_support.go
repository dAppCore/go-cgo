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

import (
	"unsafe"

	core "dappco.re/go"
)

type testCInt = C.int

func testCallReset() {
	C.cgo_test_reset()
}

func testCallSetRC(rc int) {
	C.cgo_test_set_rc(C.int(rc))
}

func testCallPtr0() unsafe.Pointer {
	return C.cgo_test_ptr_0()
}

func testCallPtr1() unsafe.Pointer {
	return C.cgo_test_ptr_1()
}

func testCallPtr2() unsafe.Pointer {
	return C.cgo_test_ptr_2()
}

func testCallPtr18() unsafe.Pointer {
	return C.cgo_test_ptr_18()
}

// testCallPtrN returns the test C function pointer whose arity is n. Each
// pointer targets a cgo_test_fn_N that writes the sum of its arguments into
// cgo_test_sum, so Call dispatch through every switch arm (0..18) can be
// exercised and verified via testCallSum.
//
//	r := Call(testCallPtrN(5), 1, 2, 3, 4, 5)
//	// testCallSum() == 15
func testCallPtrN(n int) unsafe.Pointer {
	switch n {
	case 0:
		return C.cgo_test_ptr_0()
	case 1:
		return C.cgo_test_ptr_1()
	case 2:
		return C.cgo_test_ptr_2()
	case 3:
		return C.cgo_test_ptr_3()
	case 4:
		return C.cgo_test_ptr_4()
	case 5:
		return C.cgo_test_ptr_5()
	case 6:
		return C.cgo_test_ptr_6()
	case 7:
		return C.cgo_test_ptr_7()
	case 8:
		return C.cgo_test_ptr_8()
	case 9:
		return C.cgo_test_ptr_9()
	case 10:
		return C.cgo_test_ptr_10()
	case 11:
		return C.cgo_test_ptr_11()
	case 12:
		return C.cgo_test_ptr_12()
	case 13:
		return C.cgo_test_ptr_13()
	case 14:
		return C.cgo_test_ptr_14()
	case 15:
		return C.cgo_test_ptr_15()
	case 16:
		return C.cgo_test_ptr_16()
	case 17:
		return C.cgo_test_ptr_17()
	case 18:
		return C.cgo_test_ptr_18()
	default:
		panic("testCallPtrN: unsupported arity")
	}
}

func testCallSum() uintptr {
	return uintptr(C.cgo_test_sum_value())
}

func testWithErrno(rc C.int) core.Result {
	return WithErrno(func() C.int {
		return rc
	})
}

func testMalloc(size uintptr) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

// testCharPtr returns a heap-allocated, null-terminated C char pointer so
// Call's *C.char encodeCallArg branch can be exercised. The caller frees
// the result via Free.
//
//	cp := testCharPtr()
//	defer Free(unsafe.Pointer(cp))
//	r := Call(testCallPtr1(), cp)
func testCharPtr() *C.char {
	return CString("x")
}

// testNilCharPtr returns a nil *C.char so the nil branch of the *C.char
// case in encodeCallArg (which must encode to 0) can be exercised.
func testNilCharPtr() *C.char {
	return nil
}

// testSizeTArg returns a C.size_t value so encodeCallArg's C.size_t branch
// can be exercised through Call.
func testSizeTArg(v uintptr) C.size_t {
	return C.size_t(v)
}
