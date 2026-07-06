// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 && !haiku

package runtime

// controlWord64 is the x87 FPU control word installed by asminit. Precision
// control is set to double (53-bit) to match SSE2 and the results produced on
// other platforms. See vlrt.go for the bit layout.
var controlWord64 uint16 = 0x3f + 2<<8 + 0<<10
