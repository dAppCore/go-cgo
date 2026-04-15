package cgo

import (
	"runtime"
	"testing"
)

func mustPanic(t *testing.T, want string, fn func()) {
	t.Helper()

	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("expected panic %q", want)
		}
		got := r.(string)
		if got != want {
			t.Fatalf("panic = %q, want %q", got, want)
		}
	}()

	fn()
}

func TestBuffer_NewBuffer_Good(t *testing.T) {
	buffer := NewBuffer(8)
	if buffer == nil {
		t.Fatal("buffer is nil")
	}
	if buffer.IsFreed() {
		t.Fatal("buffer should not start freed")
	}
	if buffer.Len() != 8 {
		t.Fatalf("Len() = %d, want 8", buffer.Len())
	}
	if got := len(buffer.Bytes()); got != 8 {
		t.Fatalf("Bytes len = %d, want 8", got)
	}
	if buffer.Ptr() == nil {
		t.Fatal("Ptr() is nil")
	}
	buffer.Free()
}

func TestBuffer_NewBuffer_Bad(t *testing.T) {
	mustPanic(t, "cgo.NewBuffer: size must be non-negative", func() {
		_ = NewBuffer(-1)
	})
}

func TestBuffer_Free_Good(t *testing.T) {
	buffer := NewBuffer(4)
	buffer.Free()
	if !buffer.IsFreed() {
		t.Fatal("buffer should be freed")
	}
	mustPanic(t, "cgo.Buffer.Free: double-free detected", func() {
		buffer.Free()
	})
}

func TestBuffer_Free_Nil_Good(t *testing.T) {
	var buffer *Buffer
	buffer.Free()
}

func TestBuffer_Free_Ugly(t *testing.T) {
	buffer := NewBuffer(4)
	buffer.Free()
	buffer = nil
	runtime.GC()
	runtime.GC()
}

func TestBuffer_CopyFrom_Good(t *testing.T) {
	buffer := NewBuffer(3)
	defer buffer.Free()

	if copied := buffer.CopyFrom([]byte("abcd")); copied != 3 {
		t.Fatalf("CopyFrom copied = %d, want 3", copied)
	}
	if got := string(buffer.Bytes()); got != "abc" {
		t.Fatalf("buffer contents = %q, want %q", got, "abc")
	}
}

func TestBuffer_CopyFrom_Bad(t *testing.T) {
	buffer := NewBuffer(1)
	buffer.Free()
	mustPanic(t, "cgo.Buffer: use-after-free detected", func() {
		buffer.CopyFrom([]byte("x"))
	})
}

func TestBuffer_Bytes_Good(t *testing.T) {
	buffer := NewBuffer(2)
	defer buffer.Free()

	b := buffer.Bytes()
	if len(b) != 2 {
		t.Fatalf("len(Bytes()) = %d, want 2", len(b))
	}
}

func TestBuffer_Bytes_Ugly(t *testing.T) {
	buffer := NewBuffer(2)
	buffer.Free()
	mustPanic(t, "cgo.Buffer: use-after-free detected", func() {
		_ = buffer.Bytes()
	})
}

func TestBuffer_Ptr_Good(t *testing.T) {
	buffer := NewBuffer(2)
	defer buffer.Free()

	if buffer.Ptr() == nil {
		t.Fatal("Ptr() is nil")
	}
}

func TestBuffer_Close_Good(t *testing.T) {
	buffer := NewBuffer(1)
	if err := buffer.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if !buffer.IsFreed() {
		t.Fatal("buffer should be freed after Close()")
	}
}

func TestBuffer_Len_Good(t *testing.T) {
	buffer := NewBuffer(5)
	defer buffer.Free()

	if got := buffer.Len(); got != 5 {
		t.Fatalf("Len() = %d, want 5", got)
	}
}

func TestBuffer_IsFreed_Good(t *testing.T) {
	if !((*Buffer)(nil)).IsFreed() {
		t.Fatal("nil buffer should report freed")
	}

	buffer := NewBuffer(1)
	if buffer.IsFreed() {
		t.Fatal("buffer should not be freed yet")
	}
	buffer.Free()
	if !buffer.IsFreed() {
		t.Fatal("buffer should report freed")
	}
}
