package cgo

import (
	"syscall"
	"unsafe"
)

func TestStringConversion_SizeT_Good(t *T) {
	got := SizeT(3)

	AssertEqual(t, uint64(3), uint64(got))
	AssertNotPanics(t, func() {
		_ = SizeT(1)
	})
}

func TestStringConversion_SizeT_Bad(t *T) {
	AssertPanicsWithError(t, "negative values are not representable", func() {
		_ = SizeT(-1)
	})
	AssertNotPanics(t, func() {
		_ = SizeT(0)
	})
}

func TestStringConversion_SizeT_Ugly(t *T) {
	got := SizeT(0)

	AssertEqual(t, uint64(0), uint64(got))
	AssertNotPanics(t, func() {
		_ = SizeT(int(^uint(0) >> 1))
	})
}

func TestStringConversion_Int_Good(t *T) {
	got := Int(4)

	AssertEqual(t, 4, int(got))
	AssertNotPanics(t, func() {
		_ = Int(1)
	})
}

func TestStringConversion_Int_Bad(t *T) {
	AssertPanicsWithError(t, "value exceeds C.int range", func() {
		_ = Int(-1)
	})
	AssertNotPanics(t, func() {
		_ = Int(0)
	})
}

func TestStringConversion_Int_Ugly(t *T) {
	got := Int(0)

	AssertEqual(t, 0, int(got))
	AssertNotPanics(t, func() {
		_ = Int(int(^uint32(0) >> 1))
	})
}

func TestStringConversion_Errno_Good(t *T) {
	err := Errno(0)

	AssertNoError(t, err)
	AssertNil(t, err)
}

func TestStringConversion_Errno_Bad(t *T) {
	err := Errno(5)

	AssertError(t, err)
	AssertErrorIs(t, err, syscall.Errno(5))
}

func TestStringConversion_Errno_Ugly(t *T) {
	err := Errno(-1)

	AssertError(t, err)
	AssertNotEqual(t, "", err.Error())
}

func TestStringConversion_WithErrno_Good(t *T) {
	rc, err := WithErrno(func() testCInt { return 0 })

	AssertNoError(t, err)
	AssertEqual(t, 0, rc)
}

func TestStringConversion_WithErrno_Bad(t *T) {
	AssertPanics(t, func() {
		_, _ = WithErrno(nil)
	})
}

func TestStringConversion_WithErrno_Ugly(t *T) {
	rc, err := WithErrno(func() testCInt { return 2 })

	AssertEqual(t, 2, rc)
	AssertErrorIs(t, err, syscall.Errno(2))
}

func TestStringConversion_GoString_Good(t *T) {
	ptr := CString("hello")
	defer Free(unsafe.Pointer(ptr))

	got := GoString(ptr)
	AssertEqual(t, "hello", got)
	AssertNotNil(t, ptr)
}

func TestStringConversion_GoString_Bad(t *T) {
	got := GoString(nil)

	AssertEqual(t, "", got)
	AssertNotPanics(t, func() {
		_ = GoString(nil)
	})
}

func TestStringConversion_GoString_Ugly(t *T) {
	ptr := CString("hello\x00ignored")
	defer Free(unsafe.Pointer(ptr))

	got := GoString(ptr)
	AssertEqual(t, "hello", got)
}

func TestStringConversion_CString_Good(t *T) {
	ptr := CString("hello")
	defer Free(unsafe.Pointer(ptr))

	AssertNotNil(t, ptr)
	AssertEqual(t, "hello", GoString(ptr))
}

func TestStringConversion_CString_Bad(t *T) {
	ptr := CString("")
	defer Free(unsafe.Pointer(ptr))

	AssertNotNil(t, ptr)
	AssertEqual(t, "", GoString(ptr))
}

func TestStringConversion_CString_Ugly(t *T) {
	ptr := CString("go\x00cgo")
	defer Free(unsafe.Pointer(ptr))

	AssertNotNil(t, ptr)
	AssertEqual(t, "go", GoString(ptr))
}

func TestStringConversion_Free_Good(t *T) {
	ptr := CString("hello")

	AssertNotNil(t, ptr)
	Free(unsafe.Pointer(ptr))
	AssertNotPanics(t, func() {
		Free(unsafe.Pointer(ptr))
	})
}

func TestStringConversion_Free_Bad(t *T) {
	AssertNotPanics(t, func() {
		Free(nil)
	})
	AssertEqual(t, "", GoString(nil))
}

func TestStringConversion_Free_Ugly(t *T) {
	ptr := testMalloc(8)
	AssertNotNil(t, ptr)

	Free(ptr)
	AssertNotPanics(t, func() {
		Free(ptr)
	})
}
