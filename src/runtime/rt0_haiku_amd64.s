// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// runtime_loader invokes the entry point with argc/argv/env passed in
// DI/SI/DX (SysV C ABI), so jump straight into rt0_go which expects
// the same — _rt0_amd64 would try to read argc from 0(SP).
TEXT _rt0_amd64_haiku(SB),NOSPLIT,$-8
	JMP	runtime·rt0_go(SB)

TEXT _rt0_amd64_haiku_lib(SB),NOSPLIT,$0
	JMP	_rt0_amd64_lib(SB)
