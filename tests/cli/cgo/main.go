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
	"syscall"
	"unsafe"

	core "dappco.re/go"
	corecgo "dappco.re/go/cgo"
)

func main() {
	if r := run(); !r.OK {
		core.Print(core.Stderr(), "%s", r.Error())
		core.Exit(1)
	}
}

func run() core.Result {
	if r := corecgo.Call(C.ax10_noop_ptr()); !r.OK {
		return core.Fail(core.Errorf("call noop: %w", r.Value.(error)))
	}

	var sum uintptr
	if r := corecgo.Call(C.ax10_sum_ptr(), 2, corecgo.SizeT(3), unsafe.Pointer(&sum)); !r.OK {
		return core.Fail(core.Errorf("call sum: %w", r.Value.(error)))
	}
	if sum != 5 {
		return core.Fail(core.Errorf("call sum = %d, want 5", sum))
	}

	cString := corecgo.CString("ax-10")
	defer corecgo.Free(unsafe.Pointer(cString))
	if got := corecgo.GoString(cString); got != "ax-10" {
		return core.Fail(core.Errorf("go string = %q, want %q", got, "ax-10"))
	}

	scope := corecgo.NewScope()
	input := scope.Buffer(5)
	output := scope.Buffer(5)

	if copied := input.CopyFrom([]byte("cgo!!")); copied != input.Len() {
		return core.Fail(core.Errorf("buffer copy copied %d bytes, want %d", copied, input.Len()))
	}
	if r := corecgo.Call(C.ax10_copy_ptr(), input, output, corecgo.SizeT(input.Len())); !r.OK {
		return core.Fail(core.Errorf("call copy: %w", r.Value.(error)))
	}
	if string(output.Bytes()) != "cgo!!" {
		return core.Fail(core.Errorf("copied buffer = %q, want %q", output.Bytes(), "cgo!!"))
	}

	scope.FreeAll()
	if !scope.IsFreed() {
		return core.Fail(core.NewError("scope is not marked freed"))
	}

	errno := corecgo.Errno(corecgo.Int(int(syscall.EINVAL)))
	if errno.OK || !core.Is(errno.Value.(error), syscall.EINVAL) {
		return core.Fail(core.Errorf("errno = %v, want %v", errno.Value, syscall.EINVAL))
	}

	return core.Ok(nil)
}
