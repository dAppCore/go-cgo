package cgo

import (
	"runtime"
	"testing"
)

func TestScope_NewScope_Good(t *testing.T) {
	scope := NewScope()
	if scope == nil {
		t.Fatal("scope is nil")
	}
	if scope.IsFreed() {
		t.Fatal("scope should not start freed")
	}
	scope.FreeAll()
}

func TestScope_Buffer_Good(t *testing.T) {
	scope := NewScope()
	defer scope.FreeAll()

	buffer := scope.Buffer(4)
	if buffer == nil {
		t.Fatal("buffer is nil")
	}
}

func TestScope_CString_Good(t *testing.T) {
	scope := NewScope()
	defer scope.FreeAll()

	cString := scope.CString("hello")
	if cString == nil {
		t.Fatal("CString is nil")
	}
	if got := GoString(cString); got != "hello" {
		t.Fatalf("GoString(CString) = %q, want %q", got, "hello")
	}
}

func TestScope_FreeAll_Good(t *testing.T) {
	scope := NewScope()
	buffer := scope.Buffer(2)
	cString := scope.CString("hi")

	scope.FreeAll()
	if !scope.IsFreed() {
		t.Fatal("scope should be freed")
	}
	if !buffer.IsFreed() {
		t.Fatal("buffer should be freed with scope")
	}
	if cString == nil {
		t.Fatal("CString is nil")
	}
}

func TestScope_FreeAll_Bad(t *testing.T) {
	scope := NewScope()
	scope.FreeAll()
	mustPanic(t, "cgo.Scope.FreeAll: double-free detected", func() {
		scope.FreeAll()
	})
}

func TestScope_FreeAll_Ugly(t *testing.T) {
	scope := NewScope()
	scope.Buffer(1)
	scope.CString("x")
	scope.FreeAll()
	scope = nil
	runtime.GC()
	runtime.GC()
}

func TestScope_Close_Good(t *testing.T) {
	scope := NewScope()
	if err := scope.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
}

func TestScope_IsFreed_Good(t *testing.T) {
	if !((*Scope)(nil)).IsFreed() {
		t.Fatal("nil scope should report freed")
	}

	scope := NewScope()
	if scope.IsFreed() {
		t.Fatal("scope should not be freed yet")
	}
	scope.FreeAll()
	if !scope.IsFreed() {
		t.Fatal("scope should report freed")
	}
}
