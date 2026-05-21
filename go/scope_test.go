package cgo

func TestScope_NewScope_Good(t *T) {
	scope := NewScope()
	defer scope.FreeAll()

	AssertNotNil(t, scope)
	AssertFalse(t, scope.IsFreed())
	AssertNotNil(t, scope.Buffer(1))
}

func TestScope_NewScope_Bad(t *T) {
	first := NewScope()
	second := NewScope()
	defer second.FreeAll()

	first.FreeAll()
	AssertTrue(t, first.IsFreed())
	AssertFalse(t, second.IsFreed())
}

func TestScope_NewScope_Ugly(t *T) {
	scope := NewScope()

	AssertNotPanics(t, func() {
		scope.FreeAll()
	})
	AssertTrue(t, scope.IsFreed())
}

func TestScope_Scope_Buffer_Good(t *T) {
	scope := NewScope()
	defer scope.FreeAll()

	buffer := scope.Buffer(4)
	AssertNotNil(t, buffer)
	AssertEqual(t, 4, buffer.Len())
}

func TestScope_Scope_Buffer_Bad(t *T) {
	scope := NewScope()
	scope.FreeAll()

	AssertPanicsWithError(t, "scope is already freed", func() {
		_ = scope.Buffer(1)
	})
	AssertTrue(t, scope.IsFreed())
}

func TestScope_Scope_Buffer_Ugly(t *T) {
	scope := NewScope()
	defer scope.FreeAll()

	AssertPanicsWithError(t, "size must be non-negative", func() {
		_ = scope.Buffer(-1)
	})
	AssertFalse(t, scope.IsFreed())
}

func TestScope_Scope_CString_Good(t *T) {
	scope := NewScope()
	defer scope.FreeAll()

	cString := scope.CString("hello")
	AssertNotNil(t, cString)
	AssertEqual(t, "hello", GoString(cString))
}

func TestScope_Scope_CString_Bad(t *T) {
	scope := NewScope()
	scope.FreeAll()

	AssertPanicsWithError(t, "scope is already freed", func() {
		_ = scope.CString("x")
	})
	AssertTrue(t, scope.IsFreed())
}

func TestScope_Scope_CString_Ugly(t *T) {
	scope := NewScope()
	defer scope.FreeAll()

	cString := scope.CString("")
	AssertNotNil(t, cString)
	AssertEqual(t, "", GoString(cString))
}

func TestScope_Scope_FreeAll_Good(t *T) {
	scope := NewScope()
	buffer := scope.Buffer(2)
	cString := scope.CString("hi")

	scope.FreeAll()
	AssertTrue(t, scope.IsFreed())
	AssertTrue(t, buffer.IsFreed())
	AssertNotNil(t, cString)
}

func TestScope_Scope_FreeAll_Bad(t *T) {
	scope := NewScope()
	scope.FreeAll()

	AssertPanicsWithError(t, "double-free detected", func() {
		scope.FreeAll()
	})
	AssertTrue(t, scope.IsFreed())
}

func TestScope_Scope_FreeAll_Ugly(t *T) {
	var scope *Scope

	AssertTrue(t, scope.IsFreed())
	AssertNotPanics(t, func() {
		scope.FreeAll()
	})
	AssertTrue(t, scope.IsFreed())
}

func TestScope_Scope_Close_Good(t *T) {
	scope := NewScope()
	buffer := scope.Buffer(1)
	r := scope.Close()

	AssertTrue(t, r.OK)
	AssertNil(t, r.Value)
	AssertTrue(t, scope.IsFreed())
	AssertTrue(t, buffer.IsFreed())
}

func TestScope_Scope_Close_Bad(t *T) {
	scope := NewScope()
	AssertTrue(t, scope.Close().OK)

	AssertPanicsWithError(t, "double-free detected", func() {
		_ = scope.Close()
	})
	AssertTrue(t, scope.IsFreed())
}

func TestScope_Scope_Close_Ugly(t *T) {
	var scope *Scope
	r := scope.Close()

	AssertTrue(t, r.OK)
	AssertNil(t, r.Value)
	AssertTrue(t, scope.IsFreed())
}

func TestScope_Scope_IsFreed_Good(t *T) {
	scope := NewScope()
	AssertFalse(t, scope.IsFreed())

	scope.FreeAll()
	AssertTrue(t, scope.IsFreed())
}

func TestScope_Scope_IsFreed_Bad(t *T) {
	scope := NewScope()
	defer scope.FreeAll()

	AssertFalse(t, scope.IsFreed())
	AssertNotNil(t, scope.Buffer(1))
}

func TestScope_Scope_IsFreed_Ugly(t *T) {
	var scope *Scope

	AssertTrue(t, scope.IsFreed())
	AssertNotPanics(t, func() {
		scope.FreeAll()
	})
}

func TestScope_PinIn_Good(t *T) {
	scope := NewScope()
	defer scope.FreeAll()

	slice := []int32{1, 2, 3, 4}
	view := PinIn(scope, slice)

	AssertNotNil(t, view)
	AssertTrue(t, view.Active())
	AssertEqual(t, 4, view.Len())
	AssertEqual(t, 16, view.Bytes())
}

func TestScope_PinIn_EmptySlice(t *T) {
	scope := NewScope()
	defer scope.FreeAll()

	var slice []int32
	view := PinIn(scope, slice)

	AssertNotNil(t, view)
	AssertFalse(t, view.Active())
}

func TestScope_PinIn_ReleasedOnFreeAll(t *T) {
	scope := NewScope()
	slice := []float32{1, 2, 3, 4}
	view := PinIn(scope, slice)

	AssertTrue(t, view.Active())
	scope.FreeAll()
	AssertFalse(t, view.Active())
}

func TestScope_PinIn_NilScope_Panics(t *T) {
	AssertPanicsWithError(t, "scope is nil", func() {
		PinIn[int32](nil, []int32{1, 2})
	})
}

func TestScope_PinIn_FreedScope_Panics(t *T) {
	scope := NewScope()
	scope.FreeAll()

	AssertPanicsWithError(t, "scope is already freed", func() {
		PinIn(scope, []int32{1, 2})
	})
}
