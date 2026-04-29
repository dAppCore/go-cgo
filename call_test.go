package cgo

import (
	"syscall"
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
	args := make([]interface{}, 18)
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

func TestCall_Call_UintptrLike_Good(t *T) {
	testCallReset()
	r := Call(testCallPtr1(), testUintptr(15))

	AssertTrue(t, r.OK)
	AssertEqual(t, 0, r.Value)
	AssertEqual(t, uintptr(15), testCallSum())
}
