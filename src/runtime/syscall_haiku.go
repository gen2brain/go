// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/abi"
	"unsafe"
)

// Bridges from package syscall into asmsyscall6. Errnos are Haiku's native
// values, zero-extended from 32 bits.

// Haiku has no numbered-syscall API. Syscall and RawSyscall return EINVAL
// so callers fall back to the named libc wrappers.
const einval = ^uintptr(-_EINVAL - 1)

//go:nosplit
//go:linkname syscall_Syscall
func syscall_Syscall(trap, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	return 0, 0, einval
}

//go:nosplit
//go:linkname syscall_Syscall6
func syscall_Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	return 0, 0, einval
}

//go:nosplit
//go:linkname syscall_RawSyscall
func syscall_RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	return 0, 0, einval
}

//go:nosplit
//go:linkname syscall_RawSyscall6
func syscall_RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	return 0, 0, einval
}

// syscall_syscall, syscall_syscall6 and syscall_syscall9 call the libroot
// function at fn, a libc_X_trampoline, with 3, 6 or 9 argument words. More
// than 6 are only needed on 386, where a 64-bit argument takes two words.
//
//go:nosplit
//go:cgo_unsafe_args
//go:linkname syscall_syscall
func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	entersyscallblock()
	r1, r2, err = libcCall(fn, 3, unsafe.Pointer(&a1))
	exitsyscall()
	return
}

//go:nosplit
//go:cgo_unsafe_args
//go:linkname syscall_syscall6
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	entersyscallblock()
	r1, r2, err = libcCall(fn, 6, unsafe.Pointer(&a1))
	exitsyscall()
	return
}

//go:nosplit
//go:cgo_unsafe_args
//go:linkname syscall_syscall9
func syscall_syscall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2, err uintptr) {
	entersyscallblock()
	r1, r2, err = libcCall(fn, 9, unsafe.Pointer(&a1))
	exitsyscall()
	return
}

//go:nosplit
//go:cgo_unsafe_args
//go:linkname syscall_rawSyscall
func syscall_rawSyscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	return libcCall(fn, 3, unsafe.Pointer(&a1))
}

//go:nosplit
//go:cgo_unsafe_args
//go:linkname syscall_rawSyscall6
func syscall_rawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	return libcCall(fn, 6, unsafe.Pointer(&a1))
}

//go:nosplit
func libcCall(fn, n uintptr, args unsafe.Pointer) (r1, r2, err uintptr) {
	c := libcall{
		fn:   fn,
		n:    n,
		args: uintptr(args),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
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
		fn:   abi.FuncPCABI0(libc_chdir_trampoline),
		n:    1,
		args: uintptr(unsafe.Pointer(&path)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.err
}

//go:nosplit
//go:linkname syscall_chroot
//go:cgo_unsafe_args
func syscall_chroot(path uintptr) (err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_chroot_trampoline),
		n:    1,
		args: uintptr(unsafe.Pointer(&path)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.err
}

//go:nosplit
//go:linkname syscall_close
//go:cgo_unsafe_args
func syscall_close(fd uintptr) (err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_close_trampoline),
		n:    1,
		args: uintptr(unsafe.Pointer(&fd)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.err
}

//go:nosplit
//go:linkname syscall_dup2
//go:cgo_unsafe_args
func syscall_dup2(old, new uintptr) (val, err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_dup2_trampoline),
		n:    2,
		args: uintptr(unsafe.Pointer(&old)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.r1, c.err
}

//go:nosplit
//go:linkname syscall_execve
//go:cgo_unsafe_args
func syscall_execve(path, argv, envp uintptr) (err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_execve_trampoline),
		n:    3,
		args: uintptr(unsafe.Pointer(&path)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.err
}

//go:nosplit
//go:linkname syscall_exit
//go:cgo_unsafe_args
func syscall_exit(code uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_exit_trampoline),
		n:    1,
		args: uintptr(unsafe.Pointer(&code)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
}

//go:nosplit
//go:linkname syscall_fcntl
//go:cgo_unsafe_args
func syscall_fcntl(fd, cmd, arg uintptr) (val, err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_fcntl_trampoline),
		n:    3,
		args: uintptr(unsafe.Pointer(&fd)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.r1, c.err
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
		fn:   abi.FuncPCABI0(libc_fork_trampoline),
		n:    0,
		args: uintptr(unsafe.Pointer(&flags)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	if c.err != 0 {
		return 0, c.err
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
		fn:   abi.FuncPCABI0(libc_getpid_trampoline),
		n:    0,
		args: abi.FuncPCABI0(libc_getpid_trampoline),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.r1, 0
}

//go:nosplit
//go:linkname syscall_ioctl
//go:cgo_unsafe_args
func syscall_ioctl(fd, req, arg uintptr) (err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_ioctl_trampoline),
		n:    3,
		args: uintptr(unsafe.Pointer(&fd)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.err
}

//go:nosplit
//go:linkname syscall_setgid
//go:cgo_unsafe_args
func syscall_setgid(gid uintptr) (err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_setgid_trampoline),
		n:    1,
		args: uintptr(unsafe.Pointer(&gid)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.err
}

//go:nosplit
//go:linkname syscall_setgroups
//go:cgo_unsafe_args
func syscall_setgroups(ngid, gid uintptr) (err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_setgroups_trampoline),
		n:    2,
		args: uintptr(unsafe.Pointer(&ngid)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.err
}

//go:nosplit
//go:linkname syscall_setrlimit
//go:cgo_unsafe_args
func syscall_setrlimit(which uintptr, lim unsafe.Pointer) (err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_setrlimit_trampoline),
		n:    2,
		args: uintptr(unsafe.Pointer(&which)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.err
}

//go:nosplit
//go:linkname syscall_setsid
func syscall_setsid() (pid, err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_setsid_trampoline),
		n:    0,
		args: abi.FuncPCABI0(libc_setsid_trampoline),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.r1, c.err
}

//go:nosplit
//go:linkname syscall_setuid
//go:cgo_unsafe_args
func syscall_setuid(uid uintptr) (err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_setuid_trampoline),
		n:    1,
		args: uintptr(unsafe.Pointer(&uid)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.err
}

//go:nosplit
//go:linkname syscall_setpgid
//go:cgo_unsafe_args
func syscall_setpgid(pid, pgid uintptr) (err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_setpgid_trampoline),
		n:    2,
		args: uintptr(unsafe.Pointer(&pid)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.err
}

//go:nosplit
//go:linkname syscall_write
//go:cgo_unsafe_args
func syscall_write(fd, buf, nbyte uintptr) (n, err uintptr) {
	c := libcall{
		fn:   abi.FuncPCABI0(libc_write_trampoline),
		n:    3,
		args: uintptr(unsafe.Pointer(&fd)),
	}
	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))
	return c.r1, c.err
}
