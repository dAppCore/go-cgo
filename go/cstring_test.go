package cgo

import "unsafe"

// AdoptCString tests — exercise the package-internal helper without
// importing C in the test file (Go forbids cgo in test files). The
// CString function in this package allocates real C memory we can
// hand to Adopt, so we reuse it to construct test inputs.

func TestAdoptCString_Good(t *T) {
	cStr := CString("hello, mlx")
	got := AdoptCString(unsafe.Pointer(cStr))
	AssertEqual(t, "hello, mlx", got)
}

// TestAdoptCString_ClearsTracker pins the correctness contract: after
// Adopt the package tracker MUST NOT retain a stale `freed=false` entry
// at the freed address. A stale entry causes the next CString or raw
// malloc that reuses that address to land on a freedPointers hit and
// silently skip the C.free — a real-malloc leak that only surfaces
// under address-reuse pressure.
//
// Pre-fix: AdoptCString called C.free directly, leaving the tracker
// untouched. Post-fix: AdoptCString routes through cgo.Free so the
// tracker entry is removed in lockstep with the C-side free.
func TestAdoptCString_ClearsTracker(t *T) {
	cStr := CString("trace me")
	addr := uintptr(unsafe.Pointer(cStr))
	_ = AdoptCString(unsafe.Pointer(cStr))

	if _, stillTracked := cStringAllocations.Load(addr); stillTracked {
		t.Fatalf("AdoptCString left stale cStringAllocations entry at %x; address reuse will see freed=false and let a real malloc leak", addr)
	}
}

// TestAdoptCStringN_ClearsTracker mirrors the contract for the
// length-counted Adopt variant — both paths must route through cgo.Free
// so tracker shed happens in lockstep with the C-side free.
func TestAdoptCStringN_ClearsTracker(t *T) {
	cStr := CString("trace me too")
	addr := uintptr(unsafe.Pointer(cStr))
	_ = AdoptCStringN(unsafe.Pointer(cStr), 8)

	if _, stillTracked := cStringAllocations.Load(addr); stillTracked {
		t.Fatalf("AdoptCStringN left stale cStringAllocations entry at %x; address reuse will see freed=false and let a real malloc leak", addr)
	}
}

// TestAdoptCStringN_ZeroLenClearsTracker locks the early-return branch
// (n <= 0) — that path also calls free, so tracker shed must run too.
func TestAdoptCStringN_ZeroLenClearsTracker(t *T) {
	cStr := CString("zero branch")
	addr := uintptr(unsafe.Pointer(cStr))
	_ = AdoptCStringN(unsafe.Pointer(cStr), 0)

	if _, stillTracked := cStringAllocations.Load(addr); stillTracked {
		t.Fatalf("AdoptCStringN(n=0) left stale cStringAllocations entry at %x; address reuse will see freed=false and let a real malloc leak", addr)
	}
}

func TestAdoptCString_NilSafe(t *T) {
	got := AdoptCString(nil)
	AssertEqual(t, "", got)
}

func TestAdoptCString_Empty(t *T) {
	cStr := CString("")
	got := AdoptCString(unsafe.Pointer(cStr))
	AssertEqual(t, "", got)
}

func TestAdoptCStringN_Good(t *T) {
	// CString allocates "abcd\0" (5 bytes); read only 4 to test
	// partial-buffer adoption.
	cStr := CString("abcd")
	got := AdoptCStringN(unsafe.Pointer(cStr), 4)
	AssertEqual(t, "abcd", got)
}

func TestAdoptCStringN_NilSafe(t *T) {
	got := AdoptCStringN(nil, 4)
	AssertEqual(t, "", got)
}

func TestAdoptCStringN_ZeroLen(t *T) {
	cStr := CString("ignored")
	got := AdoptCStringN(unsafe.Pointer(cStr), 0)
	AssertEqual(t, "", got)
}
