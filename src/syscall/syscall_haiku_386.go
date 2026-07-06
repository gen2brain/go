// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import "unsafe"

// syscall7 bridges to the runtime for the one libc call that needs seven
// words on 386: mmap, whose 64-bit off_t is passed as two of them.
func syscall7(trap, nargs, a1, a2, a3, a4, a5, a6, a7 uintptr) (r1, r2 uintptr, err Errno)

func (iov *Iovec) SetLen(length int) {
	iov.Len = uint32(length)
}

func mmap(addr uintptr, length uintptr, prot, flags, fd int, offset int64) (uintptr, error) {
	// off_t is 64-bit on Haiku x86; pass it as two 32-bit words.
	r0, _, e1 := syscall7(uintptr(unsafe.Pointer(&libc_mmap)), 7,
		addr, length, uintptr(prot), uintptr(flags), uintptr(fd), uintptr(offset), uintptr(offset>>32))
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return r0, nil
}

func setTimespec(sec, nsec int64) Timespec { return Timespec{Sec: int32(sec), Nsec: int32(nsec)} }
func setTimeval(sec, usec int64) Timeval   { return Timeval{Sec: int32(sec), Usec: int32(usec)} }

func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	return 0, ENOSYS
}
