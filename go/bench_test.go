package cgo

import (
	"testing"
	"unsafe"
)

// BenchmarkCString_Free measures the round-trip of allocating a C string and
// freeing it — the hot pair every cgo caller hits when passing a Go string
// into C. Watch allocs/op: tracker maps and the syscall itself dominate.
func BenchmarkCString_Free(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		p := CString("hello world")
		Free(unsafe.Pointer(p))
	}
}

// BenchmarkCString_Short exercises the small-string path in isolation.
func BenchmarkCString_Short(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		p := CString("x")
		Free(unsafe.Pointer(p))
	}
}

// BenchmarkCall_NoArgs is the variadic-call floor: arg slice allocation +
// dispatch overhead with zero arguments. Any allocs here come purely from
// the Call() machinery itself.
func BenchmarkCall_NoArgs(b *testing.B) {
	testCallReset()
	fn := testCallPtr0()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Call(fn)
	}
}

// BenchmarkCall_2Args is the typical small-arity shape. Measures the
// make([]uintptr, n) plus encodeCallArg cost.
func BenchmarkCall_2Args(b *testing.B) {
	testCallReset()
	fn := testCallPtr2()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Call(fn, uintptr(1), uintptr(2))
	}
}

// BenchmarkCall_18Args exercises the upper arity bound — the largest
// stack-array we'd reach for if we replaced make([]uintptr, n).
func BenchmarkCall_18Args(b *testing.B) {
	testCallReset()
	fn := testCallPtr18()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Call(fn,
			uintptr(1), uintptr(2), uintptr(3), uintptr(4), uintptr(5),
			uintptr(6), uintptr(7), uintptr(8), uintptr(9), uintptr(10),
			uintptr(11), uintptr(12), uintptr(13), uintptr(14), uintptr(15),
			uintptr(16), uintptr(17), uintptr(18),
		)
	}
}

// BenchmarkNewBuffer exercises malloc + finalizer registration.
func BenchmarkNewBuffer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := NewBuffer(64)
		buf.Free()
	}
}

// BenchmarkScope_CString measures the common scoped-allocation idiom that
// every cgo consumer reaches for when building a path/name argument.
func BenchmarkScope_CString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s := NewScope()
		_ = s.CString("hello")
		s.FreeAll()
	}
}
