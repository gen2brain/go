// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"errors"
	"syscall"
)

func executable() (string, error) {
	p, err := syscall.FindImagePath()
	if err != nil {
		return "", errors.New("os: find_path failed: " + err.Error())
	}
	return p, nil
}
