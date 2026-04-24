package main

/*
#include <stdint.h>
#include <string.h>

static int ax10_noop(void) {
	return 0;
}

static int ax10_sum(uintptr_t a, uintptr_t b, uintptr_t out) {
	uintptr_t *target = (uintptr_t*)out;
	*target = a + b;
	return 0;
}

static int ax10_copy(uintptr_t src, uintptr_t dst, uintptr_t n) {
	memcpy((void*)dst, (void*)src, (size_t)n);
	return 0;
}

static void* ax10_noop_ptr(void) { return (void*)ax10_noop; }
static void* ax10_sum_ptr(void) { return (void*)ax10_sum; }
static void* ax10_copy_ptr(void) { return (void*)ax10_copy; }
*/
import "C"

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	corecgo "dappco.re/go/cgo"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if err := corecgo.Call(C.ax10_noop_ptr()); err != nil {
		return fmt.Errorf("call noop: %w", err)
	}

	var sum uintptr
	if err := corecgo.Call(C.ax10_sum_ptr(), 2, corecgo.SizeT(3), unsafe.Pointer(&sum)); err != nil {
		return fmt.Errorf("call sum: %w", err)
	}
	if sum != 5 {
		return fmt.Errorf("call sum = %d, want 5", sum)
	}

	cString := corecgo.CString("ax-10")
	defer corecgo.Free(unsafe.Pointer(cString))
	if got := corecgo.GoString(cString); got != "ax-10" {
		return fmt.Errorf("go string = %q, want %q", got, "ax-10")
	}

	scope := corecgo.NewScope()
	input := scope.Buffer(5)
	output := scope.Buffer(5)

	if copied := input.CopyFrom([]byte("cgo!!")); copied != input.Len() {
		return fmt.Errorf("buffer copy copied %d bytes, want %d", copied, input.Len())
	}
	if err := corecgo.Call(C.ax10_copy_ptr(), input, output, corecgo.SizeT(input.Len())); err != nil {
		return fmt.Errorf("call copy: %w", err)
	}
	if !bytes.Equal(output.Bytes(), []byte("cgo!!")) {
		return fmt.Errorf("copied buffer = %q, want %q", output.Bytes(), "cgo!!")
	}

	scope.FreeAll()
	if !scope.IsFreed() {
		return errors.New("scope is not marked freed")
	}

	if err := corecgo.Errno(corecgo.Int(int(syscall.EINVAL))); !errors.Is(err, syscall.EINVAL) {
		return fmt.Errorf("errno = %v, want %v", err, syscall.EINVAL)
	}

	return nil
}
