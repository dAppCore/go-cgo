package cgo

import (
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
	}
}

func TestCallRejectsUnsupportedInputs(t *testing.T) {
	t.Parallel()

	assertPanics(t, "nil function pointer", func() {
		_ = Call(nil)
	})

	assertPanics(t, "unsupported argument count", func() {
		_ = Call(callFailureFunction(), 1, 2, 3, 4)
	})
}
