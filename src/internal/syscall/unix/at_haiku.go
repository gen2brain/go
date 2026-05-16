// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"syscall"
	"unsafe"
)

// Implemented as syscall_syscall6 in runtime/syscall_haiku.go.
func syscall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

// Implemented as syscall_rawSyscall6 in runtime/syscall_haiku.go.
func rawSyscall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

//go:cgo_import_dynamic libc_faccessat faccessat "libroot.so"
//go:cgo_import_dynamic libc_fchmodat fchmodat "libroot.so"
//go:cgo_import_dynamic libc_fchownat fchownat "libroot.so"
//go:cgo_import_dynamic libc_fstatat fstatat#LIBROOT_1_ALPHA1 "libroot.so"
//go:cgo_import_dynamic libc_linkat linkat "libroot.so"
//go:cgo_import_dynamic libc_openat openat "libroot.so"
//go:cgo_import_dynamic libc_renameat renameat "libroot.so"
//go:cgo_import_dynamic libc_symlinkat symlinkat "libroot.so"
//go:cgo_import_dynamic libc_unlinkat unlinkat "libroot.so"
//go:cgo_import_dynamic libc_readlinkat readlinkat "libroot.so"
//go:cgo_import_dynamic libc_mkdirat mkdirat "libroot.so"
//go:cgo_import_dynamic libc_utimensat utimensat "libroot.so"

//go:linkname procFaccessat libc_faccessat
//go:linkname procFchmodat libc_fchmodat
//go:linkname procFchownat libc_fchownat
//go:linkname procFstatat libc_fstatat
//go:linkname procLinkat libc_linkat
//go:linkname procOpenat libc_openat
//go:linkname procRenameat libc_renameat
//go:linkname procSymlinkat libc_symlinkat
//go:linkname procUnlinkat libc_unlinkat
//go:linkname procReadlinkat libc_readlinkat
//go:linkname procMkdirat libc_mkdirat
//go:linkname procUtimensat libc_utimensat

var (
	procFaccessat  uintptr
	procFchmodat   uintptr
	procFchownat   uintptr
	procFstatat    uintptr
	procLinkat     uintptr
	procOpenat     uintptr
	procRenameat   uintptr
	procSymlinkat  uintptr
	procUnlinkat   uintptr
	procReadlinkat uintptr
	procMkdirat    uintptr
	procUtimensat  uintptr
)

const (
	AT_EACCESS          = 0x08
	AT_FDCWD            = -100
	AT_REMOVEDIR        = 0x04
	AT_SYMLINK_NOFOLLOW = 0x01

	UTIME_OMIT = 1000000001
)

func faccessat(dirfd int, path string, mode uint32, flags int) error {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procFaccessat)), 4,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(mode), uintptr(flags), 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procFstatat)), 4,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(stat)), uintptr(flags), 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

func Openat(dirfd int, path string, flags int, perm uint32) (int, error) {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return 0, err
	}
	fd, _, errno := syscall6(uintptr(unsafe.Pointer(&procOpenat)), 4,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(flags), uintptr(perm), 0, 0)
	if errno != 0 {
		return 0, errno
	}
	return int(fd), nil
}

func Mkdirat(dirfd int, path string, mode uint32) error {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procMkdirat)), 3,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(mode), 0, 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

func Fchmodat(dirfd int, path string, mode uint32, flags int) error {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procFchmodat)), 4,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(mode), uintptr(flags), 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

func Fchownat(dirfd int, path string, uid, gid int, flags int) error {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procFchownat)), 5,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(uid), uintptr(gid), uintptr(flags), 0)
	if errno != 0 {
		return errno
	}
	return nil
}

func Linkat(oldDirfd int, oldPath string, newDirfd int, newPath string, flags int) error {
	op, err := syscall.BytePtrFromString(oldPath)
	if err != nil {
		return err
	}
	np, err := syscall.BytePtrFromString(newPath)
	if err != nil {
		return err
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procLinkat)), 5,
		uintptr(oldDirfd), uintptr(unsafe.Pointer(op)),
		uintptr(newDirfd), uintptr(unsafe.Pointer(np)),
		uintptr(flags), 0)
	if errno != 0 {
		return errno
	}
	return nil
}

func Renameat(oldDirfd int, oldPath string, newDirfd int, newPath string) error {
	op, err := syscall.BytePtrFromString(oldPath)
	if err != nil {
		return err
	}
	np, err := syscall.BytePtrFromString(newPath)
	if err != nil {
		return err
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procRenameat)), 4,
		uintptr(oldDirfd), uintptr(unsafe.Pointer(op)),
		uintptr(newDirfd), uintptr(unsafe.Pointer(np)),
		0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

func Symlinkat(oldPath string, newDirfd int, newPath string) error {
	op, err := syscall.BytePtrFromString(oldPath)
	if err != nil {
		return err
	}
	np, err := syscall.BytePtrFromString(newPath)
	if err != nil {
		return err
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procSymlinkat)), 3,
		uintptr(unsafe.Pointer(op)), uintptr(newDirfd), uintptr(unsafe.Pointer(np)),
		0, 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

func Readlinkat(dirfd int, path string, buf []byte) (int, error) {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return 0, err
	}
	var bufp *byte
	if len(buf) > 0 {
		bufp = &buf[0]
	}
	n, _, errno := syscall6(uintptr(unsafe.Pointer(&procReadlinkat)), 4,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)),
		uintptr(unsafe.Pointer(bufp)), uintptr(len(buf)), 0, 0)
	if errno != 0 {
		return 0, errno
	}
	return int(n), nil
}

func Unlinkat(dirfd int, path string, flags int) error {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procUnlinkat)), 3,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(flags), 0, 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

func utimensat(dirfd int, path string, times *[2]syscall.Timespec, flags int) error {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procUtimensat)), 4,
		uintptr(dirfd), uintptr(unsafe.Pointer(p)),
		uintptr(unsafe.Pointer(times)), uintptr(flags), 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}
