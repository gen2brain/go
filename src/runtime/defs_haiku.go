// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

/*
Input to cgo -godefs
GOARCH=amd64 go tool cgo -godefs defs_haiku.go > defs_haiku_amd64.go
*/

package runtime

/*
#include <sys/types.h>
#include <sys/time.h>
#include <signal.h>
#include <errno.h>
#include <fcntl.h>
#include <sys/mman.h>
#include <semaphore.h>
#include <pthread.h>
#include <unistd.h>

#include <OS.h>
*/
import "C"

const (
	_EPERM     = C.EPERM
	_ENOENT    = C.ENOENT
	_EINTR     = C.EINTR
	_EAGAIN    = C.EAGAIN
	_ENOMEM    = C.ENOMEM
	_EACCES    = C.EACCES
	_EFAULT    = C.EFAULT
	_EINVAL    = C.EINVAL
	_ETIMEDOUT = C.ETIMEDOUT

	_PROT_NONE  = C.PROT_NONE
	_PROT_READ  = C.PROT_READ
	_PROT_WRITE = C.PROT_WRITE
	_PROT_EXEC  = C.PROT_EXEC

	_MAP_ANON      = C.MAP_ANON
	_MAP_PRIVATE   = C.MAP_PRIVATE
	_MAP_FIXED     = C.MAP_FIXED
	_MADV_DONTNEED = C.MADV_DONTNEED
	_MADV_FREE     = C.MADV_FREE

	_SIGHUP    = C.SIGHUP
	_SIGINT    = C.SIGINT
	_SIGQUIT   = C.SIGQUIT
	_SIGILL    = C.SIGILL
	_SIGTRAP   = C.SIGTRAP
	_SIGABRT   = C.SIGABRT
	_SIGBUS    = C.SIGBUS
	_SIGFPE    = C.SIGFPE
	_SIGKILL   = C.SIGKILL
	_SIGUSR1   = C.SIGUSR1
	_SIGSEGV   = C.SIGSEGV
	_SIGUSR2   = C.SIGUSR2
	_SIGPIPE   = C.SIGPIPE
	_SIGALRM   = C.SIGALRM
	_SIGCHLD   = C.SIGCHLD
	_SIGCONT   = C.SIGCONT
	_SIGSTOP   = C.SIGSTOP
	_SIGTSTP   = C.SIGTSTP
	_SIGTTIN   = C.SIGTTIN
	_SIGTTOU   = C.SIGTTOU
	_SIGURG    = C.SIGURG
	_SIGXCPU   = C.SIGXCPU
	_SIGXFSZ   = C.SIGXFSZ
	_SIGVTALRM = C.SIGVTALRM
	_SIGPROF   = C.SIGPROF
	_SIGWINCH  = C.SIGWINCH
	_SIGSYS    = C.SIGSYS
	_SIGTERM   = C.SIGTERM

	_FPE_INTDIV = C.FPE_INTDIV
	_FPE_INTOVF = C.FPE_INTOVF
	_FPE_FLTDIV = C.FPE_FLTDIV
	_FPE_FLTOVF = C.FPE_FLTOVF
	_FPE_FLTUND = C.FPE_FLTUND
	_FPE_FLTRES = C.FPE_FLTRES
	_FPE_FLTINV = C.FPE_FLTINV
	_FPE_FLTSUB = C.FPE_FLTSUB

	_BUS_ADRALN = C.BUS_ADRALN
	_BUS_ADRERR = C.BUS_ADRERR
	_BUS_OBJERR = C.BUS_OBJERR

	_SEGV_MAPERR = C.SEGV_MAPERR
	_SEGV_ACCERR = C.SEGV_ACCERR

	_O_RDONLY   = C.O_RDONLY
	_O_WRONLY   = C.O_WRONLY
	_O_NONBLOCK = C.O_NONBLOCK
	_O_CREAT    = C.O_CREAT
	_O_TRUNC    = C.O_TRUNC
	_O_CLOEXEC  = C.O_CLOEXEC

	_SS_DISABLE  = C.SS_DISABLE
	_SI_USER     = C.SI_USER
	_SIG_BLOCK   = C.SIG_BLOCK
	_SIG_UNBLOCK = C.SIG_UNBLOCK
	_SIG_SETMASK = C.SIG_SETMASK

	_SA_SIGINFO = C.SA_SIGINFO
	_SA_RESTART = C.SA_RESTART
	_SA_ONSTACK = C.SA_ONSTACK

	_PTHREAD_CREATE_DETACHED = C.PTHREAD_CREATE_DETACHED

	__SC_PAGE_SIZE        = C._SC_PAGE_SIZE
	__SC_NPROCESSORS_ONLN = C._SC_NPROCESSORS_ONLN

	_F_SETFL = C.F_SETFL
	_F_GETFD = C.F_GETFD
	_F_GETFL = C.F_GETFL
	_F_SETFD = C.F_SETFD

	_FD_CLOEXEC = C.FD_CLOEXEC
)

type sigset C.sigset_t
type siginfo C.siginfo_t
type timespec C.struct_timespec
type timeval C.struct_timeval
type itimerval C.struct_itimerval

type stackt C.stack_t
type ucontext C.ucontext_t
type mcontext C.mcontext_t
type sigactiont C.struct_sigaction

type pthread C.pthread_t
type pthread_attr C.pthread_attr_t

type semt C.sem_t
