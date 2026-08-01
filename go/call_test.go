package cgo

import (
	"syscall"
	"unsafe"
)

type testUintptr uintptr

func TestCall_Call_Good(t *T) {
	testCallReset()
	r := Call(testCallPtr2(), 2, 3)

	AssertTrue(t, r.OK)
	AssertEqual(t, 0, r.Value)
	AssertEqual(t, uintptr(5), testCallSum())
}

func TestCall_Call_Bad(t *T) {
	AssertPanicsWithError(t, "function pointer is nil", func() {
		_ = Call(nil)
	})
	AssertPanicsWithError(t, "unsupported argument type at argument 1", func() {
		_ = Call(testCallPtr0(), "bad")
	})
}

func TestCall_Call_Ugly(t *T) {
	testCallReset()
	args := make([]any, 18)
	for i := range args {
		args[i] = i + 1
	}

	r := Call(testCallPtr18(), args...)
	AssertTrue(t, r.OK)
	AssertEqual(t, uintptr(171), testCallSum())

	testCallReset()
	testCallSetRC(5)
	r = Call(testCallPtr0())
	AssertFalse(t, r.OK)
	AssertErrorIs(t, r.Value.(error), syscall.Errno(5))
}

func TestCall_Call_ArgumentEncodings_Good(t *T) {
	testCallReset()
	buffer := NewBuffer(2)
	defer buffer.Free()
	copied := buffer.CopyFrom([]byte("go"))

	AssertEqual(t, 2, copied)
	r := Call(testCallPtr1(), buffer)
	AssertTrue(t, r.OK)
	AssertNotEqual(t, uintptr(0), testCallSum())
}

func TestCall_Call_EmptyBytes_Good(t *T) {
	testCallReset()
	r := Call(testCallPtr1(), []byte(nil))

	AssertTrue(t, r.OK)
	AssertEqual(t, 0, r.Value)
	AssertEqual(t, uintptr(0), testCallSum())
}

// TestCall_Call_NonEmptyBytes_Good exercises the non-empty []byte branch of
// encodeCallArg, which must encode the address of the slice's first element
// (a non-zero pointer) rather than the zero an empty slice yields.
func TestCall_Call_NonEmptyBytes_Good(t *T) {
	testCallReset()
	payload := []byte("cgo")
	r := Call(testCallPtr1(), payload)

	AssertTrue(t, r.OK)
	AssertNotEqual(t, uintptr(0), testCallSum())
}

func TestCall_Call_UintptrLike_Good(t *T) {
	testCallReset()
	r := Call(testCallPtr1(), testUintptr(15))

	AssertTrue(t, r.OK)
	AssertEqual(t, 0, r.Value)
	AssertEqual(t, uintptr(15), testCallSum())
}

// TestCall_Call_AllArities_Good walks every dispatch arm 0..18. Each test
// pointer sums its uintptr arguments into cgo_test_sum, so the expected
// result for arity n with args 1..n is the triangular number n(n+1)/2.
func TestCall_Call_AllArities_Good(t *T) {
	for n := 0; n <= 18; n++ {
		testCallReset()

		args := make([]any, n)
		for i := range args {
			args[i] = i + 1
		}

		r := Call(testCallPtrN(n), args...)
		AssertTrue(t, r.OK)
		AssertEqual(t, 0, r.Value)

		if n == 0 {
			// cgo_test_fn_0 increments the sum by one rather than summing
			// arguments, so the count is 1 after a reset.
			AssertEqual(t, uintptr(1), testCallSum())
			continue
		}

		want := uintptr(n * (n + 1) / 2)
		AssertEqual(t, want, testCallSum())
	}
}

// TestCall_Call_ArgumentTypes_Good exercises every concrete branch of
// encodeCallArg via the single-argument dispatch. Each scalar type must
// round-trip to the same uintptr the C function records.
func TestCall_Call_ArgumentTypes_Good(t *T) {
	cases := []struct {
		name string
		arg  any
		want uintptr
	}{
		{"int", int(7), 7},
		{"int32", int32(8), 8},
		{"int64", int64(9), 9},
		{"uint", uint(10), 10},
		{"uint32", uint32(11), 11},
		{"uint64", uint64(12), 12},
		{"uintptr", uintptr(13), 13},
		{"testCInt", testCInt(14), 14},
	}

	for _, tc := range cases {
		testCallReset()
		r := Call(testCallPtr1(), tc.arg)
		AssertTrue(t, r.OK)
		AssertEqual(t, tc.want, testCallSum())
	}
}

// TestCall_Call_PointerArgs_Good exercises the pointer-shaped branches of
// encodeCallArg: unsafe.Pointer, *Buffer and a nil *Buffer (which must
// encode to 0 rather than panic).
func TestCall_Call_PointerArgs_Good(t *T) {
	testCallReset()
	buffer := NewBuffer(4)
	defer buffer.Free()

	r := Call(testCallPtr1(), buffer.Ptr())
	AssertTrue(t, r.OK)
	AssertEqual(t, uintptr(buffer.Ptr()), testCallSum())

	testCallReset()
	r = Call(testCallPtr1(), buffer)
	AssertTrue(t, r.OK)
	AssertEqual(t, uintptr(buffer.Ptr()), testCallSum())

	testCallReset()
	var nilBuffer *Buffer
	r = Call(testCallPtr1(), nilBuffer)
	AssertTrue(t, r.OK)
	AssertEqual(t, uintptr(0), testCallSum())
}

// TestCall_Call_CharPtrArg_Good exercises the *C.char branch of
// encodeCallArg for both a live pointer and a nil pointer (which must
// encode to 0).
func TestCall_Call_CharPtrArg_Good(t *T) {
	testCallReset()
	cp := testCharPtr()
	defer Free(unsafe.Pointer(cp))

	r := Call(testCallPtr1(), cp)
	AssertTrue(t, r.OK)
	AssertNotEqual(t, uintptr(0), testCallSum())

	testCallReset()
	r = Call(testCallPtr1(), testNilCharPtr())
	AssertTrue(t, r.OK)
	AssertEqual(t, uintptr(0), testCallSum())
}

// TestCall_Call_SizeTArg_Good exercises the C.size_t branch of
// encodeCallArg through the single-argument dispatch.
func TestCall_Call_SizeTArg_Good(t *T) {
	testCallReset()
	r := Call(testCallPtr1(), testSizeTArg(42))

	AssertTrue(t, r.OK)
	AssertEqual(t, uintptr(42), testCallSum())
}

// TestCall_Call_ArgEncoding_Bad covers the nil-argument guard, an
// unsupported concrete type and the arity ceiling — each panics before
// any C dispatch runs.
func TestCall_Call_ArgEncoding_Bad(t *T) {
	AssertPanicsWithError(t, "unsupported argument type at argument 1", func() {
		_ = Call(testCallPtr1(), nil)
	})
	AssertPanicsWithError(t, "unsupported argument type at argument 1", func() {
		type unsupported struct{ x int }
		_ = Call(testCallPtr1(), unsupported{x: 1})
	})
	AssertPanicsWithError(t, "unsupported arity: 19", func() {
		args := make([]any, 19)
		for i := range args {
			args[i] = i + 1
		}
		_ = Call(testCallPtr0(), args...)
	})
}
