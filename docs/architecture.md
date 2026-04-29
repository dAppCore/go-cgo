# Architecture

The package is organized around the problems that appear at a Go/C boundary:
stable memory, explicit ownership, type conversion, and error translation.

## Memory Ownership

`Buffer` owns one C allocation and exposes three views over it:

- `Bytes` for Go-side reads and writes.
- `Ptr` for C APIs that need a raw pointer.
- `Len` for the fixed allocation size.

`Free` and `Close` mark the buffer as released. Any later access through
`Bytes`, `Ptr`, or `Len` panics, which turns use-after-free bugs into visible
test failures instead of silent memory corruption.

`Scope` owns a group of temporary allocations. A call path can create buffers
and C strings under a scope, defer `FreeAll`, and avoid threading individual
cleanup calls through every branch. `Scope` also detects double-free attempts in
normal use.

## Conversion Layer

`SizeT` and `Int` make range checks explicit before Go values cross into C.
`CString` allocates a NUL-terminated C string and registers it so repeated
`Free` calls on the same pointer are harmless. `GoString` treats nil C strings
as empty strings, which gives callers a simple safe default for optional C
returns.

`Errno` and `WithErrno` adapt C's integer return convention to Core's `Result`
flow. Zero means success; non-zero values become failed Results carrying
`syscall.Errno`.

## Function Pointer Calls

`Call` is the dispatcher for C function pointers. It accepts up to 18
pointer-sized arguments, converts supported Go values into `uintptr`, invokes a
small C shim for the matching arity, and returns the `Errno(rc)` Result.

Supported argument shapes include raw pointers, byte slices, `*Buffer`, C
strings, integer widths, and uintptr-like values. Unsupported types and nil
function pointers panic with package-specific messages so invalid bindings fail
early.
