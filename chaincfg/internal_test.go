// Copyright (c) 2017 The Aero Blockchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"testing"

	"github.com/abcsuite/abcd/chaincfg/chainhash"
)

func TestInvalidHashStr(t *testing.T) {
	_, err := chainhash.NewHashFromStr("banana")
	if err == nil {
		t.Error("Invalid string should fail.")
	}
}
