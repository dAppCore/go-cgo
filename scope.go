package cgo

/*
#include <stdlib.h>
*/
import "C"

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// Scope tracks multiple allocations and releases them together.
//
// scope := cgo.NewScope()
// defer scope.FreeAll()
// buffer := scope.Buffer(32)
// cString := scope.CString("hello")
type Scope struct {
	lock     sync.Mutex
	buffers  []*Buffer
	strings  []unsafe.Pointer
	freed    atomic.Bool
}

// NewScope creates a zero-value scope for grouped CGo allocations.
func NewScope() *Scope {
	return &Scope{}
}

// Buffer allocates a new managed buffer.
func (s *Scope) Buffer(size int) *Buffer {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.freed.Load() {
		panic("cgo.Scope.Buffer: scope is already freed")
	}

	buffer := NewBuffer(size)
	s.buffers = append(s.buffers, buffer)
	return buffer
}

// CString allocates a managed C string.
func (s *Scope) CString(value string) *C.char {
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
func (s *Scope) FreeAll() {
	if s == nil {
		return
	}

	if !s.freed.CompareAndSwap(false, true) {
		return
	}

	s.lock.Lock()
	buffers := s.buffers
	strings := s.strings
	s.buffers = nil
	s.strings = nil
	s.lock.Unlock()

	for _, buffer := range buffers {
		if buffer != nil && !buffer.IsFreed() {
			buffer.Free()
		}
	}

	for _, pointer := range strings {
		Free(pointer)
	}
}

// IsFreed reports whether FreeAll has been called.
func (s *Scope) IsFreed() bool {
	if s == nil {
		return true
	}
	return s.freed.Load()
}

// Close releases every allocation in the scope and implements io.Closer.
func (s *Scope) Close() error {
	s.FreeAll()
	return nil
}
