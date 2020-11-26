// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: global

//nolint:golint
package global

import (
	"fmt"

	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

// Error definitions
// =================

var (
	ErrNoNamespace       = verror.NewIDAction("v.io/x/ref/lib/discovery/global.NoNamespace", verror.NoRetry)
	ErrAdInvalidEncoding = verror.NewIDAction("v.io/x/ref/lib/discovery/global.AdInvalidEncoding", verror.NoRetry)
)

// NewErrNoNamespace returns an error with the ErrNoNamespace ID.
// Deprecated: this function will be removed in the future,
// use ErrorfNoNamespace or MessageNoNamespace instead.
func NewErrNoNamespace(ctx *context.T) error {
	return verror.New(ErrNoNamespace, ctx)
}

// ErrorfNoNamespace calls ErrNoNamespace.Errorf with the supplied arguments.
func ErrorfNoNamespace(ctx *context.T, format string) error {
	return ErrNoNamespace.Errorf(ctx, format)
}

// MessageNoNamespace calls ErrNoNamespace.Message with the supplied arguments.
func MessageNoNamespace(ctx *context.T, message string) error {
	return ErrNoNamespace.Message(ctx, message)
}

// ParamsErrNoNamespace extracts the expected parameters from the error's ParameterList.
func ParamsErrNoNamespace(argumentError error) (verrorComponent string, verrorOperation string, returnErr error) {
	params := verror.Params(argumentError)
	if params == nil {
		returnErr = fmt.Errorf("no parameters found in: %T: %v", argumentError, argumentError)
		return
	}
	iter := &paramListIterator{params: params, max: len(params)}

	if verrorComponent, verrorOperation, returnErr = iter.preamble(); returnErr != nil {
		return
	}

	return
}

// NewErrAdInvalidEncoding returns an error with the ErrAdInvalidEncoding ID.
// Deprecated: this function will be removed in the future,
// use ErrorfAdInvalidEncoding or MessageAdInvalidEncoding instead.
func NewErrAdInvalidEncoding(ctx *context.T, ad string) error {
	return verror.New(ErrAdInvalidEncoding, ctx, ad)
}

// ErrorfAdInvalidEncoding calls ErrAdInvalidEncoding.Errorf with the supplied arguments.
func ErrorfAdInvalidEncoding(ctx *context.T, format string, ad string) error {
	return ErrAdInvalidEncoding.Errorf(ctx, format, ad)
}

// MessageAdInvalidEncoding calls ErrAdInvalidEncoding.Message with the supplied arguments.
func MessageAdInvalidEncoding(ctx *context.T, message string, ad string) error {
	return ErrAdInvalidEncoding.Message(ctx, message, ad)
}

// ParamsErrAdInvalidEncoding extracts the expected parameters from the error's ParameterList.
func ParamsErrAdInvalidEncoding(argumentError error) (verrorComponent string, verrorOperation string, ad string, returnErr error) {
	params := verror.Params(argumentError)
	if params == nil {
		returnErr = fmt.Errorf("no parameters found in: %T: %v", argumentError, argumentError)
		return
	}
	iter := &paramListIterator{params: params, max: len(params)}

	if verrorComponent, verrorOperation, returnErr = iter.preamble(); returnErr != nil {
		return
	}

	var (
		tmp interface{}
		ok  bool
	)
	tmp, returnErr = iter.next()
	if ad, ok = tmp.(string); !ok {
		if returnErr != nil {
			return
		}
		returnErr = fmt.Errorf("parameter list contains the wrong type for return value ad, has %T and not string", tmp)
		return
	}

	return
}

type paramListIterator struct {
	err      error
	idx, max int
	params   []interface{}
}

func (pl *paramListIterator) next() (interface{}, error) {
	if pl.err != nil {
		return nil, pl.err
	}
	if pl.idx+1 > pl.max {
		pl.err = fmt.Errorf("too few parameters: have %v", pl.max)
		return nil, pl.err
	}
	pl.idx++
	return pl.params[pl.idx-1], nil
}

func (pl *paramListIterator) preamble() (component, operation string, err error) {
	var tmp interface{}
	if tmp, err = pl.next(); err != nil {
		return
	}
	var ok bool
	if component, ok = tmp.(string); !ok {
		return "", "", fmt.Errorf("ParamList[0]: component name is not a string: %T", tmp)
	}
	if tmp, err = pl.next(); err != nil {
		return
	}
	if operation, ok = tmp.(string); !ok {
		return "", "", fmt.Errorf("ParamList[1]: operation name is not a string: %T", tmp)
	}
	return
}

var initializeVDLCalled bool

// initializeVDL performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = initializeVDL()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func initializeVDL() struct{} {
	if initializeVDLCalled {
		return struct{}{}
	}
	initializeVDLCalled = true

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoNamespace.ID), "{1:}{2:} namespace not found")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrAdInvalidEncoding.ID), "{1:}{2:} ad ({3}) has invalid encoding")

	return struct{}{}
}