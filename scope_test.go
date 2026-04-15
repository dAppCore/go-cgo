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

func TestScope_Buffer_Bad(t *testing.T) {
	scope := NewScope()
	scope.FreeAll()
	mustPanic(t, "cgo.Scope.Buffer: scope is already freed", func() {
		_ = scope.Buffer(1)
	})
}

func TestScope_Buffer_Nil_Bad(t *testing.T) {
	mustPanic(t, "cgo.Scope.Buffer: scope is already freed", func() {
		_ = ((*Scope)(nil)).Buffer(1)
	})
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

func TestScope_CString_Bad(t *testing.T) {
	scope := NewScope()
	scope.FreeAll()
	mustPanic(t, "cgo.Scope.CString: scope is already freed", func() {
		_ = scope.CString("x")
	})
}

func TestScope_CString_Nil_Bad(t *testing.T) {
	mustPanic(t, "cgo.Scope.CString: scope is already freed", func() {
		_ = ((*Scope)(nil)).CString("x")
	})
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

func TestScope_FreeAll_Nil_Good(t *testing.T) {
	var scope *Scope
	scope.FreeAll()
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
	buffer := scope.Buffer(1)
	if err := scope.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if !scope.IsFreed() {
		t.Fatal("scope should be freed after Close()")
	}
	if !buffer.IsFreed() {
		t.Fatal("buffer should be freed by Close()")
	}
}

func TestScope_Close_Nil_Good(t *testing.T) {
	var scope *Scope
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
