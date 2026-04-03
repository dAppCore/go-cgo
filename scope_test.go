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

func TestScopeIsFreedTracksLifecycle(t *testing.T) {
	t.Parallel()

	scope := NewScope()
	if scope.IsFreed() {
		t.Fatal("expected new scope to be active")
	}

	scope.FreeAll()
	if !scope.IsFreed() {
		t.Fatal("expected scope to report freed after FreeAll")
	}

	var nilScope *Scope
	if !nilScope.IsFreed() {
		t.Fatal("expected nil scope to be treated as freed")
	}
}

func TestScopeCloseIsSafeAndIdempotent(t *testing.T) {
	t.Parallel()

	var nilScope *Scope
	if err := nilScope.Close(); err != nil {
		t.Fatalf("expected nil scope close to return nil, got %v", err)
	}

	scope := NewScope()
	buffer := scope.Buffer(4)
	cString := scope.CString("agent")
	if cString == nil {
		t.Fatal("expected CString allocation")
	}
	if copied := buffer.CopyFrom([]byte("go")); copied != 2 {
		t.Fatalf("expected 2 bytes copied, got %d", copied)
	}

	if err := scope.Close(); err != nil {
		t.Fatalf("expected close to return nil, got %v", err)
	}

	scope.Close()
	if !buffer.IsFreed() {
		t.Fatal("expected buffer to be freed after close")
	}
}
