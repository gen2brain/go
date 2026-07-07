// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// libroot wrappers. Each Go function calls a libroot symbol via the
// runtime's asmsyscall6 trampoline, bridged through syscall6.

//go:build haiku && arm64

package syscall

import "unsafe"

//go:cgo_import_dynamic libc_open_dyn open "libroot.so"
//go:cgo_import_dynamic libc_close_dyn close "libroot.so"
//go:cgo_import_dynamic libc_read_dyn read "libroot.so"
//go:cgo_import_dynamic libc_write_dyn write "libroot.so"
//go:cgo_import_dynamic libc_writev_dyn writev "libroot.so"
//go:cgo_import_dynamic libc_pread_dyn pread "libroot.so"
//go:cgo_import_dynamic libc_pwrite_dyn pwrite "libroot.so"
//go:cgo_import_dynamic libc_lseek_dyn lseek "libroot.so"
//go:cgo_import_dynamic libc_dup_dyn dup "libroot.so"
//go:cgo_import_dynamic libc_dup2_dyn dup2 "libroot.so"
//go:cgo_import_dynamic libc_pipe_dyn pipe "libroot.so"
//go:cgo_import_dynamic libc_fcntl_dyn fcntl "libroot.so"
//go:cgo_import_dynamic libc_chdir_dyn chdir "libroot.so"
//go:cgo_import_dynamic libc_fchdir_dyn fchdir "libroot.so"
//go:cgo_import_dynamic libc_chmod_dyn chmod "libroot.so"
//go:cgo_import_dynamic libc_fchmod_dyn fchmod "libroot.so"
//go:cgo_import_dynamic libc_chown_dyn chown "libroot.so"
//go:cgo_import_dynamic libc_lchown_dyn lchown "libroot.so"
//go:cgo_import_dynamic libc_fchown_dyn fchown "libroot.so"
//go:cgo_import_dynamic libc_chroot_dyn chroot "libroot.so"
//go:cgo_import_dynamic libc_link_dyn link "libroot.so"
//go:cgo_import_dynamic libc_symlink_dyn symlink "libroot.so"
//go:cgo_import_dynamic libc_readlink_dyn readlink "libroot.so"
//go:cgo_import_dynamic libc_unlink_dyn unlink "libroot.so"
//go:cgo_import_dynamic libc_unlinkat_dyn unlinkat "libroot.so"
//go:cgo_import_dynamic libc_rename_dyn rename "libroot.so"
//go:cgo_import_dynamic libc_renameat_dyn renameat "libroot.so"
//go:cgo_import_dynamic libc_mkdir_dyn mkdir "libroot.so"
//go:cgo_import_dynamic libc_mkdirat_dyn mkdirat "libroot.so"
//go:cgo_import_dynamic libc_rmdir_dyn rmdir "libroot.so"
//go:cgo_import_dynamic libc_truncate_dyn truncate "libroot.so"
//go:cgo_import_dynamic libc_ftruncate_dyn ftruncate "libroot.so"
//go:cgo_import_dynamic libc_mknod_dyn mknod "libroot.so"
//go:cgo_import_dynamic libc_stat_dyn stat#LIBROOT_1_ALPHA1 "libroot.so"
//go:cgo_import_dynamic libc_fstat_dyn fstat#LIBROOT_1_ALPHA1 "libroot.so"
//go:cgo_import_dynamic libc_lstat_dyn lstat#LIBROOT_1_ALPHA1 "libroot.so"
//go:cgo_import_dynamic libc_fstatat_dyn fstatat#LIBROOT_1_ALPHA1 "libroot.so"
//go:cgo_import_dynamic libc_fchmodat_dyn fchmodat "libroot.so"
//go:cgo_import_dynamic libc_fchownat_dyn fchownat "libroot.so"
//go:cgo_import_dynamic libc_faccessat_dyn faccessat "libroot.so"
//go:cgo_import_dynamic libc_linkat_dyn linkat "libroot.so"
//go:cgo_import_dynamic libc_symlinkat_dyn symlinkat "libroot.so"
//go:cgo_import_dynamic libc_readlinkat_dyn readlinkat "libroot.so"
//go:cgo_import_dynamic libc_openat_dyn openat "libroot.so"
//go:cgo_import_dynamic libc_utimes_dyn utimes "libroot.so"
//go:cgo_import_dynamic libc_utimensat_dyn utimensat "libroot.so"
//go:cgo_import_dynamic libc_futimes_dyn futimes "libroot.so"
//go:cgo_import_dynamic libc_fsync_dyn fsync "libroot.so"
//go:cgo_import_dynamic libc_sync_dyn sync "libroot.so"
//go:cgo_import_dynamic libc_getcwd_dyn getcwd "libroot.so"
//go:cgo_import_dynamic libc_getuid_dyn getuid "libroot.so"
//go:cgo_import_dynamic libc_geteuid_dyn geteuid "libroot.so"
//go:cgo_import_dynamic libc_getgid_dyn getgid "libroot.so"
//go:cgo_import_dynamic libc_getegid_dyn getegid "libroot.so"
//go:cgo_import_dynamic libc_getpid_dyn getpid "libroot.so"
//go:cgo_import_dynamic libc_getppid_dyn getppid "libroot.so"
//go:cgo_import_dynamic libc_getpgrp_dyn getpgrp "libroot.so"
//go:cgo_import_dynamic libc_getpgid_dyn getpgid "libroot.so"
//go:cgo_import_dynamic libc_setpgid_dyn setpgid "libroot.so"
//go:cgo_import_dynamic libc_setsid_dyn setsid "libroot.so"
//go:cgo_import_dynamic libc_setuid_dyn setuid "libroot.so"
//go:cgo_import_dynamic libc_setgid_dyn setgid "libroot.so"
//go:cgo_import_dynamic libc_setreuid_dyn setreuid "libroot.so"
//go:cgo_import_dynamic libc_setregid_dyn setregid "libroot.so"
//go:cgo_import_dynamic libc_getgroups_dyn getgroups "libroot.so"
//go:cgo_import_dynamic libc_setgroups_dyn setgroups "libroot.so"
//go:cgo_import_dynamic libc_kill_dyn kill "libroot.so"
//go:cgo_import_dynamic libc_waitpid_dyn waitpid "libroot.so"
//go:cgo_import_dynamic libc_fork_dyn fork "libroot.so"
//go:cgo_import_dynamic libc_vfork_dyn vfork "libroot.so"
//go:cgo_import_dynamic libc_execve_dyn execve "libroot.so"
//go:cgo_import_dynamic libc_exit_dyn _exit "libroot.so"
//go:cgo_import_dynamic libc_ioctl_dyn ioctl "libroot.so"
//go:cgo_import_dynamic libc_uname_dyn uname "libroot.so"
//go:cgo_import_dynamic libc_gettimeofday_dyn gettimeofday "libroot.so"
//go:cgo_import_dynamic libc_getrusage_dyn getrusage#LIBROOT_1_BETA3 "libroot.so"
//go:cgo_import_dynamic libc_getrlimit_dyn getrlimit "libroot.so"
//go:cgo_import_dynamic libc_setrlimit_dyn setrlimit "libroot.so"
//go:cgo_import_dynamic libc_kern_read_dir_dyn _kern_read_dir "libroot.so"
//go:cgo_import_dynamic libc_kern_open_dir_dyn _kern_open_dir "libroot.so"
//go:cgo_import_dynamic libc_kern_rewind_dir_dyn _kern_rewind_dir "libroot.so"
//go:cgo_import_dynamic libc_flock_dyn flock "libroot.so"
//go:cgo_import_dynamic libc_find_path_dyn find_path "libroot.so"
//go:cgo_import_dynamic libc_socket_dyn socket "libnetwork.so"
//go:cgo_import_dynamic libc_socketpair_dyn socketpair "libnetwork.so"
//go:cgo_import_dynamic libc_bind_dyn bind "libnetwork.so"
//go:cgo_import_dynamic libc_connect_dyn connect "libnetwork.so"
//go:cgo_import_dynamic libc_listen_dyn listen "libnetwork.so"
//go:cgo_import_dynamic libc_accept_dyn accept "libnetwork.so"
//go:cgo_import_dynamic libc_getsockname_dyn getsockname "libnetwork.so"
//go:cgo_import_dynamic libc_getpeername_dyn getpeername "libnetwork.so"
//go:cgo_import_dynamic libc_recvfrom_dyn recvfrom "libnetwork.so"
//go:cgo_import_dynamic libc_sendto_dyn sendto "libnetwork.so"
//go:cgo_import_dynamic libc_recvmsg_dyn recvmsg "libnetwork.so"
//go:cgo_import_dynamic libc_sendmsg_dyn sendmsg "libnetwork.so"
//go:cgo_import_dynamic libc_shutdown_dyn shutdown "libnetwork.so"
//go:cgo_import_dynamic libc_getsockopt_dyn getsockopt "libnetwork.so"
//go:cgo_import_dynamic libc_setsockopt_dyn setsockopt "libnetwork.so"
//go:cgo_import_dynamic libc_mmap_dyn mmap "libroot.so"
//go:cgo_import_dynamic libc_munmap_dyn munmap "libroot.so"
//go:cgo_import_dynamic libc_madvise_dyn madvise "libroot.so"
//go:cgo_import_dynamic libc_mprotect_dyn mprotect "libroot.so"
//go:cgo_import_dynamic libc_select_dyn select "libroot.so"
//go:cgo_import_dynamic libc_umask_dyn umask "libroot.so"
//go:cgo_import_dynamic libc_mkfifo_dyn mkfifo "libroot.so"
//go:cgo_import_dynamic libc_tcgetpgrp_dyn tcgetpgrp "libroot.so"
//go:cgo_import_dynamic libc_tcsetpgrp_dyn tcsetpgrp "libroot.so"
//go:cgo_import_dynamic libc_getpriority_dyn getpriority "libroot.so"
//go:cgo_import_dynamic libc_setpriority_dyn setpriority "libroot.so"
//go:cgo_import_dynamic libc_find_thread_dyn find_thread "libroot.so"
//go:cgo_import_dynamic libc_kill_thread_dyn kill_thread "libroot.so"
//go:cgo_import_dynamic libc_wait_for_thread_dyn wait_for_thread "libroot.so"
//go:cgo_import_dynamic libc_get_next_thread_info_dyn _get_next_thread_info "libroot.so"

//go:linkname libc_open libc_open
//go:linkname libc_close libc_close
//go:linkname libc_read libc_read
//go:linkname libc_write libc_write
//go:linkname libc_writev libc_writev
//go:linkname libc_pread libc_pread
//go:linkname libc_pwrite libc_pwrite
//go:linkname libc_lseek libc_lseek
//go:linkname libc_dup libc_dup
//go:linkname libc_dup2 libc_dup2
//go:linkname libc_pipe libc_pipe
//go:linkname libc_fcntl libc_fcntl
//go:linkname libc_chdir libc_chdir
//go:linkname libc_fchdir libc_fchdir
//go:linkname libc_chmod libc_chmod
//go:linkname libc_fchmod libc_fchmod
//go:linkname libc_chown libc_chown
//go:linkname libc_lchown libc_lchown
//go:linkname libc_fchown libc_fchown
//go:linkname libc_chroot libc_chroot
//go:linkname libc_link libc_link
//go:linkname libc_symlink libc_symlink
//go:linkname libc_readlink libc_readlink
//go:linkname libc_unlink libc_unlink
//go:linkname libc_unlinkat libc_unlinkat
//go:linkname libc_rename libc_rename
//go:linkname libc_renameat libc_renameat
//go:linkname libc_mkdir libc_mkdir
//go:linkname libc_mkdirat libc_mkdirat
//go:linkname libc_rmdir libc_rmdir
//go:linkname libc_truncate libc_truncate
//go:linkname libc_ftruncate libc_ftruncate
//go:linkname libc_mknod libc_mknod
//go:linkname libc_stat libc_stat
//go:linkname libc_fstat libc_fstat
//go:linkname libc_lstat libc_lstat
//go:linkname libc_fstatat libc_fstatat
//go:linkname libc_fchmodat libc_fchmodat
//go:linkname libc_fchownat libc_fchownat
//go:linkname libc_faccessat libc_faccessat
//go:linkname libc_linkat libc_linkat
//go:linkname libc_symlinkat libc_symlinkat
//go:linkname libc_readlinkat libc_readlinkat
//go:linkname libc_openat libc_openat
//go:linkname libc_utimes libc_utimes
//go:linkname libc_utimensat libc_utimensat
//go:linkname libc_futimes libc_futimes
//go:linkname libc_fsync libc_fsync
//go:linkname libc_sync libc_sync
//go:linkname libc_getcwd libc_getcwd
//go:linkname libc_getuid libc_getuid
//go:linkname libc_geteuid libc_geteuid
//go:linkname libc_getgid libc_getgid
//go:linkname libc_getegid libc_getegid
//go:linkname libc_getpid libc_getpid
//go:linkname libc_getppid libc_getppid
//go:linkname libc_getpgrp libc_getpgrp
//go:linkname libc_getpgid libc_getpgid
//go:linkname libc_setpgid libc_setpgid
//go:linkname libc_setsid libc_setsid
//go:linkname libc_setuid libc_setuid
//go:linkname libc_setgid libc_setgid
//go:linkname libc_setreuid libc_setreuid
//go:linkname libc_setregid libc_setregid
//go:linkname libc_getgroups libc_getgroups
//go:linkname libc_setgroups libc_setgroups
//go:linkname libc_kill libc_kill
//go:linkname libc_waitpid libc_waitpid
//go:linkname libc_fork libc_fork
//go:linkname libc_vfork libc_vfork
//go:linkname libc_execve libc_execve
//go:linkname libc_exit libc_exit
//go:linkname libc_ioctl libc_ioctl
//go:linkname libc_uname libc_uname
//go:linkname libc_gettimeofday libc_gettimeofday
//go:linkname libc_getrusage libc_getrusage
//go:linkname libc_getrlimit libc_getrlimit
//go:linkname libc_setrlimit libc_setrlimit
//go:linkname libc_kern_read_dir libc_kern_read_dir
//go:linkname libc_kern_open_dir libc_kern_open_dir
//go:linkname libc_kern_rewind_dir libc_kern_rewind_dir
//go:linkname libc_flock libc_flock
//go:linkname libc_find_path libc_find_path
//go:linkname libc_socket libc_socket
//go:linkname libc_socketpair libc_socketpair
//go:linkname libc_bind libc_bind
//go:linkname libc_connect libc_connect
//go:linkname libc_listen libc_listen
//go:linkname libc_accept libc_accept
//go:linkname libc_getsockname libc_getsockname
//go:linkname libc_getpeername libc_getpeername
//go:linkname libc_recvfrom libc_recvfrom
//go:linkname libc_sendto libc_sendto
//go:linkname libc_recvmsg libc_recvmsg
//go:linkname libc_sendmsg libc_sendmsg
//go:linkname libc_shutdown libc_shutdown
//go:linkname libc_getsockopt libc_getsockopt
//go:linkname libc_setsockopt libc_setsockopt
//go:linkname libc_mmap libc_mmap
//go:linkname libc_munmap libc_munmap
//go:linkname libc_madvise libc_madvise
//go:linkname libc_mprotect libc_mprotect
//go:linkname libc_select libc_select
//go:linkname libc_umask libc_umask
//go:linkname libc_mkfifo libc_mkfifo
//go:linkname libc_tcgetpgrp libc_tcgetpgrp
//go:linkname libc_tcsetpgrp libc_tcsetpgrp
//go:linkname libc_getpriority libc_getpriority
//go:linkname libc_setpriority libc_setpriority
//go:linkname libc_find_thread libc_find_thread
//go:linkname libc_kill_thread libc_kill_thread
//go:linkname libc_wait_for_thread libc_wait_for_thread
//go:linkname libc_get_next_thread_info libc_get_next_thread_info

type libcFunc uintptr

var (
	libc_open,
	libc_close,
	libc_read,
	libc_write,
	libc_writev,
	libc_pread,
	libc_pwrite,
	libc_lseek,
	libc_dup,
	libc_dup2,
	libc_pipe,
	libc_fcntl,
	libc_chdir,
	libc_fchdir,
	libc_chmod,
	libc_fchmod,
	libc_chown,
	libc_lchown,
	libc_fchown,
	libc_chroot,
	libc_link,
	libc_symlink,
	libc_readlink,
	libc_unlink,
	libc_unlinkat,
	libc_rename,
	libc_renameat,
	libc_mkdir,
	libc_mkdirat,
	libc_rmdir,
	libc_truncate,
	libc_ftruncate,
	libc_mknod,
	libc_stat,
	libc_fstat,
	libc_lstat,
	libc_fstatat,
	libc_fchmodat,
	libc_fchownat,
	libc_faccessat,
	libc_linkat,
	libc_symlinkat,
	libc_readlinkat,
	libc_openat,
	libc_utimes,
	libc_utimensat,
	libc_futimes,
	libc_fsync,
	libc_sync,
	libc_getcwd,
	libc_getuid,
	libc_geteuid,
	libc_getgid,
	libc_getegid,
	libc_getpid,
	libc_getppid,
	libc_getpgrp,
	libc_getpgid,
	libc_setpgid,
	libc_setsid,
	libc_setuid,
	libc_setgid,
	libc_setreuid,
	libc_setregid,
	libc_getgroups,
	libc_setgroups,
	libc_kill,
	libc_waitpid,
	libc_fork,
	libc_vfork,
	libc_execve,
	libc_exit,
	libc_ioctl,
	libc_uname,
	libc_gettimeofday,
	libc_getrusage,
	libc_getrlimit,
	libc_setrlimit,
	libc_kern_read_dir,
	libc_kern_open_dir,
	libc_kern_rewind_dir,
	libc_flock,
	libc_find_path,
	libc_socket,
	libc_socketpair,
	libc_bind,
	libc_connect,
	libc_listen,
	libc_accept,
	libc_getsockname,
	libc_getpeername,
	libc_recvfrom,
	libc_sendto,
	libc_recvmsg,
	libc_sendmsg,
	libc_shutdown,
	libc_getsockopt,
	libc_setsockopt,
	libc_mmap,
	libc_munmap,
	libc_madvise,
	libc_mprotect,
	libc_select,
	libc_umask,
	libc_mkfifo,
	libc_tcgetpgrp,
	libc_tcsetpgrp,
	libc_getpriority,
	libc_setpriority,
	libc_find_thread,
	libc_kill_thread,
	libc_wait_for_thread,
	libc_get_next_thread_info libcFunc
)

// File I/O.

func open(path string, mode int, perm uint32) (fd int, err error) {
	var p *byte
	p, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_open)), 3,
		uintptr(unsafe.Pointer(p)), uintptr(mode), uintptr(perm), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func closefd(fd int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_close)), 1,
		uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func read(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_read)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func write(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_write)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

//go:linkname writev
func writev(fd int, iovecs []Iovec) (n uintptr, err error) {
	var _p0 *Iovec
	if len(iovecs) > 0 {
		_p0 = &iovecs[0]
	}
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_writev)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(iovecs)), 0, 0, 0)
	n = r0
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func pread(fd int, p []byte, off int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_pread)), 4,
		uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(off), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func pwrite(fd int, p []byte, off int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_pwrite)), 4,
		uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(off), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func lseek(fd int, offset int64, whence int) (int64, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_lseek)), 3,
		uintptr(fd), uintptr(offset), uintptr(whence), 0, 0, 0)
	if e1 != 0 {
		// lseek rejects directory FDs with ESPIPE. Emulate POSIX's
		// rewind-to-zero via _kern_rewind_dir; fall back to ESPIPE for
		// non-directory FDs (it returns B_UNSUPPORTED).
		if e1 == ESPIPE && offset == 0 && whence == 0 {
			r1, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_kern_rewind_dir)), 1,
				uintptr(fd), 0, 0, 0, 0, 0)
			if int32(r1) >= 0 {
				return 0, nil
			}
		}
		return 0, errnoErr(e1)
	}
	return int64(r0), nil
}

func dup(fd int) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_dup)), 1,
		uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}

func dup2(old, new int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_dup2)), 2,
		uintptr(old), uintptr(new), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func pipe(p *[2]int32) error {
	_, _, e1 := rawSyscall6(uintptr(unsafe.Pointer(&libc_pipe)), 1,
		uintptr(unsafe.Pointer(p)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func fcntl(fd int, cmd int, arg int) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_fcntl)), 3,
		uintptr(fd), uintptr(cmd), uintptr(arg), 0, 0, 0)
	if e1 != 0 {
		return int(r0), errnoErr(e1)
	}
	return int(r0), nil
}

// Path operations.

func Chdir(path string) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_chdir)), 1,
		uintptr(unsafe.Pointer(p)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func fchdir(fd int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_fchdir)), 1,
		uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func chmod(path string, mode uint32) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_chmod)), 2,
		uintptr(unsafe.Pointer(p)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func fchmod(fd int, mode uint32) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_fchmod)), 2,
		uintptr(fd), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func chown(path string, uid, gid int) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_chown)), 3,
		uintptr(unsafe.Pointer(p)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func lchown(path string, uid, gid int) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_lchown)), 3,
		uintptr(unsafe.Pointer(p)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func fchown(fd int, uid, gid int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_fchown)), 3,
		uintptr(fd), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func chroot(path string) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_chroot)), 1,
		uintptr(unsafe.Pointer(p)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func link(from, to string) error {
	a, err := BytePtrFromString(from)
	if err != nil {
		return err
	}
	b, err := BytePtrFromString(to)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_link)), 2,
		uintptr(unsafe.Pointer(a)), uintptr(unsafe.Pointer(b)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func symlink(from, to string) error {
	a, err := BytePtrFromString(from)
	if err != nil {
		return err
	}
	b, err := BytePtrFromString(to)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_symlink)), 2,
		uintptr(unsafe.Pointer(a)), uintptr(unsafe.Pointer(b)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func readlink(path string, buf []byte) (int, error) {
	p, err := BytePtrFromString(path)
	if err != nil {
		return 0, err
	}
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_readlink)), 3,
		uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}

func unlink(path string) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_unlink)), 1,
		uintptr(unsafe.Pointer(p)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func unlinkat(dirfd int, path string, flags int) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_unlinkat)), 3,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(flags), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func rename(from, to string) error {
	a, err := BytePtrFromString(from)
	if err != nil {
		return err
	}
	b, err := BytePtrFromString(to)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_rename)), 2,
		uintptr(unsafe.Pointer(a)), uintptr(unsafe.Pointer(b)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func mkdir(path string, mode uint32) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_mkdir)), 2,
		uintptr(unsafe.Pointer(p)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func rmdir(path string) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_rmdir)), 1,
		uintptr(unsafe.Pointer(p)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func truncate(path string, length int64) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_truncate)), 2,
		uintptr(unsafe.Pointer(p)), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func ftruncate(fd int, length int64) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_ftruncate)), 2,
		uintptr(fd), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func mknod(path string, mode uint32, dev int) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_mknod)), 3,
		uintptr(unsafe.Pointer(p)), uintptr(mode), uintptr(dev), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func stat_(path string, stat *Stat_t) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_stat)), 2,
		uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func lstat(path string, stat *Stat_t) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_lstat)), 2,
		uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func fstat(fd int, stat *Stat_t) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_fstat)), 2,
		uintptr(fd), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func fstatat(dirfd int, path string, stat *Stat_t, flags int) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_fstatat)), 4,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(stat)), uintptr(flags), 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func faccessat(dirfd int, path string, mode uint32, flags int) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_faccessat)), 4,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(mode), uintptr(flags), 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func utimes(path string, tv *[2]Timeval) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_utimes)), 2,
		uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(tv)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func futimes(fd int, tv *[2]Timeval) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_futimes)), 2,
		uintptr(fd), uintptr(unsafe.Pointer(tv)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func utimensat(dirfd int, path string, ts *[2]Timespec, flags int) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_utimensat)), 4,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(ts)), uintptr(flags), 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func fsync(fd int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_fsync)), 1,
		uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func libcSync() {
	syscall6(uintptr(unsafe.Pointer(&libc_sync)), 0, 0, 0, 0, 0, 0, 0)
}

func getcwd(buf *byte, sz uint64) error {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_getcwd)), 2,
		uintptr(unsafe.Pointer(buf)), uintptr(sz), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	if r0 == 0 {
		return EINVAL
	}
	return nil
}

// IDs.

func getuidRaw() int32 {
	r0, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_getuid)), 0, 0, 0, 0, 0, 0, 0)
	return int32(r0)
}
func geteuid() int32 {
	r0, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_geteuid)), 0, 0, 0, 0, 0, 0, 0)
	return int32(r0)
}
func getgid() int32 {
	r0, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_getgid)), 0, 0, 0, 0, 0, 0, 0)
	return int32(r0)
}
func getegid() int32 {
	r0, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_getegid)), 0, 0, 0, 0, 0, 0, 0)
	return int32(r0)
}
func getpidRaw() int32 {
	r0, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_getpid)), 0, 0, 0, 0, 0, 0, 0)
	return int32(r0)
}
func getppid() int32 {
	r0, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_getppid)), 0, 0, 0, 0, 0, 0, 0)
	return int32(r0)
}
func getpgrp() (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_getpgrp)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}
func getpgid(pid int) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_getpgid)), 1,
		uintptr(pid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}
func Setpgid(pid, pgid int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_setpgid)), 2,
		uintptr(pid), uintptr(pgid), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}
func Setsid() (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_setsid)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}
func Setuid(uid int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_setuid)), 1,
		uintptr(uid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}
func Setgid(gid int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_setgid)), 1,
		uintptr(gid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}
func setreuid(ruid, euid int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_setreuid)), 2,
		uintptr(ruid), uintptr(euid), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}
func setregid(rgid, egid int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_setregid)), 2,
		uintptr(rgid), uintptr(egid), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}
func getgroups(ngid int, gid *_Gid_t) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_getgroups)), 2,
		uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}
func setgroups(ngid int, gid *_Gid_t) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_setgroups)), 2,
		uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func kill(pid int, sig int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_kill)), 2,
		uintptr(pid), uintptr(sig), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

// libroot has no wait4(); use waitpid() and zero the rusage. Callers
// needing real rusage must call Getrusage separately.
func wait4(pid int, status *_C_int, options int, rusage *Rusage) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_waitpid)), 3,
		uintptr(pid), uintptr(unsafe.Pointer(status)), uintptr(options),
		0, 0, 0)
	if int32(r0) < 0 {
		if e1 != 0 {
			return -1, errnoErr(e1)
		}
		return -1, EINVAL
	}
	if rusage != nil {
		*rusage = Rusage{}
	}
	return int(int32(r0)), nil
}

func uname(u *Utsname) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_uname)), 1,
		uintptr(unsafe.Pointer(u)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func gettimeofday(tv *Timeval) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_gettimeofday)), 2,
		uintptr(unsafe.Pointer(tv)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func getrusage(who int, rusage *Rusage) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_getrusage)), 2,
		uintptr(who), uintptr(unsafe.Pointer(rusage)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func getrlimitRaw(which int, lim *Rlimit) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_getrlimit)), 2,
		uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func setrlimit(which int, lim *Rlimit) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_setrlimit)), 2,
		uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

// Sockets.

func socket(domain, typ, proto int) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_socket)), 3,
		uintptr(domain), uintptr(typ), uintptr(proto), 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}

func socketpair(domain, typ, proto int, fd *[2]int32) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_socketpair)), 4,
		uintptr(domain), uintptr(typ), uintptr(proto), uintptr(unsafe.Pointer(fd)), 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func bind(s int, addr unsafe.Pointer, addrlen _Socklen) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_bind)), 3,
		uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func connect(s int, addr unsafe.Pointer, addrlen _Socklen) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_connect)), 3,
		uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_accept)), 3,
		uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_accept)), 3,
		uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}

func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_getsockname)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_getpeername)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (int, error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_recvfrom)), 6,
		uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(flags),
		uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}

func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) error {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_sendto)), 6,
		uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags),
		uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func recvmsg(fd int, msg *Msghdr, flags int) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_recvmsg)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}

func sendmsg(fd int, msg *Msghdr, flags int) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_sendmsg)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}

func getsockopt(s int, level, opt int, val unsafe.Pointer, vallen *_Socklen) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_getsockopt)), 5,
		uintptr(s), uintptr(level), uintptr(opt), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func setsockopt(s int, level, opt int, val unsafe.Pointer, vallen uintptr) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_setsockopt)), 5,
		uintptr(s), uintptr(level), uintptr(opt), uintptr(val), vallen, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func Shutdown(fd int, how int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_shutdown)), 2,
		uintptr(fd), uintptr(how), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func Select(n int, r, w, e *FdSet, timeout *Timeval) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_select)), 5,
		uintptr(n), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)), 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}

func umask(mask int) (oldmask int) {
	r0, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_umask)), 1,
		uintptr(mask), 0, 0, 0, 0, 0)
	return int(r0)
}

func mkfifo(path string, mode uint32) error {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_mkfifo)), 2,
		uintptr(unsafe.Pointer(p)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func tcgetpgrp(fd int) (pgid int32, err error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_tcgetpgrp)), 1,
		uintptr(fd), 0, 0, 0, 0, 0)
	if int32(r0) < 0 {
		return -1, errnoErr(e1)
	}
	return int32(r0), nil
}

func tcsetpgrp(fd int, pgid int32) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_tcsetpgrp)), 2,
		uintptr(fd), uintptr(pgid), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func getpriority(which, who int) (int, error) {
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_getpriority)), 2,
		uintptr(which), uintptr(who), 0, 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(int32(r0)), nil
}

func setpriority(which, who, prio int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_setpriority)), 3,
		uintptr(which), uintptr(who), uintptr(prio), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}
