package cgo_test

import (
	"reflect"
	"unsafe"

	. "dappco.re/go"
	cgo "dappco.re/go/cgo"
)

// ExampleSizeT converts Go sizes before passing them to C APIs with size_t
// parameters.
func ExampleSizeT() {
	size := cgo.SizeT(8)
	Println(uint64(size))
	// Output: 8
}

// ExampleInt converts non-negative Go integers before passing them to C APIs
// with int parameters.
func ExampleInt() {
	value := cgo.Int(3)
	Println(int(value))
	// Output: 3
}

// ExampleErrno maps the C convention of zero success and non-zero errno values
// into Go's error flow.
func ExampleErrno() {
	Println(cgo.Errno(cgo.Int(0)) == nil)
	Println(cgo.Errno(cgo.Int(2)) != nil)
	// Output:
	// true
	// true
}

// ExampleWithErrno wraps a C-style integer return function and exposes both the
// raw return code and the Go error value. Reflection keeps this external
// example runnable without importing C from a test file.
func ExampleWithErrno() {
	fnType := reflect.TypeOf(cgo.WithErrno).In(0)
	fn := reflect.MakeFunc(fnType, func(_ []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf(cgo.Int(0))}
	})

	out := reflect.ValueOf(cgo.WithErrno).Call([]reflect.Value{fn})
	Println(out[0].Interface())
	Println(out[1].IsNil())
	// Output:
	// 0
	// true
}

// ExampleGoString converts a NUL-terminated C string pointer into a Go string.
func ExampleGoString() {
	cString := cgo.CString("hello")
	defer cgo.Free(unsafe.Pointer(cString))

	Println(cgo.GoString(cString))
	// Output: hello
}

// ExampleCString allocates a C string that must later be released with Free.
func ExampleCString() {
	cString := cgo.CString("hello")
	defer cgo.Free(unsafe.Pointer(cString))

	Println(cgo.GoString(cString))
	// Output: hello
}

// ExampleFree releases memory allocated by CString and tolerates a repeated
// cleanup call for the same tracked pointer.
func ExampleFree() {
	cString := cgo.CString("release")
	Println(cgo.GoString(cString))

	cgo.Free(unsafe.Pointer(cString))
	cgo.Free(unsafe.Pointer(cString))
	Println("freed")
	// Output:
	// release
	// freed
}
