package cgo

import (
	"testing"
	"unsafe"
)

func TestStringConversion_GoString_Nil_Good(t *testing.T) {
	if got := GoString(nil); got != "" {
		t.Fatalf("GoString(nil) = %q, want empty string", got)
	}
}

func TestStringConversion_Free_Nil_Good(t *testing.T) {
	Free(nil)
}

func TestStringConversion_CString_Free_Good(t *testing.T) {
	ptr := CString("hello")
	if ptr == nil {
		t.Fatal("CString returned nil")
	}

	if got := GoString(ptr); got != "hello" {
		t.Fatalf("GoString(CString) = %q, want %q", got, "hello")
	}

	Free(unsafe.Pointer(ptr))
	Free(unsafe.Pointer(ptr))
}

func TestStringConversion_Free_GenericMalloc_Good(t *testing.T) {
	ptr := testMalloc(8)
	if ptr == nil {
		t.Fatal("testMalloc returned nil")
	}

	Free(ptr)
	Free(ptr)
}
