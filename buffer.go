package cgo

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

// Buffer owns byte memory that can be passed safely to C.
//
// buffer := NewBuffer(16)
// defer buffer.Free()
// n := buffer.CopyFrom([]byte("payload"))
// _ = buffer.Bytes()[:n]
type Buffer struct {
	data     []byte
	length   int
	pointer  unsafe.Pointer
	freed    atomic.Bool
	isPinned bool
	pinner   runtime.Pinner
}

// NewBuffer allocates memory and pins it so Ptr can be used across C boundaries.
func NewBuffer(size int) *Buffer {
	if size < 0 {
		panic("cgo.NewBuffer: size must be non-negative")
	}

	data := make([]byte, size)
	buffer := &Buffer{
		data:   data,
		length: size,
	}

	if size > 0 {
		buffer.pinner.Pin(&buffer.data[0])
		buffer.isPinned = true
		buffer.pointer = unsafe.Pointer(&buffer.data[0])
	}

	return buffer
}

// Free releases the pinned memory backing slice and marks the buffer as freed.
func (b *Buffer) Free() {
	if b == nil {
		return
	}

	if !b.freed.CompareAndSwap(false, true) {
		panic("cgo.Buffer.Free: double-free detected")
	}

	if b.isPinned {
		b.pinner.Unpin()
		b.isPinned = false
	}

	b.pointer = nil
	b.data = nil
}

// CopyFrom copies bytes from src into the buffer and returns bytes copied.
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

// Bytes returns a mutable byte slice backed by the buffer memory.
func (b *Buffer) Bytes() []byte {
	b.assertNotFreed()
	return b.data
}

// Ptr returns the raw pointer to the buffer.
func (b *Buffer) Ptr() unsafe.Pointer {
	b.assertNotFreed()
	return b.pointer
}

// Len returns the current buffer length.
func (b *Buffer) Len() int {
	b.assertNotFreed()
	return b.length
}

// IsFreed reports whether Free has already been called.
func (b *Buffer) IsFreed() bool {
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
