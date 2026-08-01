package cgo

/*
#include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"

	core "dappco.re/go"
)

// scopeInlineCap is the small-buffer-optimisation capacity for the per-kind
// tracking arrays. The common scope shape is "one or two of each" (a buffer
// + a path string + a pinned weight tensor in a kernel launch); inline
// storage at this size lets the first ~4 appends per kind stay on the
// scope's own struct, avoiding the heap alloc that fresh nil-slice growth
// would otherwise incur.
const scopeInlineCap = 4

// Scope tracks multiple C allocations and releases them together.
//
//	scope := NewScope()
//	defer scope.FreeAll()
//	buffer := scope.Buffer(32)
//	cString := scope.CString("hello")
type Scope struct {
	lock          sync.Mutex
	buffers       []*Buffer
	strings       []unsafe.Pointer
	pins          []*core.PinnedView
	buffersInline [scopeInlineCap]*Buffer
	stringsInline [scopeInlineCap]unsafe.Pointer
	pinsInline    [scopeInlineCap]*core.PinnedView
	freed         atomic.Bool
}

// NewScope creates a grouped allocator for temporary C memory.
//
//	scope := NewScope()
//	defer scope.FreeAll()
func NewScope() *Scope {
	scope := &Scope{}
	// Slice headers point into the inline arrays (len=0, cap=scopeInlineCap).
	// Appends up to scopeInlineCap stay on the Scope struct itself; growth
	// past that falls back to standard heap-backed slice doubling.
	scope.buffers = scope.buffersInline[:0:scopeInlineCap]
	scope.strings = scope.stringsInline[:0:scopeInlineCap]
	scope.pins = scope.pinsInline[:0:scopeInlineCap]
	runtime.SetFinalizer(scope, func(owned *Scope) {
		owned.freeAll(true)
	})
	return scope
}

// Buffer allocates a managed buffer and registers it for cleanup.
//
//	buffer := scope.Buffer(64)
func (s *Scope) Buffer(size int) *Buffer {
	if s == nil {
		panic("cgo.Scope.Buffer: scope is already freed")
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	if s.freed.Load() {
		panic("cgo.Scope.Buffer: scope is already freed")
	}

	// NewBufferUnmanaged: scope's freeAll already drains s.buffers via
	// buffer.free(true), and scope itself has a finalizer covering the
	// caller-forgot path. The Buffer's own finalizer is redundant here
	// — skipping it shaves the SetFinalizer cost per scope.Buffer call.
	buffer := NewBufferUnmanaged(size)
	s.buffers = append(s.buffers, buffer)
	return buffer
}

// CString allocates a managed C string and registers it for cleanup.
//
//	cString := scope.CString("hello")
func (s *Scope) CString(value string) *C.char {
	if s == nil {
		panic("cgo.Scope.CString: scope is already freed")
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	if s.freed.Load() {
		panic("cgo.Scope.CString: scope is already freed")
	}

	cString := CString(value)
	s.strings = append(s.strings, unsafe.Pointer(cString))
	return cString
}

// FreeAll releases every allocation created under this scope.
//
//	scope := NewScope()
//	defer scope.FreeAll()
func (s *Scope) FreeAll() {
	s.freeAll(false)
}

// Close releases every allocation in the scope and reports cleanup through a
// Core Result.
//
//	scope := NewScope()
//	defer scope.Close()
func (s *Scope) Close() core.Result {
	s.FreeAll()
	return core.Ok(nil)
}

// IsFreed reports whether FreeAll has been called.
//
//	if scope.IsFreed() {
//		// scope is inactive
//	}
func (s *Scope) IsFreed() bool {
	if s == nil {
		return true
	}
	return s.freed.Load()
}

func (s *Scope) freeAll(noPanic bool) bool {
	if s == nil {
		return false
	}

	if !s.freed.CompareAndSwap(false, true) {
		if noPanic {
			return false
		}
		panic("cgo.Scope.FreeAll: double-free detected")
	}
	runtime.SetFinalizer(s, nil)

	s.lock.Lock()
	buffers := s.buffers
	strings := s.strings
	pins := s.pins
	s.buffers = nil
	s.strings = nil
	s.pins = nil
	s.lock.Unlock()

	for _, buffer := range buffers {
		if buffer != nil {
			buffer.free(true)
		}
	}

	for _, pointer := range strings {
		Free(pointer)
	}

	for _, pin := range pins {
		pin.Release()
	}
	return true
}

// PinIn pins slice's backing array under scope's lifetime — the pin
// is released when scope.FreeAll runs. The returned *core.PinnedView
// can be passed to C via Ptr()/Len()/Bytes() and remains valid until
// FreeAll. Use this for slices C may retain across more than one
// cgo invocation (async kernels, mlx_array data slots, weight
// tensors) that the surrounding scope already manages.
//
//	scope := cgo.NewScope()
//	defer scope.FreeAll()
//	weights := cgo.PinIn(scope, modelWeights)
//	C.kernel_load(weights.Ptr(), C.size_t(weights.Bytes()))
//
// For one-shot calls where C consumes the pointer during the call,
// pass &slice[0] directly — the cgo runtime already prevents GC
// movement for the call's duration without a pin.
func PinIn[T any](scope *Scope, slice []T) *core.PinnedView {
	if scope == nil {
		panic("cgo.PinIn: scope is nil")
	}
	scope.lock.Lock()
	defer scope.lock.Unlock()
	if scope.freed.Load() {
		panic("cgo.PinIn: scope is already freed")
	}
	view := &core.PinnedView{}
	core.PinSlice(slice, view)
	if view.Active() {
		scope.pins = append(scope.pins, view)
	}
	return view
}
