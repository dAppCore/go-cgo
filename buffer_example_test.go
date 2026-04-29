package cgo_test

import (
	. "dappco.re/go"
	cgo "dappco.re/go/cgo"
)

// ExampleBuffer shows the owned memory wrapper used when a C API needs a stable
// byte buffer for the duration of a call.
func ExampleBuffer() {
	buffer := cgo.NewBuffer(4)
	defer buffer.Free()

	Println(buffer.Len())
	Println(buffer.IsFreed())
	// Output:
	// 4
	// false
}

// ExampleNewBuffer allocates C-backed byte storage and releases it explicitly
// once the call boundary is finished.
func ExampleNewBuffer() {
	buffer := cgo.NewBuffer(3)
	defer buffer.Free()

	Println(buffer.Len())
	Println(len(buffer.Bytes()))
	// Output:
	// 3
	// 3
}

// ExampleBuffer_Free releases the allocation and marks the wrapper unusable for
// later pointer or slice access.
func ExampleBuffer_Free() {
	buffer := cgo.NewBuffer(1)
	buffer.Free()

	Println(buffer.IsFreed())
	// Output: true
}

// ExampleBuffer_Close lets Buffer satisfy io.Closer-style cleanup paths while
// preserving the same ownership semantics as Free.
func ExampleBuffer_Close() {
	buffer := cgo.NewBuffer(1)
	err := buffer.Close()

	Println(err == nil)
	Println(buffer.IsFreed())
	// Output:
	// true
	// true
}

// ExampleBuffer_CopyFrom copies Go bytes into the C-backed allocation and
// returns the number of bytes that fit.
func ExampleBuffer_CopyFrom() {
	buffer := cgo.NewBuffer(3)
	defer buffer.Free()

	n := buffer.CopyFrom([]byte("cgo!"))
	Println(n)
	Println(string(buffer.Bytes()[:n]))
	// Output:
	// 3
	// cgo
}

// ExampleBuffer_Bytes exposes the mutable byte slice backed by the C
// allocation, useful for reading data a C function wrote into the buffer.
func ExampleBuffer_Bytes() {
	buffer := cgo.NewBuffer(2)
	defer buffer.Free()

	bytes := buffer.Bytes()
	bytes[0] = 'g'
	bytes[1] = 'o'
	Println(string(buffer.Bytes()))
	// Output: go
}

// ExampleBuffer_Ptr returns the raw pointer passed to C APIs that accept an
// output or scratch buffer.
func ExampleBuffer_Ptr() {
	buffer := cgo.NewBuffer(4)
	defer buffer.Free()

	Println(buffer.Ptr() != nil)
	// Output: true
}

// ExampleBuffer_Len reports the fixed allocation size, not the number of bytes
// most recently copied.
func ExampleBuffer_Len() {
	buffer := cgo.NewBuffer(5)
	defer buffer.Free()

	buffer.CopyFrom([]byte("go"))
	Println(buffer.Len())
	// Output: 5
}

// ExampleBuffer_IsFreed reports whether the buffer has crossed its cleanup
// boundary.
func ExampleBuffer_IsFreed() {
	buffer := cgo.NewBuffer(1)
	Println(buffer.IsFreed())

	buffer.Free()
	Println(buffer.IsFreed())
	// Output:
	// false
	// true
}
