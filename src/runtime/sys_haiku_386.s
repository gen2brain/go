// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// System calls and other sys.stuff for 386, Haiku.
// Each library symbol is invoked through asmcgocall via the libcall
// trampoline runtime·asmsyscall6.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

// runtime·settls is a no-op on Haiku. Each thread's FS register is set up
// by the kernel to point at the per-thread TLS array, so there is no need
// to install a base. The G pointer is stored at fs:[runtime·tls_g], where
// runtime·tls_g is the byte offset returned by tls_allocate() (* 4) at
// program startup.
// setldt is unused on Haiku: the kernel installs each thread's FS base and
// rt0_go skips ldt0setup. It exists only to satisfy the reference from the
// (never executed on Haiku) ldt0setup path.
TEXT runtime·setldt(SB),NOSPLIT,$0
	RET

// Call a library function using the System V i386 (cdecl) convention:
// arguments are pushed on the stack right-to-left and the caller cleans up.
// The called function can take a maximum of 6 arguments.
//
// Called by runtime·asmcgocall or runtime·cgocall, which pass the pointer
// to the libcall as the single C argument (at 0(SP) before their CALL).
// NOT USING GO CALLING CONVENTION.
//
// The i386 ABI requires SP to be 16-byte aligned at every CALL site, and
// Haiku's libroot uses SSE (MOVAPS) internally, so the argument block is
// laid out on a freshly 16-aligned stack before the call.
TEXT runtime·asmsyscall6(SB),NOSPLIT|NOFRAME,$0
	PUSHL	BP
	MOVL	SP, BP
	PUSHL	BX
	PUSHL	SI
	PUSHL	DI
	// 8(BP) holds the *libcall passed by asmcgocall. Keep it in BX, which
	// the callee preserves across the libc CALL below (cdecl callee-saved).
	MOVL	8(BP), BX

	// Clear g.m.perrno so a successful call leaves errno at 0.
	get_tls(CX)
	MOVL	g(CX), DX
	CMPL	DX, $0
	JEQ	skiperrno1
	MOVL	g_m(DX), DX
	MOVL	(m_mOS+mOS_perrno)(DX), DX
	CMPL	DX, $0
	JEQ	skiperrno1
	MOVL	$0, 0(DX)

skiperrno1:
	// Marshal up to 6 arguments onto a 16-byte aligned stack.
	ANDL	$~15, SP
	SUBL	$32, SP
	MOVL	libcall_args(BX), SI
	MOVL	libcall_n(BX), CX
	MOVL	SP, DI

copyargs:
	TESTL	CX, CX
	JEQ	docall
	MOVL	0(SI), AX
	MOVL	AX, 0(DI)
	ADDL	$4, SI
	ADDL	$4, DI
	DECL	CX
	JMP	copyargs

docall:
	MOVL	libcall_fn(BX), AX
	CALL	AX

	MOVL	AX, libcall_r1(BX)
	MOVL	DX, libcall_r2(BX)

	// Save errno; the caller inspects it after the call.
	get_tls(CX)
	MOVL	g(CX), DX
	CMPL	DX, $0
	JEQ	skiperrno2
	MOVL	g_m(DX), DX
	MOVL	(m_mOS+mOS_perrno)(DX), DX
	CMPL	DX, $0
	JEQ	skiperrno2
	MOVL	0(DX), AX
	MOVL	AX, libcall_err(BX)

skiperrno2:
	LEAL	-12(BP), SP
	POPL	DI
	POPL	SI
	POPL	BX
	POPL	BP
	RET

// uint32 tstart(M *newm);
// Entry point for new OS threads created via pthread_create. The single
// argument (newm) arrives at 4(SP) under the cdecl convention.
TEXT runtime·tstart(SB),NOSPLIT|NOFRAME,$0
	MOVL	4(SP), BX		// newm
	MOVL	m_g0(BX), DX		// g0

	// Make TLS entries point at g and m.
	get_tls(CX)
	MOVL	DX, g(CX)
	MOVL	BX, g_m(DX)

	// Layout new m scheduler stack on os stack.
	MOVL	SP, AX
	MOVL	AX, (g_stack+stack_hi)(DX)
	SUBL	$(0x100000), AX		// stack size
	MOVL	AX, (g_stack+stack_lo)(DX)
	ADDL	$const_stackGuard, AX
	MOVL	AX, g_stackguard0(DX)
	MOVL	AX, g_stackguard1(DX)

	CLD

	CALL	runtime·stackcheck(SB)
	CALL	runtime·mstart(SB)

	XORL	AX, AX			// return 0 == success
	RET

// Careful, this is invoked by libroot as a signal handler. We must preserve
// callee-saved registers as per i386 cdecl. The handler arguments
// (sig, info, ctx) arrive on the stack above our frame.
TEXT runtime·sigtramp(SB),NOSPLIT|TOPFRAME|NOFRAME,$0
	// Executing on the alternate signal stack, so allocate a frame ourselves.
	SUBL	$52, SP
	MOVL	BX, 36(SP)
	MOVL	BP, 40(SP)
	MOVL	SI, 44(SP)
	MOVL	DI, 48(SP)

	get_tls(BX)
	MOVL	g(BX), BP
	CMPL	BP, $0
	JNE	allgood

	// No g: report a bad signal on a foreign thread. Build a sigctxt
	// {info, ctx} on our frame and pass its address.
	MOVL	60(SP), AX		// info
	MOVL	AX, 12(SP)
	MOVL	64(SP), AX		// ctx
	MOVL	AX, 16(SP)
	MOVL	56(SP), AX		// sig
	MOVL	AX, 0(SP)
	LEAL	12(SP), AX
	MOVL	AX, 4(SP)
	CALL	runtime·badsignal(SB)
	JMP	exit

allgood:
	// Save m->libcall. We might be interrupting runtime·asmcgocall.
	MOVL	g_m(BP), BX
	LEAL	(m_mOS+mOS_libcall)(BX), DI
	MOVL	libcall_fn(DI), AX
	MOVL	AX, 12(SP)
	MOVL	libcall_args(DI), AX
	MOVL	AX, 16(SP)
	MOVL	libcall_n(DI), AX
	MOVL	AX, 20(SP)
	MOVL	libcall_r1(DI), AX
	MOVL	AX, 24(SP)
	MOVL	libcall_r2(DI), AX
	MOVL	AX, 28(SP)

	// Save errno; signal handlers may clobber it.
	MOVL	(m_mOS+mOS_perrno)(BX), AX
	MOVL	0(AX), AX
	MOVL	AX, 32(SP)

	// sigtrampgo(sig, info, ctx)
	MOVL	56(SP), AX
	MOVL	AX, 0(SP)
	MOVL	60(SP), AX
	MOVL	AX, 4(SP)
	MOVL	64(SP), AX
	MOVL	AX, 8(SP)
	CALL	runtime·sigtrampgo(SB)

	get_tls(BX)
	MOVL	g(BX), BP
	MOVL	g_m(BP), BX

	// Restore m->libcall.
	LEAL	(m_mOS+mOS_libcall)(BX), DI
	MOVL	12(SP), AX
	MOVL	AX, libcall_fn(DI)
	MOVL	16(SP), AX
	MOVL	AX, libcall_args(DI)
	MOVL	20(SP), AX
	MOVL	AX, libcall_n(DI)
	MOVL	24(SP), AX
	MOVL	AX, libcall_r1(DI)
	MOVL	28(SP), AX
	MOVL	AX, libcall_r2(DI)

	// Restore errno.
	MOVL	(m_mOS+mOS_perrno)(BX), DI
	MOVL	32(SP), AX
	MOVL	AX, 0(DI)

exit:
	MOVL	36(SP), BX
	MOVL	40(SP), BP
	MOVL	44(SP), SI
	MOVL	48(SP), DI
	ADDL	$52, SP
	RET

// void sigfwd(void (*fn)(int, siginfo*, void*), int sig, siginfo *info, void *ctx)
TEXT runtime·sigfwd(SB),NOSPLIT,$0-16
	MOVL	fn+0(FP),   AX
	MOVL	sig+4(FP),  CX
	MOVL	info+8(FP), DX
	MOVL	ctx+12(FP), SI
	MOVL	SP, BX			// callee-saved across the CALL
	ANDL	$~15, SP
	SUBL	$16, SP
	MOVL	CX, 0(SP)
	MOVL	DX, 4(SP)
	MOVL	SI, 8(SP)
	CALL	AX
	MOVL	BX, SP
	RET

// Called from runtime·usleep (Go). Can be called on Go stack, on OS stack,
// can also be called in cgo callback path without a g->m.
TEXT runtime·usleep1(SB),NOSPLIT,$0
	MOVL	us+0(FP), SI		// preserved across the stack switch
	MOVL	$usleep2<>(SB), AX

	// Execute call on m->g0.
	get_tls(CX)
	CMPL	CX, $0
	JE	noswitch

	MOVL	g(CX), DX
	CMPL	DX, $0
	JE	noswitch
	MOVL	g_m(DX), DX
	CMPL	DX, $0
	JE	noswitch

	MOVL	m_g0(DX), BX
	MOVL	g(CX), DI
	CMPL	DI, BX
	JNE	switch
	// already on m->g0
	CALL	AX
	RET

switch:
	// Switch to m->g0 stack and back.
	MOVL	(g_sched+gobuf_sp)(BX), BX
	MOVL	SP, -4(BX)
	LEAL	-4(BX), SP
	CALL	AX
	MOVL	0(SP), SP
	RET

noswitch:
	// Not a Go-managed thread. Do not switch stack.
	CALL	AX
	RET

// Runs on OS stack. duration (in microseconds) is kept in SI by usleep1.
TEXT usleep2<>(SB),NOSPLIT,$0
	MOVL	SP, BX
	ANDL	$~15, SP
	SUBL	$16, SP
	MOVL	SI, 0(SP)
	LEAL	libc_usleep(SB), AX
	CALL	AX
	MOVL	BX, SP
	RET

// Runs on OS stack, called from runtime·osyield.
TEXT runtime·osyield1(SB),NOSPLIT,$0
	MOVL	SP, BX
	ANDL	$~15, SP
	LEAL	libc_sched_yield(SB), AX
	CALL	AX
	MOVL	BX, SP
	RET

// Trivial wrappers used when no g is available (newosproc0/exit/write1).

TEXT runtime·exit1(SB),NOSPLIT,$0-4
	MOVL	code+0(FP), CX
	MOVL	SP, BX
	ANDL	$~15, SP
	SUBL	$16, SP
	MOVL	CX, 0(SP)
	LEAL	libc_exit(SB), AX
	CALL	AX
	MOVL	BX, SP
	RET

TEXT runtime·write2(SB),NOSPLIT,$0-16
	MOVL	fd+0(FP), CX
	MOVL	p+4(FP),  DX
	MOVL	n+8(FP),  SI
	MOVL	SP, BX
	ANDL	$~15, SP
	SUBL	$16, SP
	MOVL	CX, 0(SP)
	MOVL	DX, 4(SP)
	MOVL	SI, 8(SP)
	LEAL	libc_write(SB), AX
	CALL	AX
	MOVL	BX, SP
	MOVL	AX, ret+12(FP)
	RET

TEXT runtime·sigaction1(SB),NOSPLIT,$0-12
	MOVL	sig+0(FP), CX
	MOVL	new+4(FP), DX
	MOVL	old+8(FP), SI
	MOVL	SP, BX
	ANDL	$~15, SP
	SUBL	$16, SP
	MOVL	CX, 0(SP)
	MOVL	DX, 4(SP)
	MOVL	SI, 8(SP)
	LEAL	libc_sigaction(SB), AX
	CALL	AX
	MOVL	BX, SP
	RET

TEXT runtime·sigprocmask1(SB),NOSPLIT,$0-12
	MOVL	how+0(FP), CX
	MOVL	new+4(FP), DX
	MOVL	old+8(FP), SI
	MOVL	SP, BX
	ANDL	$~15, SP
	SUBL	$16, SP
	MOVL	CX, 0(SP)
	MOVL	DX, 4(SP)
	MOVL	SI, 8(SP)
	LEAL	libc_sigprocmask(SB), AX
	CALL	AX
	MOVL	BX, SP
	RET

TEXT runtime·pthread_attr_init1(SB),NOSPLIT,$0-8
	MOVL	attr+0(FP), CX
	MOVL	SP, BX
	ANDL	$~15, SP
	SUBL	$16, SP
	MOVL	CX, 0(SP)
	LEAL	libc_pthread_attr_init(SB), AX
	CALL	AX
	MOVL	BX, SP
	MOVL	AX, ret+4(FP)
	RET

TEXT runtime·pthread_attr_setdetachstate1(SB),NOSPLIT,$0-12
	MOVL	attr+0(FP),  CX
	MOVL	state+4(FP), DX
	MOVL	SP, BX
	ANDL	$~15, SP
	SUBL	$16, SP
	MOVL	CX, 0(SP)
	MOVL	DX, 4(SP)
	LEAL	libc_pthread_attr_setdetachstate(SB), AX
	CALL	AX
	MOVL	BX, SP
	MOVL	AX, ret+8(FP)
	RET

// size is declared uint64 for the shared syscall path; on 386 pthread_attr_setstacksize
// takes a 32-bit size_t, so only the low word is passed.
TEXT runtime·pthread_attr_setstacksize1(SB),NOSPLIT,$0-16
	MOVL	attr+0(FP), CX
	MOVL	size_lo+4(FP), DX
	MOVL	SP, BX
	ANDL	$~15, SP
	SUBL	$16, SP
	MOVL	CX, 0(SP)
	MOVL	DX, 4(SP)
	LEAL	libc_pthread_attr_setstacksize(SB), AX
	CALL	AX
	MOVL	BX, SP
	MOVL	AX, ret+12(FP)
	RET

TEXT runtime·pthread_create1(SB),NOSPLIT,$0-20
	MOVL	tid+0(FP),  CX
	MOVL	attr+4(FP), DX
	MOVL	fn+8(FP),   SI
	MOVL	arg+12(FP), DI
	MOVL	SP, BX
	ANDL	$~15, SP
	SUBL	$16, SP
	MOVL	CX, 0(SP)
	MOVL	DX, 4(SP)
	MOVL	SI, 8(SP)
	MOVL	DI, 12(SP)
	LEAL	libc_pthread_create(SB), AX
	CALL	AX
	MOVL	BX, SP
	MOVL	AX, ret+16(FP)
	RET

// haikuTlsInit allocates a per-process TLS slot via libroot's tls_allocate()
// and stores the resulting byte offset (slot index * 4) into runtime.tls_g.
// Called once from rt0_go before any FS-relative TLS access.
TEXT runtime·haikuTlsInit(SB),NOSPLIT,$0
	MOVL	SP, BX
	ANDL	$~15, SP
	LEAL	libc_tls_allocate(SB), AX
	CALL	AX			// AX = tls_allocate() (slot index, int32)
	MOVL	BX, SP
	SHLL	$2, AX			// byte offset = slot * 4
	MOVL	AX, runtime·tls_g(SB)
	RET

// runtime·tls_g holds the byte offset (slot index * 4) of Go's G pointer
// inside the per-thread TLS array. It is filled in at startup by
// runtime·haikuTlsInit which calls tls_allocate().
GLOBL runtime·tls_g+0(SB), NOPTR, $4

// runtime_loader inspects these to identify the binary's ABI.
// 0x00040000 = B_HAIKU_ABI_GCC_4. 0x00010000 = B_HAIKU_VERSION_1.
DATA  _gSharedObjectHaikuABI+0(SB)/4, $0x00040000
GLOBL _gSharedObjectHaikuABI(SB), NOPTR, $4
DATA  _gSharedObjectHaikuVersion+0(SB)/4, $0x00010000
GLOBL _gSharedObjectHaikuVersion(SB), NOPTR, $4
