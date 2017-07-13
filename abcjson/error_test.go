// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2017 The Aero Blockchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package abcjson_test

import (
	"testing"

	"github.com/abcsuite/abcd/abcjson"
)

// TestErrorCodeStringer tests the stringized output for the ErrorCode type.
func TestErrorCodeStringer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   abcjson.ErrorCode
		want string
	}{
		{abcjson.ErrDuplicateMethod, "ErrDuplicateMethod"},
		{abcjson.ErrInvalidUsageFlags, "ErrInvalidUsageFlags"},
		{abcjson.ErrInvalidType, "ErrInvalidType"},
		{abcjson.ErrEmbeddedType, "ErrEmbeddedType"},
		{abcjson.ErrUnexportedField, "ErrUnexportedField"},
		{abcjson.ErrUnsupportedFieldType, "ErrUnsupportedFieldType"},
		{abcjson.ErrNonOptionalField, "ErrNonOptionalField"},
		{abcjson.ErrNonOptionalDefault, "ErrNonOptionalDefault"},
		{abcjson.ErrMismatchedDefault, "ErrMismatchedDefault"},
		{abcjson.ErrUnregisteredMethod, "ErrUnregisteredMethod"},
		{abcjson.ErrNumParams, "ErrNumParams"},
		{abcjson.ErrMissingDescription, "ErrMissingDescription"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	// Detect additional error codes that don't have the stringer added.
	if len(tests)-1 != int(abcjson.TstNumErrorCodes) {
		t.Errorf("It appears an error code was added without adding an " +
			"associated stringer test")
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestError tests the error output for the Error type.
func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   abcjson.Error
		want string
	}{
		{
			abcjson.Error{Message: "some error"},
			"some error",
		},
		{
			abcjson.Error{Message: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
