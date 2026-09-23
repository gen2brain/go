// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Additional libroot symbols used by the runtime side of the exec_libc.go
// forkAndExecInChild path, mirroring runtime/syscall_solaris.go.

package runtime

func libc_chdir_trampoline()
func libc_chroot_trampoline()
func libc_dup2_trampoline()
func libc_execve_trampoline()
func libc_fork_trampoline()
func libc_ioctl_trampoline()
func libc_setgid_trampoline()
func libc_setgroups_trampoline()
func libc_setpgid_trampoline()
func libc_setrlimit_trampoline()
func libc_setsid_trampoline()
func libc_setuid_trampoline()
