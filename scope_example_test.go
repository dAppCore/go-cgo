package cgo_test

import (
	. "dappco.re/go"
	cgo "dappco.re/go/cgo"
)

// ExampleScope shows a short-lived allocation group for C calls that need more
// than one temporary value.
func ExampleScope() {
	scope := cgo.NewScope()
	defer scope.FreeAll()

	buffer := scope.Buffer(2)
	buffer.CopyFrom([]byte("go"))
	Println(string(buffer.Bytes()))
	// Output: go
}

// ExampleNewScope creates an allocation owner that releases every registered C
// string and buffer together.
func ExampleNewScope() {
	scope := cgo.NewScope()
	defer scope.FreeAll()

	Println(scope.IsFreed())
	Println(scope.Buffer(1).Len())
	// Output:
	// false
	// 1
}

// ExampleScope_Buffer registers a managed Buffer with the scope so FreeAll
// cleans it up with the rest of the call's temporary memory.
func ExampleScope_Buffer() {
	scope := cgo.NewScope()
	defer scope.FreeAll()

	buffer := scope.Buffer(4)
	Println(buffer.Len())
	Println(buffer.IsFreed())
	// Output:
	// 4
	// false
}

// ExampleScope_CString registers a managed C string with the scope and converts
// it back to Go for inspection.
func ExampleScope_CString() {
	scope := cgo.NewScope()
	defer scope.FreeAll()

	cString := scope.CString("hello")
	Println(cgo.GoString(cString))
	// Output: hello
}

// ExampleScope_FreeAll releases every allocation created by the scope.
func ExampleScope_FreeAll() {
	scope := cgo.NewScope()
	buffer := scope.Buffer(1)

	scope.FreeAll()
	Println(scope.IsFreed())
	Println(buffer.IsFreed())
	// Output:
	// true
	// true
}

// ExampleScope_Close uses the io.Closer-compatible cleanup hook for callers
// that collect resources behind a common interface.
func ExampleScope_Close() {
	scope := cgo.NewScope()
	buffer := scope.Buffer(1)

	err := scope.Close()
	Println(err == nil)
	Println(buffer.IsFreed())
	// Output:
	// true
	// true
}

// ExampleScope_IsFreed reports whether the scope has already released its
// owned allocations.
func ExampleScope_IsFreed() {
	scope := cgo.NewScope()
	Println(scope.IsFreed())

	scope.FreeAll()
	Println(scope.IsFreed())
	// Output:
	// false
	// true
}
