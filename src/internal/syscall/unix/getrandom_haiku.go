// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import "unsafe"

//go:cgo_import_dynamic libc_getentropy getentropy "libroot.so"

//go:linkname procGetentropy libc_getentropy

var procGetentropy uintptr

// GetRandomFlag is a flag supported by the getrandom system call.
type GetRandomFlag uintptr

const (
	GRND_NONBLOCK GetRandomFlag = 0x0001
	GRND_RANDOM   GetRandomFlag = 0x0002
)

// GetRandom calls Haiku's getentropy. Haiku rejects requests larger than
// 256 bytes; the caller in crypto/internal/sysrand chunks accordingly.
func GetRandom(p []byte, flags GetRandomFlag) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procGetentropy)),
		2,
		uintptr(unsafe.Pointer(&p[0])),
		uintptr(len(p)),
		0, 0, 0, 0)
	if errno != 0 {
		return 0, errno
	}
	return len(p), nil
}
