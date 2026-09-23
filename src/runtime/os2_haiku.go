// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Each syscall is made by calling its libc symbol via asmsyscall6, defined
// in sys_haiku_$GOARCH.s.

package runtime

import (
	"internal/abi"
	"internal/goarch"
	"internal/runtime/sys"
	"unsafe"
)

// Export the ABI markers so runtime_loader can find them; see
// sys_haiku_amd64.s for the values.

//go:cgo_export_dynamic _gSharedObjectHaikuABI
//go:cgo_export_dynamic _gSharedObjectHaikuVersion

// Syscalls
//
// The cgo_import_dynamic directives binding the libc_* names to libroot
// symbols live in os2_haiku_import.go. Each libc_X is reached through the
// libc_X_trampoline in sys_libc_haiku.s, so no code takes the address of an
// imported symbol.

func libc__errnop_trampoline()
func libc_clock_gettime_trampoline()
func libc_close_trampoline()
func libc_exit_trampoline()
func libc_fcntl_trampoline()
func libc_find_thread_trampoline()
func libc_getpid_trampoline()
func libc_kill_trampoline()
func libc_madvise_trampoline()
func libc_malloc_trampoline()
func libc_mmap_trampoline()
func libc_mprotect_trampoline()
func libc_munmap_trampoline()
func libc_open_trampoline()
func libc_pipe_trampoline()
func libc_poll_trampoline()
func libc_raise_trampoline()
func libc_read_trampoline()
func libc_sched_yield_trampoline()
func libc_sem_init_trampoline()
func libc_sem_post_trampoline()
func libc_sem_timedwait_trampoline()
func libc_sem_wait_trampoline()
func libc_send_signal_trampoline()
func libc_setitimer_trampoline()
func libc_shutdown_trampoline()
func libc_sigaction_trampoline()
func libc_sigaltstack_trampoline()
func libc_sigprocmask_trampoline()
func libc_sysconf_trampoline()
func libc_system_time_nsecs_trampoline()
func libc_usleep_trampoline()
func libc_write_trampoline()
func libc_getuid_trampoline()
func libc_geteuid_trampoline()
func libc_getgid_trampoline()
func libc_getegid_trampoline()
func libc_pthread_attr_destroy_trampoline()
func libc_pthread_attr_init_trampoline()
func libc_pthread_attr_setstacksize_trampoline()
func libc_pthread_attr_setdetachstate_trampoline()
func libc_pthread_create_trampoline()
func libc_pthread_self_trampoline()
func libc_tls_allocate_trampoline()

// asmsyscall6 calls the libc symbol using the C ABI.
// It is defined in sys_haiku_$GOARCH.s.
func asmsyscall6()

// syscallN must be called on a goroutine with a valid m: it uses
// g.m.libcall to pass arguments to asmcgocall. Code paths without g/m
// must call the matching helper in sys_haiku_$GOARCH.s instead.

//go:nowritebarrier
//go:nosplit
func syscall0(fn uintptr) (r, err uintptr) {
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
		fn:   fn,
		n:    0,
		args: uintptr(unsafe.Pointer(&fn)), // unused but must be non-nil
	}

	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
func syscall1(fn uintptr, a0 uintptr) (r, err uintptr) {
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
		fn:   fn,
		n:    1,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscall2(fn uintptr, a0, a1 uintptr) (r, err uintptr) {
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
		fn:   fn,
		n:    2,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscall3(fn uintptr, a0, a1, a2 uintptr) (r, err uintptr) {
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
		fn:   fn,
		n:    3,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscall4(fn uintptr, a0, a1, a2, a3 uintptr) (r, err uintptr) {
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
		fn:   fn,
		n:    4,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscall5(fn uintptr, a0, a1, a2, a3, a4 uintptr) (r, err uintptr) {
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
		fn:   fn,
		n:    5,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	return c.r1, c.err
}

//go:nowritebarrier
//go:nosplit
//go:cgo_unsafe_args
func syscall6(fn uintptr, a0, a1, a2, a3, a4, a5 uintptr) (r, err uintptr) {
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
		fn:   fn,
		n:    6,
		args: uintptr(unsafe.Pointer(&a0)),
	}

	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))

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
		syscall1(abi.FuncPCABI0(libc_exit_trampoline), uintptr(code))
		return
	}
	exit1(code)
}

func write2(fd, p uintptr, n int32) int32

//go:nosplit
func write1(fd uintptr, p unsafe.Pointer, n int32) int32 {
	gp := getg()
	if gp != nil {
		r, errno := syscall3(abi.FuncPCABI0(libc_write_trampoline), uintptr(fd), uintptr(p), uintptr(n))
		if int32(r) < 0 {
			return int32(errno)
		}
		return int32(r)
	}
	return write2(fd, uintptr(p), n)
}

//go:nosplit
func read(fd int32, p unsafe.Pointer, n int32) int32 {
	r, errno := syscall3(abi.FuncPCABI0(libc_read_trampoline), uintptr(fd), uintptr(p), uintptr(n))
	if int32(r) < 0 {
		return int32(errno)
	}
	return int32(r)
}

//go:nosplit
func open(name *byte, mode, perm int32) int32 {
	r, _ := syscall3(abi.FuncPCABI0(libc_open_trampoline), uintptr(unsafe.Pointer(name)), uintptr(mode), uintptr(perm))
	return int32(r)
}

//go:nosplit
func closefd(fd int32) int32 {
	r, _ := syscall1(abi.FuncPCABI0(libc_close_trampoline), uintptr(fd))
	return int32(r)
}

//go:nosplit
func pipe() (r, w int32, errno int32) {
	var p [2]int32
	_, err := syscall1(abi.FuncPCABI0(libc_pipe_trampoline), uintptr(noescape(unsafe.Pointer(&p[0]))))
	return p[0], p[1], int32(err)
}

//go:nosplit
func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) (unsafe.Pointer, int) {
	r, err0 := syscall6(abi.FuncPCABI0(libc_mmap_trampoline), uintptr(addr), uintptr(n), uintptr(prot), uintptr(flags), uintptr(fd), uintptr(off))
	// err0 holds a raw libroot errno (an INT_MIN+N value stored
	// zero-extended); sign-extend it so callers see the negative B_* code
	// that runtime's _E* constants use.
	if r == ^uintptr(0) {
		return nil, int(int32(err0))
	}
	return unsafe.Pointer(r), int(int32(err0))
}

//go:nosplit
func mprotect(addr unsafe.Pointer, n uintptr, prot int32) (unsafe.Pointer, int) {
	r, err0 := syscall3(abi.FuncPCABI0(libc_mprotect_trampoline), uintptr(addr), uintptr(n), uintptr(prot))
	if r == ^uintptr(0) {
		return nil, int(err0)
	}
	return unsafe.Pointer(r), int(err0)
}

//go:nosplit
func munmap(addr unsafe.Pointer, n uintptr) {
	r, err := syscall2(abi.FuncPCABI0(libc_munmap_trampoline), uintptr(addr), uintptr(n))
	if int32(r) == -1 {
		println("syscall munmap failed: ", hex(err))
		throw("syscall munmap")
	}
}

//go:nosplit
func madvise(addr unsafe.Pointer, n uintptr, flags int32) {
	syscall3(abi.FuncPCABI0(libc_madvise_trampoline), uintptr(addr), uintptr(n), uintptr(flags))
	// madvise on Haiku may legitimately fail with EINVAL for some
	// advice values; ignore errors as on other Unix ports.
}

func sigaction1(sig, new, old uintptr)

//go:nosplit
func sigaction(sig uintptr, new, old *sigactiont) {
	gp := getg()
	if gp != nil {
		r, err := syscall3(abi.FuncPCABI0(libc_sigaction_trampoline), sig, uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
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
	r, err := syscall2(abi.FuncPCABI0(libc_sigaltstack_trampoline), uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
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
	r, err := syscall1(abi.FuncPCABI0(libc_usleep_trampoline), uintptr(us))
	if int32(r) == -1 {
		println("syscall usleep failed: ", hex(err))
		throw("syscall usleep")
	}
}

//go:nowritebarrier
//go:nosplit
func system_time_nsecs() int64 {
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

	fn := abi.FuncPCABI0(libc_system_time_nsecs_trampoline)
	c := libcall{
		fn:   fn,
		n:    0,
		args: uintptr(unsafe.Pointer(&fn)), // unused but must be non-nil
	}

	asmcgocall(unsafe.Pointer(abi.FuncPCABI0(asmsyscall6)), unsafe.Pointer(&c))

	if resetLibcall {
		mp.libcallsp = 0
	}

	if goarch.PtrSize == 4 {
		return int64(uint64(c.r1) | uint64(c.r2)<<32)
	}
	return int64(c.r1)
}

//go:nosplit
func clock_gettime(clockid int32, tp *timespec) int32 {
	r, _ := syscall2(abi.FuncPCABI0(libc_clock_gettime_trampoline), uintptr(clockid), uintptr(unsafe.Pointer(tp)))
	return int32(r)
}

//go:nosplit
func setitimer(mode int32, new, old *itimerval) {
	r, err := syscall3(abi.FuncPCABI0(libc_setitimer_trampoline), uintptr(mode), uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
	if int32(r) == -1 {
		println("syscall setitimer failed: ", hex(err))
		throw("syscall setitimer")
	}
}

//go:nosplit
func malloc(size uintptr) unsafe.Pointer {
	r, _ := syscall1(abi.FuncPCABI0(libc_malloc_trampoline), size)
	return unsafe.Pointer(r)
}

//go:nosplit
func sem_init(sem *semt, pshared int32, value uint32) int32 {
	r, _ := syscall3(abi.FuncPCABI0(libc_sem_init_trampoline), uintptr(unsafe.Pointer(sem)), uintptr(pshared), uintptr(value))
	return int32(r)
}

//go:nosplit
func sem_wait(sem *semt) (int32, int32) {
	r, err := syscall1(abi.FuncPCABI0(libc_sem_wait_trampoline), uintptr(unsafe.Pointer(sem)))
	return int32(r), int32(err)
}

//go:nosplit
func sem_post(sem *semt) int32 {
	r, _ := syscall1(abi.FuncPCABI0(libc_sem_post_trampoline), uintptr(unsafe.Pointer(sem)))
	return int32(r)
}

//go:nosplit
func sem_timedwait(sem *semt, timeout *timespec) (int32, int32) {
	r, err := syscall2(abi.FuncPCABI0(libc_sem_timedwait_trampoline), uintptr(unsafe.Pointer(sem)), uintptr(unsafe.Pointer(timeout)))
	return int32(r), int32(err)
}

//go:nosplit
func raise(sig uint32) {
	r, err := syscall1(abi.FuncPCABI0(libc_raise_trampoline), uintptr(sig))
	if int32(r) == -1 {
		println("syscall raise failed: ", hex(err))
		throw("syscall raise")
	}
}

//go:nosplit
func raiseproc(sig uint32) {
	pid, err := syscall0(abi.FuncPCABI0(libc_getpid_trampoline))
	if int32(pid) == -1 {
		println("syscall getpid failed: ", hex(err))
		throw("syscall raiseproc")
	}
	syscall2(abi.FuncPCABI0(libc_kill_trampoline), pid, uintptr(sig))
}

func osyield1()

//go:nosplit
func osyield_no_g() {
	osyield1()
}

//go:nosplit
func osyield() {
	r, err := syscall0(abi.FuncPCABI0(libc_sched_yield_trampoline))
	if int32(r) == -1 {
		println("syscall osyield failed: ", hex(err))
		throw("syscall osyield")
	}
}

//go:nosplit
func sysconf(name int32) uintptr {
	r, _ := syscall1(abi.FuncPCABI0(libc_sysconf_trampoline), uintptr(name))
	if int32(r) == -1 {
		throw("syscall sysconf")
	}
	return r
}

// pthread functions return their error code in the main return value;
// errno is not used.

//go:nosplit
func pthread_attr_destroy(attr *pthread_attr) int32 {
	r, _ := syscall1(abi.FuncPCABI0(libc_pthread_attr_destroy_trampoline), uintptr(unsafe.Pointer(attr)))
	return int32(r)
}

func pthread_attr_init1(attr uintptr) int32

//go:nosplit
func pthread_attr_init(attr *pthread_attr) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall1(abi.FuncPCABI0(libc_pthread_attr_init_trampoline), uintptr(unsafe.Pointer(attr)))
		return int32(r)
	}
	return pthread_attr_init1(uintptr(unsafe.Pointer(attr)))
}

func pthread_attr_setdetachstate1(attr uintptr, state int32) int32

//go:nosplit
func pthread_attr_setdetachstate(attr *pthread_attr, state int32) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall2(abi.FuncPCABI0(libc_pthread_attr_setdetachstate_trampoline), uintptr(unsafe.Pointer(attr)), uintptr(state))
		return int32(r)
	}
	return pthread_attr_setdetachstate1(uintptr(unsafe.Pointer(attr)), state)
}

func pthread_attr_setstacksize1(attr uintptr, size uint64) int32

//go:nosplit
func pthread_attr_setstacksize(attr *pthread_attr, size uint64) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall2(abi.FuncPCABI0(libc_pthread_attr_setstacksize_trampoline), uintptr(unsafe.Pointer(attr)), uintptr(size))
		return int32(r)
	}
	return pthread_attr_setstacksize1(uintptr(unsafe.Pointer(attr)), size)
}

func pthread_create1(tid, attr, fn, arg uintptr) int32

//go:nosplit
func pthread_create(tid *pthread, attr *pthread_attr, fn uintptr, arg unsafe.Pointer) int32 {
	gp := getg()
	if gp != nil {
		r, _ := syscall4(abi.FuncPCABI0(libc_pthread_create_trampoline), uintptr(unsafe.Pointer(tid)), uintptr(unsafe.Pointer(attr)), fn, uintptr(arg))
		return int32(r)
	}
	return pthread_create1(uintptr(unsafe.Pointer(tid)), uintptr(unsafe.Pointer(attr)), fn, uintptr(arg))
}

func sigprocmask1(how, new, old uintptr)

//go:nosplit
func sigprocmask(how int32, new, old *sigset) {
	gp := getg()
	if gp != nil && gp.m != nil {
		r, err := syscall3(abi.FuncPCABI0(libc_sigprocmask_trampoline), uintptr(how), uintptr(unsafe.Pointer(new)), uintptr(unsafe.Pointer(old)))
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
	r, _ := syscall0(abi.FuncPCABI0(libc_pthread_self_trampoline))
	return pthread(r)
}

//go:nosplit
func find_thread(name *byte) int32 {
	r, _ := syscall1(abi.FuncPCABI0(libc_find_thread_trampoline), uintptr(unsafe.Pointer(name)))
	return int32(r)
}

//go:nosplit
func send_signal(thread int32, sig uint32) int32 {
	r, _ := syscall2(abi.FuncPCABI0(libc_send_signal_trampoline), uintptr(thread), uintptr(sig))
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
	syscall2(abi.FuncPCABI0(libc_send_signal_trampoline), uintptr(mp.procid), uintptr(sig))
}

//go:nosplit
func fcntl(fd, cmd, arg int32) (int32, int32) {
	r, errno := syscall3(abi.FuncPCABI0(libc_fcntl_trampoline), uintptr(fd), uintptr(cmd), uintptr(arg))
	return int32(r), int32(errno)
}
