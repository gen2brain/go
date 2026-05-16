// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build haiku

package runtime

// Errno values are libroot's INT_MIN+offset codes. Untyped so they
// compose with both int32 (libroot returns) and int (runtime tests).
const (
	_EPERM     = -2147483633 // B_NOT_ALLOWED
	_ENOENT    = -2147459069 // B_ENTRY_NOT_FOUND
	_EINTR     = -2147483638 // B_INTERRUPTED
	_EAGAIN    = -2147483637 // B_WOULD_BLOCK
	_ENOMEM    = -2147483648 // B_NO_MEMORY
	_EACCES    = -2147483646 // B_PERMISSION_DENIED
	_EFAULT    = -2147478783 // B_BAD_ADDRESS
	_EINVAL    = -2147483643 // B_BAD_VALUE
	_ETIMEDOUT = -2147483639 // B_TIMED_OUT

	_PROT_NONE  = 0x0
	_PROT_READ  = 0x1
	_PROT_WRITE = 0x2
	_PROT_EXEC  = 0x4

	_MAP_ANON      = 0x8
	_MAP_PRIVATE   = 0x2
	_MAP_FIXED     = 0x4
	_MAP_NORESERVE = 0x10
	_MADV_DONTNEED = 0x5
	_MADV_FREE     = 0x6

	_SIGHUP    = 0x1
	_SIGINT    = 0x2
	_SIGQUIT   = 0x3
	_SIGILL    = 0x4
	_SIGCHLD   = 0x5
	_SIGABRT   = 0x6
	_SIGPIPE   = 0x7
	_SIGFPE    = 0x8
	_SIGKILL   = 0x9
	_SIGSTOP   = 0xa
	_SIGSEGV   = 0xb
	_SIGCONT   = 0xc
	_SIGTSTP   = 0xd
	_SIGALRM   = 0xe
	_SIGTERM   = 0xf
	_SIGTTIN   = 0x10
	_SIGTTOU   = 0x11
	_SIGUSR1   = 0x12
	_SIGUSR2   = 0x13
	_SIGWINCH  = 0x14
	_SIGTRAP   = 0x16
	_SIGPOLL   = 0x17
	_SIGPROF   = 0x18
	_SIGSYS    = 0x19
	_SIGURG    = 0x1a
	_SIGVTALRM = 0x1b
	_SIGXCPU   = 0x1c
	_SIGXFSZ   = 0x1d
	_SIGBUS    = 0x1e

	_FPE_INTDIV = 0x14
	_FPE_INTOVF = 0x15
	_FPE_FLTDIV = 0x16
	_FPE_FLTOVF = 0x17
	_FPE_FLTUND = 0x18
	_FPE_FLTRES = 0x19
	_FPE_FLTINV = 0x1a
	_FPE_FLTSUB = 0x1b

	_BUS_ADRALN = 0x28
	_BUS_ADRERR = 0x29
	_BUS_OBJERR = 0x2a

	_SEGV_MAPERR = 0x1e
	_SEGV_ACCERR = 0x1f

	_O_RDONLY   = 0x0
	_O_WRONLY   = 0x1
	_O_NONBLOCK = 0x80
	_O_CREAT    = 0x200
	_O_TRUNC    = 0x400
	_O_CLOEXEC  = 0x40

	_SS_DISABLE  = 0x2
	_SI_USER     = 0x0
	_SIG_BLOCK   = 0x1
	_SIG_UNBLOCK = 0x2
	_SIG_SETMASK = 0x3

	_SA_SIGINFO = 0x40
	_SA_RESTART = 0x10
	_SA_ONSTACK = 0x20
	_SA_NODEFER = 0x08

	_PTHREAD_CREATE_DETACHED = 0x1

	__SC_PAGE_SIZE        = 0x1b
	__SC_NPROCESSORS_ONLN = 0x23

	_F_GETFD = 0x2
	_F_SETFD = 0x4
	_F_GETFL = 0x8
	_F_SETFL = 0x10

	_FD_CLOEXEC = 0x1

	_ITIMER_REAL    = 1
	_ITIMER_VIRTUAL = 2
	_ITIMER_PROF    = 3
)

// sigset_t is a 64-bit integer in Haiku.
type sigset uint64

var sigset_all = sigset(^uint64(0))

type siginfo struct {
	si_signo  int32
	si_code   int32
	si_errno  int32
	pad_cgo_0 [4]byte
	si_pid    int32
	si_uid    uint32
	si_addr   uintptr
	si_status int32
	pad_cgo_1 [4]byte
	si_band   int64
	si_value  [8]byte // union sigval (sival_int / sival_ptr)
}

type timespec struct {
	tv_sec  int64
	tv_nsec int64
}

//go:nosplit
func (ts *timespec) setNsec(ns int64) {
	ts.tv_sec = ns / 1e9
	ts.tv_nsec = ns % 1e9
}

type timeval struct {
	tv_sec    int64
	tv_usec   int32
	pad_cgo_0 [4]byte
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = x
}

type itimerval struct {
	it_interval timeval
	it_value    timeval
}

type stackt struct {
	ss_sp     uintptr
	ss_size   uintptr
	ss_flags  int32
	pad_cgo_0 [4]byte
}

// fpu_state is the FXSAVE area used by Haiku x86_64. Layout matches
// arch/x86_64/signal.h struct fpu_state plus xstate_hdr and 16 ymm regs.
type fpu_state struct {
	control    uint16
	status     uint16
	tag        uint16
	opcode     uint16
	rip        uint64
	rdp        uint64
	mxcsr      uint32
	mxcsr_mask uint32
	fp_or_mmx  [128]byte // 8 * 16-byte fp/mmx registers
	xmm        [256]byte // 16 * 16-byte xmm registers
	_reserved  [96]byte
}

type xstate_hdr struct {
	bv        uint64
	xcomp_bv  uint64
	_reserved [48]byte
}

type savefpu struct {
	fp_fxsave fpu_state
	fp_xstate xstate_hdr
	fp_ymm    [256]byte // 16 * 16-byte high YMM
}

// vregs matches struct vregs in posix/arch/x86_64/signal.h.
// Used as both mcontext_t and uc_mcontext.
type mcontext struct {
	rax    uint64
	rbx    uint64
	rcx    uint64
	rdx    uint64
	rdi    uint64
	rsi    uint64
	rbp    uint64
	r8     uint64
	r9     uint64
	r10    uint64
	r11    uint64
	r12    uint64
	r13    uint64
	r14    uint64
	r15    uint64
	rsp    uint64
	rip    uint64
	rflags uint64
	fpu    savefpu
}

type ucontext struct {
	uc_link     *ucontext
	uc_sigmask  sigset
	uc_stack    stackt
	uc_mcontext mcontext
}

// sigaction matches struct sigaction in posix/signal.h.
// The first member is a union of __sighandler_t and __siginfo_handler_t,
// represented here as a single uintptr.
type sigactiont struct {
	sa_handler  uintptr
	sa_mask     sigset
	sa_flags    int32
	pad_cgo_0   [4]byte
	sa_userdata uintptr
}

type pthread uintptr      // typedef struct _pthread_thread *pthread_t
type pthread_attr uintptr // typedef struct _pthread_attr *pthread_attr_t

// sem_t struct from posix/semaphore.h: 4 * int32 = 16 bytes.
type semt struct {
	stype   int32
	id      int32
	padding [2]int32
}
