// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build haiku && 386

package runtime

// Haiku's libm (roundf, lroundf and friends) relies on the x87 FPU running at
// 64-bit extended precision, the i386 platform default. Go generates SSE2 for
// its own float64 arithmetic, so the x87 precision only affects cgo C code;
// keep it at extended precision (bits 8-9 = 3) so the Haiku C runtime behaves
// the same inside a Go process as it does outside one. See vlrt.go for the bit
// layout.
var controlWord64 uint16 = 0x3f + 3<<8 + 0<<10
