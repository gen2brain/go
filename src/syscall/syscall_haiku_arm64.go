// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

func (iov *Iovec) SetLen(length int) {
	iov.Len = uint64(length)
}

func setTimespec(sec, nsec int64) Timespec { return Timespec{Sec: sec, Nsec: nsec} }
func setTimeval(sec, usec int64) Timeval   { return Timeval{Sec: sec, Usec: int32(usec)} }

func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	return 0, ENOSYS
}
