package cgo

import (
	"errors"
	"syscall"
	"testing"
)

func TestSizeTAndIntConversions(t *testing.T) {
	t.Parallel()

	const exampleLength = 12
	if got := int(SizeT(exampleLength)); got != exampleLength {
		t.Fatalf("expected SizeT(%d) to be %d, got %d", exampleLength, exampleLength, got)
	}

	if got := Int(-1); got != -1 {
		t.Fatalf("expected Int(-1) to be -1, got %v", got)
	}

	if got := Int(4096); got != 4096 {
		t.Fatalf("expected Int(4096) to be 4096, got %v", got)
	}

	cIntBits := cIntBitSize()
	maxInt := int64(int(^uint(0) >> 1))
	minInt := -maxInt - 1
	cIntMax := int64(1)<<(cIntBits-1) - 1
	cIntMin := -cIntMax - 1

	// Verify exact C.int boundaries for this platform.
	if cIntMax <= maxInt {
		got := Int(int(cIntMax))
		if int(got) != int(cIntMax) {
			t.Fatalf("expected Int(%d) to be %d, got %v", cIntMax, cIntMax, got)
		}
	}
	if cIntMin >= minInt {
		got := Int(int(cIntMin))
		if int(got) != int(cIntMin) {
			t.Fatalf("expected Int(%d) to be %d, got %v", cIntMin, cIntMin, got)
		}
	}

	// Int should panic when value is outside C.int range.
	if cIntMax < maxInt {
		assertPanics(t, "value exceeds C.int range", func() {
			_ = Int(int(cIntMax + 1))
		})
	}
	if cIntMin > minInt {
		assertPanics(t, "value exceeds C.int range", func() {
			_ = Int(int(cIntMin - 1))
		})
	}
}

func TestCallWrapsZeroAndNonZeroReturns(t *testing.T) {
	t.Parallel()

	payload := []byte("agent")
	buffer := NewBuffer(len(payload))
	defer buffer.Free()
	buffer.CopyFrom(payload)

	if err := Call(callSumLengthFunction(), buffer.Ptr(), SizeT(len(payload))); err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}

	if err := Call(callFailureFunction()); err == nil {
		t.Fatalf("expected error, got nil")
	} else if !errors.Is(err, syscall.Errno(13)) {
		t.Fatalf("expected errno error for return code 13, got %T %v", err, err)
	}
}

func TestCallRejectsUnsupportedInputs(t *testing.T) {
	t.Parallel()

	assertPanics(t, "nil function pointer", func() {
		_ = Call(nil)
	})

	assertPanics(t, "unsupported argument count", func() {
		_ = Call(callFailureFunction(), 1, 2, 3, 4, 5, 6, 7)
	})
}

func TestCallSupportsSixArguments(t *testing.T) {
	t.Parallel()

	if err := Call(callSixArgumentFunction(), 10, 11, 12, 13, 14, 15); err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
}

func TestCallSupportsBufferArgument(t *testing.T) {
	t.Parallel()

	buffer := NewBuffer(4)
	defer buffer.Free()

	if err := Call(callBufferArgumentFunction(), buffer); err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
}

func TestErrnoMapping(t *testing.T) {
	t.Parallel()

	if err := Errno(0); err != nil {
		t.Fatalf("expected nil error for zero, got %v", err)
	}

	if err := Errno(2); err == nil {
		t.Fatal("expected non-nil error for non-zero errno")
	} else if !errors.Is(err, syscall.Errno(2)) {
		t.Fatalf("expected error type %v, got %T %v", syscall.Errno(2), err, err)
	}
}

func TestWithErrnoReturnsBothResultAndError(t *testing.T) {
	t.Parallel()

	result, err := callWithErrnoZero()
	if result != 0 {
		t.Fatalf("expected result 0, got %d", result)
	}
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	result, err = callWithErrnoFailure()
	if result != 2 {
		t.Fatalf("expected result 2, got %d", result)
	}
	if err == nil {
		t.Fatal("expected error, got nil")
	} else if !errors.Is(err, syscall.Errno(2)) {
		t.Fatalf("expected errno error, got %v", err)
	}
}
