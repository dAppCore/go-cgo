package cgo

import core "dappco.re/go"

type T = core.T

var (
	AssertEmpty           = core.AssertEmpty
	AssertEqual           = core.AssertEqual
	AssertError           = core.AssertError
	AssertErrorIs         = core.AssertErrorIs
	AssertFalse           = core.AssertFalse
	AssertLen             = core.AssertLen
	AssertNil             = core.AssertNil
	AssertNoError         = core.AssertNoError
	AssertNotEqual        = core.AssertNotEqual
	AssertNotNil          = core.AssertNotNil
	AssertNotPanics       = core.AssertNotPanics
	AssertPanics          = core.AssertPanics
	AssertPanicsWithError = core.AssertPanicsWithError
	AssertTrue            = core.AssertTrue
	RequireNoError        = core.RequireNoError
)
