# go-cgo

`dappco.re/go/cgo` is the Core cgo utility package. It provides the small set of
allocation, conversion, and function-pointer helpers that Core packages use when
they need to cross from Go into C without repeating ownership boilerplate.

The package focuses on four jobs:

- `Buffer` owns C-backed byte memory and exposes it as a Go slice plus a raw
  pointer for C APIs.
- `Scope` groups temporary `Buffer` and C string allocations so one cleanup call
  releases everything from a C interaction.
- `CString`, `GoString`, `SizeT`, `Int`, `Errno`, and `WithErrno` make common C
  boundary conversions explicit.
- `Call` invokes C function pointers with up to 18 pointer-sized arguments and
  maps non-zero return codes into failed Core `Result`s.

## Install

```sh
go get dappco.re/go/cgo
```

The package uses cgo, so local builds need a working C toolchain and
`CGO_ENABLED=1`.

## Quick Start

```go
package main

import (
	"unsafe"

	core "dappco.re/go"
	corecgo "dappco.re/go/cgo"
)

func main() {
	scope := corecgo.NewScope()
	defer scope.FreeAll()

	input := scope.Buffer(5)
	input.CopyFrom([]byte("hello"))

	cString := corecgo.CString("core")
	defer corecgo.Free(unsafe.Pointer(cString))

	core.Println(input.Len())
	core.Println(corecgo.GoString(cString))
}
```

Use `Scope` for temporary values owned by a single call path. Use standalone
`NewBuffer`, `CString`, and `Free` when a value must outlive a local scope or be
managed by a higher-level owner.

## Verification

Before committing changes, run the same compliance gate used by this repository:

```sh
GOWORK=off go mod tidy
GOWORK=off go vet ./...
GOWORK=off go test -count=1 ./...
gofmt -l .
bash /Users/snider/Code/core/go/tests/cli/v090-upgrade/audit.sh .
```

The audit must finish with `verdict: COMPLIANT`.
