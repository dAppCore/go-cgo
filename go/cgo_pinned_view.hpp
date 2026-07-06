// SPDX-License-Identifier: EUPL-1.2
//
// Companion to dappco.re/go/cgo PinIn (and core.PinnedView). Provides
// a C++23 mdspan-shaped view over Go-pinned memory, so kernels written
// in C++ can read pinned Go buffers as typed multi-dimensional arrays
// without copying or re-wrapping.
//
// The Go side pins a slice (runtime.Pinner), pushes the raw pointer
// + shape + strides across the cgo boundary, and the C++ side
// reconstructs an mdspan view on demand. The pin guarantees the
// pointer stays valid for the duration of every C++ kernel that
// borrows it; Release happens on the Go side after every kernel that
// holds the pointer has returned.
//
// Header-only by design — including this in a cgo C++ preamble adds
// only the mdspan wrapping machinery, no link-time symbols. The
// shape/stride parameters are passed as plain C arrays so the C ABI
// stays simple and Go-side conversion is trivial (a single
// runtime.Pinner pin + an unsafe.Pointer).
//
// Usage from a consumer's cgo block:
//
//   /*
//   #cgo CXXFLAGS: -std=gnu++23
//   #include "cgo_pinned_view.hpp"
//
//   static double sum_2d(const double* data, size_t rows, size_t cols,
//                        ptrdiff_t row_stride, ptrdiff_t col_stride) {
//     auto view = lthn::cgo::pinned_view_2d<const double>(
//         data, rows, cols, row_stride, col_stride);
//     double s = 0;
//     for (std::size_t i = 0; i < view.extent(0); ++i)
//       for (std::size_t j = 0; j < view.extent(1); ++j)
//         s += view[i, j];
//     return s;
//   }
//   */
//   import "C"
//
//   var view core.PinnedView
//   core.PinSlice(matrix, &view)
//   defer view.Release()
//   sum := C.sum_2d((*C.double)(view.Ptr()), C.size_t(rows), C.size_t(cols),
//                   C.ptrdiff_t(rowStride), C.ptrdiff_t(colStride))
//
// Strides are in *elements*, not bytes — matches mdspan's
// layout_stride convention. Callers compute strides from the Go
// slice's logical shape; for contiguous row-major: row_stride =
// cols, col_stride = 1.

#pragma once

#include <mdspan>
#include <cstddef>
#include <array>

namespace lthn::cgo {

// pinned_view_1d wraps a 1-D pinned buffer as a strided mdspan.
// Use for token streams, embedding rows, sample buffers — anywhere
// the Go side hands over a flat slice that the C++ kernel walks
// linearly.
template <typename T>
[[nodiscard]] auto pinned_view_1d(
    T* data,
    std::size_t length,
    std::ptrdiff_t stride = 1) noexcept {
    using ext_t = std::dextents<std::size_t, 1>;
    using map_t = std::layout_stride::mapping<ext_t>;
    return std::mdspan<T, ext_t, std::layout_stride>(
        data,
        map_t{ext_t{length}, std::array<std::size_t, 1>{static_cast<std::size_t>(stride)}});
}

// pinned_view_2d wraps a 2-D pinned buffer (e.g. [rows × cols]
// embedding matrix, attention scores) with explicit per-dim strides.
template <typename T>
[[nodiscard]] auto pinned_view_2d(
    T* data,
    std::size_t rows,
    std::size_t cols,
    std::ptrdiff_t row_stride,
    std::ptrdiff_t col_stride) noexcept {
    using ext_t = std::dextents<std::size_t, 2>;
    using map_t = std::layout_stride::mapping<ext_t>;
    return std::mdspan<T, ext_t, std::layout_stride>(
        data,
        map_t{ext_t{rows, cols},
              std::array<std::size_t, 2>{
                  static_cast<std::size_t>(row_stride),
                  static_cast<std::size_t>(col_stride)}});
}

// pinned_view_3d wraps a 3-D pinned buffer (e.g. [batch × seq × dim]
// hidden states) with explicit per-dim strides.
template <typename T>
[[nodiscard]] auto pinned_view_3d(
    T* data,
    std::size_t d0,
    std::size_t d1,
    std::size_t d2,
    std::ptrdiff_t s0,
    std::ptrdiff_t s1,
    std::ptrdiff_t s2) noexcept {
    using ext_t = std::dextents<std::size_t, 3>;
    using map_t = std::layout_stride::mapping<ext_t>;
    return std::mdspan<T, ext_t, std::layout_stride>(
        data,
        map_t{ext_t{d0, d1, d2},
              std::array<std::size_t, 3>{
                  static_cast<std::size_t>(s0),
                  static_cast<std::size_t>(s1),
                  static_cast<std::size_t>(s2)}});
}

// pinned_view_4d wraps a 4-D pinned buffer (e.g. [batch × heads ×
// seq × head_dim] KV cache tensors, [layers × heads × seq × dim]
// MLX storage layouts) with explicit per-dim strides.
template <typename T>
[[nodiscard]] auto pinned_view_4d(
    T* data,
    std::size_t d0,
    std::size_t d1,
    std::size_t d2,
    std::size_t d3,
    std::ptrdiff_t s0,
    std::ptrdiff_t s1,
    std::ptrdiff_t s2,
    std::ptrdiff_t s3) noexcept {
    using ext_t = std::dextents<std::size_t, 4>;
    using map_t = std::layout_stride::mapping<ext_t>;
    return std::mdspan<T, ext_t, std::layout_stride>(
        data,
        map_t{ext_t{d0, d1, d2, d3},
              std::array<std::size_t, 4>{
                  static_cast<std::size_t>(s0),
                  static_cast<std::size_t>(s1),
                  static_cast<std::size_t>(s2),
                  static_cast<std::size_t>(s3)}});
}

// contiguous_strides_1d / _2d / _3d / _4d compute row-major strides
// for the given shape, returning a std::array suitable for passing
// to pinned_view_*. Use when the Go side stored the buffer in row-
// major layout (the common case).

[[nodiscard]] constexpr auto contiguous_strides_2d(
    std::size_t, std::size_t cols) noexcept {
    return std::array<std::ptrdiff_t, 2>{
        static_cast<std::ptrdiff_t>(cols),
        1};
}

[[nodiscard]] constexpr auto contiguous_strides_3d(
    std::size_t, std::size_t d1, std::size_t d2) noexcept {
    return std::array<std::ptrdiff_t, 3>{
        static_cast<std::ptrdiff_t>(d1 * d2),
        static_cast<std::ptrdiff_t>(d2),
        1};
}

[[nodiscard]] constexpr auto contiguous_strides_4d(
    std::size_t, std::size_t d1, std::size_t d2, std::size_t d3) noexcept {
    return std::array<std::ptrdiff_t, 4>{
        static_cast<std::ptrdiff_t>(d1 * d2 * d3),
        static_cast<std::ptrdiff_t>(d2 * d3),
        static_cast<std::ptrdiff_t>(d3),
        1};
}

} // namespace lthn::cgo
