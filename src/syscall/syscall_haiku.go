// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Haiku system calls.

package syscall

import (
	"unsafe"
)

// Syscall, Syscall6, RawSyscall, RawSyscall6 are implemented in
// asm_haiku_amd64.s; they jump to runtime helpers. Haiku exposes no
// numbered syscall API, so Syscall always returns EINVAL.
func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// syscall6 / rawSyscall6 are bridged to the runtime's libcall machinery.
// They are used by the auto-generated wrappers in zsyscall_haiku_amd64.go.
func syscall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func rawSyscall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// haikuErrnoToPosix maps a raw libroot errno to the POSIX-positive value
// the syscall package exposes. Implemented in runtime/syscall_haiku.go.
func haikuErrnoToPosix(raw uintptr) uintptr

const ImplementsGetwd = true

// SockaddrDatalink is a placeholder; Haiku does not expose AF_LINK
// the same way BSDs do.
type SockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [46]int8
	raw    RawSockaddrDatalink
}

// dirent helpers used by package os and our own ParseDirent.
func direntIno(buf []byte) (uint64, bool) {
	return readInt(buf, unsafe.Offsetof(Dirent{}.Ino), unsafe.Sizeof(Dirent{}.Ino))
}
func direntReclen(buf []byte) (uint64, bool) {
	return readInt(buf, unsafe.Offsetof(Dirent{}.Reclen), unsafe.Sizeof(Dirent{}.Reclen))
}
func direntNamlen(buf []byte) (uint64, bool) {
	reclen, ok := direntReclen(buf)
	if !ok {
		return 0, false
	}
	return reclen - uint64(unsafe.Offsetof(Dirent{}.Name)), true
}

// Sockaddr methods.

func (sa *SockaddrInet4) sockaddr() (unsafe.Pointer, _Socklen, error) {
	if sa.Port < 0 || sa.Port > 0xFFFF {
		return nil, 0, EINVAL
	}
	sa.raw.Len = SizeofSockaddrInet4
	sa.raw.Family = AF_INET
	p := (*[2]byte)(unsafe.Pointer(&sa.raw.Port))
	p[0] = byte(sa.Port >> 8)
	p[1] = byte(sa.Port)
	for i := 0; i < len(sa.Addr); i++ {
		sa.raw.Addr[i] = sa.Addr[i]
	}
	return unsafe.Pointer(&sa.raw), SizeofSockaddrInet4, nil
}

func (sa *SockaddrInet6) sockaddr() (unsafe.Pointer, _Socklen, error) {
	if sa.Port < 0 || sa.Port > 0xFFFF {
		return nil, 0, EINVAL
	}
	sa.raw.Len = SizeofSockaddrInet6
	sa.raw.Family = AF_INET6
	p := (*[2]byte)(unsafe.Pointer(&sa.raw.Port))
	p[0] = byte(sa.Port >> 8)
	p[1] = byte(sa.Port)
	sa.raw.Scope_id = sa.ZoneId
	for i := 0; i < len(sa.Addr); i++ {
		sa.raw.Addr[i] = sa.Addr[i]
	}
	return unsafe.Pointer(&sa.raw), SizeofSockaddrInet6, nil
}

func (sa *SockaddrUnix) sockaddr() (unsafe.Pointer, _Socklen, error) {
	name := sa.Name
	n := len(name)
	if n >= len(sa.raw.Path) {
		return nil, 0, EINVAL
	}
	sa.raw.Family = AF_UNIX
	for i := 0; i < n; i++ {
		sa.raw.Path[i] = int8(name[i])
	}
	if n > 0 {
		sa.raw.Path[n] = 0
	}
	sa.raw.Len = byte(2 + n + 1)
	return unsafe.Pointer(&sa.raw), _Socklen(sa.raw.Len), nil
}

func anyToSockaddr(rsa *RawSockaddrAny) (Sockaddr, error) {
	switch rsa.Addr.Family {
	case AF_INET:
		pp := (*RawSockaddrInet4)(unsafe.Pointer(rsa))
		sa := new(SockaddrInet4)
		p := (*[2]byte)(unsafe.Pointer(&pp.Port))
		sa.Port = int(p[0])<<8 + int(p[1])
		for i := 0; i < len(sa.Addr); i++ {
			sa.Addr[i] = pp.Addr[i]
		}
		return sa, nil

	case AF_INET6:
		pp := (*RawSockaddrInet6)(unsafe.Pointer(rsa))
		sa := new(SockaddrInet6)
		p := (*[2]byte)(unsafe.Pointer(&pp.Port))
		sa.Port = int(p[0])<<8 + int(p[1])
		sa.ZoneId = pp.Scope_id
		for i := 0; i < len(sa.Addr); i++ {
			sa.Addr[i] = pp.Addr[i]
		}
		return sa, nil

	case AF_UNIX:
		pp := (*RawSockaddrUnix)(unsafe.Pointer(rsa))
		sa := new(SockaddrUnix)
		n := 0
		for n < len(pp.Path) && pp.Path[n] != 0 {
			n++
		}
		bytes := (*[len(pp.Path)]byte)(unsafe.Pointer(&pp.Path[0]))
		sa.Name = string(bytes[0:n])
		return sa, nil
	}
	return nil, EAFNOSUPPORT
}

// Wrappers in addition to the //sys-generated functions.

func Pipe(p []int) error {
	if len(p) != 2 {
		return EINVAL
	}
	var pp [2]int32
	err := pipe(&pp)
	if err == nil {
		p[0] = int(pp[0])
		p[1] = int(pp[1])
	}
	return err
}

func Pipe2(p []int, flags int) error {
	// Haiku has no pipe2; emulate by setting flags after pipe().
	if len(p) != 2 {
		return EINVAL
	}
	if err := Pipe(p); err != nil {
		return err
	}
	if flags&O_NONBLOCK != 0 {
		if err := SetNonblock(p[0], true); err != nil {
			return err
		}
		if err := SetNonblock(p[1], true); err != nil {
			return err
		}
	}
	if flags&O_CLOEXEC != 0 {
		fcntl(p[0], F_SETFD, FD_CLOEXEC)
		fcntl(p[1], F_SETFD, FD_CLOEXEC)
	}
	return nil
}

func Readlink(path string, buf []byte) (n int, err error) {
	return readlink(path, buf)
}

func Utimes(path string, tv []Timeval) error {
	if len(tv) != 2 {
		return EINVAL
	}
	return utimes(path, (*[2]Timeval)(unsafe.Pointer(&tv[0])))
}

func UtimesNano(path string, ts []Timespec) error {
	if len(ts) != 2 {
		return EINVAL
	}
	return utimensat(AT_FDCWD, path, (*[2]Timespec)(unsafe.Pointer(&ts[0])), 0)
}

func Futimes(fd int, tv []Timeval) error {
	if len(tv) != 2 {
		return EINVAL
	}
	return futimes(fd, (*[2]Timeval)(unsafe.Pointer(&tv[0])))
}

func Getwd() (string, error) {
	for sz := 4096; ; sz *= 2 {
		b := make([]byte, sz)
		err := getcwd(&b[0], uint64(sz))
		if err == nil {
			n := clen(b)
			return string(b[:n]), nil
		}
		if err != ERANGE {
			return "", err
		}
	}
}

func Getcwd(buf []byte) (n int, err error) {
	err = getcwd(&buf[0], uint64(len(buf)))
	if err == nil {
		i := 0
		for buf[i] != 0 {
			i++
		}
		n = i + 1
	}
	return
}

func Getgroups() ([]int, error) {
	n, err := getgroups(0, nil)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	if n < 0 || n > 1024 {
		return nil, EINVAL
	}
	a := make([]_Gid_t, n)
	n, err = getgroups(n, &a[0])
	if err != nil {
		return nil, err
	}
	gids := make([]int, n)
	for i, v := range a[0:n] {
		gids[i] = int(v)
	}
	return gids, nil
}

func Setgroups(gids []int) error {
	if len(gids) == 0 {
		return setgroups(0, nil)
	}
	a := make([]_Gid_t, len(gids))
	for i, v := range gids {
		a[i] = _Gid_t(v)
	}
	return setgroups(len(a), &a[0])
}

// FindImagePath returns the filesystem path of the running executable.
// It calls Haiku's find_path() with B_FIND_PATH_IMAGE_PATH and a code
// pointer inside this binary; the loader resolves that to the image path.
func FindImagePath() (string, error) {
	const B_FIND_PATH_IMAGE_PATH = 1000
	var buf [4096]byte
	codePtr := uintptr(unsafe.Pointer(&libc_find_path))
	r0, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_find_path)), 5,
		codePtr, B_FIND_PATH_IMAGE_PATH, 0,
		uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)), 0)
	if int32(r0) < 0 {
		return "", Errno(haikuErrnoToPosix(uintptr(uint32(int32(r0)))))
	}
	n := 0
	for n < len(buf) && buf[n] != 0 {
		n++
	}
	return string(buf[:n]), nil
}

// B_UNSUPPORTED returned by _kern_read_dir when the fd is a regular file
// descriptor rather than a directory descriptor. We promote on demand.
const _haikuBUnsupported = -2147459058

// ReadDirent reads directory entries. _kern_read_dir requires a dir FD;
// Go opens directories with open(), so on the first call we promote the
// plain FD via _kern_open_dir + dup2.
func ReadDirent(fd int, buf []byte) (n int, err error) {
	if len(buf) == 0 {
		return 0, nil
	}
	count, e := readDirRaw(fd, buf)
	if e == _haikuBUnsupported {
		if err := promoteToDirFD(fd); err != nil {
			return 0, err
		}
		count, e = readDirRaw(fd, buf)
	}
	if e != 0 {
		return 0, Errno(haikuErrnoToPosix(uintptr(uint32(int32(e)))))
	}
	if count == 0 {
		return 0, nil
	}
	off := 0
	for i := int32(0); i < count; i++ {
		if off+int(unsafe.Offsetof(Dirent{}.Reclen))+2 > len(buf) {
			break
		}
		reclen := *(*uint16)(unsafe.Pointer(&buf[off+int(unsafe.Offsetof(Dirent{}.Reclen))]))
		if reclen == 0 || off+int(reclen) > len(buf) {
			break
		}
		off += int(reclen)
	}
	return off, nil
}

func readDirRaw(fd int, buf []byte) (count int32, status int32) {
	r0, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_kern_read_dir)), 4,
		uintptr(fd), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)),
		0xFFFF, 0, 0)
	c := int32(r0)
	if c < 0 {
		return 0, c
	}
	return c, 0
}

func promoteToDirFD(fd int) error {
	r0, _, _ := syscall6(uintptr(unsafe.Pointer(&libc_kern_open_dir)), 2,
		uintptr(fd), 0, 0, 0, 0, 0)
	dirFD := int32(r0)
	if dirFD < 0 {
		return Errno(haikuErrnoToPosix(uintptr(uint32(dirFD))))
	}
	if err := Dup2(int(dirFD), fd); err != nil {
		Close(int(dirFD))
		return err
	}
	Close(int(dirFD))
	return nil
}

func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (int, error) {
	var status _C_int
	wpid, err := wait4(pid, &status, options, rusage)
	if wstatus != nil {
		*wstatus = WaitStatus(status)
	}
	return wpid, err
}

func Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error) {
	return mapper.Mmap(fd, offset, length, prot, flags)
}

func Munmap(b []byte) (err error) {
	return mapper.Munmap(b)
}

var mapper = &mmapper{
	active: make(map[*byte][]byte),
	mmap:   mmap,
	munmap: munmap,
}

// mmap is exposed via linkname (see linkname_unix.go) for hall-of-shame
// third-party packages. The signature must not change. See
// go.dev/issue/67401.

func munmap(addr uintptr, length uintptr) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_munmap)), 2, addr, length, 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func Madvise(b []byte, behav int) error {
	if len(b) == 0 {
		return nil
	}
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_madvise)), 3,
		uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), uintptr(behav), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

// Mprotect helper for memory mapping.
func Mprotect(b []byte, prot int) (err error) {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_mprotect)), 3,
		uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), uintptr(prot), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Mlock(b []byte) error     { return ENOSYS }
func Munlock(b []byte) error   { return ENOSYS }
func Mlockall(flags int) error { return ENOSYS }
func Munlockall() error        { return ENOSYS }

func Sysctl(name string) (string, error)       { return "", ENOSYS }
func SysctlUint32(name string) (uint32, error) { return 0, ENOSYS }
func Faccessat(dirfd int, path string, mode uint32, flags int) error {
	return faccessat(dirfd, path, mode, flags)
}
func Access(path string, mode uint32) error { return faccessat(AT_FDCWD, path, mode, 0) }

func Getsockname(fd int) (Sockaddr, error) {
	var rsa RawSockaddrAny
	var addrlen _Socklen = SizeofSockaddrAny
	if err := getsockname(fd, &rsa, &addrlen); err != nil {
		return nil, err
	}
	return anyToSockaddr(&rsa)
}

func Listen(fd, n int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_listen)), 2, uintptr(fd), uintptr(n), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

func Accept(fd int) (nfd int, sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	var len _Socklen = SizeofSockaddrAny
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_accept)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(&rsa)), uintptr(unsafe.Pointer(&len)), 0, 0, 0)
	if e1 != 0 {
		return 0, nil, errnoErr(e1)
	}
	nfd = int(r0)
	sa, err = anyToSockaddr(&rsa)
	if err != nil {
		Close(nfd)
		return 0, nil, err
	}
	return nfd, sa, nil
}
func Accept4(fd, flags int) (nfd int, sa Sockaddr, err error) {
	return Accept(fd)
}

func GetsockoptIPv6MTUInfo(fd, level, opt int) (*IPv6MTUInfo, error) {
	return nil, ENOSYS
}
func GetsockoptUcred(fd, level, opt int) (*Ucred, error) { return nil, ENOSYS }

func Open(path string, mode int, perm uint32) (fd int, err error) {
	return open(path, mode, perm)
}

func Close(fd int) error { return closefd(fd) }

// Read/Write/Pread/Pwrite are provided by syscall_unix.go and route to
// the lowercase read/write/pread/pwrite defined in zsyscall_haiku_amd64.go.

func Seek(fd int, offset int64, whence int) (int64, error) {
	return lseek(fd, offset, whence)
}

func Stat(path string, stat *Stat_t) error  { return stat_(path, stat) }
func Lstat(path string, stat *Stat_t) error { return lstat(path, stat) }
func Fstat(fd int, stat *Stat_t) error      { return fstat(fd, stat) }
func Fstatat(dirfd int, path string, stat *Stat_t, flags int) error {
	return fstatat(dirfd, path, stat, flags)
}
func Statfs(path string, buf *Statfs_t) error { return ENOSYS }
func Fstatfs(fd int, buf *Statfs_t) error     { return ENOSYS }

func Chmod(path string, mode uint32) error           { return chmod(path, mode) }
func Fchmod(fd int, mode uint32) error               { return fchmod(fd, mode) }
func Chown(path string, uid, gid int) error          { return chown(path, uid, gid) }
func Lchown(path string, uid, gid int) error         { return lchown(path, uid, gid) }
func Fchown(fd int, uid, gid int) error              { return fchown(fd, uid, gid) }
func Truncate(path string, length int64) error       { return truncate(path, length) }
func Ftruncate(fd int, length int64) error           { return ftruncate(fd, length) }
func Mkdir(path string, mode uint32) error           { return mkdir(path, mode) }
func Rmdir(path string) error                        { return rmdir(path) }
func Unlink(path string) error                       { return unlink(path) }
func Rename(from, to string) error                   { return rename(from, to) }
func Link(from, to string) error                     { return link(from, to) }
func Symlink(from, to string) error                  { return symlink(from, to) }
func Fchdir(fd int) error                            { return fchdir(fd) }
func Mknod(path string, mode uint32, dev int) error  { return mknod(path, mode, dev) }
func Chroot(path string) error                       { return chroot(path) }
func Sync()                                          { libcSync() }
func Fsync(fd int) error                             { return fsync(fd) }
func Dup(oldfd int) (int, error)                     { return dup(oldfd) }
func Dup2(oldfd, newfd int) error                    { return dup2(oldfd, newfd) }
func Dup3(oldfd, newfd, flags int) error             { return dup2(oldfd, newfd) }
func Getdents(fd int, buf []byte) (n int, err error) { return read(fd, buf) }
func Mkfifo(path string, mode uint32) error          { return mkfifo(path, mode) }
func Umask(mask int) (oldmask int)                   { return umask(mask) }
func Tcgetpgrp(fd int) (pgid int32, err error)       { return tcgetpgrp(fd) }
func Tcsetpgrp(fd int, pgid int32) error             { return tcsetpgrp(fd, pgid) }
func Getpriority(which, who int) (int, error)        { return getpriority(which, who) }
func Setpriority(which, who, prio int) error         { return setpriority(which, who, prio) }

// FcntlFlock performs a fcntl syscall for the [F_GETLK], [F_SETLK] or [F_SETLKW] command.
func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_fcntl)), 3,
		fd, uintptr(cmd), uintptr(unsafe.Pointer(lk)), 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

// IDs / process info.

func Getuid() int  { return int(getuidRaw()) }
func Geteuid() int { return int(geteuid()) }
func Getgid() int  { return int(getgid()) }
func Getegid() int { return int(getegid()) }
func Getpid() int  { return int(getpidRaw()) }
func Getppid() int { return int(getppid()) }
func Getpgrp() int { pid, _ := getpgrp(); return pid }
func Getpgid(pid int) (int, error) {
	r, err := getpgid(pid)
	return r, err
}
func Setreuid(ruid, euid int) error { return setreuid(ruid, euid) }
func Setregid(rgid, egid int) error { return setregid(rgid, egid) }
func Getrlimit(which int, lim *Rlimit) error {
	return getrlimitRaw(which, lim)
}
func Kill(pid int, sig Signal) error {
	return kill(pid, int(sig))
}
func Gettimeofday(tv *Timeval) error { return gettimeofday(tv) }
func Getrusage(who int, rusage *Rusage) error {
	return getrusage(who, rusage)
}

// Hostname.
func Gethostname() (string, error) {
	var u Utsname
	err := uname(&u)
	if err != nil {
		return "", err
	}
	n := 0
	for n < len(u.Nodename) && u.Nodename[n] != 0 {
		n++
	}
	return string(u.Nodename[:n]), nil
}

func Uname(u *Utsname) error { return uname(u) }

// SetNonblock and CloseOnExec are provided by exec_unix.go.

func Flock(fd int, how int) error {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_flock)), 2,
		uintptr(fd), uintptr(how), 0, 0, 0, 0)
	if e1 != 0 {
		return errnoErr(e1)
	}
	return nil
}

// Statfs_t and Ucred placeholders for tools that import them.
type Statfs_t struct {
	Type   int64
	Bsize  int64
	Blocks uint64
	Bfree  uint64
	Bavail uint64
	Files  uint64
	Ffree  uint64
	Fsid   [2]int32
	Spare  [4]int64
}

type Ucred struct {
	Pid int32
	Uid uint32
	Gid uint32
}

// PtraceRegs / Tms placeholders for cross-platform sources.
type PtraceRegs struct{}
type Tms struct{}

// WaitStatus follows Haiku's <sys/wait.h> macro layout, which packs every
// field into separate byte ranges:
//
//	bits  0..7  exit code (WIFEXITED when bytes 8..23 are all zero)
//	bits  8..15 termination signal (WIFSIGNALED when nonzero)
//	bit   16    "core dumped" indicator (WIFCORED)
//	bit   17    "continued" indicator (WIFCONTINUED)
//	bits 16..23 stop signal (WIFSTOPPED); note this overlaps the cored/
//	            continued bits, so we must order our checks: a status that
//	            is signaled-with-core also has bits 16..23 nonzero.
type WaitStatus uint32

func (w WaitStatus) Exited() bool {
	return w&0xffff00 == 0
}

func (w WaitStatus) ExitStatus() int {
	if !w.Exited() {
		return -1
	}
	return int(w & 0xff)
}

func (w WaitStatus) Signaled() bool {
	return (w>>8)&0xff != 0
}

func (w WaitStatus) Signal() Signal {
	if !w.Signaled() {
		return -1
	}
	return Signal((w >> 8) & 0xff)
}

func (w WaitStatus) CoreDump() bool {
	return w.Signaled() && w&0x10000 != 0
}

func (w WaitStatus) Stopped() bool {
	return !w.Exited() && !w.Signaled() && !w.Continued() && (w>>16)&0xff != 0
}

func (w WaitStatus) StopSignal() Signal {
	if !w.Stopped() {
		return -1
	}
	return Signal((w >> 16) & 0xff)
}

func (w WaitStatus) Continued() bool {
	return w&0x20000 != 0
}

func (w WaitStatus) TrapCause() int {
	return -1
}

// SetLen sets the Cmsghdr length field.
func (cmsg *Cmsghdr) SetLen(length int) { cmsg.Len = uint32(length) }

const SCM_RIGHTS = 0x1

// recvmsgRaw / sendmsgN feed the public Recvmsg / Sendmsg in syscall_unix.go.
func recvmsgRaw(fd int, p, oob []byte, flags int, rsa *RawSockaddrAny) (n, oobn int, recvflags int, err error) {
	var msg Msghdr
	msg.Name = (*byte)(unsafe.Pointer(rsa))
	msg.Namelen = uint32(SizeofSockaddrAny)
	var iov Iovec
	if len(p) > 0 {
		iov.Base = (*byte)(unsafe.Pointer(&p[0]))
		iov.SetLen(len(p))
	}
	var dummy byte
	if len(oob) > 0 {
		// receive at least one normal byte to retrieve cmsghdrs
		if len(p) == 0 {
			iov.Base = &dummy
			iov.SetLen(1)
		}
		msg.Control = (*byte)(unsafe.Pointer(&oob[0]))
		msg.Controllen = uint32(len(oob))
	}
	msg.Iov = &iov
	msg.Iovlen = 1
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_recvmsg)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(&msg)), uintptr(flags), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
		return
	}
	n = int(r0)
	oobn = int(msg.Controllen)
	recvflags = int(msg.Flags)
	return
}

func sendmsgN(fd int, p, oob []byte, ptr unsafe.Pointer, salen _Socklen, flags int) (n int, err error) {
	var msg Msghdr
	msg.Name = (*byte)(ptr)
	msg.Namelen = uint32(salen)
	var iov Iovec
	if len(p) > 0 {
		iov.Base = (*byte)(unsafe.Pointer(&p[0]))
		iov.SetLen(len(p))
	}
	var dummy byte
	if len(oob) > 0 {
		if len(p) == 0 {
			iov.Base = &dummy
			iov.SetLen(1)
		}
		msg.Control = (*byte)(unsafe.Pointer(&oob[0]))
		msg.Controllen = uint32(len(oob))
	}
	msg.Iov = &iov
	msg.Iovlen = 1
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_sendmsg)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(&msg)), uintptr(flags), 0, 0, 0)
	if e1 != 0 {
		return 0, errnoErr(e1)
	}
	return int(r0), nil
}

// Misc constants exec_unix.go needs.

const (
	RLIMIT_NOFILE = 8
	RLIMIT_AS     = 6
	RLIMIT_CORE   = 4
	RLIMIT_CPU    = 0
	RLIMIT_DATA   = 2
	RLIMIT_FSIZE  = 1
	RLIMIT_NPROC  = 7
	RLIMIT_RSS    = 5
	RLIMIT_STACK  = 3
)

const (
	PRIO_PROCESS = 0
	PRIO_PGRP    = 1
	PRIO_USER    = 2
)

const (
	TCIFLUSH  = 0x01
	TCOFLUSH  = 0x02
	TCIOFLUSH = 0x03
)

const (
	RUSAGE_SELF     = 0
	RUSAGE_CHILDREN = -1
	RUSAGE_THREAD   = 1
)

const SYS_EXECVE = 0


// readlen wraps read for exec_unix.go.
func readlen(fd int, buf *byte, nbuf int) (int, error) {
	var p []byte = unsafe.Slice(buf, nbuf)
	r0, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_read)), 3,
		uintptr(fd), uintptr(unsafe.Pointer(&p[0])), uintptr(nbuf), 0, 0, 0)
	if e1 != 0 {
		return int(r0), e1
	}
	return int(r0), nil
}

// Errno aliases for symbols other Unix flavors export but Haiku does not
// have at the libc layer. Mapped to the closest Haiku error so that
// errors.Is checks in net behave reasonably.
const (
	ENOTSUPP        = ENOTSUP
	EPROTOOPT       = ENOPROTOOPT
	EREMOTE         = ENOTSUP
	ESOCKTNOSUPPORT = EPROTONOSUPPORT
	ETOOMANYREFS    = ENOBUFS
	EUSERS          = ENOMEM
)
