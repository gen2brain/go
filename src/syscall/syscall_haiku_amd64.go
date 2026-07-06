// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import "unsafe"

func (iov *Iovec) SetLen(length int) {
	iov.Len = uint64(length)
}

func mmap(addr uintptr, length uintptr, prot, flags, fd int, offset int64) (uintptr, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_mmap)), 6,
		addr, length, uintptr(prot), uintptr(flags), uintptr(fd), uintptr(offset))
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return r0, nil
}

func setTimespec(sec, nsec int64) Timespec { return Timespec{Sec: sec, Nsec: nsec} }
func setTimeval(sec, usec int64) Timeval   { return Timeval{Sec: sec, Usec: int32(usec)} }

func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	return 0, ENOSYS
}
