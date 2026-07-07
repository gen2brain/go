// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// System calls and other sys.stuff for ARM64, Haiku.
// Each library symbol is invoked through asmcgocall via the libcall
// trampoline runtime·asmsyscall6. The G pointer lives in R28 and is
// preserved across libroot calls (AAPCS64 makes X28 callee-saved), so no
// per-call TLS reload is needed. runtime·tls_g (declared in tls_arm64.s) only
// serves save_g/load_g on the cgo path.

#include "go_asm.h"
#include "go_tls.h"
#include "tls_arm64.h"
#include "funcdata.h"
#include "textflag.h"
#include "cgo/abi_arm64.h"

// runtime·settls is a no-op on Haiku. Each thread's TPIDR_EL0 is set up by the
// kernel to point at the per-thread TLS array.
TEXT runtime·settls(SB),NOSPLIT,$0
	RET

// Call a library function with the AAPCS64 calling convention.
// The called function can take a maximum of 6 INTEGER class arguments.
//
// Called by runtime·asmcgocall or runtime·cgocall, which passes the pointer
// to the libcall struct in R0. NOT USING GO CALLING CONVENTION.
TEXT runtime·asmsyscall6(SB),NOSPLIT,$16-0
	MOVD	R0, 8(RSP)		// save libcall ptr across the C call
	MOVD	libcall_fn(R0), R9	// fn
	MOVD	libcall_args(R0), R10	// args pointer

	// clear errno via g.m.mOS.perrno
	CBZ	g, clearerrno_done
	MOVD	g_m(g), R11
	MOVD	(m_mOS+mOS_perrno)(R11), R11
	CBZ	R11, clearerrno_done
	MOVW	ZR, (R11)
clearerrno_done:

	CBZ	R10, noargs
	MOVD	0(R10), R0
	MOVD	8(R10), R1
	MOVD	16(R10), R2
	MOVD	24(R10), R3
	MOVD	32(R10), R4
	MOVD	40(R10), R5
noargs:
	CALL	(R9)

	MOVD	8(RSP), R10		// reload libcall ptr
	MOVD	R0, libcall_r1(R10)
	MOVD	R1, libcall_r2(R10)

	// read errno into libcall.err
	CBZ	g, readerrno_done
	MOVD	g_m(g), R11
	MOVD	(m_mOS+mOS_perrno)(R11), R11
	CBZ	R11, readerrno_done
	MOVWU	(R11), R0		// zero-extend errno (a 32-bit B_ code), matching
	MOVD	R0, libcall_err(R10)	// amd64's MOVL, so the errno table matches
readerrno_done:
	RET

// uint32 tstart(M *newm);
// Entry point for new OS threads created via pthread_create. The M pointer is
// passed in R0 per AAPCS64.
TEXT runtime·tstart(SB),NOSPLIT,$0
	MOVD	m_g0(R0), g		// g = newm.g0
	MOVD	R0, g_m(g)		// g0.m = newm

	// Layout the scheduler stack bounds over the OS stack.
	MOVD	RSP, R1
	MOVD	R1, (g_stack+stack_hi)(g)
	SUB	$(0x100000), R1, R1	// stack size
	MOVD	R1, (g_stack+stack_lo)(g)
	ADD	$const_stackGuard, R1, R1
	MOVD	R1, g_stackguard0(g)
	MOVD	R1, g_stackguard1(g)

	BL	runtime·save_g(SB)
	BL	runtime·mstart(SB)

	MOVD	$0, R0			// return 0 == success
	RET

// Careful, this is invoked by libroot as a signal handler. Callee-saved
// registers are preserved for the signal-forwarding path.
// C ABI: R0 sig, R1 info, R2 ctx.
TEXT runtime·sigtramp(SB),NOSPLIT|TOPFRAME,$240
	// Save callee-save registers in case of signal forwarding.
	SAVE_R19_TO_R28(8*4)
	SAVE_F8_TO_F15(8*14)

	// Recover g. It may be unset if the signal hit a non-Go thread or, on
	// the cgo path, C code that clobbered R28. load_g clobbers R0 and R27
	// but leaves R1 (info) and R2 (ctx) untouched.
	MOVW	R0, 176(RSP)		// save signum
	MOVBU	runtime·iscgo(SB), R0
	CBZ	R0, 2(PC)
	BL	runtime·load_g(SB)

	// Save m->libcall and errno: a signal may interrupt runtime·asmsyscall6
	// mid libcall, and sigtrampgo may itself issue libcalls.
	CBZ	g, nosave
	MOVD	g_m(g), R19
	ADD	$(m_mOS+mOS_libcall), R19, R20
	MOVD	libcall_fn(R20), R3
	MOVD	R3, 184(RSP)
	MOVD	libcall_args(R20), R3
	MOVD	R3, 192(RSP)
	MOVD	libcall_n(R20), R3
	MOVD	R3, 200(RSP)
	MOVD	libcall_r1(R20), R3
	MOVD	R3, 208(RSP)
	MOVD	libcall_r2(R20), R3
	MOVD	R3, 216(RSP)
	MOVD	(m_mOS+mOS_perrno)(R19), R3
	MOVW	(R3), R4
	MOVW	R4, 224(RSP)
nosave:

	MOVW	176(RSP), R0		// sig; R1=info, R2=ctx still intact
	MOVD	$runtime·sigtrampgo<ABIInternal>(SB), R3
	BL	(R3)

	// Restore m->libcall and errno (g/R28 preserved across the Go call).
	CBZ	g, norestore
	MOVD	g_m(g), R19
	ADD	$(m_mOS+mOS_libcall), R19, R20
	MOVD	184(RSP), R3
	MOVD	R3, libcall_fn(R20)
	MOVD	192(RSP), R3
	MOVD	R3, libcall_args(R20)
	MOVD	200(RSP), R3
	MOVD	R3, libcall_n(R20)
	MOVD	208(RSP), R3
	MOVD	R3, libcall_r1(R20)
	MOVD	216(RSP), R3
	MOVD	R3, libcall_r2(R20)
	MOVD	(m_mOS+mOS_perrno)(R19), R3
	MOVW	224(RSP), R4
	MOVW	R4, (R3)
norestore:

	RESTORE_R19_TO_R28(8*4)
	RESTORE_F8_TO_F15(8*14)
	RET

// void sigfwd(void (*fn)(int, siginfo*, void*), int sig, siginfo *info, void *ctx)
TEXT runtime·sigfwd(SB),NOSPLIT,$0-32
	MOVD	fn+0(FP),   R11
	MOVW	sig+8(FP),  R0
	MOVD	info+16(FP), R1
	MOVD	ctx+24(FP), R2
	CALL	(R11)
	RET

// Called from runtime·usleep (Go).
TEXT runtime·usleep1(SB),NOSPLIT,$0
	MOVW	us+0(FP), R0
	MOVD	$usleep2<>(SB), R9
	CALL	(R9)
	RET

// Runs on OS stack. duration (in microseconds) is in R0.
TEXT usleep2<>(SB),NOSPLIT,$0
	MOVD	$libc_usleep(SB), R11
	CALL	(R11)
	RET

// Runs on OS stack, called from runtime·osyield.
TEXT runtime·osyield1(SB),NOSPLIT,$0
	MOVD	$libc_sched_yield(SB), R11
	CALL	(R11)
	RET

// Trivial wrappers used when no g is available (newosproc0/exit/write1).

TEXT runtime·exit1(SB),NOSPLIT,$0-4
	MOVW	code+0(FP), R0
	MOVD	$libc_exit(SB), R11
	CALL	(R11)
	RET

TEXT runtime·write2(SB),NOSPLIT,$0-28
	MOVD	fd+0(FP), R0
	MOVD	p+8(FP),  R1
	MOVW	n+16(FP), R2
	MOVD	$libc_write(SB), R11
	CALL	(R11)
	MOVW	R0, ret+24(FP)
	RET

TEXT runtime·sigaction1(SB),NOSPLIT,$0-24
	MOVD	sig+0(FP),  R0
	MOVD	new+8(FP),  R1
	MOVD	old+16(FP), R2
	MOVD	$libc_sigaction(SB), R11
	CALL	(R11)
	RET

TEXT runtime·sigprocmask1(SB),NOSPLIT,$0-24
	MOVD	how+0(FP), R0
	MOVD	new+8(FP), R1
	MOVD	old+16(FP), R2
	MOVD	$libc_sigprocmask(SB), R11
	CALL	(R11)
	RET

TEXT runtime·pthread_attr_init1(SB),NOSPLIT,$0-12
	MOVD	attr+0(FP), R0
	MOVD	$libc_pthread_attr_init(SB), R11
	CALL	(R11)
	MOVW	R0, ret+8(FP)
	RET

TEXT runtime·pthread_attr_setdetachstate1(SB),NOSPLIT,$0-20
	MOVD	attr+0(FP),  R0
	MOVW	state+8(FP), R1
	MOVD	$libc_pthread_attr_setdetachstate(SB), R11
	CALL	(R11)
	MOVW	R0, ret+16(FP)
	RET

TEXT runtime·pthread_attr_setstacksize1(SB),NOSPLIT,$0-20
	MOVD	attr+0(FP), R0
	MOVD	size+8(FP), R1
	MOVD	$libc_pthread_attr_setstacksize(SB), R11
	CALL	(R11)
	MOVW	R0, ret+16(FP)
	RET

TEXT runtime·pthread_create1(SB),NOSPLIT,$0-36
	MOVD	tid+0(FP),  R0
	MOVD	attr+8(FP), R1
	MOVD	fn+16(FP),  R2
	MOVD	arg+24(FP), R3
	MOVD	$libc_pthread_create(SB), R11
	CALL	(R11)
	MOVW	R0, ret+32(FP)
	RET

// haikuTlsInit allocates a per-process TLS slot via libroot's tls_allocate()
// and stores the resulting byte offset (slot index * 8) into runtime.tls_g.
// Called once from rt0_go before any TPIDR-relative TLS access.
TEXT runtime·haikuTlsInit(SB),NOSPLIT,$0
	MOVD	$libc_tls_allocate(SB), R11
	CALL	(R11)			// R0 = tls_allocate() (slot index, int32)
	SXTW	R0, R0			// sign-extend int32 to 64 bits
	LSL	$3, R0, R0		// byte offset = slot * 8
	MOVD	R0, runtime·tls_g(SB)
	RET

// runtime_loader inspects these to identify the binary's ABI.
// 0x00040000 = B_HAIKU_ABI_GCC_4. 0x00010000 = B_HAIKU_VERSION_1.
DATA  _gSharedObjectHaikuABI+0(SB)/4, $0x00040000
GLOBL _gSharedObjectHaikuABI(SB), NOPTR, $4
DATA  _gSharedObjectHaikuVersion+0(SB)/4, $0x00010000
GLOBL _gSharedObjectHaikuVersion(SB), NOPTR, $4
