// Code generated for haiku/arm64 libroot trampolines. DO NOT EDIT.

#include "textflag.h"

TEXT libc__errnop(SB),NOSPLIT,$0
	JMP	libc__errnop_dyn(SB)

TEXT libc_clock_gettime(SB),NOSPLIT,$0
	JMP	libc_clock_gettime_dyn(SB)

TEXT libc_close(SB),NOSPLIT,$0
	JMP	libc_close_dyn(SB)

TEXT libc_exit(SB),NOSPLIT,$0
	JMP	libc_exit_dyn(SB)

TEXT libc_fcntl(SB),NOSPLIT,$0
	JMP	libc_fcntl_dyn(SB)

TEXT libc_find_thread(SB),NOSPLIT,$0
	JMP	libc_find_thread_dyn(SB)

TEXT libc_getpid(SB),NOSPLIT,$0
	JMP	libc_getpid_dyn(SB)

TEXT libc_kill(SB),NOSPLIT,$0
	JMP	libc_kill_dyn(SB)

TEXT libc_madvise(SB),NOSPLIT,$0
	JMP	libc_madvise_dyn(SB)

TEXT libc_malloc(SB),NOSPLIT,$0
	JMP	libc_malloc_dyn(SB)

TEXT libc_mmap(SB),NOSPLIT,$0
	JMP	libc_mmap_dyn(SB)

TEXT libc_mprotect(SB),NOSPLIT,$0
	JMP	libc_mprotect_dyn(SB)

TEXT libc_munmap(SB),NOSPLIT,$0
	JMP	libc_munmap_dyn(SB)

TEXT libc_open(SB),NOSPLIT,$0
	JMP	libc_open_dyn(SB)

TEXT libc_pipe(SB),NOSPLIT,$0
	JMP	libc_pipe_dyn(SB)

TEXT libc_poll(SB),NOSPLIT,$0
	JMP	libc_poll_dyn(SB)

TEXT libc_raise(SB),NOSPLIT,$0
	JMP	libc_raise_dyn(SB)

TEXT libc_read(SB),NOSPLIT,$0
	JMP	libc_read_dyn(SB)

TEXT libc_sched_yield(SB),NOSPLIT,$0
	JMP	libc_sched_yield_dyn(SB)

TEXT libc_sem_init(SB),NOSPLIT,$0
	JMP	libc_sem_init_dyn(SB)

TEXT libc_sem_post(SB),NOSPLIT,$0
	JMP	libc_sem_post_dyn(SB)

TEXT libc_sem_timedwait(SB),NOSPLIT,$0
	JMP	libc_sem_timedwait_dyn(SB)

TEXT libc_sem_wait(SB),NOSPLIT,$0
	JMP	libc_sem_wait_dyn(SB)

TEXT libc_send_signal(SB),NOSPLIT,$0
	JMP	libc_send_signal_dyn(SB)

TEXT libc_setitimer(SB),NOSPLIT,$0
	JMP	libc_setitimer_dyn(SB)

TEXT libc_shutdown(SB),NOSPLIT,$0
	JMP	libc_shutdown_dyn(SB)

TEXT libc_sigaction(SB),NOSPLIT,$0
	JMP	libc_sigaction_dyn(SB)

TEXT libc_sigaltstack(SB),NOSPLIT,$0
	JMP	libc_sigaltstack_dyn(SB)

TEXT libc_sigprocmask(SB),NOSPLIT,$0
	JMP	libc_sigprocmask_dyn(SB)

TEXT libc_sysconf(SB),NOSPLIT,$0
	JMP	libc_sysconf_dyn(SB)

TEXT libc_usleep(SB),NOSPLIT,$0
	JMP	libc_usleep_dyn(SB)

TEXT libc_write(SB),NOSPLIT,$0
	JMP	libc_write_dyn(SB)

TEXT libc_getuid(SB),NOSPLIT,$0
	JMP	libc_getuid_dyn(SB)

TEXT libc_geteuid(SB),NOSPLIT,$0
	JMP	libc_geteuid_dyn(SB)

TEXT libc_getgid(SB),NOSPLIT,$0
	JMP	libc_getgid_dyn(SB)

TEXT libc_getegid(SB),NOSPLIT,$0
	JMP	libc_getegid_dyn(SB)

TEXT libc_pthread_attr_destroy(SB),NOSPLIT,$0
	JMP	libc_pthread_attr_destroy_dyn(SB)

TEXT libc_pthread_attr_init(SB),NOSPLIT,$0
	JMP	libc_pthread_attr_init_dyn(SB)

TEXT libc_pthread_attr_setstacksize(SB),NOSPLIT,$0
	JMP	libc_pthread_attr_setstacksize_dyn(SB)

TEXT libc_pthread_attr_setdetachstate(SB),NOSPLIT,$0
	JMP	libc_pthread_attr_setdetachstate_dyn(SB)

TEXT libc_pthread_create(SB),NOSPLIT,$0
	JMP	libc_pthread_create_dyn(SB)

TEXT libc_pthread_self(SB),NOSPLIT,$0
	JMP	libc_pthread_self_dyn(SB)

TEXT libc_tls_allocate(SB),NOSPLIT,$0
	JMP	libc_tls_allocate_dyn(SB)

TEXT libc_chdir(SB),NOSPLIT,$0
	JMP	libc_chdir_dyn(SB)

TEXT libc_chroot(SB),NOSPLIT,$0
	JMP	libc_chroot_dyn(SB)

TEXT libc_dup2(SB),NOSPLIT,$0
	JMP	libc_dup2_dyn(SB)

TEXT libc_execve(SB),NOSPLIT,$0
	JMP	libc_execve_dyn(SB)

TEXT libc_fork(SB),NOSPLIT,$0
	JMP	libc_fork_dyn(SB)

TEXT libc_ioctl(SB),NOSPLIT,$0
	JMP	libc_ioctl_dyn(SB)

TEXT libc_setgid(SB),NOSPLIT,$0
	JMP	libc_setgid_dyn(SB)

TEXT libc_setgroups(SB),NOSPLIT,$0
	JMP	libc_setgroups_dyn(SB)

TEXT libc_setpgid(SB),NOSPLIT,$0
	JMP	libc_setpgid_dyn(SB)

TEXT libc_setrlimit(SB),NOSPLIT,$0
	JMP	libc_setrlimit_dyn(SB)

TEXT libc_setsid(SB),NOSPLIT,$0
	JMP	libc_setsid_dyn(SB)

TEXT libc_setuid(SB),NOSPLIT,$0
	JMP	libc_setuid_dyn(SB)


// syscall-package libroot imports (unique to package syscall)
TEXT libc_writev(SB),NOSPLIT,$0
	JMP	libc_writev_dyn(SB)

TEXT libc_pread(SB),NOSPLIT,$0
	JMP	libc_pread_dyn(SB)

TEXT libc_pwrite(SB),NOSPLIT,$0
	JMP	libc_pwrite_dyn(SB)

TEXT libc_lseek(SB),NOSPLIT,$0
	JMP	libc_lseek_dyn(SB)

TEXT libc_dup(SB),NOSPLIT,$0
	JMP	libc_dup_dyn(SB)

TEXT libc_fchdir(SB),NOSPLIT,$0
	JMP	libc_fchdir_dyn(SB)

TEXT libc_chmod(SB),NOSPLIT,$0
	JMP	libc_chmod_dyn(SB)

TEXT libc_fchmod(SB),NOSPLIT,$0
	JMP	libc_fchmod_dyn(SB)

TEXT libc_chown(SB),NOSPLIT,$0
	JMP	libc_chown_dyn(SB)

TEXT libc_lchown(SB),NOSPLIT,$0
	JMP	libc_lchown_dyn(SB)

TEXT libc_fchown(SB),NOSPLIT,$0
	JMP	libc_fchown_dyn(SB)

TEXT libc_link(SB),NOSPLIT,$0
	JMP	libc_link_dyn(SB)

TEXT libc_symlink(SB),NOSPLIT,$0
	JMP	libc_symlink_dyn(SB)

TEXT libc_readlink(SB),NOSPLIT,$0
	JMP	libc_readlink_dyn(SB)

TEXT libc_unlink(SB),NOSPLIT,$0
	JMP	libc_unlink_dyn(SB)

TEXT libc_unlinkat(SB),NOSPLIT,$0
	JMP	libc_unlinkat_dyn(SB)

TEXT libc_rename(SB),NOSPLIT,$0
	JMP	libc_rename_dyn(SB)

TEXT libc_renameat(SB),NOSPLIT,$0
	JMP	libc_renameat_dyn(SB)

TEXT libc_mkdir(SB),NOSPLIT,$0
	JMP	libc_mkdir_dyn(SB)

TEXT libc_mkdirat(SB),NOSPLIT,$0
	JMP	libc_mkdirat_dyn(SB)

TEXT libc_rmdir(SB),NOSPLIT,$0
	JMP	libc_rmdir_dyn(SB)

TEXT libc_truncate(SB),NOSPLIT,$0
	JMP	libc_truncate_dyn(SB)

TEXT libc_ftruncate(SB),NOSPLIT,$0
	JMP	libc_ftruncate_dyn(SB)

TEXT libc_mknod(SB),NOSPLIT,$0
	JMP	libc_mknod_dyn(SB)

TEXT libc_stat(SB),NOSPLIT,$0
	JMP	libc_stat_dyn(SB)

TEXT libc_fstat(SB),NOSPLIT,$0
	JMP	libc_fstat_dyn(SB)

TEXT libc_lstat(SB),NOSPLIT,$0
	JMP	libc_lstat_dyn(SB)

TEXT libc_fstatat(SB),NOSPLIT,$0
	JMP	libc_fstatat_dyn(SB)

TEXT libc_fchmodat(SB),NOSPLIT,$0
	JMP	libc_fchmodat_dyn(SB)

TEXT libc_fchownat(SB),NOSPLIT,$0
	JMP	libc_fchownat_dyn(SB)

TEXT libc_faccessat(SB),NOSPLIT,$0
	JMP	libc_faccessat_dyn(SB)

TEXT libc_linkat(SB),NOSPLIT,$0
	JMP	libc_linkat_dyn(SB)

TEXT libc_symlinkat(SB),NOSPLIT,$0
	JMP	libc_symlinkat_dyn(SB)

TEXT libc_readlinkat(SB),NOSPLIT,$0
	JMP	libc_readlinkat_dyn(SB)

TEXT libc_openat(SB),NOSPLIT,$0
	JMP	libc_openat_dyn(SB)

TEXT libc_utimes(SB),NOSPLIT,$0
	JMP	libc_utimes_dyn(SB)

TEXT libc_utimensat(SB),NOSPLIT,$0
	JMP	libc_utimensat_dyn(SB)

TEXT libc_futimes(SB),NOSPLIT,$0
	JMP	libc_futimes_dyn(SB)

TEXT libc_fsync(SB),NOSPLIT,$0
	JMP	libc_fsync_dyn(SB)

TEXT libc_sync(SB),NOSPLIT,$0
	JMP	libc_sync_dyn(SB)

TEXT libc_getcwd(SB),NOSPLIT,$0
	JMP	libc_getcwd_dyn(SB)

TEXT libc_getppid(SB),NOSPLIT,$0
	JMP	libc_getppid_dyn(SB)

TEXT libc_getpgrp(SB),NOSPLIT,$0
	JMP	libc_getpgrp_dyn(SB)

TEXT libc_getpgid(SB),NOSPLIT,$0
	JMP	libc_getpgid_dyn(SB)

TEXT libc_setreuid(SB),NOSPLIT,$0
	JMP	libc_setreuid_dyn(SB)

TEXT libc_setregid(SB),NOSPLIT,$0
	JMP	libc_setregid_dyn(SB)

TEXT libc_getgroups(SB),NOSPLIT,$0
	JMP	libc_getgroups_dyn(SB)

TEXT libc_waitpid(SB),NOSPLIT,$0
	JMP	libc_waitpid_dyn(SB)

TEXT libc_vfork(SB),NOSPLIT,$0
	JMP	libc_vfork_dyn(SB)

TEXT libc_uname(SB),NOSPLIT,$0
	JMP	libc_uname_dyn(SB)

TEXT libc_gettimeofday(SB),NOSPLIT,$0
	JMP	libc_gettimeofday_dyn(SB)

TEXT libc_getrusage(SB),NOSPLIT,$0
	JMP	libc_getrusage_dyn(SB)

TEXT libc_getrlimit(SB),NOSPLIT,$0
	JMP	libc_getrlimit_dyn(SB)

TEXT libc_kern_read_dir(SB),NOSPLIT,$0
	JMP	libc_kern_read_dir_dyn(SB)

TEXT libc_kern_open_dir(SB),NOSPLIT,$0
	JMP	libc_kern_open_dir_dyn(SB)

TEXT libc_kern_rewind_dir(SB),NOSPLIT,$0
	JMP	libc_kern_rewind_dir_dyn(SB)

TEXT libc_flock(SB),NOSPLIT,$0
	JMP	libc_flock_dyn(SB)

TEXT libc_find_path(SB),NOSPLIT,$0
	JMP	libc_find_path_dyn(SB)

TEXT libc_socket(SB),NOSPLIT,$0
	JMP	libc_socket_dyn(SB)

TEXT libc_socketpair(SB),NOSPLIT,$0
	JMP	libc_socketpair_dyn(SB)

TEXT libc_bind(SB),NOSPLIT,$0
	JMP	libc_bind_dyn(SB)

TEXT libc_connect(SB),NOSPLIT,$0
	JMP	libc_connect_dyn(SB)

TEXT libc_listen(SB),NOSPLIT,$0
	JMP	libc_listen_dyn(SB)

TEXT libc_accept(SB),NOSPLIT,$0
	JMP	libc_accept_dyn(SB)

TEXT libc_getsockname(SB),NOSPLIT,$0
	JMP	libc_getsockname_dyn(SB)

TEXT libc_getpeername(SB),NOSPLIT,$0
	JMP	libc_getpeername_dyn(SB)

TEXT libc_recvfrom(SB),NOSPLIT,$0
	JMP	libc_recvfrom_dyn(SB)

TEXT libc_sendto(SB),NOSPLIT,$0
	JMP	libc_sendto_dyn(SB)

TEXT libc_recvmsg(SB),NOSPLIT,$0
	JMP	libc_recvmsg_dyn(SB)

TEXT libc_sendmsg(SB),NOSPLIT,$0
	JMP	libc_sendmsg_dyn(SB)

TEXT libc_getsockopt(SB),NOSPLIT,$0
	JMP	libc_getsockopt_dyn(SB)

TEXT libc_setsockopt(SB),NOSPLIT,$0
	JMP	libc_setsockopt_dyn(SB)

TEXT libc_select(SB),NOSPLIT,$0
	JMP	libc_select_dyn(SB)

TEXT libc_umask(SB),NOSPLIT,$0
	JMP	libc_umask_dyn(SB)

TEXT libc_mkfifo(SB),NOSPLIT,$0
	JMP	libc_mkfifo_dyn(SB)

TEXT libc_tcgetpgrp(SB),NOSPLIT,$0
	JMP	libc_tcgetpgrp_dyn(SB)

TEXT libc_tcsetpgrp(SB),NOSPLIT,$0
	JMP	libc_tcsetpgrp_dyn(SB)

TEXT libc_getpriority(SB),NOSPLIT,$0
	JMP	libc_getpriority_dyn(SB)

TEXT libc_setpriority(SB),NOSPLIT,$0
	JMP	libc_setpriority_dyn(SB)

TEXT libc_kill_thread(SB),NOSPLIT,$0
	JMP	libc_kill_thread_dyn(SB)

TEXT libc_wait_for_thread(SB),NOSPLIT,$0
	JMP	libc_wait_for_thread_dyn(SB)

TEXT libc_get_next_thread_info(SB),NOSPLIT,$0
	JMP	libc_get_next_thread_info_dyn(SB)

