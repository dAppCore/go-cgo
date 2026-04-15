package cgo

import "testing"

func TestCall_SizeT_Good(t *testing.T) {
	if got := SizeT(3); got != 3 {
		t.Fatalf("SizeT(3) = %d, want 3", got)
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

func TestCall_UnsupportedArgType_Bad(t *testing.T) {
	mustPanic(t, "cgo.Call: unsupported argument type at argument 1: string", func() {
		_ = Call(testCallPtr0(), "bad")
	})
}

func TestCall_NilArg_Bad(t *testing.T) {
	mustPanic(t, "cgo.Call: unsupported argument type at argument 1: <nil>", func() {
		_ = Call(testCallPtr0(), nil)
	})
}
