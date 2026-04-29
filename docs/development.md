# Development

This repository requires a working Go toolchain with cgo enabled and a C
compiler available on `PATH`. CI uses `CGO_ENABLED=1` and installs a build
toolchain before running tests.

## Local Workflow

Run commands with workspace mode disabled so this module is checked on its own:

```sh
GOWORK=off go mod tidy
GOWORK=off go vet ./...
GOWORK=off go test -count=1 ./...
```

Then check formatting and compliance:

```sh
gofmt -l .
bash /Users/snider/Code/core/go/tests/cli/v090-upgrade/audit.sh .
```

The audit must end with `verdict: COMPLIANT`.

## Test Shape

Each production file has a matching `_test.go` file with
`Test<File>_<Symbol>_{Good,Bad,Ugly}` triplets. Tests use Core assertions and
exercise the named symbol directly.

Each production file also has a matching `_example_test.go` file. Examples must
be runnable, must include accurate `// Output:` comments, and must print through
`Println` or `Printf` from `. "dappco.re/go"`.

Go does not support importing `C` from test files. Shared cgo test fixtures live
in ordinary package test support files, and external examples use exported
helpers or runnable guard paths instead of importing `C`.

## Contribution Notes

Keep ownership behavior explicit. If a change allocates memory, the cleanup path
should be visible in the same function or owned by `Scope`. If a change crosses
the C boundary, prefer checked conversion helpers over ad hoc casts.

Do not add dependencies for convenience. This package exists to keep the cgo
surface small and predictable for downstream Core packages.
