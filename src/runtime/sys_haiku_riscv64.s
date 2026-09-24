// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// System calls and other sys.stuff for riscv64, Haiku.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"
#include "cgo/abi_riscv64.h"

// Called by runtime·asmcgocall with the libcall pointer in A0. NOT USING GO CALLING CONVENTION.
TEXT runtime·asmsyscall6(SB),NOSPLIT,$8
	MOV	A0, 8(X2)
	MOV	libcall_fn(A0), X5
	MOV	libcall_args(A0), X6

	BEQZ	g, clearerrno_done
	MOV	g_m(g), X7
	MOV	(m_mOS+mOS_perrno)(X7), X7
	BEQZ	X7, clearerrno_done
	MOVW	ZERO, (X7)
clearerrno_done:

	BEQZ	X6, noargs
	MOV	0(X6), A0
	MOV	8(X6), A1
	MOV	16(X6), A2
	MOV	24(X6), A3
	MOV	32(X6), A4
	MOV	40(X6), A5
noargs:
	JALR	RA, X5

	MOV	8(X2), X6
	MOV	A0, libcall_r1(X6)
	MOV	A1, libcall_r2(X6)

	BEQZ	g, readerrno_done
	MOV	g_m(g), X7
	MOV	(m_mOS+mOS_perrno)(X7), X7
	BEQZ	X7, readerrno_done
	MOVW	(X7), A0
	MOV	A0, libcall_err(X6)
readerrno_done:
	RET

// Entry point for new OS threads created via pthread_create, with the M pointer in A0.
TEXT runtime·tstart(SB),NOSPLIT|TOPFRAME,$200
	SAVE_GPR(8)
	SAVE_FPR(104)

	MOV	m_g0(A0), g
	MOV	A0, g_m(g)

	MOV	X2, X5
	MOV	X5, (g_stack+stack_hi)(g)
	MOV	$(0x100000), X6
	SUB	X6, X5
	MOV	X5, (g_stack+stack_lo)(g)
	ADD	$const_stackGuard, X5
	MOV	X5, g_stackguard0(g)
	MOV	X5, g_stackguard1(g)

	CALL	runtime·save_g(SB)
	CALL	runtime·mstart(SB)

	RESTORE_GPR(8)
	RESTORE_FPR(104)
	MOV	ZERO, A0
	RET

// Invoked by libroot as a signal handler: A0 sig, A1 info, A2 ctx.
TEXT runtime·sigtramp(SB),NOSPLIT|TOPFRAME,$272
	SAVE_GPR(32)
	SAVE_FPR(128)

	MOVW	A0, 8(X2)
	MOV	A1, 16(X2)
	MOV	A2, 24(X2)

	CALL	runtime·load_g(SB)

	BEQZ	g, nosave
	MOV	g_m(g), X5
	ADD	$(m_mOS+mOS_libcall), X5, X6
	MOV	libcall_fn(X6), X7
	MOV	X7, 224(X2)
	MOV	libcall_args(X6), X7
	MOV	X7, 232(X2)
	MOV	libcall_n(X6), X7
	MOV	X7, 240(X2)
	MOV	libcall_r1(X6), X7
	MOV	X7, 248(X2)
	MOV	libcall_r2(X6), X7
	MOV	X7, 256(X2)
	MOV	(m_mOS+mOS_perrno)(X5), X7
	BEQZ	X7, nosave
	MOVW	(X7), X7
	MOVW	X7, 264(X2)
nosave:

	CALL	runtime·sigtrampgo(SB)

	BEQZ	g, norestore
	MOV	g_m(g), X5
	ADD	$(m_mOS+mOS_libcall), X5, X6
	MOV	224(X2), X7
	MOV	X7, libcall_fn(X6)
	MOV	232(X2), X7
	MOV	X7, libcall_args(X6)
	MOV	240(X2), X7
	MOV	X7, libcall_n(X6)
	MOV	248(X2), X7
	MOV	X7, libcall_r1(X6)
	MOV	256(X2), X7
	MOV	X7, libcall_r2(X6)
	MOV	(m_mOS+mOS_perrno)(X5), X7
	BEQZ	X7, norestore
	MOVW	264(X2), X6
	MOVW	X6, (X7)
norestore:

	RESTORE_GPR(32)
	RESTORE_FPR(128)
	RET

// The wrappers below run without a g and align SP for the C call.

TEXT runtime·sigfwd(SB),NOSPLIT,$8-32
	MOVW	sig+8(FP), A0
	MOV	info+16(FP), A1
	MOV	ctx+24(FP), A2
	MOV	fn+0(FP), X5
	MOV	X2, X9
	AND	$-16, X2
	JALR	RA, X5
	MOV	X9, X2
	RET

TEXT runtime·usleep1(SB),NOSPLIT,$8-4
	MOVWU	us+0(FP), A0
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_usleep(SB)
	MOV	X9, X2
	RET

TEXT runtime·osyield1(SB),NOSPLIT,$8
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_sched_yield(SB)
	MOV	X9, X2
	RET

TEXT runtime·exit1(SB),NOSPLIT,$8-4
	MOVW	code+0(FP), A0
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_exit(SB)
	MOV	X9, X2
	RET

TEXT runtime·write2(SB),NOSPLIT,$8-28
	MOV	fd+0(FP), A0
	MOV	p+8(FP), A1
	MOVW	n+16(FP), A2
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_write(SB)
	MOV	X9, X2
	MOVW	A0, ret+24(FP)
	RET

TEXT runtime·sigaction1(SB),NOSPLIT,$8-24
	MOV	sig+0(FP), A0
	MOV	new+8(FP), A1
	MOV	old+16(FP), A2
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_sigaction(SB)
	MOV	X9, X2
	RET

TEXT runtime·sigprocmask1(SB),NOSPLIT,$8-24
	MOV	how+0(FP), A0
	MOV	new+8(FP), A1
	MOV	old+16(FP), A2
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_sigprocmask(SB)
	MOV	X9, X2
	RET

TEXT runtime·pthread_attr_init1(SB),NOSPLIT,$8-12
	MOV	attr+0(FP), A0
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_pthread_attr_init(SB)
	MOV	X9, X2
	MOVW	A0, ret+8(FP)
	RET

TEXT runtime·pthread_attr_setdetachstate1(SB),NOSPLIT,$8-20
	MOV	attr+0(FP), A0
	MOVW	state+8(FP), A1
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_pthread_attr_setdetachstate(SB)
	MOV	X9, X2
	MOVW	A0, ret+16(FP)
	RET

TEXT runtime·pthread_attr_setstacksize1(SB),NOSPLIT,$8-20
	MOV	attr+0(FP), A0
	MOV	size+8(FP), A1
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_pthread_attr_setstacksize(SB)
	MOV	X9, X2
	MOVW	A0, ret+16(FP)
	RET

TEXT runtime·pthread_create1(SB),NOSPLIT,$8-36
	MOV	tid+0(FP), A0
	MOV	attr+8(FP), A1
	MOV	fn+16(FP), A2
	MOV	arg+24(FP), A3
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_pthread_create(SB)
	MOV	X9, X2
	MOVW	A0, ret+32(FP)
	RET

// Stores the byte offset of a new TLS slot from TP in runtime·tls_g.
TEXT runtime·haikuTlsInit(SB),NOSPLIT,$8
	MOV	X2, X9
	AND	$-16, X2
	CALL	libc_tls_allocate(SB)
	MOV	X9, X2
	SLL	$3, A0
	MOV	A0, runtime·tls_g(SB)
	RET

// B_HAIKU_ABI_GCC_4 and B_HAIKU_VERSION_1, checked by runtime_loader.
DATA  _gSharedObjectHaikuABI+0(SB)/4, $0x00040000
GLOBL _gSharedObjectHaikuABI(SB), NOPTR, $4
DATA  _gSharedObjectHaikuVersion+0(SB)/4, $0x00010000
GLOBL _gSharedObjectHaikuVersion(SB), NOPTR, $4
