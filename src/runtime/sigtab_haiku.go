// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// 1..32 are the standard signals; 33..40 are realtime SIGRTMIN..SIGRTMAX.
var sigtable = [...]sigTabT{
	0:        {0, "SIGNONE: no trap"},
	_SIGHUP:  {_SigNotify + _SigKill, "SIGHUP: terminal line hangup"},
	_SIGINT:  {_SigNotify + _SigKill, "SIGINT: interrupt"},
	_SIGQUIT: {_SigNotify + _SigThrow, "SIGQUIT: quit"},
	_SIGILL:  {_SigThrow + _SigUnblock, "SIGILL: illegal instruction"},
	// SIGCHLD must stay at SIG_DFL: installing any handler tells Haiku to
	// auto-reap child processes, breaking waitpid() in the parent.
	_SIGCHLD:   {0, "SIGCHLD: child status has changed"},
	_SIGABRT:   {_SigNotify + _SigThrow, "SIGABRT: abort"},
	_SIGPIPE:   {_SigNotify, "SIGPIPE: write to broken pipe"},
	_SIGFPE:    {_SigPanic + _SigUnblock, "SIGFPE: floating-point exception"},
	_SIGKILL:   {0, "SIGKILL: kill"},
	_SIGSTOP:   {0, "SIGSTOP: stop"},
	_SIGSEGV:   {_SigPanic + _SigUnblock, "SIGSEGV: segmentation violation"},
	_SIGCONT:   {_SigNotify + _SigDefault, "SIGCONT: continue"},
	_SIGTSTP:   {_SigNotify + _SigDefault, "SIGTSTP: keyboard stop"},
	_SIGALRM:   {_SigNotify, "SIGALRM: alarm clock"},
	_SIGTERM:   {_SigNotify + _SigKill, "SIGTERM: termination"},
	_SIGTTIN:   {_SigNotify + _SigDefault, "SIGTTIN: background read from tty"},
	_SIGTTOU:   {_SigNotify + _SigDefault, "SIGTTOU: background write to tty"},
	_SIGUSR1:   {_SigNotify, "SIGUSR1: user-defined signal 1"},
	_SIGUSR2:   {_SigNotify, "SIGUSR2: user-defined signal 2"},
	_SIGWINCH:  {_SigNotify, "SIGWINCH: window size change"},
	21:         {0, "SIGKILLTHR: kill thread"},
	_SIGTRAP:   {_SigThrow + _SigUnblock, "SIGTRAP: trace trap"},
	_SIGPOLL:   {_SigNotify, "SIGPOLL: pollable event"},
	_SIGPROF:   {_SigNotify + _SigUnblock, "SIGPROF: profiling alarm clock"},
	_SIGSYS:    {_SigThrow, "SIGSYS: bad system call"},
	_SIGURG:    {_SigNotify, "SIGURG: urgent condition on socket"},
	_SIGVTALRM: {_SigNotify, "SIGVTALRM: virtual alarm clock"},
	_SIGXCPU:   {_SigNotify, "SIGXCPU: cpu limit exceeded"},
	_SIGXFSZ:   {_SigNotify, "SIGXFSZ: file size limit exceeded"},
	_SIGBUS:    {_SigPanic + _SigUnblock, "SIGBUS: bus error"},
	31:         {_SigNotify, "SIGRESERVED1: reserved 1"},
	32:         {_SigNotify, "SIGRESERVED2: reserved 2"},
	33:         {_SigNotify, "signal 33"},
	34:         {_SigNotify, "signal 34"},
	35:         {_SigNotify, "signal 35"},
	36:         {_SigNotify, "signal 36"},
	37:         {_SigNotify, "signal 37"},
	38:         {_SigNotify, "signal 38"},
	39:         {_SigNotify, "signal 39"},
	40:         {_SigNotify, "signal 40"},
}
