package cgo_test

import (
	. "github.com/dappcore/go"
	cgo "github.com/dappcore/go-cgo/go"
)

// ExampleCall_panic shows the guard that prevents dispatch through a nil C
// function pointer. Real C call sites pass a pointer obtained from C code.
func ExampleCall_panic() {
	defer func() {
		Println(recover())
	}()

	_ = cgo.Call(nil)
	// Output:
	// cgo.Call: function pointer is nil
}
