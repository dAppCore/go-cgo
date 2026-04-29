# Agent Notes

This repository is a compact cgo boundary package for the Core Go ecosystem.
New agents should treat it as infrastructure code: small public surface, strict
ownership rules, and compliance enforced by the v0.9.0 audit script.

## Code Structure

- `doc.go` contains the package overview shown by `go doc`.
- `buffer.go` defines `Buffer`, a C-backed byte allocation with `Free`,
  `Close`, `CopyFrom`, `Bytes`, `Ptr`, `Len`, and `IsFreed`.
- `scope.go` defines `Scope`, a grouped owner for temporary buffers and C
  strings. `FreeAll` is the canonical cleanup path and `Close` adapts it to
  `io.Closer`.
- `string_conversion.go` contains C boundary conversions: `SizeT`, `Int`,
  `Errno`, `WithErrno`, `GoString`, `CString`, and `Free`.
- `call.go` contains the function-pointer dispatcher. It accepts pointer-sized
  arguments and converts the integer return code through `Errno`.
- `*_test.go` files contain AX-style triplet tests for each public callable.
- `*_example_test.go` files contain runnable examples with `Println` from
  `dappco.re/go`.
- `tests/cli/cgo/main.go` is the small AX-10 binary scenario used to exercise
  real cgo calls.

## Working Rules

Do not edit `BRIEF.md`; it is an untracked task file. Do not touch `.git/`,
`.codex/`, or any `third_party/` directory.

Keep tests and examples file-aware. A production file named `buffer.go` owns
`buffer_test.go` and `buffer_example_test.go`; do not create monolithic or
version-suffixed test files.

The compliance audit bans direct imports of `fmt`, `errors`, `strings`, `path`,
`path/filepath`, `os`, `os/exec`, `io/ioutil`, `log`, `encoding/json`, and
`bytes` in every Go file, including examples. Use Core wrappers from
`dappco.re/go` instead.

## Verification

Run the full gate before claiming completion:

```sh
GOWORK=off go mod tidy
GOWORK=off go vet ./...
GOWORK=off go test -count=1 ./...
gofmt -l .
bash /Users/snider/Code/core/go/tests/cli/v090-upgrade/audit.sh .
```

The final audit line must be `verdict: COMPLIANT`.
