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
