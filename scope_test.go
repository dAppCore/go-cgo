package cgo

import "testing"

func TestScopeFreesAllResources(t *testing.T) {
	t.Parallel()

	scope := NewScope()
	buffer := scope.Buffer(4)
	cString := scope.CString("agent")
	if cString == nil {
		t.Fatal("expected CString allocation")
	}

	if copied := buffer.CopyFrom([]byte("go")); copied != 2 {
		t.Fatalf("expected 2 bytes copied, got %d", copied)
	}

	scope.FreeAll()
	scope.FreeAll() // idempotent

	if !buffer.IsFreed() {
		t.Fatal("expected buffer to be freed")
	}
}

func TestScopePanicsAfterFreeAll(t *testing.T) {
	t.Parallel()

	scope := NewScope()
	scope.FreeAll()

	assertPanics(t, "scope is already freed", func() {
		scope.Buffer(1)
	})
	assertPanics(t, "scope is already freed", func() {
		scope.CString("nope")
	})
}
