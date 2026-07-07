// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build haiku && arm64

package runtime

// On arm64 the address of a libroot import (&libc_X) can't be materialized
// with a PC-relative adrp+add in the ELF DYN output produced by external
// linking. Import each function under a _dyn name and expose libc_X as a
// local trampoline (in sys_haiku_arm64.s) that tail-jumps to it via a
// direct branch, which the (internal or host) linker routes through the PLT.
//go:cgo_import_dynamic libc__errnop_dyn _errnop "libroot.so"
//go:cgo_import_dynamic libc_clock_gettime_dyn clock_gettime "libroot.so"
//go:cgo_import_dynamic libc_close_dyn close "libroot.so"
//go:cgo_import_dynamic libc_exit_dyn _exit "libroot.so"
//go:cgo_import_dynamic libc_fcntl_dyn fcntl "libroot.so"
//go:cgo_import_dynamic libc_find_thread_dyn find_thread "libroot.so"
//go:cgo_import_dynamic libc_getpid_dyn getpid "libroot.so"
//go:cgo_import_dynamic libc_kill_dyn kill "libroot.so"
//go:cgo_import_dynamic libc_madvise_dyn madvise "libroot.so"
//go:cgo_import_dynamic libc_malloc_dyn malloc "libroot.so"
//go:cgo_import_dynamic libc_mmap_dyn mmap "libroot.so"
//go:cgo_import_dynamic libc_mprotect_dyn mprotect "libroot.so"
//go:cgo_import_dynamic libc_munmap_dyn munmap "libroot.so"
//go:cgo_import_dynamic libc_open_dyn open "libroot.so"
//go:cgo_import_dynamic libc_pipe_dyn pipe "libroot.so"
//go:cgo_import_dynamic libc_poll_dyn poll "libroot.so"
//go:cgo_import_dynamic libc_raise_dyn raise "libroot.so"
//go:cgo_import_dynamic libc_read_dyn read "libroot.so"
//go:cgo_import_dynamic libc_sched_yield_dyn sched_yield "libroot.so"
//go:cgo_import_dynamic libc_sem_init_dyn sem_init "libroot.so"
//go:cgo_import_dynamic libc_sem_post_dyn sem_post "libroot.so"
//go:cgo_import_dynamic libc_sem_timedwait_dyn sem_timedwait "libroot.so"
//go:cgo_import_dynamic libc_sem_wait_dyn sem_wait "libroot.so"
//go:cgo_import_dynamic libc_send_signal_dyn send_signal "libroot.so"
//go:cgo_import_dynamic libc_setitimer_dyn setitimer "libroot.so"
//go:cgo_import_dynamic libc_shutdown_dyn shutdown "libnetwork.so"
//go:cgo_import_dynamic libc_sigaction_dyn sigaction#LIBROOT_1_ALPHA4 "libroot.so"
//go:cgo_import_dynamic libc_sigaltstack_dyn sigaltstack "libroot.so"
//go:cgo_import_dynamic libc_sigprocmask_dyn sigprocmask#LIBROOT_1_ALPHA4 "libroot.so"
//go:cgo_import_dynamic libc_sysconf_dyn sysconf#LIBROOT_1_ALPHA4 "libroot.so"
//go:cgo_import_dynamic libc_usleep_dyn usleep "libroot.so"
//go:cgo_import_dynamic libc_write_dyn write "libroot.so"
//go:cgo_import_dynamic libc_getuid_dyn getuid "libroot.so"
//go:cgo_import_dynamic libc_geteuid_dyn geteuid "libroot.so"
//go:cgo_import_dynamic libc_getgid_dyn getgid "libroot.so"
//go:cgo_import_dynamic libc_getegid_dyn getegid "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_destroy_dyn pthread_attr_destroy "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_init_dyn pthread_attr_init "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_setstacksize_dyn pthread_attr_setstacksize "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_setdetachstate_dyn pthread_attr_setdetachstate "libroot.so"
//go:cgo_import_dynamic libc_pthread_create_dyn pthread_create "libroot.so"
//go:cgo_import_dynamic libc_pthread_self_dyn pthread_self "libroot.so"
//go:cgo_import_dynamic libc_tls_allocate_dyn tls_allocate "libroot.so"
//go:cgo_import_dynamic libc_chdir_dyn chdir "libroot.so"
//go:cgo_import_dynamic libc_chroot_dyn chroot "libroot.so"
//go:cgo_import_dynamic libc_dup2_dyn dup2 "libroot.so"
//go:cgo_import_dynamic libc_execve_dyn execve "libroot.so"
//go:cgo_import_dynamic libc_fork_dyn fork "libroot.so"
//go:cgo_import_dynamic libc_ioctl_dyn ioctl "libroot.so"
//go:cgo_import_dynamic libc_setgid_dyn setgid "libroot.so"
//go:cgo_import_dynamic libc_setgroups_dyn setgroups "libroot.so"
//go:cgo_import_dynamic libc_setpgid_dyn setpgid "libroot.so"
//go:cgo_import_dynamic libc_setrlimit_dyn setrlimit "libroot.so"
//go:cgo_import_dynamic libc_setsid_dyn setsid "libroot.so"
//go:cgo_import_dynamic libc_setuid_dyn setuid "libroot.so"
