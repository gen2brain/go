// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// Haiku's i386 PLT needs EBX set to the GOT; jump to the address instead.
#ifdef GOARCH_386
#define LIBCALL(fn) LEAL fn(SB), AX; JMP AX
#else
#define LIBCALL(fn) JMP fn(SB)
#endif

TEXT ·libc_faccessat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_faccessat)

TEXT ·libc_fchmodat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_fchmodat)

TEXT ·libc_fchownat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_fchownat)

TEXT ·libc_fstatat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_fstatat)

TEXT ·libc_linkat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_linkat)

TEXT ·libc_openat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_openat)

TEXT ·libc_renameat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_renameat)

TEXT ·libc_symlinkat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_symlinkat)

TEXT ·libc_unlinkat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_unlinkat)

TEXT ·libc_readlinkat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_readlinkat)

TEXT ·libc_mkdirat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_mkdirat)

TEXT ·libc_utimensat_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_utimensat)

TEXT ·libc_getentropy_trampoline(SB),NOSPLIT,$0-0
	LIBCALL(libc_getentropy)
