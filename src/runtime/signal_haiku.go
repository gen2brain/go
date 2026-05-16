// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "internal/abi"

// Haiku's kernel resets sa_handler to SIG_DFL on its way out of certain
// user handlers regardless of sa_flags. The two helpers below re-install
// the handler so the next signal is still delivered to runtime code.

func rearmSighandlerOnEntry(sig uint32) {
	switch sig {
	case _SIGHUP, _SIGINT, _SIGQUIT, _SIGTERM, _SIGUSR1, _SIGUSR2,
		_SIGCHLD, _SIGWINCH, _SIGALRM, _SIGVTALRM, _SIGPIPE,
		_SIGURG, _SIGTSTP:
		setsig(sig, abi.FuncPCABIInternal(sighandler))
	}
}

func rearmFaultSignalAfterSigpanic(sig uint32) {
	switch sig {
	case _SIGSEGV, _SIGBUS, _SIGFPE, _SIGILL, _SIGTRAP:
		setsig(sig, abi.FuncPCABIInternal(sighandler))
	}
}
