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
