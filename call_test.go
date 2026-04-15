package cgo

import (
	"testing"
	"unsafe"
)

type testUintptr uintptr

func TestCall_SizeT_Good(t *testing.T) {
	if got := SizeT(3); got != 3 {
		t.Fatalf("SizeT(3) = %d, want 3", got)
	}
}

func TestCall_SizeT_Boundary_Good(t *testing.T) {
	max := int(^uint32(0) >> 1)

	if got := SizeT(max); uint64(got) != uint64(max) {
		t.Fatalf("SizeT(max) = %d, want %d", got, max)
	}
}

func TestCall_SizeT_Bad(t *testing.T) {
	mustPanic(t, "cgo.SizeT: negative values are not representable as C.size_t", func() {
		_ = SizeT(-1)
	})
}

func TestCall_Int_Good(t *testing.T) {
	if got := Int(4); got != 4 {
		t.Fatalf("Int(4) = %d, want 4", got)
	}
}

func TestCall_Int_Boundary_Good(t *testing.T) {
	max := int(^uint32(0) >> 1)

	if got := Int(max); int(got) != max {
		t.Fatalf("Int(max) = %d, want %d", got, max)
	}
}

func TestCall_Int_Bad(t *testing.T) {
	mustPanic(t, "cgo.Int: value exceeds C.int range", func() {
		_ = Int(int(^uint32(0)>>1) + 1)
	})
}

func TestCall_Int_Negative_Bad(t *testing.T) {
	mustPanic(t, "cgo.Int: value exceeds C.int range", func() {
		_ = Int(-1)
	})
}

func TestCall_0Args_Good(t *testing.T) {
	testCallReset()
	if err := Call(testCallPtr0()); err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if got := testCallSum(); got != 1 {
		t.Fatalf("sum = %d, want 1", got)
	}
}

func TestCall_1Args_Good(t *testing.T) {
	testCallReset()
	if err := Call(testCallPtr1(), 7); err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if got := testCallSum(); got != 7 {
		t.Fatalf("sum = %d, want 7", got)
	}
}

func TestCall_BytesArg_Good(t *testing.T) {
	testCallReset()

	payload := []byte("go")
	if err := Call(testCallPtr1(), payload); err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if got := testCallSum(); got == 0 {
		t.Fatal("sum = 0, want non-zero pointer value")
	}
}

func TestCall_BytesNil_Good(t *testing.T) {
	testCallReset()

	if err := Call(testCallPtr1(), []byte(nil)); err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if got := testCallSum(); got != 0 {
		t.Fatalf("sum = %d, want 0", got)
	}
}

func TestCall_BufferArg_Good(t *testing.T) {
	testCallReset()

	buffer := NewBuffer(2)
	defer buffer.Free()

	if copied := buffer.CopyFrom([]byte("go")); copied != 2 {
		t.Fatalf("CopyFrom copied = %d, want 2", copied)
	}
	if err := Call(testCallPtr1(), buffer); err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if got := testCallSum(); got == 0 {
		t.Fatal("sum = 0, want non-zero pointer value")
	}
}

func TestCall_BufferNil_Good(t *testing.T) {
	testCallReset()

	if err := Call(testCallPtr1(), (*Buffer)(nil)); err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if got := testCallSum(); got != 0 {
		t.Fatalf("sum = %d, want 0", got)
	}
}

func TestCall_CStringArg_Good(t *testing.T) {
	testCallReset()

	ptr := CString("hello")
	defer Free(unsafe.Pointer(ptr))

	if err := Call(testCallPtr1(), ptr); err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if got := testCallSum(); got == 0 {
		t.Fatal("sum = 0, want non-zero pointer value")
	}
}

func TestCall_UnsafePointerArg_Good(t *testing.T) {
	testCallReset()

	value := 42
	if err := Call(testCallPtr1(), unsafe.Pointer(&value)); err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if got := testCallSum(); got == 0 {
		t.Fatal("sum = 0, want non-zero pointer value")
	}
}

func TestCall_18Args_Good(t *testing.T) {
	testCallReset()
	args := make([]interface{}, 18)
	for i := range args {
		args[i] = i + 1
	}
	if err := Call(testCallPtr18(), args...); err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if got := testCallSum(); got != 171 {
		t.Fatalf("sum = %d, want 171", got)
	}
}

func TestCall_19Args_Bad(t *testing.T) {
	args := make([]interface{}, 19)
	for i := range args {
		args[i] = i
	}
	mustPanic(t, "cgo.Call: unsupported arity: 19", func() {
		_ = Call(testCallPtr0(), args...)
	})
}

func TestCall_CTypes_Good(t *testing.T) {
	testCallReset()

	if err := Call(testCallPtr1(), SizeT(11)); err != nil {
		t.Fatalf("Call(SizeT) returned error: %v", err)
	}
	if got := testCallSum(); got != 11 {
		t.Fatalf("sum(SizeT) = %d, want 11", got)
	}

	testCallReset()
	if err := Call(testCallPtr1(), Int(12)); err != nil {
		t.Fatalf("Call(Int) returned error: %v", err)
	}
	if got := testCallSum(); got != 12 {
		t.Fatalf("sum(Int) = %d, want 12", got)
	}
}

func TestCall_CSizeT_Good(t *testing.T) {
	testCallReset()

	if err := testCallSizeT(13); err != nil {
		t.Fatalf("Call(C.size_t) returned error: %v", err)
	}
	if got := testCallSum(); got != 13 {
		t.Fatalf("sum(C.size_t) = %d, want 13", got)
	}
}

func TestCall_CInt_Good(t *testing.T) {
	testCallReset()

	if err := testCallCInt(14); err != nil {
		t.Fatalf("Call(C.int) returned error: %v", err)
	}
	if got := testCallSum(); got != 14 {
		t.Fatalf("sum(C.int) = %d, want 14", got)
	}
}

func TestCall_CUintptr_Good(t *testing.T) {
	testCallReset()

	if err := testCallUintptr(13); err != nil {
		t.Fatalf("Call(C.uintptr_t) returned error: %v", err)
	}
	if got := testCallSum(); got != 13 {
		t.Fatalf("sum(C.uintptr_t) = %d, want 13", got)
	}
}

func TestCall_UintptrLike_Good(t *testing.T) {
	testCallReset()

	if err := Call(testCallPtr1(), testUintptr(15)); err != nil {
		t.Fatalf("Call(uintptr-like) returned error: %v", err)
	}
	if got := testCallSum(); got != 15 {
		t.Fatalf("sum(uintptr-like) = %d, want 15", got)
	}
}

func TestCall_Errno_Good(t *testing.T) {
	if err := Errno(0); err != nil {
		t.Fatalf("Errno(0) = %v, want nil", err)
	}
	if err := Errno(5); err == nil {
		t.Fatal("Errno(5) = nil, want error")
	}
}

func TestCall_WithErrno_Good(t *testing.T) {
	rc, err := testWithErrno(0)
	if err != nil {
		t.Fatalf("WithErrno success returned error: %v", err)
	}
	if rc != 0 {
		t.Fatalf("WithErrno rc = %d, want 0", rc)
	}

	rc, err = testWithErrno(2)
	if rc != 2 {
		t.Fatalf("WithErrno rc = %d, want 2", rc)
	}
	if err == nil {
		t.Fatal("WithErrno error is nil, want error")
	}
}

func TestCall_NilFunction_Bad(t *testing.T) {
	mustPanic(t, "cgo.Call: function pointer is nil", func() {
		_ = Call(nil)
	})
}

func TestCall_UnsupportedArgType_Bad(t *testing.T) {
	mustPanic(t, "cgo.Call: unsupported argument type at argument 1", func() {
		_ = Call(testCallPtr0(), "bad")
	})
}

func TestCall_UnsupportedIntegerType_Bad(t *testing.T) {
	mustPanic(t, "cgo.Call: unsupported argument type at argument 1", func() {
		_ = Call(testCallPtr0(), int8(1))
	})

	mustPanic(t, "cgo.Call: unsupported argument type at argument 1", func() {
		_ = Call(testCallPtr0(), uint16(1))
	})
}

func TestCall_NilArg_Bad(t *testing.T) {
	mustPanic(t, "cgo.Call: unsupported argument type at argument 1", func() {
		_ = Call(testCallPtr0(), nil)
	})
}
