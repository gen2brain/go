// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Haiku system calls.

package syscall

import (
	"internal/abi"
	"unsafe"
)

// Syscall, Syscall6, RawSyscall and RawSyscall6 jump to runtime helpers
// (asm_haiku.s). Haiku exposes no numbered syscall API, so they return EINVAL.
func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// haikuErrnoToPosix maps a raw libroot errno to the POSIX-positive value
// the syscall package exposes. Implemented in runtime/syscall_haiku.go.
func haikuErrnoToPosix(raw uintptr) uintptr

// syscall, syscall6, syscall9, rawSyscall and rawSyscall6 call the libroot
// function at fn, a libc_X_trampoline address. They are used by the wrappers
// generated into zsyscall_haiku_$GOARCH.go and are implemented in the runtime.
func syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)
func rawSyscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func rawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

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
	var pp [2]_C_int
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
		err := getcwd(b)
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
	err = getcwd(buf)
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
	status := findPath(abi.FuncPCABI0(libc_find_path_trampoline), B_FIND_PATH_IMAGE_PATH, nil, &buf[0], uintptr(len(buf)))
	if status < 0 {
		return "", Errno(haikuErrnoToPosix(uintptr(uint32(status))))
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
	c := kernReadDir(fd, buf, 0xFFFF)
	if c < 0 {
		return 0, c
	}
	return c, 0
}

func promoteToDirFD(fd int) error {
	dirFD := kernOpenDir(fd, nil)
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
	wpid, err := waitpid(pid, &status, options)
	if err != nil {
		return -1, err
	}
	if wstatus != nil {
		*wstatus = WaitStatus(status)
	}
	if rusage != nil {
		*rusage = Rusage{}
	}
	return wpid, nil
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

func Madvise(b []byte, behav int) error {
	if len(b) == 0 {
		return nil
	}
	return madvise(b, behav)
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

func Accept(fd int) (nfd int, sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	var len _Socklen = SizeofSockaddrAny
	nfd, err = accept(fd, &rsa, &len)
	if err != nil {
		return 0, nil, err
	}
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

func GetsockoptUcred(fd, level, opt int) (*Ucred, error) { return nil, ENOSYS }

// Seek rewinds directory descriptors through _kern_rewind_dir; Haiku's lseek rejects them.
func Seek(fd int, offset int64, whence int) (int64, error) {
	off, err := lseek(fd, offset, whence)
	if err == ESPIPE && offset == 0 && whence == 0 && kernRewindDir(fd) >= 0 {
		return 0, nil
	}
	return off, err
}

func Fstatat(dirfd int, path string, stat *Stat_t, flags int) error {
	return fstatat(dirfd, path, stat, flags)
}
func Statfs(path string, buf *Statfs_t) error { return ENOSYS }
func Fstatfs(fd int, buf *Statfs_t) error     { return ENOSYS }

func Dup3(oldfd, newfd, flags int) error             { return Dup2(oldfd, newfd) }
func Getdents(fd int, buf []byte) (n int, err error) { return read(fd, buf) }

func Getpriority(which, who int) (int, error) {
	prio, err := getpriority(which, who)
	return int(prio), err
}

// FcntlFlock performs a fcntl syscall for the [F_GETLK], [F_SETLK] or [F_SETLKW] command.
func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) error {
	return fcntlFlock(fd, cmd, lk)
}

func Getpgrp() int {
	pgid, _ := getpgrp()
	return pgid
}

func Kill(pid int, sig Signal) error {
	return kill(pid, int(sig))
}

// Hostname.
func Gethostname() (string, error) {
	var u Utsname
	err := Uname(&u)
	if err != nil {
		return "", err
	}
	b := make([]byte, 0, len(u.Nodename))
	for _, c := range u.Nodename {
		if c == 0 {
			break
		}
		b = append(b, byte(c))
	}
	return string(b), nil
}

// SetNonblock and CloseOnExec are provided by exec_unix.go.

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
	if n, err = recvmsg(fd, &msg, flags); err != nil {
		return
	}
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
	return sendmsg(fd, &msg, flags)
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

/*
 * Wrapped
 */

//sys	open(path string, mode int, perm uint32) (fd int, err error)
//sys	read(fd int, p []byte) (n int, err error)
//sys	readlen(fd int, buf *byte, nbuf int) (n int, err error) = read
//sys	write(fd int, p []byte) (n int, err error)
//sys	writev(fd int, iovecs []Iovec) (n uintptr, err error)
//sys	pread(fd int, p []byte, offset int64) (n int, err error)
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error)
//sys	lseek(fd int, offset int64, whence int) (newoffset int64, err error)
//sysnb	pipe(p *[2]_C_int) (err error)
//sys	fcntl(fd int, cmd int, arg int) (val int, err error)
//sys	fcntlFlock(fd uintptr, cmd int, lk *Flock_t) (err error) = fcntl
//sys	readlink(path string, buf []byte) (n int, err error)
//sys	unlinkat(dirfd int, path string, flags int) (err error)
//sys	fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = fstatat#LIBROOT_1_ALPHA1
//sys	faccessat(dirfd int, path string, mode uint32, flags int) (err error)
//sys	utimes(path string, timeval *[2]Timeval) (err error)
//sys	futimes(fd int, timeval *[2]Timeval) (err error)
//sys	utimensat(dirfd int, path string, times *[2]Timespec, flags int) (err error)
//sys	getcwd(buf []byte) (err error)
//sysnb	getpgrp() (pgid int, err error)
//sysnb	getgroups(ngid int, gid *_Gid_t) (n int, err error)
//sysnb	setgroups(ngid int, gid *_Gid_t) (err error)
//sys	kill(pid int, signum int) (err error)
//sys	waitpid(pid int, status *_C_int, options int) (wpid int, err error)
//sysnb	setrlimit(which int, lim *Rlimit) (err error)
//sys	getpriority(which int, who int) (prio int32, err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error)
//sys	munmap(addr uintptr, length uintptr) (err error)
//sys	madvise(b []byte, behav int) (err error)
//sysnb	socket(domain int, typ int, proto int) (fd int, err error) = libnetwork.socket
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) = libnetwork.socketpair
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) = libnetwork.bind
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) = libnetwork.connect
//sys	accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) = libnetwork.accept
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) = libnetwork.getsockname
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) = libnetwork.getpeername
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) = libnetwork.recvfrom
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) = libnetwork.sendto
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error) = libnetwork.recvmsg
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error) = libnetwork.sendmsg
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) = libnetwork.getsockopt
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) = libnetwork.setsockopt
//sys	findPath(codePointer uintptr, baseDirectory int, subPath *byte, pathBuffer *byte, bufferSize uintptr) (status int32) = find_path
//sys	kernReadDir(fd int, buf []byte, maxCount uint32) (count int32) = _kern_read_dir
//sys	kernOpenDir(fd int, path *byte) (dirfd int32) = _kern_open_dir
//sys	kernRewindDir(fd int) (status int32) = _kern_rewind_dir

/*
 * Exposed directly
 */

//sys	Open(path string, mode int, perm uint32) (fd int, err error)
//sys	Close(fd int) (err error)
//sys	Dup(fd int) (nfd int, err error)
//sys	Dup2(from int, to int) (err error)
//sys	Fsync(fd int) (err error)
//sys	Sync()
//sys	Chdir(path string) (err error)
//sys	Fchdir(fd int) (err error)
//sys	Chmod(path string, mode uint32) (err error)
//sys	Fchmod(fd int, mode uint32) (err error)
//sys	Chown(path string, uid int, gid int) (err error)
//sys	Lchown(path string, uid int, gid int) (err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	Chroot(path string) (err error)
//sys	Link(path string, link string) (err error)
//sys	Symlink(path string, link string) (err error)
//sys	Unlink(path string) (err error)
//sys	Rename(from string, to string) (err error)
//sys	Mkdir(path string, mode uint32) (err error)
//sys	Rmdir(path string) (err error)
//sys	Mknod(path string, mode uint32, dev int) (err error)
//sys	Mkfifo(path string, mode uint32) (err error)
//sys	Truncate(path string, length int64) (err error)
//sys	Ftruncate(fd int, length int64) (err error)
//sys	Stat(path string, stat *Stat_t) (err error) = stat#LIBROOT_1_ALPHA1
//sys	Lstat(path string, stat *Stat_t) (err error) = lstat#LIBROOT_1_ALPHA1
//sys	Fstat(fd int, stat *Stat_t) (err error) = fstat#LIBROOT_1_ALPHA1
//sysnb	Umask(newmask int) (oldmask int)
//sysnb	Getuid() (uid int)
//sysnb	Geteuid() (uid int)
//sysnb	Getgid() (gid int)
//sysnb	Getegid() (gid int)
//sysnb	Getpid() (pid int)
//sysnb	Getppid() (ppid int)
//sysnb	Getpgid(pid int) (pgid int, err error)
//sysnb	Setpgid(pid int, pgid int) (err error)
//sysnb	Setsid() (pid int, err error)
//sysnb	Setuid(uid int) (err error)
//sysnb	Setgid(gid int) (err error)
//sysnb	Setreuid(ruid int, euid int) (err error)
//sysnb	Setregid(rgid int, egid int) (err error)
//sysnb	Getrlimit(which int, lim *Rlimit) (err error)
//sysnb	Getrusage(who int, rusage *Rusage) (err error) = getrusage#LIBROOT_1_BETA3
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sysnb	Uname(buf *Utsname) (err error)
//sys	Tcgetpgrp(fd int) (pgid int32, err error)
//sys	Tcsetpgrp(fd int, pgid int32) (err error)
//sys	Setpriority(which int, who int, prio int) (err error)
//sys	Mprotect(b []byte, prot int) (err error)
//sys	Flock(fd int, how int) (err error)
//sys	Listen(s int, backlog int) (err error) = libnetwork.listen
//sys	Shutdown(s int, how int) (err error) = libnetwork.shutdown
//sys	Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error)
