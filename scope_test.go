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
	err := scope.Close()

	AssertNoError(t, err)
	AssertTrue(t, scope.IsFreed())
	AssertTrue(t, buffer.IsFreed())
}

func TestScope_Scope_Close_Bad(t *T) {
	scope := NewScope()
	RequireNoError(t, scope.Close())

	AssertPanicsWithError(t, "double-free detected", func() {
		_ = scope.Close()
	})
	AssertTrue(t, scope.IsFreed())
}

func TestScope_Scope_Close_Ugly(t *T) {
	var scope *Scope
	err := scope.Close()

	AssertNoError(t, err)
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
