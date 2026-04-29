# Documentation

`dappco.re/go/cgo` is the Core package for small, repeatable cgo boundary work.
It keeps memory ownership, string conversion, errno handling, and C function
pointer dispatch in one place so consumer packages do not need to invent their
own conventions for each binding.

Start here:

- [Architecture](architecture.md) explains the package shape and ownership
  model.
- [Development](development.md) explains local build, test, audit, and cgo
  requirements.
- [README](../README.md) is the human-facing quick start.
- [AGENTS](../AGENTS.md) is the handoff guide for automated coding agents.
- [CLAUDE](../CLAUDE.md) records conventions for Claude Code agents.

The public surface is intentionally small. `Buffer` and `Scope` manage lifetime,
`string_conversion.go` handles type conversion, and `Call` is the low-level
function-pointer bridge.
