// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// System calls and other sys.stuff for AMD64, Haiku.
// Each library symbol is invoked through asmcgocall via the libcall
// trampoline runtime·asmsyscall6.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

// runtime·settls is a no-op on Haiku. Each thread's FS register is set up
// by the kernel to point at the per-thread TLS array, so there is no need
// to install a base. The G pointer is stored at fs:[runtime·tls_g], where
// runtime·tls_g is the byte offset returned by tls_allocate() (* 8) at
// program startup.
TEXT runtime·settls(SB),NOSPLIT,$0
	RET

// Call a library function with SysV calling conventions.
// The called function can take a maximum of 6 INTEGER class arguments.
//
// Called by runtime·asmcgocall or runtime·cgocall.
// NOT USING GO CALLING CONVENTION.
//
// NOFRAME prevents the assembler from auto-inserting a frame-pointer push,
// which would shift RSP by an odd amount (8 bytes) and misalign the stack
// for the libc CALL below. The SysV ABI requires RSP to be 16-byte aligned
// at every call site (callee sees RSP%16 == 8); Haiku's libroot snprintf
// uses MOVAPS internally and #GP faults if this is violated.
TEXT runtime·asmsyscall6(SB),NOSPLIT|NOFRAME,$0
	// asmcgocall puts the argument pointer into DI.
	PUSHQ	DI			// save for later
	MOVQ	libcall_fn(DI), AX
	MOVQ	libcall_args(DI), R11
	MOVQ	libcall_n(DI), R10

	get_tls(CX)
	MOVQ	g(CX), BX
	CMPQ	BX, $0
	JEQ	skiperrno1
	MOVQ	g_m(BX), BX
	MOVQ	(m_mOS+mOS_perrno)(BX), DX
	CMPQ	DX, $0
	JEQ	skiperrno1
	MOVL	$0, 0(DX)

skiperrno1:
	CMPQ	R11, $0
	JEQ	skipargs
	MOVQ	0(R11), DI
	MOVQ	8(R11), SI
	MOVQ	16(R11), DX
	MOVQ	24(R11), CX
	MOVQ	32(R11), R8
	MOVQ	40(R11), R9
skipargs:

	CALL	AX

	POPQ	DI
	MOVQ	AX, libcall_r1(DI)
	MOVQ	DX, libcall_r2(DI)

	get_tls(CX)
	MOVQ	g(CX), BX
	CMPQ	BX, $0
	JEQ	skiperrno2
	MOVQ	g_m(BX), BX
	MOVQ	(m_mOS+mOS_perrno)(BX), AX
	CMPQ	AX, $0
	JEQ	skiperrno2
	MOVL	0(AX), AX
	MOVQ	AX, libcall_err(DI)

skiperrno2:
	RET

// uint32 tstart(M *newm);
// Entry point for new OS threads created via pthread_create.
TEXT runtime·tstart(SB),NOSPLIT,$0
	// DI contains first arg newm
	MOVQ	m_g0(DI), DX		// g

	// Make TLS entries point at g and m.
	get_tls(BX)
	MOVQ	DX, g(BX)
	MOVQ	DI, g_m(DX)

	// Layout new m scheduler stack on os stack.
	MOVQ	SP, AX
	MOVQ	AX, (g_stack+stack_hi)(DX)
	SUBQ	$(0x100000), AX		// stack size
	MOVQ	AX, (g_stack+stack_lo)(DX)
	ADDQ	$const_stackGuard, AX
	MOVQ	AX, g_stackguard0(DX)
	MOVQ	AX, g_stackguard1(DX)

	CLD

	CALL	runtime·stackcheck(SB)
	CALL	runtime·mstart(SB)

	XORL	AX, AX			// return 0 == success
	RET

// Careful, this is invoked by libroot as a signal handler. We must preserve
// callee-saved registers as per AMD64 SysV ABI.
TEXT runtime·sigtramp(SB),NOSPLIT|TOPFRAME|NOFRAME,$0
	// Note that we are executing on altsigstack here, so we have
	// more stack available than NOSPLIT would have us believe.
	// Allocate a frame ourselves.
	SUBQ	$168, SP
	// save callee-saved registers
	MOVQ	BX,  24(SP)
	MOVQ	BP,  32(SP)
	MOVQ	R12, 40(SP)
	MOVQ	R13, 48(SP)
	MOVQ	R14, 56(SP)
	MOVQ	R15, 64(SP)

	get_tls(BX)
	// check that g exists
	MOVQ	g(BX), R10
	CMPQ	R10, $0
	JNE	allgood
	MOVQ	SI, 72(SP)
	MOVQ	DX, 80(SP)
	LEAQ	72(SP), AX
	MOVQ	DI, 0(SP)
	MOVQ	AX, 8(SP)
	MOVQ	$runtime·badsignal(SB), AX
	CALL	AX
	JMP	exit

allgood:
	// Save m->libcall. We need to do this because we
	// might get interrupted by a signal in runtime·asmcgocall.
	MOVQ	g_m(R10), BP
	LEAQ	(m_mOS+mOS_libcall)(BP), R11
	MOVQ	libcall_fn(R11), R10
	MOVQ	R10, 72(SP)
	MOVQ	libcall_args(R11), R10
	MOVQ	R10, 80(SP)
	MOVQ	libcall_n(R11), R10
	MOVQ	R10, 88(SP)
	MOVQ	libcall_r1(R11), R10
	MOVQ	R10, 152(SP)
	MOVQ	libcall_r2(R11), R10
	MOVQ	R10, 160(SP)

	// save errno; signal handlers may clobber it.
	MOVQ	(m_mOS+mOS_perrno)(BP), R10
	MOVL	0(R10), R10
	MOVQ	R10, 144(SP)

	// prepare call: sigtrampgo(sig, info, ctx)
	MOVQ	DI, 0(SP)
	MOVQ	SI, 8(SP)
	MOVQ	DX, 16(SP)
	CALL	runtime·sigtrampgo(SB)

	get_tls(BX)
	MOVQ	g(BX), BP
	MOVQ	g_m(BP), BP

	// restore libcall
	LEAQ	(m_mOS+mOS_libcall)(BP), R11
	MOVQ	72(SP), R10
	MOVQ	R10, libcall_fn(R11)
	MOVQ	80(SP), R10
	MOVQ	R10, libcall_args(R11)
	MOVQ	88(SP), R10
	MOVQ	R10, libcall_n(R11)
	MOVQ	152(SP), R10
	MOVQ	R10, libcall_r1(R11)
	MOVQ	160(SP), R10
	MOVQ	R10, libcall_r2(R11)

	// restore errno
	MOVQ	(m_mOS+mOS_perrno)(BP), R11
	MOVQ	144(SP), R10
	MOVL	R10, 0(R11)

exit:
	MOVQ	24(SP), BX
	MOVQ	32(SP), BP
	MOVQ	40(SP), R12
	MOVQ	48(SP), R13
	MOVQ	56(SP), R14
	MOVQ	64(SP), R15
	ADDQ	$168, SP
	RET

TEXT runtime·sigfwd(SB),NOSPLIT,$0-32
	MOVQ	fn+0(FP),    AX
	MOVL	sig+8(FP),   DI
	MOVQ	info+16(FP), SI
	MOVQ	ctx+24(FP),  DX
	MOVQ	SP, BX		// callee-saved
	ANDQ	$~15, SP	// alignment for x86_64 ABI
	CALL	AX
	MOVQ	BX, SP
	RET

// Called from runtime·usleep (Go). Can be called on Go stack, on OS stack,
// can also be called in cgo callback path without a g->m.
TEXT runtime·usleep1(SB),NOSPLIT,$0
	MOVL	us+0(FP), DI
	MOVQ	$usleep2<>(SB), AX

	// Execute call on m->g0.
	get_tls(R15)
	CMPQ	R15, $0
	JE	noswitch

	MOVQ	g(R15), R13
	CMPQ	R13, $0
	JE	noswitch
	MOVQ	g_m(R13), R13
	CMPQ	R13, $0
	JE	noswitch

	MOVQ	m_g0(R13), R14
	CMPQ	g(R15), R14
	JNE	switch
	// already on m->g0
	CALL	AX
	RET

switch:
	// Switch to m->g0 stack and back.
	MOVQ	(g_sched+gobuf_sp)(R14), R14
	MOVQ	SP, -8(R14)
	LEAQ	-8(R14), SP
	CALL	AX
	MOVQ	0(SP), SP
	RET

noswitch:
	// Not a Go-managed thread. Do not switch stack.
	CALL	AX
	RET

// Runs on OS stack. duration (in µs units) is in DI.
// The SysV ABI requires RSP to be 16-byte aligned at every CALL site
// (so the callee sees RSP%16 == 8). Go's calling convention does not
// always provide that when calling these ABI0 trampolines, so we realign
// before invoking libc, otherwise libc routines that use SSE (e.g.
// snprintf via vfprintf) take a #GP fault on MOVAPS.
TEXT usleep2<>(SB),NOSPLIT,$0
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_usleep(SB), AX
	CALL	AX
	MOVQ	BX, SP
	RET

// Runs on OS stack, called from runtime·osyield.
TEXT runtime·osyield1(SB),NOSPLIT,$0
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_sched_yield(SB), AX
	CALL	AX
	MOVQ	BX, SP
	RET

// Trivial wrappers used when no g is available (newosproc0/exit/write1).

TEXT runtime·exit1(SB),NOSPLIT,$0-4
	MOVL	code+0(FP), DI
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_exit(SB), AX
	CALL	AX
	MOVQ	BX, SP
	RET

TEXT runtime·write2(SB),NOSPLIT,$0-28
	MOVQ	fd+0(FP),  DI
	MOVQ	p+8(FP),   SI
	MOVL	n+16(FP),  DX
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_write(SB), AX
	CALL	AX
	MOVQ	BX, SP
	MOVL	AX, ret+24(FP)
	RET

TEXT runtime·sigaction1(SB),NOSPLIT,$0-24
	MOVQ	sig+0(FP),  DI
	MOVQ	new+8(FP),  SI
	MOVQ	old+16(FP), DX
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_sigaction(SB), AX
	CALL	AX
	MOVQ	BX, SP
	RET

TEXT runtime·sigprocmask1(SB),NOSPLIT,$0-24
	MOVQ	how+0(FP), DI
	MOVQ	new+8(FP), SI
	MOVQ	old+16(FP), DX
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_sigprocmask(SB), AX
	CALL	AX
	MOVQ	BX, SP
	RET

TEXT runtime·pthread_attr_init1(SB),NOSPLIT,$0-12
	MOVQ	attr+0(FP), DI
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_pthread_attr_init(SB), AX
	CALL	AX
	MOVQ	BX, SP
	MOVL	AX, ret+8(FP)
	RET

TEXT runtime·pthread_attr_setdetachstate1(SB),NOSPLIT,$0-20
	MOVQ	attr+0(FP),  DI
	MOVL	state+8(FP), SI
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_pthread_attr_setdetachstate(SB), AX
	CALL	AX
	MOVQ	BX, SP
	MOVL	AX, ret+16(FP)
	RET

TEXT runtime·pthread_attr_setstacksize1(SB),NOSPLIT,$0-20
	MOVQ	attr+0(FP), DI
	MOVQ	size+8(FP), SI
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_pthread_attr_setstacksize(SB), AX
	CALL	AX
	MOVQ	BX, SP
	MOVL	AX, ret+16(FP)
	RET

TEXT runtime·pthread_create1(SB),NOSPLIT,$0-36
	MOVQ	tid+0(FP),  DI
	MOVQ	attr+8(FP), SI
	MOVQ	fn+16(FP),  DX
	MOVQ	arg+24(FP), CX
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_pthread_create(SB), AX
	CALL	AX
	MOVQ	BX, SP
	MOVL	AX, ret+32(FP)
	RET

// haikuTlsInit allocates a per-process TLS slot via libroot's tls_allocate()
// and stores the resulting byte offset (slot index * 8) into runtime.tls_g.
// Called once from rt0_go before any FS-relative TLS access.
TEXT runtime·haikuTlsInit(SB),NOSPLIT,$0
	MOVQ	SP, BX
	ANDQ	$~15, SP
	LEAQ	libc_tls_allocate(SB), AX
	CALL	AX			// AX = tls_allocate() (slot index, int32)
	MOVQ	BX, SP
	MOVLQSX	AX, AX			// sign-extend int32 result to 64 bits
	SHLQ	$3, AX			// byte offset = slot * 8
	MOVQ	AX, runtime·tls_g(SB)
	RET

// runtime·tls_g holds the byte offset (slot index * 8) of Go's G pointer
// inside the per-thread TLS array. It is filled in at startup by
// runtime·haikuTlsInit which calls tls_allocate().
GLOBL runtime·tls_g+0(SB), NOPTR, $8

// runtime_loader inspects these to identify the binary's ABI.
// 0x00040000 = B_HAIKU_ABI_GCC_4. 0x00010000 = B_HAIKU_VERSION_1.
DATA  _gSharedObjectHaikuABI+0(SB)/4, $0x00040000
GLOBL _gSharedObjectHaikuABI(SB), NOPTR, $4
DATA  _gSharedObjectHaikuVersion+0(SB)/4, $0x00010000
GLOBL _gSharedObjectHaikuVersion(SB), NOPTR, $4
