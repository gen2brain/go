// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// haikuMiniterrnoChild is implemented in package runtime; after fork() in
// the child it re-resolves the per-thread errno location, which differs
// between threads on Haiku.
func haikuMiniterrnoChild()
