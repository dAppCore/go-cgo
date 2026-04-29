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

// Scope tracks multiple C allocations and releases them together.
//
//	scope := NewScope()
//	defer scope.FreeAll()
//	buffer := scope.Buffer(32)
//	cString := scope.CString("hello")
type Scope struct {
	lock    sync.Mutex
	buffers []*Buffer
	strings []unsafe.Pointer
	freed   atomic.Bool
}

// NewScope creates a grouped allocator for temporary C memory.
//
//	scope := NewScope()
//	defer scope.FreeAll()
func NewScope() *Scope {
	scope := &Scope{}
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

	buffer := NewBuffer(size)
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
	s.buffers = nil
	s.strings = nil
	s.lock.Unlock()

	for _, buffer := range buffers {
		if buffer != nil {
			buffer.free(true)
		}
	}

	for _, pointer := range strings {
		Free(pointer)
	}
	return true
}
