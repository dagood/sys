// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package sound_test

import (
	"testing"

	"golang.org/x/sys/windows/sound"
)

func TestSound(t *testing.T) {
	x := sound.AcmGetVersion()
	t.Logf("ACM version: 0x%x", x)
}
