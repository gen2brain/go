// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo && !netgo

package net

/*
#include <sys/types.h>
#include <sys/socket.h>

#include <netdb.h>
*/
import "C"

import (
	"syscall"
	"unsafe"
)

const cgoAddrInfoFlags = (C.AI_CANONNAME | C.AI_V4MAPPED | C.AI_ALL) & C.AI_MASK

func cgoNameinfoPTR(b []byte, sa *C.struct_sockaddr, salen C.socklen_t) (int, error) {
	gerrno, err := C.getnameinfo(sa, salen, (*C.char)(unsafe.Pointer(&b[0])), C.socklen_t(len(b)), nil, 0, C.NI_NAMEREQD)
	return int(gerrno), err
}

func cgoSockaddrInet4(ip IP) *C.struct_sockaddr {
	sa := syscall.RawSockaddrInet4{Len: syscall.SizeofSockaddrInet4, Family: syscall.AF_INET}
	copy(sa.Addr[:], ip)
	return (*C.struct_sockaddr)(unsafe.Pointer(&sa))
}

func cgoSockaddrInet6(ip IP, zone int) *C.struct_sockaddr {
	sa := syscall.RawSockaddrInet6{Len: syscall.SizeofSockaddrInet6, Family: syscall.AF_INET6, Scope_id: uint32(zone)}
	copy(sa.Addr[:], ip)
	return (*C.struct_sockaddr)(unsafe.Pointer(&sa))
}
