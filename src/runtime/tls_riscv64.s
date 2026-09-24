// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"

// If !iscgo, this is a no-op.
//
// NOTE: mcall() assumes this clobbers only X31 (REG_TMP).
TEXT runtime·save_g(SB),NOSPLIT|NOFRAME,$0-0
#ifndef GOOS_openbsd
#ifndef GOOS_haiku
	MOVB	runtime·iscgo(SB), X31
	BEQZ	X31, nocgo
#endif
#endif
#ifdef GOOS_haiku
	MOV	runtime·tls_g(SB), X31
	ADD	TP, X31
	MOV	g, (X31)
#else
	MOV	g, runtime·tls_g(SB)
#endif
nocgo:
	RET

TEXT runtime·load_g(SB),NOSPLIT|NOFRAME,$0-0
#ifdef GOOS_haiku
	MOV	runtime·tls_g(SB), X31
	ADD	TP, X31
	MOV	(X31), g
#else
	MOV	runtime·tls_g(SB), g
#endif
	RET

#ifdef GOOS_haiku
GLOBL runtime·tls_g(SB), NOPTR, $8
#else
GLOBL runtime·tls_g(SB), TLSBSS, $8
#endif
