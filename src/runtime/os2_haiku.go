// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Each syscall is made by calling its libc symbol via asmsyscall6, defined
// in sys_haiku_$GOARCH.s.

package runtime

import (
	"internal/runtime/sys"
	"unsafe"
)

// Export the ABI markers so runtime_loader can find them; see
// sys_haiku_amd64.s for the values.

//go:cgo_export_dynamic _gSharedObjectHaikuABI
//go:cgo_export_dynamic _gSharedObjectHaikuVersion

// Syscalls

//go:cgo_import_dynamic libc__errnop _errnop "libroot.so"
//go:cgo_import_dynamic libc_clock_gettime clock_gettime "libroot.so"
//go:cgo_import_dynamic libc_close close "libroot.so"
//go:cgo_import_dynamic libc_exit _exit "libroot.so"
//go:cgo_import_dynamic libc_fcntl fcntl "libroot.so"
//go:cgo_import_dynamic libc_find_thread find_thread "libroot.so"
//go:cgo_import_dynamic libc_getpid getpid "libroot.so"
//go:cgo_import_dynamic libc_kill kill "libroot.so"
//go:cgo_import_dynamic libc_madvise madvise "libroot.so"
//go:cgo_import_dynamic libc_malloc malloc "libroot.so"
//go:cgo_import_dynamic libc_mmap mmap "libroot.so"
//go:cgo_import_dynamic libc_mprotect mprotect "libroot.so"
//go:cgo_import_dynamic libc_munmap munmap "libroot.so"
//go:cgo_import_dynamic libc_open open "libroot.so"
//go:cgo_import_dynamic libc_pipe pipe "libroot.so"
//go:cgo_import_dynamic libc_poll poll "libroot.so"
//go:cgo_import_dynamic libc_raise raise "libroot.so"
//go:cgo_import_dynamic libc_read read "libroot.so"
//go:cgo_import_dynamic libc_sched_yield sched_yield "libroot.so"
//go:cgo_import_dynamic libc_sem_init sem_init "libroot.so"
//go:cgo_import_dynamic libc_sem_post sem_post "libroot.so"
//go:cgo_import_dynamic libc_sem_timedwait sem_timedwait "libroot.so"
//go:cgo_import_dynamic libc_sem_wait sem_wait "libroot.so"
//go:cgo_import_dynamic libc_send_signal send_signal "libroot.so"
//go:cgo_import_dynamic libc_setitimer setitimer "libroot.so"
//go:cgo_import_dynamic libc_shutdown shutdown "libnetwork.so"
//go:cgo_import_dynamic libc_sigaction sigaction#LIBROOT_1_ALPHA4 "libroot.so"
//go:cgo_import_dynamic libc_sigaltstack sigaltstack "libroot.so"
//go:cgo_import_dynamic libc_sigprocmask sigprocmask#LIBROOT_1_ALPHA4 "libroot.so"
//go:cgo_import_dynamic libc_sysconf sysconf#LIBROOT_1_ALPHA4 "libroot.so"
//go:cgo_import_dynamic libc_usleep usleep "libroot.so"
//go:cgo_import_dynamic libc_write write "libroot.so"
//go:cgo_import_dynamic libc_getuid getuid "libroot.so"
//go:cgo_import_dynamic libc_geteuid geteuid "libroot.so"
//go:cgo_import_dynamic libc_getgid getgid "libroot.so"
//go:cgo_import_dynamic libc_getegid getegid "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_destroy pthread_attr_destroy "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_init pthread_attr_init "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_setstacksize pthread_attr_setstacksize "libroot.so"
//go:cgo_import_dynamic libc_pthread_attr_setdetachstate pthread_attr_setdetachstate "libroot.so"
//go:cgo_import_dynamic libc_pthread_create pthread_create "libroot.so"
//go:cgo_import_dynamic libc_pthread_self pthread_self "libroot.so"
//go:cgo_import_dynamic libc_tls_allocate tls_allocate "libroot.so"

//go:linkname libc__errnop libc__errnop
//go:linkname libc_clock_gettime libc_clock_gettime
//go:linkname libc_close libc_close
//go:linkname libc_exit libc_exit
//go:linkname libc_fcntl libc_fcntl
//go:linkname libc_find_thread libc_find_thread
//go:linkname libc_getpid libc_getpid
//go:linkname libc_kill libc_kill
//go:linkname libc_madvise libc_madvise
//go:linkname libc_malloc libc_malloc
//go:linkname libc_mmap libc_mmap
//go:linkname libc_mprotect libc_mprotect
//go:linkname libc_munmap libc_munmap
//go:linkname libc_open libc_open
//go:linkname libc_pipe libc_pipe
//go:linkname libc_poll libc_poll
//go:linkname libc_raise libc_raise
//go:linkname libc_read libc_read
//go:linkname libc_sched_yield libc_sched_yield
//go:linkname libc_sem_init libc_sem_init
//go:linkname libc_sem_post libc_sem_post
//go:linkname libc_sem_timedwait libc_sem_timedwait
//go:linkname libc_sem_wait libc_sem_wait
//go:linkname libc_send_signal libc_send_signal
//go:linkname libc_setitimer libc_setitimer
//go:linkname libc_shutdown libc_shutdown
//go:linkname libc_sigaction libc_sigaction
//go:linkname libc_sigaltstack libc_sigaltstack
//go:linkname libc_sigprocmask libc_sigprocmask
//go:linkname libc_sysconf libc_sysconf
//go:linkname libc_usleep libc_usleep
//go:linkname libc_write libc_write
//go:linkname libc_getuid libc_getuid
//go:linkname libc_geteuid libc_geteuid
//go:linkname libc_getgid libc_getgid
//go:linkname libc_getegid libc_getegid
//go:linkname libc_pthread_attr_destroy libc_pthread_attr_destroy
//go:linkname libc_pthread_attr_init libc_pthread_attr_init
//go:linkname libc_pthread_attr_setstacksize libc_pthread_attr_setstacksize
//go:linkname libc_pthread_attr_setdetachstate libc_pthread_attr_setdetachstate
//go:linkname libc_pthread_create libc_pthread_create
//go:linkname libc_pthread_self libc_pthread_self
//go:linkname libc_tls_allocate libc_tls_allocate

var (
	libc__errnop,
	libc_clock_gettime,
	libc_close,
	libc_exit,
	libc_fcntl,
	libc_find_thread,
	libc_getpid,
	libc_kill,
	libc_madvise,
	libc_malloc,
	libc_mmap,
	libc_mprotect,
	libc_munmap,
	libc_open,
	libc_pipe,
	libc_poll,
	libc_raise,
	libc_read,
	libc_sched_yield,
	libc_sem_init,
	libc_sem_post,
	libc_sem_timedwait,
	libc_sem_wait,
	libc_send_signal,
	libc_setitimer,
	libc_shutdown,
	libc_sigaction,
	libc_sigaltstack,
	libc_sigprocmask,
	libc_sysconf,
	libc_usleep,
	libc_write,
	libc_getuid,
	libc_geteuid,
	libc_getgid,
	libc_getegid,
	libc_pthread_attr_destroy,
	libc_pthread_attr_init,
	libc_pthread_attr_setstacksize,
	libc_pthread_attr_setdetachstate,
	libc_pthread_create,
	libc_pthread_self,
	libc_tls_allocate libFunc
)

type libFunc uintptr

// asmsyscall6 calls the libc symbol using the C ABI.
// It is defined in sys_haiku_$GOARCH.s.
var asmsyscall6 libFunc

// syscallN must be called on a goroutine with a valid m: it uses
// g.m.libcall to pass arguments to asmcgocall. Code paths without g/m
// must call the matching helper in sys_haiku_$GOARCH.s instead.

//go:nowritebarrier
//go:nosplit
func syscall0(fn *libFunc) (r, err uintptr) {
	gp := getg()
	mp := gp.m
	resetLibcall := true
	if mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = sys.GetCallerPC()
		mp.libcallsp = sys.GetCallerSP()
	} else {
		resetLibcall = false
	}

	c := libcall{
		fn:   uintptr(unsafe.Pointer(fn)),
		n:    0,
		args: uintptr(unsafe.Pointer(&fn)), // unused but must be non-nil
	}

	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
func syscall1(fn *libFunc, a0 uintptr) (r, err uintptr) {
	gp := getg()
	mp := gp.m
	resetLibcall := true
	if mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = sys.GetCallerPC()
		mp.libcallsp = sys.GetCallerSP()
	} else {
		resetLibcall = false
	}

	c := libcall{
		fn:   uintptr(unsafe.Pointer(fn)),
		n:    1,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscall2(fn *libFunc, a0, a1 uintptr) (r, err uintptr) {
	gp := getg()
	mp := gp.m
	resetLibcall := true
	if mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = sys.GetCallerPC()
		mp.libcallsp = sys.GetCallerSP()
	} else {
		resetLibcall = false
	}

	c := libcall{
		fn:   uintptr(unsafe.Pointer(fn)),
		n:    2,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscall3(fn *libFunc, a0, a1, a2 uintptr) (r, err uintptr) {
	gp := getg()
	mp := gp.m
	resetLibcall := true
	if mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = sys.GetCallerPC()
		mp.libcallsp = sys.GetCallerSP()
	} else {
		resetLibcall = false
	}

	c := libcall{
		fn:   uintptr(unsafe.Pointer(fn)),
		n:    3,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscall4(fn *libFunc, a0, a1, a2, a3 uintptr) (r, err uintptr) {
	gp := getg()
	mp := gp.m
	resetLibcall := true
	if mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = sys.GetCallerPC()
		mp.libcallsp = sys.GetCallerSP()
	} else {
		resetLibcall = false
	}

	c := libcall{
		fn:   uintptr(unsafe.Pointer(fn)),
		n:    4,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscall5(fn *libFunc, a0, a1, a2, a3, a4 uintptr) (r, err uintptr) {
	gp := getg()
	mp := gp.m
	resetLibcall := true
	if mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = sys.GetCallerPC()
		mp.libcallsp = sys.GetCallerSP()
	} else {
		resetLibcall = false
	}

	c := libcall{
		fn:   uintptr(unsafe.Pointer(fn)),
		n:    5,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscall6(fn *libFunc, a0, a1, a2, a3, a4, a5 uintptr) (r, err uintptr) {
	gp := getg()
	mp := gp.m
	resetLibcall := true
	if mp.libcallsp == 0 {
		mp.libcallg.set(gp)
		mp.libcallpc = sys.GetCallerPC()
		mp.libcallsp = sys.GetCallerSP()
	} else {
		resetLibcall = false
	}

	c := libcall{
		fn:   uintptr(unsafe.Pointer(fn)),
		n:    6,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(&asmsyscall6), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

func exit1(code int32)

//go:nosplit
func exit(code int32) {
	gp := getg()
	if gp != nil {
		syscall1(&libc_exit, uintptr(code))
		return
	}
	exit1(code)
}

func write2(fd, p uintptr, n int32) int32

//go:nosplit
func write1(fd uintptr, p unsafe.Pointer, n int32) int32 {
	gp := getg()
	if gp != nil {
		r, errno := syscall3(&libc_write, uintptr(fd), uintptr(p), uintptr(n))
		if int32(r) < 0 {
			return -int32(errno)
		}
		return int32(r)
	}
	return write2(fd, uintptr(p), n)
}

//go:nosplit
func read(fd int32, p unsafe.Pointer, n int32) int32 {
	r, errno := syscall3(&libc_read, uintptr(fd), uintptr(p), uintptr(n))
	if int32(r) < 0 {
		return -int32(errno)
	}
	return int32(r)
}

//go:nosplit
func open(name *byte, mode, perm int32) int32 {
	r, _ := syscall3(&libc_open, uintptr(unsafe.Pointer(name)), uintptr(mode), uintptr(perm))
	return int32(r)
}

//go:nosplit
func closefd(fd int32) int32 {
	r, _ := syscall1(&libc_close, uintptr(fd))
	return int32(r)
}

//go:nosplit
func pipe() (r, w int32, errno int32) {
	var p [2]int32
	_, err := syscall1(&libc_pipe, uintptr(noescape(unsafe.Pointer(&p[0]))))
	return p[0], p[1], int32(err)
}

//go:nosplit
func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) (unsafe.Pointer, int) {
	r, err0 := syscall6(&libc_mmap, uintptr(addr), uintptr(n), uintptr(prot), uintptr(flags), uintptr(fd), uintptr(off))
	if r == ^uintptr(0) {
		return nil, int(err0)
	}
	return unsafe.Pointer(r), int(err0)
}

//go:nosplit
func mprotect(addr unsafe.Pointer, n uintptr, prot int32) (unsafe.Pointer, int) {
	r, err0 := syscall3(&libc_mprotect, uintptr(addr), uintptr(n), uintptr(prot))
	if r == ^uintptr(0) {
		return nil, int(err0)
	}
	return unsafe.Pointer(r), int(err0)
}

//go:nosplit
func munmap(addr unsafe.Pointer, n uintptr) {
	r, err := syscall2(&libc_munmap, uintptr(addr), uintptr(n))
	if int32(r) == -1 {
		println("syscall munmap failed: ", hex(err))
		throw("syscall munmap")
	}
}

//go:nosplit
func madvise(addr unsafe.Pointer, n uintptr, flags int32) {
	syscall3(&libc_madvise, uintptr(addr), uintptr(n), uintptr(flags))
	// madvise on Haiku may legitimately fail with EINVAL for some
	// advice values; ignore errors as on other Unix ports.
}

func sigaction1(sig, new, old uintptr)

//go:nosplit
func sigaction(sig uintptr, new, old *sigactiont) {
	gp := getg()
	if gp != nil {
		r, err := syscall3(&libc_sigaction, sig, uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
		if int32(r) == -1 {
			println("Sigaction failed for sig: ", sig, " with error:", hex(err))
			throw("syscall sigaction")
		}
		return
	}
	sigaction1(sig, uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
}

//go:nosplit
func sigaltstack(new, old *stackt) {
	r, err := syscall2(&libc_sigaltstack, uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
	if int32(r) == -1 {
		println("syscall sigaltstack failed: ", hex(err))
		throw("syscall sigaltstack")
	}
}

func usleep1(us uint32)

//go:nosplit
func usleep_no_g(us uint32) {
	usleep1(us)
}

func haikuTlsInit()

//go:nosplit
func usleep(us uint32) {
	r, err := syscall1(&libc_usleep, uintptr(us))
	if int32(r) == -1 {
		println("syscall usleep failed: ", hex(err))
		throw("syscall usleep")
	}
}

//go:nosplit
func clock_gettime(clockid int32, tp *timespec) int32 {
	r, _ := syscall2(&libc_clock_gettime, uintptr(clockid), uintptr(unsafe.Pointer(tp)))
	return int32(r)
}

//go:nosplit
func setitimer(mode int32, new, old *itimerval) {
	r, err := syscall3(&libc_setitimer, uintptr(mode), uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
	if int32(r) == -1 {
		println("syscall setitimer failed: ", hex(err))
		throw("syscall setitimer")
	}
}

//go:nosplit
func malloc(size uintptr) unsafe.Pointer {
	r, _ := syscall1(&libc_malloc, size)
	return unsafe.Pointer(r)
}

//go:nosplit
func sem_init(sem *semt, pshared int32, value uint32) int32 {
	r, _ := syscall3(&libc_sem_init, uintptr(unsafe.Pointer(sem)), uintptr(pshared), uintptr(value))
	return int32(r)
}

//go:nosplit
func sem_wait(sem *semt) (int32, int32) {
	r, err := syscall1(&libc_sem_wait, uintptr(unsafe.Pointer(sem)))
	return int32(r), int32(err)
}

//go:nosplit
func sem_post(sem *semt) int32 {
	r, _ := syscall1(&libc_sem_post, uintptr(unsafe.Pointer(sem)))
	return int32(r)
}

//go:nosplit
func sem_timedwait(sem *semt, timeout *timespec) (int32, int32) {
	r, err := syscall2(&libc_sem_timedwait, uintptr(unsafe.Pointer(sem)), uintptr(unsafe.Pointer(timeout)))
	return int32(r), int32(err)
}

//go:nosplit
func raise(sig uint32) {
	r, err := syscall1(&libc_raise, uintptr(sig))
	if int32(r) == -1 {
		println("syscall raise failed: ", hex(err))
		throw("syscall raise")
	}
}

//go:nosplit
func raiseproc(sig uint32) {
	pid, err := syscall0(&libc_getpid)
	if int32(pid) == -1 {
		println("syscall getpid failed: ", hex(err))
		throw("syscall raiseproc")
	}
	syscall2(&libc_kill, pid, uintptr(sig))
}

func osyield1()

//go:nosplit
func osyield_no_g() {
	osyield1()
}

//go:nosplit
func osyield() {
	r, err := syscall0(&libc_sched_yield)
	if int32(r) == -1 {
		println("syscall osyield failed: ", hex(err))
		throw("syscall osyield")
	}
}

//go:nosplit
func sysconf(name int32) uintptr {
	r, _ := syscall1(&libc_sysconf, uintptr(name))
	if int32(r) == -1 {
		throw("syscall sysconf")
	}
	return r
}

// pthread functions return their error code in the main return value;
// errno is not used.

//go:nosplit
func pthread_attr_destroy(attr *pthread_attr) int32 {
	r, _ := syscall1(&libc_pthread_attr_destroy, uintptr(unsafe.Pointer(attr)))
	return int32(r)
}

func pthread_attr_init1(attr uintptr) int32

//go:nosplit
func pthread_attr_init(attr *pthread_attr) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall1(&libc_pthread_attr_init, uintptr(unsafe.Pointer(attr)))
		return int32(r)
	}
	return pthread_attr_init1(uintptr(unsafe.Pointer(attr)))
}

func pthread_attr_setdetachstate1(attr uintptr, state int32) int32

//go:nosplit
func pthread_attr_setdetachstate(attr *pthread_attr, state int32) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall2(&libc_pthread_attr_setdetachstate, uintptr(unsafe.Pointer(attr)), uintptr(state))
		return int32(r)
	}
	return pthread_attr_setdetachstate1(uintptr(unsafe.Pointer(attr)), state)
}

func pthread_attr_setstacksize1(attr uintptr, size uint64) int32

//go:nosplit
func pthread_attr_setstacksize(attr *pthread_attr, size uint64) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall2(&libc_pthread_attr_setstacksize, uintptr(unsafe.Pointer(attr)), uintptr(size))
		return int32(r)
	}
	return pthread_attr_setstacksize1(uintptr(unsafe.Pointer(attr)), size)
}

func pthread_create1(tid, attr, fn, arg uintptr) int32

//go:nosplit
func pthread_create(tid *pthread, attr *pthread_attr, fn uintptr, arg unsafe.Pointer) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall4(&libc_pthread_create, uintptr(unsafe.Pointer(tid)), uintptr(unsafe.Pointer(attr)), fn, uintptr(arg))
		return int32(r)
	}
	return pthread_create1(uintptr(unsafe.Pointer(tid)), uintptr(unsafe.Pointer(attr)), fn, uintptr(arg))
}

func sigprocmask1(how, new, old uintptr)

//go:nosplit
func sigprocmask(how int32, new, old *sigset) {
	gp := getg()
	if gp != nil && gp.m != nil {
		r, err := syscall3(&libc_sigprocmask, uintptr(how), uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
		if int32(r) != 0 {
			println("syscall sigprocmask failed: ", hex(err))
			throw("syscall sigprocmask")
		}
		return
	}
	sigprocmask1(uintptr(how), uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
}

//go:nosplit
func pthread_self() pthread {
	r, _ := syscall0(&libc_pthread_self)
	return pthread(r)
}

//go:nosplit
func find_thread(name *byte) int32 {
	r, _ := syscall1(&libc_find_thread, uintptr(unsafe.Pointer(name)))
	return int32(r)
}

//go:nosplit
func send_signal(thread int32, sig uint32) int32 {
	r, _ := syscall2(&libc_send_signal, uintptr(thread), uintptr(sig))
	return int32(r)
}

// procid is the kernel thread_id (find_thread): pthread_self collapses to a
// shared &sMainThread for spawn_thread'd threads (BLooper et al.).
//
//go:nosplit
func signalM(mp *m, sig int) {
	if mp.procid == 0 {
		return
	}
	syscall2(&libc_send_signal, uintptr(mp.procid), uintptr(sig))
}

//go:nosplit
func fcntl(fd, cmd, arg int32) (int32, int32) {
	r, errno := syscall3(&libc_fcntl, uintptr(fd), uintptr(cmd), uintptr(arg))
	return int32(r), int32(errno)
}
