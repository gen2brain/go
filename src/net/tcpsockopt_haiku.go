// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"syscall"
	"time"
)

func setKeepAliveIdle(fd *netFD, d time.Duration) error {
	return syscall.ENOPROTOOPT
}

func setKeepAliveInterval(fd *netFD, d time.Duration) error {
	return syscall.ENOPROTOOPT
}

func setKeepAliveCount(fd *netFD, n int) error {
	return syscall.ENOPROTOOPT
}
