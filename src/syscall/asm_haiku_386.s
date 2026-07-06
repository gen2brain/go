// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

//
// System calls for 386 on Haiku are implemented in
// ../runtime/syscall_haiku.go. The trampolines below dispatch the
// exported syscall entry points to the matching runtime helpers.
//

TEXT ·syscall6(SB),NOSPLIT,$0
	JMP	runtime·syscall_syscall6(SB)

TEXT ·syscall7(SB),NOSPLIT,$0
	JMP	runtime·syscall_syscall7(SB)

TEXT ·rawSyscall6(SB),NOSPLIT,$0
	JMP	runtime·syscall_rawSyscall6(SB)

TEXT ·RawSyscall(SB),NOSPLIT,$0
	JMP	runtime·syscall_RawSyscall(SB)

TEXT ·Syscall(SB),NOSPLIT,$0
	JMP	runtime·syscall_Syscall(SB)

// Trampolines for the libc-syscall path used by ../syscall/exec_libc.go's
// forkAndExecInChild. Each forwards to a //go:nosplit runtime helper that
// invokes libroot via asmcgocall(&asmsyscall6).

TEXT ·chdir(SB),NOSPLIT,$0
	JMP	runtime·syscall_chdir(SB)

TEXT ·chroot1(SB),NOSPLIT,$0
	JMP	runtime·syscall_chroot(SB)

TEXT ·closeFD(SB),NOSPLIT,$0
	JMP	runtime·syscall_close(SB)

TEXT ·dup2child(SB),NOSPLIT,$0
	JMP	runtime·syscall_dup2(SB)

TEXT ·execve(SB),NOSPLIT,$0
	JMP	runtime·syscall_execve(SB)

TEXT ·exit(SB),NOSPLIT,$0
	JMP	runtime·syscall_exit(SB)

TEXT ·fcntl1(SB),NOSPLIT,$0
	JMP	runtime·syscall_fcntl(SB)

TEXT ·forkx(SB),NOSPLIT,$0
	JMP	runtime·syscall_forkx(SB)

TEXT ·getpid(SB),NOSPLIT,$0
	JMP	runtime·syscall_getpid(SB)

TEXT ·ioctl(SB),NOSPLIT,$0
	JMP	runtime·syscall_ioctl(SB)

TEXT ·setgid(SB),NOSPLIT,$0
	JMP	runtime·syscall_setgid(SB)

TEXT ·setgroups1(SB),NOSPLIT,$0
	JMP	runtime·syscall_setgroups(SB)

TEXT ·setrlimit1(SB),NOSPLIT,$0
	JMP	runtime·syscall_setrlimit(SB)

TEXT ·setsid(SB),NOSPLIT,$0
	JMP	runtime·syscall_setsid(SB)

TEXT ·setuid(SB),NOSPLIT,$0
	JMP	runtime·syscall_setuid(SB)

TEXT ·setpgid(SB),NOSPLIT,$0
	JMP	runtime·syscall_setpgid(SB)

TEXT ·write1(SB),NOSPLIT,$0
	JMP	runtime·syscall_write(SB)
