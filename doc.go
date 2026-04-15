// Package cgo provides a standard CGo harness for Core packages.
//
// It centralises C-backed buffer allocation, scoped cleanup, safe string
// conversion, and function-pointer call helpers with panic guards for
// double-free and use-after-free bugs.
package cgo
