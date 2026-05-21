# `cgo_pinned_view.hpp` — C++23 mdspan companion to `core.PinnedView`

This header provides the C++ side of the zero-copy Go→C tensor handoff
primitive. The Go side (`core.PinnedView` + `cgo.PinIn`) pins a slice
so its first-element pointer stays at a stable address; the C++ side
(this header) wraps that pointer with `std::mdspan` so kernels can
walk the memory as a typed multi-dimensional array.

The combination is the killer pattern from IDEAS.md — zero-copy
graph injection from Go-allocated buffers (model weights, KV cache
snapshots, mapped `.mp4` state) directly into C++ kernels, with the
GC prevented from moving the backing memory for the duration of the
borrow.

## Why mdspan + Pinner

Three things compose:

1. **`runtime.Pinner`** (Go 1.21+) — keeps the Go-owned backing array
   at a fixed address during the C call. Without it, the Go runtime
   would be free to relocate the buffer between cgo entries; with
   it, the C side can retain the pointer across multiple kernels.

2. **The C ABI** — pointer + shape (`size_t[]`) + strides
   (`ptrdiff_t[]`) is the lowest-common-denominator transport that
   crosses the cgo boundary without invoking C++ ABI surprises.

3. **`std::mdspan`** (C++23) — non-owning, multi-dimensional, strided
   view over a raw pointer. Zero-overhead wrapper; the compiler folds
   the layout-stride mapping into pointer arithmetic at the kernel
   call site.

Combined, you get: Go slice on the heap → Pin once → pass pointer +
shape ints to C → reconstruct mdspan view per kernel call → typed
indexing that compiles to the same instructions a plain C array
would generate. No memcpy from Go-space to C-space. No reflective
shape lookup at the C boundary. No layout assumption bake-in (mdspan
respects the strides you pass).

## Consumer pattern (from a cgo file in go-mlx or similar)

```go
package metal

/*
#cgo CXXFLAGS: -std=gnu++23
#cgo CPPFLAGS: -I${SRCDIR}/../../external/go-cgo/go

#include <stddef.h>
#include <stdint.h>
#include "cgo_pinned_view.hpp"

extern "C" void apply_layer_norm_2d(
    const float* in_data,
    float* out_data,
    size_t rows,
    size_t cols,
    ptrdiff_t row_stride,
    ptrdiff_t col_stride) {
    auto in_view  = lthn::cgo::pinned_view_2d<const float>(
        in_data, rows, cols, row_stride, col_stride);
    auto out_view = lthn::cgo::pinned_view_2d<float>(
        out_data, rows, cols, row_stride, col_stride);
    for (std::size_t i = 0; i < in_view.extent(0); ++i) {
        float sum = 0;
        for (std::size_t j = 0; j < in_view.extent(1); ++j)
            sum += in_view[i, j];
        float mean = sum / static_cast<float>(cols);
        for (std::size_t j = 0; j < in_view.extent(1); ++j)
            out_view[i, j] = in_view[i, j] - mean;
    }
}
*/
import "C"

import (
    core "dappco.re/go"
    "unsafe"
)

func LayerNormPinned(in, out []float32, rows, cols int) {
    var inView, outView core.PinnedView
    core.PinSlice(in, &inView)
    core.PinSlice(out, &outView)
    defer inView.Release()
    defer outView.Release()

    C.apply_layer_norm_2d(
        (*C.float)(inView.Ptr()),
        (*C.float)(outView.Ptr()),
        C.size_t(rows), C.size_t(cols),
        C.ptrdiff_t(cols), C.ptrdiff_t(1),
    )
}
```

The include path (`-I${SRCDIR}/../../external/go-cgo/go`) assumes
the consumer has `dappco.re/go/cgo` available as a git submodule at
`external/go-cgo` — the canonical layout go-mlx and go-inference
already follow. For Go-module-only consumers (no submodule), copy
`cgo_pinned_view.hpp` into the consumer's own cgo source tree; the
header is intentionally standalone (no link symbols, no companion
.cpp file).

## Why no Go-side wrapper

A Go function that takes `unsafe.Pointer + shape ints` and calls C
would force the C++ helper into go-cgo's package. That ties the
kernel API to go-cgo's compilation unit, which means every new
kernel shape (1-D, 2-D, 3-D, 4-D, every dtype) accretes here.

Instead, consumers write their own kernel functions in their own
cgo blocks and `#include "cgo_pinned_view.hpp"` for the mdspan
helpers. The header is the substrate; the kernels live with their
consumers. This keeps go-cgo's surface small and avoids the
template instantiation explosion that would otherwise land here.

## Available helpers

- `pinned_view_1d<T>(data, length, stride = 1)` — token streams,
  embedding rows, sample buffers.
- `pinned_view_2d<T>(data, rows, cols, row_stride, col_stride)` —
  attention scores, embedding matrices.
- `pinned_view_3d<T>(data, d0, d1, d2, s0, s1, s2)` —
  `[batch, seq, dim]` hidden states.
- `pinned_view_4d<T>(data, d0, d1, d2, d3, s0..s3)` —
  `[batch, heads, seq, head_dim]` KV cache, `[layers, heads, seq, dim]`
  MLX storage layouts.
- `contiguous_strides_2d/3d/4d(...)` — compute row-major strides
  from a shape when the Go buffer is contiguous (the common case).
- All strides in **elements**, not bytes — matches mdspan's
  `layout_stride` convention.

## Safety contract

Inherits from `core.PinnedView`:

- The Go slice must NOT contain Go pointers (cgo rules — pinning
  panics in race mode if violated).
- The pinned slice must outlive every C++ use of the mdspan view.
  `defer view.Release()` after every Pin handles this; missing
  release leaks the pin until process exit.
- The strides passed to the mdspan helpers must match the actual
  layout of the buffer. Wrong strides produce wrong reads, not
  crashes — verify with a small test before shipping a kernel.

## Testing

Round-trip and bench tests live with the consumer (where the cgo
C++ blocks actually compile against this header). Test files in
go-cgo itself can't use cgo (Go restriction) — see the consumer's
`pkg/.../pinned_view_test.go` for the canonical pattern.
