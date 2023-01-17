// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build generate
// +build generate

package sound

import "windows"

//go:generate genwinmdsigs -o syscall.go -source C:\git\go-winmd\testdata\Windows.Win32.winmd -template genwinmdsigstemplate.tmpl -filter .*Windows.Win32.Media.Audio::Apis.*
//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscall.go
