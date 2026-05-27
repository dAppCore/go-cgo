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

// BenchmarkScope_Buffer measures the second composing path — scope-managed
// Buffer allocation, the typical kernel-launch shape.
func BenchmarkScope_Buffer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s := NewScope()
		_ = s.Buffer(64)
		s.FreeAll()
	}
}

// BenchmarkScope_Empty measures the bare-scope lifecycle cost — what every
// scoped cgo block pays even before allocating anything.
func BenchmarkScope_Empty(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s := NewScope()
		s.FreeAll()
	}
}

// BenchmarkBuffer_CopyFrom measures the data-transfer hot path on a reused
// Buffer — the shape every ROCm/MLX/CUDA kernel launch hits per H2D copy.
func BenchmarkBuffer_CopyFrom(b *testing.B) {
	buf := NewBuffer(1024)
	defer buf.Free()
	src := make([]byte, 1024)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = buf.CopyFrom(src)
	}
}

// BenchmarkAdoptCString measures the C→Go string adoption shape used for
// adopting error messages from mlx/sqlite/curl.
func BenchmarkAdoptCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		p := CString("error: kernel launch failed")
		_ = AdoptCString(unsafe.Pointer(p))
	}
}

// BenchmarkWithErrno measures the lambda-wrapped C-call shape used by
// callers that want one-line cgo + Result mapping. Hot per kernel invoke.
func BenchmarkWithErrno(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = WithErrno(func() testCInt { return 0 })
	}
}

// BenchmarkErrno measures the bare rc→Result mapping, the floor underneath
// WithErrno + Call.
func BenchmarkErrno(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Errno(0)
	}
}
