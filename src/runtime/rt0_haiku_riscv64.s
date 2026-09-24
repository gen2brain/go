// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// runtime_loader passes argc and argv in A0 and A1, as rt0_go expects.
TEXT _rt0_riscv64_haiku(SB),NOSPLIT|NOFRAME,$0
	JMP	main(SB)

// When building with -buildmode=c-shared, this symbol is called when the
// shared library is loaded.
TEXT _rt0_riscv64_haiku_lib(SB),NOSPLIT|NOFRAME,$0
	JMP	_rt0_riscv64_lib(SB)

TEXT main(SB),NOSPLIT|NOFRAME,$0
	MOV	$runtime·rt0_go(SB), T0
	JALR	ZERO, T0
