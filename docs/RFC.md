---
module: dappco.re/go/cgo
repo: core/go-cgo
lang: go
tier: lib
depends:
  - code/core/go
tags:
  - cgo
  - c-interop
  - bindings
  - ffi
---
# go-cgo — Standard CGo Harness

> Shared CGo boilerplate for all Core packages that cross the Go/C boundary.
> One harness, zero use-after-free bugs, every CGo package uses it.

**Module:** `dappco.re/go/cgo`
**Repository:** `core/go-cgo`

---

## 1. Overview

Multiple Core packages use CGo: go-blockchain/crypto, go-mlx, go-rocm. Each rolls its own buffer allocation, string conversion, and error mapping. This package centralises that into tested, safe primitives.

---

## 2. Buffer Management

```go
// NewBuffer allocates a C buffer with automatic cleanup tracking.
//
//   buf := cgo.NewBuffer(32)
//   defer buf.Free()
//   buf.CopyFrom(goSlice)
//   result := buf.Bytes()
type Buffer struct { /* ... */ }

func NewBuffer(size int) *Buffer
func (b *Buffer) Free()
func (b *Buffer) CopyFrom(src []byte) int
func (b *Buffer) Bytes() []byte
func (b *Buffer) Ptr() unsafe.Pointer
func (b *Buffer) Len() int
func (b *Buffer) IsFreed() bool
```

### 2.1 Safety

- Double-free panics with descriptive message
- Use-after-free panics with descriptive message
- `IsFreed()` allows conditional cleanup in complex flows

---

## 3. C Function Calls

```go
// Call wraps a C function call with Go error mapping.
//
//   err := cgo.Call(C.my_function, buf.Ptr(), cgo.SizeT(len))
//   // Returns Go error if C returns non-zero
func Call(fn unsafe.Pointer, args ...interface{}) error

// SizeT converts Go int to C size_t safely.
//
//   n := cgo.SizeT(len(data))
func SizeT(n int) C.size_t

// Int converts Go int to C int safely.
//
//   rc := cgo.Int(returnCode)
func Int(n int) C.int
```

---

## 4. String Conversion

```go
// GoString converts a C string to Go string safely.
//
//   goStr := cgo.GoString(cStr)
func GoString(cs *C.char) string

// CString converts a Go string to C string. Caller must free.
//
//   cStr := cgo.CString(goStr)
//   defer cgo.Free(unsafe.Pointer(cStr))
func CString(s string) *C.char

// Free releases C-allocated memory.
//
//   cgo.Free(ptr)
func Free(ptr unsafe.Pointer)
```

---

## 5. Scoped Allocations

```go
// Scope tracks multiple C allocations and frees them all at once.
//
//   scope := cgo.NewScope()
//   defer scope.FreeAll()
//   buf1 := scope.Buffer(32)
//   buf2 := scope.Buffer(64)
//   str := scope.CString("hello")
//   // All freed on scope.FreeAll()
type Scope struct { /* ... */ }

func NewScope() *Scope
func (s *Scope) Buffer(size int) *Buffer
func (s *Scope) CString(str string) *C.char
func (s *Scope) FreeAll()
```

---

## 6. Error Mapping

```go
// Errno converts a C errno value to a Go error.
//
//   err := cgo.Errno(rc)
func Errno(rc C.int) error

// WithErrno calls a C function and returns the errno as a Go error.
//
//   result, err := cgo.WithErrno(func() C.int {
//       return C.my_function(args...)
//   })
func WithErrno(fn func() C.int) (int, error)
```

---

## 7. Implementation Priority

1. Buffer (NewBuffer, Free, CopyFrom, Bytes, Ptr, Len, IsFreed)
2. String conversion (GoString, CString, Free)
3. Scope (NewScope, Buffer, CString, FreeAll)
4. Call wrapper (Call, SizeT, Int)
5. Error mapping (Errno, WithErrno)
6. Safety panics (double-free, use-after-free)
7. Tests for all of the above

---

## 8. Consumers

| Package | Current Boilerplate | Saves |
|---------|-------------------|-------|
| go-blockchain/crypto | Manual buffer alloc + free | ~50 lines |
| go-mlx | Manual C array management | ~40 lines |
| go-rocm | Manual kernel buffer handling | ~60 lines |
| CorePy (future) | CPython embedding via cgo | ~80 lines |

---

## Changelog

- 2026-04-02: Initial RFC from Charon's request. Standard CGo harness for all Core packages.
