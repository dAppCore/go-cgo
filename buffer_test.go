package cgo

import "testing"

func TestBufferLifecycleAndCopy(t *testing.T) {
	t.Parallel()

	const capacity = 8
	buffer := NewBuffer(capacity)
	defer buffer.Free()

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

func assertPanics(t *testing.T, want string, fn func()) {
	t.Helper()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for %s, got none", want)
		}
	}()
	fn()
}
