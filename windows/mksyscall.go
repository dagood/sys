// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build generate
// +build generate

package windows

//go:generate genwinmdsigs -o sound_windows.go -source C:\git\go-winmd\testdata\Windows.Win32.winmd -template genwinmdsigstemplate.tmpl -filter .*Windows.Win32.Media.Audio::Apis.*
//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go eventlog.go service.go syscall_windows.go security_windows.go setupapi_windows.go sound_windows.go
