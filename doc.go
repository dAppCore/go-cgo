// Package cgo is a focused harness for safe Go-to-C interop helpers.
//
// Common usage:
//
//	buffer := NewBuffer(32)
//	n := buffer.CopyFrom([]byte("agent"))
//	defer buffer.Free()
//
//	err := Call(unsafe.Pointer(C.some_function), buffer.Ptr(), SizeT(n))
//	if err != nil {
//		// handle mapped C errno
//	}
//
//	scope := NewScope()
//	defer scope.Close()
//	cString := scope.CString("hello")
//	buffer := scope.Buffer(16)
//
//	if err := Call(unsafe.Pointer(C.another_function), cString, buffer.Ptr()); err != nil {
//		// scope.Close() will release C string + buffer
//	}
package cgo

