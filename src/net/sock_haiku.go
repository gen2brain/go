// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

func maxListenerBacklog() int {
	// SOMAXCONN is 32 on Haiku, which is too small in practice; the
	// kernel accepts larger values, so use the same default as Linux.
	return 4096
}
