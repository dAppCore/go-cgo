package cgo

import (
	"testing"
	"unsafe"
)

// allocBudget asserts the average allocations-per-call for fn stays at or
// below ceiling. Ceilings are set just above the post-optimisation steady
// state, so any regression that adds even one alloc fails the gate.
func allocBudget(t *testing.T, name string, ceiling float64, fn func()) {
	t.Helper()
	got := testing.AllocsPerRun(50, fn)
	if got > ceiling {
		t.Fatalf("%s: alloc budget exceeded: got %.2f allocs/op, ceiling %.2f", name, got, ceiling)
	}
}

// TestAllocBudget_Call_18Args locks the stack-array win in call.go.
// Pre-fix: 1 alloc (make([]uintptr, 18) escapes to heap).
// Post-fix: 0 allocs (var scratch [18]uintptr stays on stack).
//
// Ceiling = 0 to detect any regression that resurrects the heap slice.
func TestAllocBudget_Call_18Args(t *testing.T) {
	testCallReset()
	fn := testCallPtr18()
	allocBudget(t, "Call(18 args)", 0, func() {
		_ = Call(fn,
			uintptr(1), uintptr(2), uintptr(3), uintptr(4), uintptr(5),
			uintptr(6), uintptr(7), uintptr(8), uintptr(9), uintptr(10),
			uintptr(11), uintptr(12), uintptr(13), uintptr(14), uintptr(15),
			uintptr(16), uintptr(17), uintptr(18),
		)
	})
}

// TestAllocBudget_Call_NoArgs locks the no-args floor — must stay zero.
func TestAllocBudget_Call_NoArgs(t *testing.T) {
	testCallReset()
	fn := testCallPtr0()
	allocBudget(t, "Call(0 args)", 0, func() {
		_ = Call(fn)
	})
}

// TestAllocBudget_Call_2Args locks the small-arity floor — must stay zero.
func TestAllocBudget_Call_2Args(t *testing.T) {
	testCallReset()
	fn := testCallPtr2()
	allocBudget(t, "Call(2 args)", 0, func() {
		_ = Call(fn, uintptr(1), uintptr(2))
	})
}

// TestAllocBudget_CString_Free locks the current CString+Free round-trip
// at its established baseline. Reducing this further requires structural
// rework of the sync.Map tracker pair (see CString docs); the budget gate
// catches accidental regressions inside that envelope.
//
// Baseline: 5 allocs (malloc + &cStringAllocation + sync.Map entries).
func TestAllocBudget_CString_Free(t *testing.T) {
	allocBudget(t, "CString+Free", 5, func() {
		p := CString("hello world")
		Free(unsafe.Pointer(p))
	})
}

// TestAllocBudget_NewBuffer locks Buffer construction at 1 alloc
// (the *Buffer struct itself — finalizer-bound so heap is required).
func TestAllocBudget_NewBuffer(t *testing.T) {
	allocBudget(t, "NewBuffer", 1, func() {
		buf := NewBuffer(64)
		buf.Free()
	})
}

// TestAllocBudget_NewBufferUnmanaged locks the no-finalizer variant at
// 1 alloc — same struct cost as NewBuffer, with the per-call latency
// drop coming from the eliminated runtime.SetFinalizer call, not from
// fewer allocations.
func TestAllocBudget_NewBufferUnmanaged(t *testing.T) {
	allocBudget(t, "NewBufferUnmanaged", 1, func() {
		buf := NewBufferUnmanaged(64)
		buf.Free()
	})
}

// TestAllocBudget_Free_Raw locks the dedup win on Free's raw-pointer path.
// Pre-fix: 2 allocs (LoadOrStore + redundant Store in freedPointers).
// Post-fix: 1 alloc (LoadOrStore alone — Store was a no-op re-write).
//
// Ceiling = 1 to detect resurrection of the redundant Store call.
func TestAllocBudget_Free_Raw(t *testing.T) {
	allocBudget(t, "Free (raw pointer)", 1, func() {
		ptr := testMalloc(64)
		Free(ptr)
	})
}

// TestAllocBudget_Scope_Buffer locks the SBO win on Scope.Buffer.
// Pre-SBO: 3 allocs (Scope + Buffer + first append into nil slice).
// Post-SBO: 2 allocs (Scope + Buffer; append fits in inline array).
//
// Ceiling = 2 to detect regression that resurrects the first-append heap
// allocation (e.g. inline arrays accidentally bypassed in NewScope).
func TestAllocBudget_Scope_Buffer(t *testing.T) {
	allocBudget(t, "Scope.Buffer", 2, func() {
		s := NewScope()
		_ = s.Buffer(64)
		s.FreeAll()
	})
}

// TestAllocBudget_Scope_Empty locks the bare scope lifecycle at 1 alloc
// (the *Scope struct itself, finalizer-bound).
func TestAllocBudget_Scope_Empty(t *testing.T) {
	allocBudget(t, "Scope (empty)", 1, func() {
		s := NewScope()
		s.FreeAll()
	})
}

// TestAllocBudget_AdoptCString locks the C→Go error-message transport at
// its measured baseline. AdoptCString allocates the Go string (1) + reuses
// the CString tracker entries on Free (5 from the matching CString side).
//
// Baseline: 6 allocs (1 Go string copy + 5 cgo.Free tracker round-trip).
// Ceiling = 6 to detect regressions in the tracker-aware Free path or
// the C.GoString allocation shape.
func TestAllocBudget_AdoptCString(t *testing.T) {
	allocBudget(t, "AdoptCString", 6, func() {
		p := CString("error: kernel launch failed")
		_ = AdoptCString(unsafe.Pointer(p))
	})
}

// TestAllocBudget_CStringPtr_Free locks the Go→C string transport at the
// CString+Free baseline. CStringPtr is a thin wrapper that returns the
// CString result as unsafe.Pointer for cross-package transport — it adds
// no allocs over the underlying CString+Free pair.
//
// Baseline: 5 allocs (matches CString+Free; CStringPtr is a zero-cost
// re-cast). Ceiling = 5 to detect accidental introduction of any tracker
// or transport overhead inside the wrapper.
func TestAllocBudget_CStringPtr_Free(t *testing.T) {
	allocBudget(t, "CStringPtr+Free", 5, func() {
		p := CStringPtr("hello world")
		Free(p)
	})
}

// TestAllocBudget_Buffer_CopyFrom locks the H2D copy hot path at zero
// allocs — the shape every kernel launch hits per Host→Device transfer.
// Buffer.CopyFrom is a pure `copy()` over the unsafe.Slice header pinned
// at NewBuffer time; nothing on this path should ever allocate.
//
// Baseline: 0 allocs (pure memmove). Ceiling = 0 to catch any future
// edit that materialises a temporary slice or escapes src to the heap.
func TestAllocBudget_Buffer_CopyFrom(t *testing.T) {
	buf := NewBuffer(1024)
	defer buf.Free()
	src := make([]byte, 1024)
	allocBudget(t, "Buffer.CopyFrom", 0, func() {
		_ = buf.CopyFrom(src)
	})
}

// TestAllocBudget_Buffer_Ptr locks the kernel-arg-encoding accessor at
// zero allocs. Buffer.Ptr returns the cached unsafe.Pointer from the
// receiver; the only path-side cost is the assertNotFreed atomic load.
// Any alloc on this path would be a regression.
//
// Baseline: 0 allocs. Ceiling = 0.
func TestAllocBudget_Buffer_Ptr(t *testing.T) {
	buf := NewBuffer(64)
	defer buf.Free()
	allocBudget(t, "Buffer.Ptr", 0, func() {
		_ = buf.Ptr()
	})
}

// TestAllocBudget_SizeT locks the bounds-checked Go int → C.size_t
// conversion at zero allocs. The non-overflowing path is a constant
// switch + cast; only the panic branches would allocate, and those
// are not on the hot path.
//
// Baseline: 0 allocs. Ceiling = 0.
func TestAllocBudget_SizeT(t *testing.T) {
	allocBudget(t, "SizeT", 0, func() {
		_ = SizeT(1024)
	})
}

// TestAllocBudget_Errno locks the bare rc → Result mapping at zero
// allocs on the success path (rc == 0). Errno is the floor underneath
// WithErrno and is hit on every cgo return-code translation.
//
// Baseline: 0 allocs (core.Ok of an int is value-stack only). Ceiling
// = 0 to catch any future edit that boxes the int into an interface
// on the hot success path.
func TestAllocBudget_Errno(t *testing.T) {
	allocBudget(t, "Errno", 0, func() {
		_ = Errno(0)
	})
}

// TestAllocBudget_WithErrno locks the lambda-wrapped C-call shape at
// zero allocs on the success path. The fn-as-closure must not escape;
// inlining + escape analysis turn the call site into a direct dispatch.
//
// Baseline: 0 allocs (closure stays on the stack, success Result is
// value-stack only). Ceiling = 0 to catch any future edit that lets
// the closure escape to the heap on the hot kernel-invoke path.
func TestAllocBudget_WithErrno(t *testing.T) {
	allocBudget(t, "WithErrno", 0, func() {
		_ = WithErrno(func() testCInt { return 0 })
	})
}

// TestAllocBudget_Scope_CString locks the scope-managed Go→C string
// idiom — the typical kernel-launch shape that pairs a path/name with
// scope cleanup. Pays the underlying CString tracker cost (5 allocs)
// + the scope struct + inline-array SBO covers the first append.
//
// Baseline: 6 allocs (1 Scope struct + 5 CString round-trip; the
// strings inline-array absorbs the first append). Ceiling = 6 to
// detect regression in the SBO path or accidental tracker bloat.
func TestAllocBudget_Scope_CString(t *testing.T) {
	allocBudget(t, "Scope.CString", 6, func() {
		s := NewScope()
		_ = s.CString("hello")
		s.FreeAll()
	})
}

// TestAllocBudget_PinIn locks the scope-managed slice-pin hot path at
// its measured baseline — the shape every weight-tensor / async-kernel
// handoff hits. Allocates the *Scope struct + the *PinnedView; the
// pins inline-array SBO absorbs the first append at no extra alloc.
//
// Baseline: 2 allocs (1 *Scope + 1 *core.PinnedView). Ceiling = 2 to
// detect regression in the SBO path or any new bookkeeping inside
// core.PinSlice / scope.pins capture.
func TestAllocBudget_PinIn(t *testing.T) {
	slice := make([]float32, 64)
	allocBudget(t, "PinIn", 2, func() {
		s := NewScope()
		_ = PinIn(s, slice)
		s.FreeAll()
	})
}

// TestAllocBudget_AdoptCStringN locks the length-known sibling of
// AdoptCString at parity with the null-terminated path. Same shape:
// 1 Go string copy + 5 tracker round-trip from the matching CString.
//
// Baseline: 6 allocs. Ceiling = 6 to detect any divergence between the
// AdoptCString and AdoptCStringN paths (both should route Free through
// the same tracker-aware idiom).
func TestAllocBudget_AdoptCStringN(t *testing.T) {
	const payload = "error: kernel launch failed"
	allocBudget(t, "AdoptCStringN", 6, func() {
		p := CString(payload)
		_ = AdoptCStringN(unsafe.Pointer(p), len(payload))
	})
}

// TestAllocBudget_Int locks the bounds-checked Go int → C.int
// conversion at zero allocs. Same shape as SizeT — non-overflow path
// is a constant switch + cast, panic branches off the hot path.
//
// Baseline: 0 allocs. Ceiling = 0.
func TestAllocBudget_Int(t *testing.T) {
	allocBudget(t, "Int", 0, func() {
		_ = Int(1024)
	})
}
