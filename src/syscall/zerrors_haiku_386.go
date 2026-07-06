// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Errno, signal, and constant table for Haiku x86.

package syscall

const (
	// File / open() flags
	O_RDONLY  = 0x0
	O_WRONLY  = 0x1
	O_RDWR    = 0x2
	O_ACCMODE = 0x3

	O_CLOEXEC   = 0x40
	O_NONBLOCK  = 0x80
	O_EXCL      = 0x100
	O_CREAT     = 0x200
	O_TRUNC     = 0x400
	O_APPEND    = 0x800
	O_NOCTTY    = 0x1000
	O_SYNC      = 0x10000
	O_RSYNC     = 0x20000
	O_DSYNC     = 0x40000
	O_NOFOLLOW  = 0x80000
	O_DIRECT    = 0x100000
	O_DIRECTORY = 0x200000

	F_DUPFD         = 0x1
	F_GETFD         = 0x2
	F_SETFD         = 0x4
	F_GETFL         = 0x8
	F_SETFL         = 0x10
	F_GETLK         = 0x20
	F_SETLK         = 0x80
	F_SETLKW        = 0x100
	F_DUPFD_CLOEXEC = 0x200

	F_RDLCK = 0x40
	F_UNLCK = 0x200
	F_WRLCK = 0x400

	FD_CLOEXEC = 1

	// *at constants from <fcntl.h>.
	AT_FDCWD            = -100
	AT_EACCESS          = 0x08
	AT_SYMLINK_NOFOLLOW = 0x01
	AT_SYMLINK_FOLLOW   = 0x02
	AT_REMOVEDIR        = 0x04

	// flock(2) operations.
	LOCK_SH = 0x1
	LOCK_EX = 0x2
	LOCK_NB = 0x4
	LOCK_UN = 0x8

	// stat / mode bits (POSIX standard, Haiku follows)
	S_IFMT   = 0xf000
	S_IFIFO  = 0x1000
	S_IFCHR  = 0x2000
	S_IFDIR  = 0x4000
	S_IFBLK  = 0x6000
	S_IFREG  = 0x8000
	S_IFLNK  = 0xa000
	S_IFSOCK = 0xc000

	S_ISUID = 0x800
	S_ISGID = 0x400
	S_ISVTX = 0x200

	S_IRUSR = 0x100
	S_IWUSR = 0x80
	S_IXUSR = 0x40
	S_IRGRP = 0x20
	S_IWGRP = 0x10
	S_IXGRP = 0x8
	S_IROTH = 0x4
	S_IWOTH = 0x2
	S_IXOTH = 0x1

	// mmap protections / flags (mirrored from defs_haiku_amd64.go).
	PROT_NONE  = 0x0
	PROT_READ  = 0x1
	PROT_WRITE = 0x2
	PROT_EXEC  = 0x4

	MAP_ANON      = 0x8
	MAP_PRIVATE   = 0x2
	MAP_FIXED     = 0x4
	MAP_SHARED    = 0x1
	MAP_FILE      = 0x0
	MAP_ANONYMOUS = MAP_ANON
	MAP_NORESERVE = 0x10

	// signals (from Haiku <signal.h>)
	SIGHUP    = Signal(0x1)
	SIGINT    = Signal(0x2)
	SIGQUIT   = Signal(0x3)
	SIGILL    = Signal(0x4)
	SIGCHLD   = Signal(0x5)
	SIGABRT   = Signal(0x6)
	SIGIOT    = Signal(0x6)
	SIGPIPE   = Signal(0x7)
	SIGFPE    = Signal(0x8)
	SIGKILL   = Signal(0x9)
	SIGSTOP   = Signal(0xa)
	SIGSEGV   = Signal(0xb)
	SIGCONT   = Signal(0xc)
	SIGTSTP   = Signal(0xd)
	SIGALRM   = Signal(0xe)
	SIGTERM   = Signal(0xf)
	SIGTTIN   = Signal(0x10)
	SIGTTOU   = Signal(0x11)
	SIGUSR1   = Signal(0x12)
	SIGUSR2   = Signal(0x13)
	SIGWINCH  = Signal(0x14)
	SIGTRAP   = Signal(0x16)
	SIGPOLL   = Signal(0x17)
	SIGIO     = Signal(0x17)
	SIGPROF   = Signal(0x18)
	SIGSYS    = Signal(0x19)
	SIGURG    = Signal(0x1a)
	SIGVTALRM = Signal(0x1b)
	SIGXCPU   = Signal(0x1c)
	SIGXFSZ   = Signal(0x1d)
	SIGBUS    = Signal(0x1e)

	// Address families.
	AF_UNSPEC    = 0
	AF_INET      = 1
	AF_APPLETALK = 2
	AF_ROUTE     = 3
	AF_LINK      = 4
	AF_INET6     = 5
	AF_DLI       = 6
	AF_IPX       = 7
	AF_NOTIFY    = 8
	AF_LOCAL     = 9
	AF_UNIX      = AF_LOCAL
	AF_BLUETOOTH = 10
	AF_MAX       = 11

	// Socket types.
	SOCK_STREAM    = 1
	SOCK_DGRAM     = 2
	SOCK_RAW       = 3
	SOCK_SEQPACKET = 5
	SOCK_MISC      = 255
	SOCK_CLOEXEC   = 0x100
	SOCK_NONBLOCK  = 0x80

	// Protocols.
	IPPROTO_IP      = 0
	IPPROTO_HOPOPTS = 0
	IPPROTO_ICMP    = 1
	IPPROTO_IGMP    = 2
	IPPROTO_TCP     = 6
	IPPROTO_UDP     = 17
	IPPROTO_IPV6    = 41
	IPPROTO_ESP     = 50
	IPPROTO_AH      = 51
	IPPROTO_ICMPV6  = 58
	IPPROTO_RAW     = 255

	// Setsockopt levels and options (subset).
	SOL_SOCKET = -1

	SO_ACCEPTCONN   = 0x00000001
	SO_BROADCAST    = 0x00000002
	SO_DEBUG        = 0x00000004
	SO_DONTROUTE    = 0x00000008
	SO_KEEPALIVE    = 0x00000010
	SO_OOBINLINE    = 0x00000020
	SO_REUSEADDR    = 0x00000040
	SO_REUSEPORT    = 0x00000080
	SO_USELOOPBACK  = 0x00000100
	SO_LINGER       = 0x00000200
	SO_SNDBUF       = 0x40000001
	SO_SNDLOWAT     = 0x40000002
	SO_SNDTIMEO     = 0x40000003
	SO_RCVBUF       = 0x40000004
	SO_RCVLOWAT     = 0x40000005
	SO_RCVTIMEO     = 0x40000006
	SO_ERROR        = 0x40000007
	SO_TYPE         = 0x40000008
	SO_NONBLOCK     = 0x40000009
	SO_BINDTODEVICE = 0x4000000a
	SO_PEERCRED     = 0x4000000b

	// Shutdown.
	SHUT_RD   = 0
	SHUT_WR   = 1
	SHUT_RDWR = 2

	SOMAXCONN = 32

	// TCP-level options.
	TCP_NODELAY   = 0x01
	TCP_MAXSEG    = 0x02
	TCP_NOPUSH    = 0x04
	TCP_NOOPT     = 0x08
	TCP_KEEPALIVE = 0x10
	TCP_KEEPIDLE  = 0x100
	TCP_KEEPINTVL = 0x101
	TCP_KEEPCNT   = 0x102

	// IP options.
	IP_OPTIONS         = 1
	IP_HDRINCL         = 2
	IP_TOS             = 3
	IP_TTL             = 4
	IP_RECVOPTS        = 5
	IP_RECVRETOPTS     = 6
	IP_RECVDSTADDR     = 7
	IP_RETOPTS         = 8
	IP_MULTICAST_IF    = 9
	IP_MULTICAST_TTL   = 10
	IP_MULTICAST_LOOP  = 11
	IP_ADD_MEMBERSHIP  = 12
	IP_DROP_MEMBERSHIP = 13

	// IPv6 options.
	IPV6_MULTICAST_IF   = 24
	IPV6_MULTICAST_HOPS = 25
	IPV6_MULTICAST_LOOP = 26
	IPV6_UNICAST_HOPS   = 27
	IPV6_JOIN_GROUP     = 28
	IPV6_LEAVE_GROUP    = 29
	IPV6_V6ONLY         = 30
	IPV6_HOPLIMIT       = 31
	IPV6_HOPOPTS        = 32
	IPV6_PKTINFO        = 33
	IPV6_TCLASS         = 35
	IPV6_RECVPKTINFO    = 38
	IPV6_RECVHOPLIMIT   = 40
	IPV6_RECVHOPOPTS    = 41
	IPV6_RECVRTHDR      = 42
	IPV6_RECVTCLASS     = 44

	// MSG_* recv/send flags.
	MSG_PEEK      = 0x2
	MSG_OOB       = 0x1
	MSG_DONTROUTE = 0x4
	MSG_EOR       = 0x8
	MSG_TRUNC     = 0x10
	MSG_CTRUNC    = 0x20
	MSG_WAITALL   = 0x40
	MSG_DONTWAIT  = 0x80
	MSG_BCAST     = 0x100
	MSG_MCAST     = 0x200
	MSG_EOF       = 0x400
	MSG_NOSIGNAL      = 0x800
	MSG_CMSG_CLOEXEC  = 0x1000

	// Wait flags.
	WNOHANG   = 0x1
	WUNTRACED = 0x2

	// Interface flags.
	IFF_UP          = 0x0001
	IFF_BROADCAST   = 0x0002
	IFF_LOOPBACK    = 0x0008
	IFF_POINTOPOINT = 0x0010
	IFF_NOARP       = 0x0040
	IFF_AUTOUP      = 0x0080
	IFF_PROMISC     = 0x0100
	IFF_ALLMULTI    = 0x0200
	IFF_SIMPLEX     = 0x0800
	IFF_LINK        = 0x1000
	IFF_MULTICAST   = 0x8000

	// Routing.
	RTM_ADD    = 0x1
	RTM_DELETE = 0x2
	RTM_CHANGE = 0x3
	RTM_GET    = 0x4

	// Terminal ioctls (TCGETA = 0x8000 base, see <termios.h>).
	TIOCGPGRP = 0x800f
	TIOCSPGRP = 0x8010
	TIOCSCTTY = 0x8011
	// Haiku has no TIOCNOTTY; exec_libc.go treats zero as "unsupported"
	// and returns ENOSYS when SysProcAttr.Noctty is set.
	TIOCNOTTY = 0

	// _F_DUP2FD_CLOEXEC is only referenced from the illumos/solaris branch
	// in exec_libc.go but must resolve at compile time on every platform
	// in the shared file. Haiku never takes that branch.
	_F_DUP2FD_CLOEXEC = 0
)

// POSIX-positive values exposed to Go. The runtime translates libroot's
// negative B_* errnos to these at the syscall boundary.
const (
	EPERM           = Errno(1)
	ENOENT          = Errno(2)
	ESRCH           = Errno(3)
	EINTR           = Errno(4)
	EIO             = Errno(5)
	ENXIO           = Errno(6)
	E2BIG           = Errno(7)
	ENOEXEC         = Errno(8)
	EBADF           = Errno(9)
	EBADFD          = EBADF // no distinct Haiku errno; matches the BSD convention
	ECHILD          = Errno(10)
	EAGAIN          = Errno(11)
	EWOULDBLOCK     = EAGAIN
	ENOMEM          = Errno(12)
	EACCES          = Errno(13)
	EFAULT          = Errno(14)
	EBUSY           = Errno(16)
	EEXIST          = Errno(17)
	EXDEV           = Errno(18)
	ENODEV          = Errno(19)
	ENOTDIR         = Errno(20)
	EISDIR          = Errno(21)
	EINVAL          = Errno(22)
	ENFILE          = Errno(23)
	EMFILE          = Errno(24)
	ENOTTY          = Errno(25)
	ETXTBSY         = Errno(26)
	EFBIG           = Errno(27)
	ENOSPC          = Errno(28)
	ESPIPE          = Errno(29)
	EROFS           = Errno(30)
	EMLINK          = Errno(31)
	EPIPE           = Errno(32)
	EDOM            = Errno(33)
	ERANGE          = Errno(34)
	EDEADLK         = Errno(35)
	ENAMETOOLONG    = Errno(36)
	ENOLCK          = Errno(37)
	ENOSYS          = Errno(38)
	ENOTEMPTY       = Errno(39)
	ELOOP           = Errno(40)
	ENOMSG          = Errno(42)
	EIDRM           = Errno(43)
	ENOLINK         = Errno(67)
	EPROTO          = Errno(71)
	EMULTIHOP       = Errno(72)
	EBADMSG         = Errno(74)
	EOVERFLOW       = Errno(75)
	EILSEQ          = Errno(84)
	ENOTSOCK        = Errno(88)
	EDESTADDRREQ    = Errno(89)
	EMSGSIZE        = Errno(90)
	EPROTOTYPE      = Errno(91)
	ENOPROTOOPT     = Errno(92)
	EPROTONOSUPPORT = Errno(93)
	EOPNOTSUPP      = Errno(95)
	ENOTSUP         = EOPNOTSUPP
	EPFNOSUPPORT    = Errno(96)
	EAFNOSUPPORT    = Errno(97)
	EADDRINUSE      = Errno(98)
	EADDRNOTAVAIL   = Errno(99)
	ENETDOWN        = Errno(100)
	ENETUNREACH     = Errno(101)
	ENETRESET       = Errno(102)
	ECONNABORTED    = Errno(103)
	ECONNRESET      = Errno(104)
	ENOBUFS         = Errno(105)
	EISCONN         = Errno(106)
	ENOTCONN        = Errno(107)
	ESHUTDOWN       = Errno(108)
	ETIMEDOUT       = Errno(110)
	ECONNREFUSED    = Errno(111)
	EHOSTDOWN       = Errno(112)
	EHOSTUNREACH    = Errno(113)
	EALREADY        = Errno(114)
	EINPROGRESS     = Errno(115)
	ESTALE          = Errno(116)
	EDQUOT          = Errno(122)
	ECANCELED       = Errno(125)
	EOWNERDEAD      = Errno(130)
	ENOTRECOVERABLE = Errno(131)
)

// errors and signals are looked up by index into Errno.Error and
// Signal.String.

var errors = [...]string{
	1:   "operation not permitted",
	2:   "no such file or directory",
	3:   "no such process",
	4:   "interrupted system call",
	5:   "input/output error",
	6:   "no such device or address",
	7:   "argument list too long",
	8:   "exec format error",
	9:   "bad file descriptor",
	10:  "no child processes",
	11:  "resource temporarily unavailable",
	12:  "cannot allocate memory",
	13:  "permission denied",
	14:  "bad address",
	16:  "device or resource busy",
	17:  "file exists",
	18:  "invalid cross-device link",
	19:  "no such device",
	20:  "not a directory",
	21:  "is a directory",
	22:  "invalid argument",
	23:  "too many open files in system",
	24:  "too many open files",
	25:  "inappropriate ioctl for device",
	26:  "text file busy",
	27:  "file too large",
	28:  "no space left on device",
	29:  "illegal seek",
	30:  "read-only file system",
	31:  "too many links",
	32:  "broken pipe",
	33:  "numerical argument out of domain",
	34:  "numerical result out of range",
	35:  "resource deadlock avoided",
	36:  "file name too long",
	37:  "no locks available",
	38:  "function not implemented",
	39:  "directory not empty",
	40:  "too many levels of symbolic links",
	42:  "no message of desired type",
	43:  "identifier removed",
	67:  "link has been severed",
	71:  "protocol error",
	72:  "multihop attempted",
	74:  "bad message",
	75:  "value too large for defined data type",
	84:  "invalid or incomplete multibyte or wide character",
	88:  "socket operation on non-socket",
	89:  "destination address required",
	90:  "message too long",
	91:  "protocol wrong type for socket",
	92:  "protocol not available",
	93:  "protocol not supported",
	95:  "operation not supported",
	96:  "protocol family not supported",
	97:  "address family not supported by protocol",
	98:  "address already in use",
	99:  "cannot assign requested address",
	100: "network is down",
	101: "network is unreachable",
	102: "network dropped connection on reset",
	103: "software caused connection abort",
	104: "connection reset by peer",
	105: "no buffer space available",
	106: "transport endpoint is already connected",
	107: "transport endpoint is not connected",
	108: "cannot send after transport endpoint shutdown",
	110: "connection timed out",
	111: "connection refused",
	112: "host is down",
	113: "no route to host",
	114: "operation already in progress",
	115: "operation now in progress",
	116: "stale file handle",
	122: "disk quota exceeded",
	125: "operation canceled",
	130: "owner died",
	131: "state not recoverable",
}

var signals = [...]string{
	SIGHUP:    "hangup",
	SIGINT:    "interrupt",
	SIGQUIT:   "quit",
	SIGILL:    "illegal instruction",
	SIGCHLD:   "child status change",
	SIGABRT:   "aborted",
	SIGPIPE:   "broken pipe",
	SIGFPE:    "floating-point exception",
	SIGKILL:   "killed",
	SIGSTOP:   "stopped (signal)",
	SIGSEGV:   "segmentation fault",
	SIGCONT:   "continued",
	SIGTSTP:   "stopped",
	SIGALRM:   "alarm clock",
	SIGTERM:   "terminated",
	SIGTTIN:   "stopped (tty input)",
	SIGTTOU:   "stopped (tty output)",
	SIGUSR1:   "user defined signal 1",
	SIGUSR2:   "user defined signal 2",
	SIGWINCH:  "window size changes",
	SIGTRAP:   "trace/breakpoint trap",
	SIGPOLL:   "pollable event occurred",
	SIGPROF:   "profiling timer expired",
	SIGSYS:    "bad system call",
	SIGURG:    "urgent socket condition",
	SIGVTALRM: "virtual timer expired",
	SIGXCPU:   "cpu limit exceeded",
	SIGXFSZ:   "file size limit exceeded",
	SIGBUS:    "bus error",
}
