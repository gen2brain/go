// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// Haiku stores network configuration under B_SYSTEM_SETTINGS_DIRECTORY
// rather than /etc.
func init() {
	hostsFilePath = "/boot/system/settings/network/hosts"
	resolvFilePath = "/boot/system/settings/network/resolv.conf"
	nssConfigPath = "/boot/system/settings/network/nsswitch.conf"
}
