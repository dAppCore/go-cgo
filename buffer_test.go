package cgo

import (
	"testing"
	"unsafe"
)

func TestBufferLifecycleAndCopy(t *testing.T) {
	t.Parallel()

	const capacity = 8
	buffer := NewBuffer(capacity)
	defer buffer.Free()
	if buffer.Ptr() == nil {
		t.Fatal("expected non-nil pointer for non-zero allocation")
	}

	copied := buffer.CopyFrom([]byte{1, 2, 3, 4, 5})
	if copied != 5 {
		t.Fatalf("expected 5 bytes copied, got %d", copied)
	}

	if got := buffer.Len(); got != capacity {
		t.Fatalf("expected buffer length %d, got %d", capacity, got)
	}

	if got, want := buffer.Bytes()[:5], []byte{1, 2, 3, 4, 5}; string(got) != string(want) {
		t.Fatalf("expected bytes %v, got %v", want, got)
	}
}

func TestBufferCopyClipsToCapacity(t *testing.T) {
	t.Parallel()

	buffer := NewBuffer(2)
	defer buffer.Free()

	copied := buffer.CopyFrom([]byte("abc"))
	if copied != 2 {
		t.Fatalf("expected 2 bytes copied, got %d", copied)
	}
	if got, want := buffer.Bytes(), []byte("ab"); string(got) != string(want) {
		t.Fatalf("expected bytes %q, got %q", string(want), string(got))
	}
}

func TestBufferDoubleFreePanics(t *testing.T) {
	t.Parallel()

	buffer := NewBuffer(1)
	buffer.Free()
	assertPanics(t, "double-free", func() {
		buffer.Free()
	})
}

func TestBufferCloseReleasesMemory(t *testing.T) {
	t.Parallel()

	buffer := NewBuffer(4)
	if err := buffer.Close(); err != nil {
		t.Fatalf("expected close to return nil, got %v", err)
	}
	if !buffer.IsFreed() {
		t.Fatal("expected buffer to be freed by Close")
	}
}

func TestBufferClosePanicsOnDoubleClose(t *testing.T) {
	t.Parallel()

	buffer := NewBuffer(1)
	if err := buffer.Close(); err != nil {
		t.Fatalf("expected close to return nil, got %v", err)
	}

	assertPanics(t, "double-free", func() {
		_ = buffer.Close()
	})
}

func TestBufferCloseNil(t *testing.T) {
	t.Parallel()

	var buffer *Buffer
	if err := buffer.Close(); err != nil {
		t.Fatalf("expected nil buffer close to return nil, got %v", err)
	}
}

func TestBufferUseAfterFreePanics(t *testing.T) {
	t.Parallel()

	buffer := NewBuffer(4)
	buffer.Free()

	assertPanics(t, "use-after-free", func() {
		_ = buffer.Len()
	})
	assertPanics(t, "use-after-free", func() {
		buffer.CopyFrom([]byte("x"))
	})
	assertPanics(t, "use-after-free", func() {
		_ = buffer.Bytes()
	})
	assertPanics(t, "use-after-free", func() {
		_ = buffer.Ptr()
	})
}

func TestBufferIsFreedIsNilSafe(t *testing.T) {
	t.Parallel()

	var nilBuffer *Buffer
	if !nilBuffer.IsFreed() {
		t.Fatal("expected nil buffer to report freed")
	}
}

func TestCStringRoundTrip(t *testing.T) {
	t.Parallel()

	source := "hello"
	cString := CString(source)
	defer Free(unsafe.Pointer(cString))

	converted := GoString(cString)
	if converted != source {
		t.Fatalf("expected %q, got %q", source, converted)
	}
}

func TestGoStringNilIsEmpty(t *testing.T) {
	t.Parallel()

	got := GoString(nil)
	if got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
}

func TestFreeNilDoesNotPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Free(nil) panicked: %v", r)
		}
	}()
	Free(nil)
}

func assertPanics(t *testing.T, want string, fn func()) {
	t.Helper()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for %s, got none", want)
		}
	}()
	fn()
}
