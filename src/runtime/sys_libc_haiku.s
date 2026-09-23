// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// Haiku's i386 PLT needs EBX set to the GOT; jump to the address instead.
#ifdef GOARCH_386
#define LIBCALL(fn) LEAL fn(SB), AX; JMP AX
#else
#define LIBCALL(fn) JMP fn(SB)
#endif

TEXT runtime·libc__errnop_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc__errnop)

TEXT runtime·libc_clock_gettime_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_clock_gettime)

TEXT runtime·libc_close_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_close)

TEXT runtime·libc_exit_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_exit)

TEXT runtime·libc_fcntl_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_fcntl)

TEXT runtime·libc_find_thread_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_find_thread)

TEXT runtime·libc_getpid_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_getpid)

TEXT runtime·libc_kill_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_kill)

TEXT runtime·libc_madvise_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_madvise)

TEXT runtime·libc_malloc_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_malloc)

TEXT runtime·libc_mmap_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_mmap)

TEXT runtime·libc_mprotect_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_mprotect)

TEXT runtime·libc_munmap_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_munmap)

TEXT runtime·libc_open_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_open)

TEXT runtime·libc_pipe_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_pipe)

TEXT runtime·libc_poll_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_poll)

TEXT runtime·libc_raise_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_raise)

TEXT runtime·libc_read_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_read)

TEXT runtime·libc_sched_yield_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_sched_yield)

TEXT runtime·libc_sem_init_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_sem_init)

TEXT runtime·libc_sem_post_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_sem_post)

TEXT runtime·libc_sem_timedwait_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_sem_timedwait)

TEXT runtime·libc_sem_wait_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_sem_wait)

TEXT runtime·libc_send_signal_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_send_signal)

TEXT runtime·libc_setitimer_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_setitimer)

TEXT runtime·libc_shutdown_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_shutdown)

TEXT runtime·libc_sigaction_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_sigaction)

TEXT runtime·libc_sigaltstack_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_sigaltstack)

TEXT runtime·libc_sigprocmask_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_sigprocmask)

TEXT runtime·libc_sysconf_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_sysconf)

TEXT runtime·libc_system_time_nsecs_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_system_time_nsecs)

TEXT runtime·libc_usleep_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_usleep)

TEXT runtime·libc_write_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_write)

TEXT runtime·libc_getuid_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_getuid)

TEXT runtime·libc_geteuid_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_geteuid)

TEXT runtime·libc_getgid_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_getgid)

TEXT runtime·libc_getegid_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_getegid)

TEXT runtime·libc_pthread_attr_destroy_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_pthread_attr_destroy)

TEXT runtime·libc_pthread_attr_init_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_pthread_attr_init)

TEXT runtime·libc_pthread_attr_setstacksize_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_pthread_attr_setstacksize)

TEXT runtime·libc_pthread_attr_setdetachstate_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_pthread_attr_setdetachstate)

TEXT runtime·libc_pthread_create_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_pthread_create)

TEXT runtime·libc_pthread_self_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_pthread_self)

TEXT runtime·libc_tls_allocate_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_tls_allocate)

TEXT runtime·libc_chdir_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_chdir)

TEXT runtime·libc_chroot_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_chroot)

TEXT runtime·libc_dup2_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_dup2)

TEXT runtime·libc_execve_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_execve)

TEXT runtime·libc_fork_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_fork)

TEXT runtime·libc_ioctl_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_ioctl)

TEXT runtime·libc_setgid_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_setgid)

TEXT runtime·libc_setgroups_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_setgroups)

TEXT runtime·libc_setpgid_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_setpgid)

TEXT runtime·libc_setrlimit_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_setrlimit)

TEXT runtime·libc_setsid_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_setsid)

TEXT runtime·libc_setuid_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_setuid)
