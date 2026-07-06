package cgo

/*
#include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"sync/atomic"
	"unsafe"

	core "dappco.re/go"
)

// Buffer owns byte memory that can be passed safely to C.
//
//	buffer := NewBuffer(16)
//	defer buffer.Free()
//	n := buffer.CopyFrom([]byte("payload"))
//	_ = buffer.Bytes()[:n]
type Buffer struct {
	data         []byte
	length       int
	pointer      unsafe.Pointer
	freed        atomic.Bool
	hasFinalizer bool // true when NewBuffer registered a GC finalizer
}

// NewBuffer allocates a C-backed byte buffer for C interop.
//
//	input := make([]byte, 32)
//	buffer := NewBuffer(len(input))
//	defer buffer.Free()
//
// NewBuffer installs a GC finalizer that calls Free if the caller drops
// the *Buffer without calling Free explicitly — the safe default. For
// hot-loop allocators where Free is guaranteed (defer in the same
// function, or scope-owned), see NewBufferUnmanaged for a ~60 ns faster
// path that skips the finalizer.
func NewBuffer(size int) *Buffer {
	buffer := NewBufferUnmanaged(size)
	runtime.SetFinalizer(buffer, func(owned *Buffer) {
		owned.free(true)
	})
	buffer.hasFinalizer = true
	return buffer
}

// NewBufferUnmanaged allocates a C-backed byte buffer WITHOUT a GC
// finalizer. The caller MUST call Free — either directly or via an
// owning collector — or the underlying malloc leaks. In return, the
// per-call cost drops by ~60 ns (the runtime.SetFinalizer cost on
// NewBuffer's hot path).
//
//	// Safe shape: defer in same scope as construction.
//	buffer := cgo.NewBufferUnmanaged(size)
//	defer buffer.Free()
//	C.kernel_launch(buffer.Ptr(), C.size_t(buffer.Len()))
//
// Safe to use when ANY of these hold:
//   - the *Buffer is paired with a `defer buffer.Free()` in the same
//     function (defer guarantees Free runs even on panic)
//   - the *Buffer is handed to an owning collector that calls Free in
//     its own cleanup (e.g. NewScope().Buffer uses this internally)
//   - the caller has an audited Free path on every code branch
//
// NOT safe when:
//   - the *Buffer is returned from a constructor without finalizer
//     coverage at the OUTERMOST owner
//   - the *Buffer is stored in a long-lived data structure whose
//     lifetime is unclear at allocation time
//   - the surrounding code might panic between alloc and Free without
//     a defer-Free guard
//
// When in doubt, use NewBuffer — the GC safety net is cheap insurance
// at typical (non-hot-loop) call rates.
func NewBufferUnmanaged(size int) *Buffer {
	if size < 0 {
		panic("cgo.NewBuffer: size must be non-negative")
	}

	var pointer unsafe.Pointer
	var data []byte
	if size > 0 {
		cMemory := C.malloc(C.size_t(size))
		if cMemory == nil {
			panic("cgo.NewBuffer: C allocation failed")
		}
		pointer = cMemory
		data = unsafe.Slice((*byte)(pointer), size)
	}

	return &Buffer{
		data:    data,
		length:  size,
		pointer: pointer,
	}
}

// Free releases the pinned memory backing slice and marks the buffer as freed.
//
//	buffer := NewBuffer(8)
//	defer buffer.Free()
func (b *Buffer) Free() {
	b.free(false)
}

// Close releases the buffer and reports cleanup through a Core Result.
//
//	buffer := NewBuffer(16)
//	defer buffer.Close()
func (b *Buffer) Close() core.Result {
	b.Free()
	return core.Ok(nil)
}

// CopyFrom copies bytes from src into the buffer and returns bytes copied.
//
//	buffer := NewBuffer(3)
//	buffer.CopyFrom([]byte("abc"))
func (b *Buffer) CopyFrom(src []byte) int {
	b.assertNotFreed()
	if len(src) == 0 || b.length == 0 {
		return 0
	}

	copied := len(src)
	if copied > b.length {
		copied = b.length
	}

	copy(b.data[:copied], src[:copied])
	return copied
}

// Bytes returns the mutable byte slice backed by the buffer memory.
//
//	buffer := NewBuffer(4)
//	n := buffer.CopyFrom([]byte("go"))
//	_ = buffer.Bytes()[:n]
func (b *Buffer) Bytes() []byte {
	b.assertNotFreed()
	return b.data
}

// Ptr returns the raw pointer to the buffer.
//
//	buffer := NewBuffer(4)
//	_ = buffer.Ptr()
func (b *Buffer) Ptr() unsafe.Pointer {
	b.assertNotFreed()
	return b.pointer
}

// Len returns the allocated byte length of the buffer.
//
//	buffer := NewBuffer(4)
//	if buffer.Len() == 4 {
//		// preallocated 4-byte buffer
//	}
func (b *Buffer) Len() int {
	b.assertNotFreed()
	return b.length
}

// IsFreed reports whether Free has already been called.
//
//	buffer := NewBuffer(4)
//	if buffer.IsFreed() {
//		return
//	}
func (b *Buffer) IsFreed() bool {
	if b == nil {
		return true
	}
	return b.freed.Load()
}

func (b *Buffer) assertNotFreed() {
	if b == nil {
		panic("cgo.Buffer: use-after-free detected: buffer is nil")
	}
	if b.freed.Load() {
		panic("cgo.Buffer: use-after-free detected")
	}
}

func (b *Buffer) free(noPanic bool) bool {
	if b == nil {
		return false
	}

	if !b.freed.CompareAndSwap(false, true) {
		if noPanic {
			return false
		}
		panic("cgo.Buffer.Free: double-free detected")
	}

	// Only clear the finalizer when NewBuffer set one — calling
	// runtime.SetFinalizer(b, nil) on a buffer that never had a finalizer
	// is a ~24 ns no-op on M3 Ultra. Skipping it on unmanaged + scope-
	// owned buffers shaves that cost from every Free.
	if b.hasFinalizer {
		runtime.SetFinalizer(b, nil)
	}
	C.free(b.pointer)
	b.pointer = nil
	b.data = nil
	return true
}
