package cgo

import "syscall"

func TestStringConversion_WithErrno_Good(t *T) {
	rc, err := testWithErrno(0)

	AssertNoError(t, err)
	AssertEqual(t, 0, rc)
}

func TestStringConversion_WithErrno_Bad(t *T) {
	AssertPanics(t, func() {
		_, _ = WithErrno(nil)
	})
	AssertNotPanics(t, func() {
		_, _ = testWithErrno(0)
	})
}

func TestStringConversion_WithErrno_Ugly(t *T) {
	rc, err := testWithErrno(2)

	AssertEqual(t, 2, rc)
	AssertErrorIs(t, err, syscall.Errno(2))
}
