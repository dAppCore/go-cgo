package cgo

func TestBuffer_NewBuffer_Good(t *T) {
	buffer := NewBuffer(8)
	defer buffer.Free()

	AssertNotNil(t, buffer)
	AssertFalse(t, buffer.IsFreed())
	AssertEqual(t, 8, buffer.Len())
	AssertLen(t, buffer.Bytes(), 8)
	AssertNotNil(t, buffer.Ptr())
}

func TestBuffer_NewBuffer_Bad(t *T) {
	AssertPanicsWithError(t, "size must be non-negative", func() {
		_ = NewBuffer(-1)
	})
	AssertNotPanics(t, func() {
		buffer := NewBuffer(1)
		buffer.Free()
	})
}

func TestBuffer_NewBuffer_Ugly(t *T) {
	buffer := NewBuffer(0)
	defer buffer.Free()

	AssertEqual(t, 0, buffer.Len())
	AssertEmpty(t, buffer.Bytes())
	AssertEqual(t, uintptr(0), uintptr(buffer.Ptr()))
}

func TestBuffer_Buffer_Free_Good(t *T) {
	buffer := NewBuffer(4)
	buffer.Free()

	AssertTrue(t, buffer.IsFreed())
	AssertPanicsWithError(t, "double-free detected", func() {
		buffer.Free()
	})
}

func TestBuffer_Buffer_Free_Bad(t *T) {
	buffer := NewBuffer(4)
	buffer.Free()

	AssertPanicsWithError(t, "double-free detected", func() {
		buffer.Free()
	})
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_Free_Ugly(t *T) {
	var buffer *Buffer

	AssertTrue(t, buffer.IsFreed())
	AssertNotPanics(t, func() {
		buffer.Free()
	})
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_Close_Good(t *T) {
	buffer := NewBuffer(1)
	r := buffer.Close()

	AssertTrue(t, r.OK)
	AssertNil(t, r.Value)
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_Close_Bad(t *T) {
	buffer := NewBuffer(1)
	AssertTrue(t, buffer.Close().OK)

	AssertPanicsWithError(t, "double-free detected", func() {
		_ = buffer.Close()
	})
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_Close_Ugly(t *T) {
	var buffer *Buffer
	r := buffer.Close()

	AssertTrue(t, r.OK)
	AssertNil(t, r.Value)
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_CopyFrom_Good(t *T) {
	buffer := NewBuffer(3)
	defer buffer.Free()

	copied := buffer.CopyFrom([]byte("abcd"))
	AssertEqual(t, 3, copied)
	AssertEqual(t, "abc", string(buffer.Bytes()))
}

func TestBuffer_Buffer_CopyFrom_Bad(t *T) {
	buffer := NewBuffer(1)
	buffer.Free()

	AssertPanicsWithError(t, "use-after-free detected", func() {
		_ = buffer.CopyFrom([]byte("x"))
	})
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_CopyFrom_Ugly(t *T) {
	buffer := NewBuffer(0)
	defer buffer.Free()

	copied := buffer.CopyFrom([]byte("abc"))
	AssertEqual(t, 0, copied)
	AssertEmpty(t, buffer.Bytes())
}

func TestBuffer_Buffer_Bytes_Good(t *T) {
	buffer := NewBuffer(2)
	defer buffer.Free()

	bytes := buffer.Bytes()
	bytes[0] = 'g'
	bytes[1] = 'o'
	AssertEqual(t, "go", string(buffer.Bytes()))
}

func TestBuffer_Buffer_Bytes_Bad(t *T) {
	buffer := NewBuffer(2)
	buffer.Free()

	AssertPanicsWithError(t, "use-after-free detected", func() {
		_ = buffer.Bytes()
	})
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_Bytes_Ugly(t *T) {
	var buffer *Buffer

	AssertPanicsWithError(t, "buffer is nil", func() {
		_ = buffer.Bytes()
	})
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_Ptr_Good(t *T) {
	buffer := NewBuffer(2)
	defer buffer.Free()

	pointer := buffer.Ptr()
	AssertNotNil(t, pointer)
	AssertFalse(t, buffer.IsFreed())
}

func TestBuffer_Buffer_Ptr_Bad(t *T) {
	buffer := NewBuffer(2)
	buffer.Free()

	AssertPanicsWithError(t, "use-after-free detected", func() {
		_ = buffer.Ptr()
	})
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_Ptr_Ugly(t *T) {
	buffer := NewBuffer(0)
	defer buffer.Free()

	AssertEqual(t, uintptr(0), uintptr(buffer.Ptr()))
	AssertEqual(t, 0, buffer.Len())
}

func TestBuffer_Buffer_Len_Good(t *T) {
	buffer := NewBuffer(5)
	defer buffer.Free()

	length := buffer.Len()
	AssertEqual(t, 5, length)
	AssertLen(t, buffer.Bytes(), 5)
}

func TestBuffer_Buffer_Len_Bad(t *T) {
	buffer := NewBuffer(1)
	buffer.Free()

	AssertPanicsWithError(t, "use-after-free detected", func() {
		_ = buffer.Len()
	})
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_Len_Ugly(t *T) {
	buffer := NewBuffer(0)
	defer buffer.Free()

	AssertEqual(t, 0, buffer.Len())
	AssertEmpty(t, buffer.Bytes())
}

func TestBuffer_Buffer_IsFreed_Good(t *T) {
	buffer := NewBuffer(1)
	AssertFalse(t, buffer.IsFreed())

	buffer.Free()
	AssertTrue(t, buffer.IsFreed())
}

func TestBuffer_Buffer_IsFreed_Bad(t *T) {
	buffer := NewBuffer(0)
	defer buffer.Free()

	AssertFalse(t, buffer.IsFreed())
	AssertEqual(t, 0, buffer.Len())
}

func TestBuffer_Buffer_IsFreed_Ugly(t *T) {
	var buffer *Buffer

	AssertTrue(t, buffer.IsFreed())
	AssertNotPanics(t, func() {
		buffer.Free()
	})
}
