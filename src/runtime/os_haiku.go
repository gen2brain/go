// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build haiku

package runtime

import (
	"internal/abi"
	"internal/runtime/atomic"
	"unsafe"
)

const (
	threadStackSize = 0x100000 // size of a thread stack allocated by OS

	// Haiku clockid_t values from <time.h>.
	_CLOCK_MONOTONIC = 0
	_CLOCK_REALTIME  = -1
)

type mOS struct {
	waitsema uintptr // semaphore for parking on locks
	perrno   uintptr // pointer to libroot.so errno (returned by _errnop)
	libcall  libcall
}

//go:nosplit
func semacreate(mp *m) {
	if mp.waitsema != 0 {
		return
	}

	var sem *semt

	// Call libc's malloc rather than mallocgc to allocate space on the
	// C heap; mallocgc could deadlock here.
	sem = (*semt)(malloc(unsafe.Sizeof(*sem)))
	if sem_init(sem, 0, 0) != 0 {
		throw("sem_init")
	}
	mp.waitsema = uintptr(unsafe.Pointer(sem))
}

//go:nosplit
func semasleep(ns int64) int32 {
	mp := getg().m
	if ns >= 0 {
		var ts timespec

		if clock_gettime(_CLOCK_REALTIME, &ts) != 0 {
			throw("clock_gettime")
		}
		ts.tv_sec += ns / 1e9
		ts.tv_nsec += ns % 1e9
		if ts.tv_nsec >= 1e9 {
			ts.tv_sec++
			ts.tv_nsec -= 1e9
		}

		if r, err := sem_timedwait((*semt)(unsafe.Pointer(mp.waitsema)), &ts); r != 0 {
			if err == _ETIMEDOUT || err == _EAGAIN || err == _EINTR {
				return -1
			}
			throw("sem_timedwait")
		}
		return 0
	}
	for {
		r1, err := sem_wait((*semt)(unsafe.Pointer(mp.waitsema)))
		if r1 == 0 {
			break
		}
		if err == _EINTR {
			continue
		}
		throw("sem_wait")
	}
	return 0
}

//go:nosplit
func semawakeup(mp *m) {
	if sem_post((*semt)(unsafe.Pointer(mp.waitsema))) != 0 {
		throw("sem_post")
	}
}

func osinit() {
	// Locate the per-thread errno slot before any syscalls run.
	miniterrno()

	numCPUStartup = getCPUCount()
	physPageSize = sysconf(__SC_PAGE_SIZE)
}

func getCPUCount() int32 {
	return int32(sysconf(__SC_NPROCESSORS_ONLN))
}

// newosproc0 is a version of newosproc that can be called before the runtime
// is initialized.
//
// This function is not safe to use after initialization as it does not pass
// an M as fnarg.
//
//go:nosplit
func newosproc0(stacksize uintptr, fn uintptr) {
	var (
		attr pthread_attr
		oset sigset
		tid  pthread
	)

	if pthread_attr_init(&attr) != 0 {
		writeErrStr(failthreadcreate)
		exit(1)
	}

	if pthread_attr_setstacksize(&attr, uint64(stacksize)) != 0 {
		writeErrStr(failthreadcreate)
		exit(1)
	}

	if pthread_attr_setdetachstate(&attr, _PTHREAD_CREATE_DETACHED) != 0 {
		writeErrStr(failthreadcreate)
		exit(1)
	}

	// Disable signals during create, so that the new thread starts
	// with signals disabled. It will enable them in minit.
	sigprocmask(_SIG_SETMASK, &sigset_all, &oset)
	var ret int32
	for tries := 0; tries < 20; tries++ {
		// pthread_create can fail with EAGAIN; retry a few times.
		ret = pthread_create(&tid, &attr, fn, nil)
		if ret != _EAGAIN {
			break
		}
		usleep(uint32(tries+1) * 1000)
	}
	sigprocmask(_SIG_SETMASK, &oset, nil)
	if ret != 0 {
		writeErrStr(failthreadcreate)
		exit(1)
	}
}

// Called to do synchronous initialization of Go code built with
// -buildmode=c-archive or -buildmode=c-shared.
// None of the Go runtime is initialized.
//
//go:nosplit
//go:nowritebarrierrec
func libpreinit() {
	initsig(true)
}

func mpreinit(mp *m) {
	mp.gsignal = malg(32 * 1024)
	mp.gsignal.m = mp
}

// errno is per-thread on Haiku and accessible via _errnop, which returns
// a pointer to the calling thread's errno location.
func miniterrno() {
	mp := getg().m
	r, _ := syscall0(&libc__errnop)
	mp.perrno = r
}

// syscall_haikuMiniterrnoChild is called by package syscall from the
// child after fork() to refresh m.perrno. fork() inherits the parent's
// TLS contents, but each Haiku thread has its own errno address; without
// this re-resolve, asmsyscall6 reads errno from the parent's stale
// per-thread location and always sees 0.
//
//go:linkname syscall_haikuMiniterrnoChild syscall.haikuMiniterrnoChild
//go:nosplit
func syscall_haikuMiniterrnoChild() {
	miniterrno()
}

func minit() {
	miniterrno()
	minitSignals()
	getg().m.procid = uint64(uint32(find_thread(nil)))
}

func unminit() {
	unminitSignals()
	mp := getg().m
	mp.procid = 0
	// Stale perrno from the previous thread would crash asmsyscall6's errno
	// clear when the M is rebound; reset so miniterrno can re-resolve.
	mp.perrno = 0
}

//go:nowritebarrierrec
func mdestroy(mp *m) {
}

// tstart is the entry point for new OS threads. Defined in sys_haiku_$GOARCH.s.
func tstart()

func newosproc(mp *m) {
	var (
		attr pthread_attr
		oset sigset
		tid  pthread
	)

	if pthread_attr_init(&attr) != 0 {
		throw("pthread_attr_init")
	}

	if pthread_attr_setstacksize(&attr, threadStackSize) != 0 {
		throw("pthread_attr_setstacksize")
	}

	if pthread_attr_setdetachstate(&attr, _PTHREAD_CREATE_DETACHED) != 0 {
		throw("pthread_attr_setdetachstate")
	}

	// Disable signals during create, so that the new thread starts
	// with signals disabled. It will enable them in minit.
	sigprocmask(_SIG_SETMASK, &sigset_all, &oset)
	ret := retryOnEAGAIN(func() int32 {
		return pthread_create(&tid, &attr, abi.FuncPCABI0(tstart), unsafe.Pointer(mp))
	})
	sigprocmask(_SIG_SETMASK, &oset, nil)
	if ret != 0 {
		print("runtime: failed to create new OS thread (have ", mcount(), " already; errno=", ret, ")\n")
		if ret == _EAGAIN {
			print("runtime: may need to increase max user processes (ulimit -u)\n")
		}
		throw("newosproc")
	}
}

func exitThread(wait *atomic.Uint32) {
	// Haiku's libroot cleans up threads; we should not reach here.
	throw("exitThread")
}

var urandom_dev = []byte("/dev/urandom\x00")

//go:nosplit
func readRandom(r []byte) int {
	fd := open(&urandom_dev[0], 0 /* O_RDONLY */, 0)
	n := read(fd, unsafe.Pointer(&r[0]), int32(len(r)))
	closefd(fd)
	return int(n)
}

func goenvs() {
	goenvs_unix()
}

/* SIGNAL */

// _NSIG is one greater than the highest signal number accepted by sigaction.
const (
	_NSIG = 41
)

func sigtramp()

//go:nosplit
//go:nowritebarrierrec
func setsig(i uint32, fn uintptr) {
	var sa sigactiont
	sa.sa_flags = _SA_SIGINFO | _SA_ONSTACK | _SA_RESTART
	// SA_NODEFER lets a fault signal recur inside its own handler, which
	// sigpanic relies on. For other signals it interacts badly with Haiku's
	// per-handler reset, so apply it only to fault signals.
	switch i {
	case _SIGSEGV, _SIGBUS, _SIGFPE, _SIGILL, _SIGTRAP:
		sa.sa_flags |= _SA_NODEFER
	}
	sa.sa_mask = sigset_all
	if fn == abi.FuncPCABIInternal(sighandler) {
		fn = abi.FuncPCABI0(sigtramp)
	}
	sa.sa_handler = fn
	sigaction(uintptr(i), &sa, nil)
}

//go:nosplit
//go:nowritebarrierrec
func setsigstack(i uint32) {
	var sa sigactiont
	sigaction(uintptr(i), nil, &sa)
	if sa.sa_flags&_SA_ONSTACK != 0 {
		return
	}
	sa.sa_flags |= _SA_ONSTACK
	sigaction(uintptr(i), &sa, nil)
}

//go:nosplit
//go:nowritebarrierrec
func getsig(i uint32) uintptr {
	var sa sigactiont
	sigaction(uintptr(i), nil, &sa)
	return sa.sa_handler
}

//go:nosplit
func setSignalstackSP(s *stackt, sp uintptr) {
	*(*uintptr)(unsafe.Pointer(&s.ss_sp)) = sp
}

//go:nosplit
func (c *sigctxt) fixsigcode(sig uint32) {
}

//go:nosplit
//go:nowritebarrierrec
func sigaddset(mask *sigset, i int) {
	*mask |= 1 << (uint(i) - 1)
}

func sigdelset(mask *sigset, i int) {
	*mask &^= 1 << (uint(i) - 1)
}

// CPU profiling is disabled on Haiku: ITIMER_PROF + multi-M SIGPROF
// delivery races the kernel's signal-handler-reset path and panics
// the kernel. pprof.StartCPUProfile produces no samples here.
func setProcessCPUProfiler(hz int32) {
	_ = hz
}

func setThreadCPUProfiler(hz int32) {
	_ = hz
}

//go:nosplit
func validSIGPROF(mp *m, c *sigctxt) bool {
	return true
}

//go:nosplit
func nanotime1() int64 {
	tp := &timespec{}
	if clock_gettime(_CLOCK_MONOTONIC, tp) != 0 {
		throw("syscall clock_gettime failed")
	}
	return tp.tv_sec*1e9 + tp.tv_nsec
}

func walltime() (sec int64, nsec int32) {
	ts := &timespec{}
	if clock_gettime(_CLOCK_REALTIME, ts) != 0 {
		throw("syscall clock_gettime failed")
	}
	return ts.tv_sec, int32(ts.tv_nsec)
}

//go:nosplit
func setNonblock(fd int32) {
	flags, _ := fcntl(fd, _F_GETFL, 0)
	if flags != -1 {
		fcntl(fd, _F_SETFL, flags|_O_NONBLOCK)
	}
}

//go:nosplit
func closeonexec(fd int32) {
	fcntl(fd, _F_SETFD, _FD_CLOEXEC)
}

// sigPerThreadSyscall is only used on linux, so we assign a bogus signal
// number.
const sigPerThreadSyscall = 1 << 31

//go:nosplit
func runPerThreadSyscall() {
	throw("runPerThreadSyscall only valid on linux")
}

//go:nosplit
func getuid() int32 {
	r, errno := syscall0(&libc_getuid)
	if errno != 0 {
		print("getuid failed ", errno)
		throw("getuid")
	}
	return int32(r)
}

//go:nosplit
func geteuid() int32 {
	r, errno := syscall0(&libc_geteuid)
	if errno != 0 {
		print("geteuid failed ", errno)
		throw("geteuid")
	}
	return int32(r)
}

//go:nosplit
func getgid() int32 {
	r, errno := syscall0(&libc_getgid)
	if errno != 0 {
		print("getgid failed ", errno)
		throw("getgid")
	}
	return int32(r)
}

//go:nosplit
func getegid() int32 {
	r, errno := syscall0(&libc_getegid)
	if errno != 0 {
		print("getegid failed ", errno)
		throw("getegid")
	}
	return int32(r)
}
