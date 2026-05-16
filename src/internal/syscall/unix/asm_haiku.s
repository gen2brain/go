// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// System calls for amd64 on Haiku are implemented in
// runtime/syscall_haiku.go via the asmsyscall6 trampoline.

TEXT ·syscall6(SB),NOSPLIT,$0-88
	JMP	syscall·syscall6(SB)

TEXT ·rawSyscall6(SB),NOSPLIT,$0-88
	JMP	syscall·rawSyscall6(SB)
