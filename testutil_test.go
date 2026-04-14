package cgo

import "testing"

func assertPanics(t *testing.T, want string, fn func()) {
	t.Helper()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for %s, got none", want)
		}
	}()

	fn()
}
