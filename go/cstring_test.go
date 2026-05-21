package cgo

// AdoptCString tests — exercise the package-internal helper without
// importing C in the test file (Go forbids cgo in test files). The
// CString function in this package allocates real C memory we can
// hand to Adopt, so we reuse it to construct test inputs.

func TestAdoptCString_Good(t *T) {
	cStr := CString("hello, mlx")
	got := AdoptCString(cStr)
	AssertEqual(t, "hello, mlx", got)
}

func TestAdoptCString_NilSafe(t *T) {
	got := AdoptCString(nil)
	AssertEqual(t, "", got)
}

func TestAdoptCString_Empty(t *T) {
	cStr := CString("")
	got := AdoptCString(cStr)
	AssertEqual(t, "", got)
}

func TestAdoptCStringN_Good(t *T) {
	// CString allocates "abcd\0" (5 bytes); read only 4 to test
	// partial-buffer adoption.
	cStr := CString("abcd")
	got := AdoptCStringN(cStr, 4)
	AssertEqual(t, "abcd", got)
}

func TestAdoptCStringN_NilSafe(t *T) {
	got := AdoptCStringN(nil, 4)
	AssertEqual(t, "", got)
}

func TestAdoptCStringN_ZeroLen(t *T) {
	cStr := CString("ignored")
	got := AdoptCStringN(cStr, 0)
	AssertEqual(t, "", got)
}
