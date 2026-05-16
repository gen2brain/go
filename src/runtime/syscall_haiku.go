// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

// Bridges from package syscall into asmsyscall6. Errnos returned to
// package syscall are POSIX-positive; see haikuErrnoToPosix.

// Haiku has no numbered-syscall API. Syscall and RawSyscall return EINVAL
// so callers fall back to the named libc wrappers.
const posixEINVAL = 22

//go:nosplit
//go:linkname syscall_Syscall
func syscall_Syscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	return 0, 0, posixEINVAL
}

//go:linkname syscall_RawSyscall
func syscall_RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	return 0, 0, posixEINVAL
}

// syscall_syscall6 invokes a libc function via the runtime's asmsyscall6
// trampoline. fn is the address of a libFunc populated by cgo_import_dynamic.
//
//go:nosplit
//go:cgo_unsafe_args
//go:linkname syscall_syscall6
func syscall_syscall6(fn, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	c := libcall{
		fn:   fn,
		n:    nargs,
		args: uintptr(unsafe.Pointer(&a1)),
	}

	entersyscallblock()
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	exitsyscall()
	if c.err != 0 {
		c.err = haikuErrnoToPosix(c.err)
	}
	return c.r1, c.r2, c.err
}

//go:nosplit
//go:cgo_unsafe_args
//go:linkname syscall_rawSyscall6
func syscall_rawSyscall6(fn, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	c := libcall{
		fn:   fn,
		n:    nargs,
		args: uintptr(unsafe.Pointer(&a1)),
	}

	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))

	if c.err != 0 {
		c.err = haikuErrnoToPosix(c.err)
	}
	return c.r1, c.r2, c.err
}

// Hand-crafted libc calls used between fork and exec by the shared
// syscall/exec_libc.go forkAndExecInChild. They mirror the Solaris
// versions and run with asmcgocall so they are safe in the child without
// scheduler involvement.

//go:nosplit
//go:linkname syscall_chdir
//go:cgo_unsafe_args
func syscall_chdir(path uintptr) (err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_chdir)),
		n:    1,
		args: uintptr(unsafe.Pointer(&path)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return haikuErrnoToPosix(c.err)
	}
	return 0
}

//go:nosplit
//go:linkname syscall_chroot
//go:cgo_unsafe_args
func syscall_chroot(path uintptr) (err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_chroot)),
		n:    1,
		args: uintptr(unsafe.Pointer(&path)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return haikuErrnoToPosix(c.err)
	}
	return 0
}

//go:nosplit
//go:linkname syscall_close
//go:cgo_unsafe_args
func syscall_close(fd uintptr) (err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_close)),
		n:    1,
		args: uintptr(unsafe.Pointer(&fd)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return haikuErrnoToPosix(c.err)
	}
	return 0
}

//go:nosplit
//go:linkname syscall_dup2
//go:cgo_unsafe_args
func syscall_dup2(old, new uintptr) (val, err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_dup2)),
		n:    2,
		args: uintptr(unsafe.Pointer(&old)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return c.r1, haikuErrnoToPosix(c.err)
	}
	return c.r1, 0
}

//go:nosplit
//go:linkname syscall_execve
//go:cgo_unsafe_args
func syscall_execve(path, argv, envp uintptr) (err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_execve)),
		n:    3,
		args: uintptr(unsafe.Pointer(&path)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return haikuErrnoToPosix(c.err)
	}
	return 0
}

//go:nosplit
//go:linkname syscall_exit
//go:cgo_unsafe_args
func syscall_exit(code uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_exit)),
		n:    1,
		args: uintptr(unsafe.Pointer(&code)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
}

//go:nosplit
//go:linkname syscall_fcntl
//go:cgo_unsafe_args
func syscall_fcntl(fd, cmd, arg uintptr) (val, err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_fcntl)),
		n:    3,
		args: uintptr(unsafe.Pointer(&fd)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return c.r1, haikuErrnoToPosix(c.err)
	}
	return c.r1, 0
}

// syscall_forkx calls libroot fork(). The flags argument exists only for
// source compatibility with the Solaris exec_libc.go path and is ignored.
// In the child we re-resolve the per-thread errno location.
//
//go:nosplit
//go:linkname syscall_forkx
//go:cgo_unsafe_args
func syscall_forkx(flags uintptr) (pid, err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_fork)),
		n:    0,
		args: uintptr(unsafe.Pointer(&flags)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return 0, haikuErrnoToPosix(c.err)
	}
	if c.r1 == 0 {
		miniterrno()
	}
	return c.r1, 0
}

//go:nosplit
//go:linkname syscall_getpid
func syscall_getpid() (pid, err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_getpid)),
		n:    0,
		args: uintptr(unsafe.Pointer(&libc_getpid)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	return c.r1, 0
}

//go:nosplit
//go:linkname syscall_ioctl
//go:cgo_unsafe_args
func syscall_ioctl(fd, req, arg uintptr) (err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_ioctl)),
		n:    3,
		args: uintptr(unsafe.Pointer(&fd)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return haikuErrnoToPosix(c.err)
	}
	return 0
}

//go:nosplit
//go:linkname syscall_setgid
//go:cgo_unsafe_args
func syscall_setgid(gid uintptr) (err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_setgid)),
		n:    1,
		args: uintptr(unsafe.Pointer(&gid)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return haikuErrnoToPosix(c.err)
	}
	return 0
}

//go:nosplit
//go:linkname syscall_setgroups
//go:cgo_unsafe_args
func syscall_setgroups(ngid, gid uintptr) (err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_setgroups)),
		n:    2,
		args: uintptr(unsafe.Pointer(&ngid)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return haikuErrnoToPosix(c.err)
	}
	return 0
}

//go:nosplit
//go:linkname syscall_setrlimit
//go:cgo_unsafe_args
func syscall_setrlimit(which uintptr, lim unsafe.Pointer) (err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_setrlimit)),
		n:    2,
		args: uintptr(unsafe.Pointer(&which)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return haikuErrnoToPosix(c.err)
	}
	return 0
}

//go:nosplit
//go:linkname syscall_setsid
func syscall_setsid() (pid, err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_setsid)),
		n:    0,
		args: uintptr(unsafe.Pointer(&libc_setsid)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return c.r1, haikuErrnoToPosix(c.err)
	}
	return c.r1, 0
}

//go:nosplit
//go:linkname syscall_setuid
//go:cgo_unsafe_args
func syscall_setuid(uid uintptr) (err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_setuid)),
		n:    1,
		args: uintptr(unsafe.Pointer(&uid)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return haikuErrnoToPosix(c.err)
	}
	return 0
}

//go:nosplit
//go:linkname syscall_setpgid
//go:cgo_unsafe_args
func syscall_setpgid(pid, pgid uintptr) (err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_setpgid)),
		n:    2,
		args: uintptr(unsafe.Pointer(&pid)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return haikuErrnoToPosix(c.err)
	}
	return 0
}

//go:nosplit
//go:linkname syscall_write
//go:cgo_unsafe_args
func syscall_write(fd, buf, nbyte uintptr) (n, err uintptr) {
	c := libcall{
		fn:   uintptr(unsafe.Pointer(&libc_write)),
		n:    3,
		args: uintptr(unsafe.Pointer(&fd)),
	}
	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))
	if c.err != 0 {
		return c.r1, haikuErrnoToPosix(c.err)
	}
	return c.r1, 0
}

// haikuErrnoToPosix maps a raw libroot errno (an INT_MIN+N int32, stored
// zero-extended in a uintptr) to the POSIX-positive value the syscall
// package exposes. Unknown negative codes fall through to EIO.
//
//go:nosplit
//go:linkname haikuErrnoToPosix syscall.haikuErrnoToPosix
func haikuErrnoToPosix(raw uintptr) uintptr {
	switch raw {
	case 0x8000000f:
		return 1 // EPERM
	case 0x80006003:
		return 2 // ENOENT
	case 0x8000700d:
		return 3 // ESRCH
	case 0x8000000a:
		return 4 // EINTR
	case 0x80000001:
		return 5 // EIO
	case 0x8000700b:
		return 6 // ENXIO
	case 0x80007001:
		return 7 // E2BIG
	case 0x80001302:
		return 8 // ENOEXEC
	case 0x80006000:
		return 9 // EBADF
	case 0x80007002:
		return 10 // ECHILD
	case 0x8000000b:
		return 11 // EAGAIN/EWOULDBLOCK
	case 0x80000000:
		return 12 // ENOMEM
	case 0x80000002:
		return 13 // EACCES
	case 0x80001301:
		return 14 // EFAULT
	case 0x8000000e:
		return 16 // EBUSY
	case 0x80006002:
		return 17 // EEXIST
	case 0x8000600b:
		return 18 // EXDEV
	case 0x80007007:
		return 19 // ENODEV
	case 0x80006005:
		return 20 // ENOTDIR
	case 0x80006009:
		return 21 // EISDIR
	case 0x80000005:
		return 22 // EINVAL
	case 0x80007006:
		return 23 // ENFILE
	case 0x8000600a:
		return 24 // EMFILE
	case 0x8000700a:
		return 25 // ENOTTY
	case 0x8000703b:
		return 26 // ETXTBSY
	case 0x80007004:
		return 27 // EFBIG
	case 0x80006007:
		return 28 // ENOSPC
	case 0x8000700c:
		return 29 // ESPIPE
	case 0x80006008:
		return 30 // EROFS
	case 0x80007005:
		return 31 // EMLINK
	case 0x8000600d:
		return 32 // EPIPE
	case 0x80007010:
		return 33 // EDOM
	case 0x80007011:
		return 34 // ERANGE
	case 0x80007003:
		return 35 // EDEADLK
	case 0x80006004:
		return 36 // ENAMETOOLONG
	case 0x80007008:
		return 37 // ENOLCK
	case 0x80007009:
		return 38 // ENOSYS
	case 0x80006006:
		return 39 // ENOTEMPTY
	case 0x8000600c:
		return 40 // ELOOP
	case 0x80007027:
		return 42 // ENOMSG
	case 0x80007032:
		return 43 // EIDRM
	case 0x80007035:
		return 67 // ENOLINK
	case 0x80007039:
		return 71 // EPROTO
	case 0x80007033:
		return 72 // EMULTIHOP
	case 0x8000702e:
		return 74 // EBADMSG
	case 0x80007029:
		return 75 // EOVERFLOW
	case 0x80007026:
		return 84 // EILSEQ
	case 0x8000702c:
		return 88 // ENOTSOCK
	case 0x80007030:
		return 89 // EDESTADDRREQ
	case 0x8000702a:
		return 90 // EMSGSIZE
	case 0x80007012:
		return 91 // EPROTOTYPE
	case 0x80007022:
		return 92 // ENOPROTOOPT
	case 0x80007013:
		return 93 // EPROTONOSUPPORT
	case 0x8000702b, 0x80007038:
		return 95 // EOPNOTSUPP / ENOTSUP
	case 0x80007014:
		return 96 // EPFNOSUPPORT
	case 0x80007015:
		return 97 // EAFNOSUPPORT
	case 0x80007016:
		return 98 // EADDRINUSE
	case 0x80007017:
		return 99 // EADDRNOTAVAIL
	case 0x80007018:
		return 100 // ENETDOWN
	case 0x80007019:
		return 101 // ENETUNREACH
	case 0x8000701a:
		return 102 // ENETRESET
	case 0x8000701b:
		return 103 // ECONNABORTED
	case 0x8000701c:
		return 104 // ECONNRESET
	case 0x80007023:
		return 105 // ENOBUFS
	case 0x8000701d:
		return 106 // EISCONN
	case 0x8000701e:
		return 107 // ENOTCONN
	case 0x8000701f:
		return 108 // ESHUTDOWN
	case 0x80000009:
		return 110 // ETIMEDOUT
	case 0x80007020:
		return 111 // ECONNREFUSED
	case 0x8000702d:
		return 112 // EHOSTDOWN
	case 0x80007021:
		return 113 // EHOSTUNREACH
	case 0x80007025:
		return 114 // EALREADY
	case 0x80007024:
		return 115 // EINPROGRESS
	case 0x80007028:
		return 116 // ESTALE
	case 0x80007031:
		return 122 // EDQUOT
	case 0x8000702f:
		return 125 // ECANCELED
	case 0x8000703e:
		return 130 // EOWNERDEAD
	case 0x8000703d:
		return 131 // ENOTRECOVERABLE
	}
	if int32(raw) >= 0 {
		return raw
	}
	return 5 // EIO fallback for unmapped negative codes
}
