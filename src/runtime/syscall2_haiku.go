// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Additional libroot symbols used by the runtime side of the exec_libc.go
// forkAndExecInChild path. The matching var libFunc declarations live
// here alongside, mirroring runtime/syscall_solaris.go.

package runtime

import _ "unsafe" // for go:linkname

//go:linkname libc_chdir libc_chdir
//go:linkname libc_chroot libc_chroot
//go:linkname libc_dup2 libc_dup2
//go:linkname libc_execve libc_execve
//go:linkname libc_fork libc_fork
//go:linkname libc_ioctl libc_ioctl
//go:linkname libc_setgid libc_setgid
//go:linkname libc_setgroups libc_setgroups
//go:linkname libc_setpgid libc_setpgid
//go:linkname libc_setrlimit libc_setrlimit
//go:linkname libc_setsid libc_setsid
//go:linkname libc_setuid libc_setuid

var (
	libc_chdir,
	libc_chroot,
	libc_dup2,
	libc_execve,
	libc_fork,
	libc_ioctl,
	libc_setgid,
	libc_setgroups,
	libc_setpgid,
	libc_setrlimit,
	libc_setsid,
	libc_setuid libFunc
)
