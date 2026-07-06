// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// runtime_loader invokes the entry point as a C function
//	entry(int argc, char **argv, char **env)
// so argc is at 4(SP) and the argv pointer at 8(SP), above the loader's
// return address at 0(SP). rt0_go expects argc at 0(SP) and the argv
// pointer at 4(SP), so shift them down before jumping in. The loader's
// return address is discarded; a Go program exits through _exit and never
// returns to the loader.
TEXT _rt0_386_haiku(SB),NOSPLIT,$0
	MOVL	4(SP), AX	// argc
	MOVL	8(SP), BX	// argv
	MOVL	AX, 0(SP)
	MOVL	BX, 4(SP)
	JMP	runtime·rt0_go(SB)

TEXT _rt0_386_haiku_lib(SB),NOSPLIT,$0
	JMP	_rt0_386_lib(SB)

// main is the symbol the C startup (start_dyn.o) calls under external
// linking. It is called main(argc, argv, env) via the cdecl convention,
// so drop the return address to leave argc at 0(SP) and argv at 4(SP)
// for rt0_go.
TEXT main(SB),NOSPLIT,$0
	ADDL	$4, SP
	JMP	runtime·rt0_go(SB)
