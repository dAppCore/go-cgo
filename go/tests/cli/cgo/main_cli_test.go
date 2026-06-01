package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// AX-10: the CLI artifact is validated by building and running the driver
// binary, mirroring the Taskfile path (go build ./tests/cli/cgo then run it).
// The cgo driver self-validates every package primitive — Call dispatch,
// CString/GoString round-trip, Scope-managed Buffer cleanup and Errno
// mapping — and exits non-zero on the first failed assertion, so a clean
// exit is the artifact's pass signal.

// buildDriver compiles the CLI driver once per test into a temp binary.
//
//	bin := buildDriver(t)
//	out, err := exec.Command(bin).CombinedOutput()
func buildDriver(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "core-cgo-driver")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	cmd.Env = os.Environ()
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build driver: %v\n%s", err, out)
	}
	return bin
}

// TestCLIDriver_Run_Good builds the driver and runs it; a clean exit means
// every cgo primitive the driver exercises behaved as specified.
func TestCLIDriver_Run_Good(t *testing.T) {
	bin := buildDriver(t)
	cmd := exec.Command(bin)
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("cgo driver exited non-zero: %v\n%s", err, out)
	}
	if len(out) != 0 {
		t.Fatalf("cgo driver wrote unexpected output on success: %q", out)
	}
}

// TestCLIDriver_Run_Repeatable runs the built driver twice to confirm the
// validation sequence is deterministic — Scope/Buffer cleanup and the
// CString allocation tracker must not leave global state that changes the
// second run's outcome.
func TestCLIDriver_Run_Repeatable(t *testing.T) {
	bin := buildDriver(t)
	for i := 0; i < 2; i++ {
		cmd := exec.Command(bin)
		cmd.Env = os.Environ()
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("cgo driver run %d exited non-zero: %v\n%s", i+1, err, out)
		}
	}
}
