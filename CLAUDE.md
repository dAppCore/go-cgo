# Claude Code Conventions

This repository follows the Core v0.9.0 shape. Claude Code agents should keep
changes narrow, preserve the existing file-aware layout, and use the audit
binary as the source of truth for compliance.

## Coding Style

Use Core wrappers instead of banned standard-library imports. In examples,
import `. "dappco.re/go"` and print with `Println` or `Printf`; do not use
`fmt`. Keep package documentation and comments practical: explain ownership,
lifetime, or call-boundary behavior rather than restating symbol names.

Do not add helper abstractions around `Buffer`, `Scope`, or `Call` unless the
public API requires them. Most changes should be local to the source file and
its matching test or example file.

## Tests And Examples

Each production file owns its matching tests:

- `call.go` -> `call_test.go` and `call_example_test.go`
- `buffer.go` -> `buffer_test.go` and `buffer_example_test.go`
- `scope.go` -> `scope_test.go` and `scope_example_test.go`
- `string_conversion.go` -> `string_conversion_test.go` and
  `string_conversion_example_test.go`

Triplet tests use `Test<File>_<Symbol>_{Good,Bad,Ugly}` and real assertions.
Examples use Go's runnable `// Output:` form and should demonstrate concrete
usage or a concrete panic guard.

Go test files cannot import `C`, so external examples that need C-only types
must either use exported conversion helpers or document the panic/guard path in
a runnable way.

## Completion Gate

Before committing or reporting completion, run:

```sh
GOWORK=off go mod tidy
GOWORK=off go vet ./...
GOWORK=off go test -count=1 ./...
gofmt -l .
bash /Users/snider/Code/core/go/tests/cli/v090-upgrade/audit.sh .
```

Commit only after the audit prints `verdict: COMPLIANT`. Do not push unless the
user explicitly asks.
