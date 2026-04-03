package cgo

/*
#include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

// Buffer owns byte memory that can be passed safely to C.
//
//	buffer := NewBuffer(16)
//	defer buffer.Free()
//	n := buffer.CopyFrom([]byte("payload"))
//	_ = buffer.Bytes()[:n]
type Buffer struct {
	data     []byte
	length   int
	pointer  unsafe.Pointer
	freed    atomic.Bool
}

// NewBuffer allocates a C-backed buffer that can be passed to C safely.
//
//	buffer := NewBuffer(32)
//	defer buffer.Free()
func NewBuffer(size int) *Buffer {
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

	buffer := &Buffer{
		data:    data,
		length:  size,
		pointer: pointer,
	}

	runtime.SetFinalizer(buffer, func(owned *Buffer) {
		owned.Free()
	})

	return buffer
}

// Free releases the pinned memory backing slice and marks the buffer as freed.
//
//	buffer := NewBuffer(8)
//	defer buffer.Free()
func (b *Buffer) Free() {
	if b == nil {
		return
	}

	if !b.freed.CompareAndSwap(false, true) {
		panic("cgo.Buffer.Free: double-free detected")
	}

	C.free(b.pointer)
	b.pointer = nil
	b.data = nil
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
//	call := buffer.Ptr()
func (b *Buffer) Ptr() unsafe.Pointer {
	b.assertNotFreed()
	return b.pointer
}

// Len returns the allocated byte length of the buffer.
//
//	buffer := NewBuffer(4)
//	if buffer.Len() == 4 { ... }
func (b *Buffer) Len() int {
	b.assertNotFreed()
	return b.length
}

// IsFreed reports whether Free has already been called.
//
//	buffer := NewBuffer(4)
//	if buffer.IsFreed() { ... }
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
