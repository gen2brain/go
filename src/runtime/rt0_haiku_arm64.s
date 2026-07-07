// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// runtime_loader invokes the entry point with argc/argv/env passed in
// R0/R1/R2 (AAPCS64), which is exactly what rt0_go expects (R0 = argc,
// R1 = argv), so jump straight in rather than reading them off the stack.
TEXT _rt0_arm64_haiku(SB),NOSPLIT,$0
	JMP	runtime·rt0_go(SB)

// When building with -buildmode=c-shared, this symbol is called when the
// shared library is loaded.
TEXT _rt0_arm64_haiku_lib(SB),NOSPLIT,$0
	JMP	_rt0_arm64_lib(SB)
