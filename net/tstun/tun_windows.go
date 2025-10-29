// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package tstun

import (
	"os"

	"github.com/tailscale/wireguard-go/tun"
	"golang.org/x/sys/windows"
	"golang.zx2c4.com/wireguard/windows/tunnel/winipcfg"
)

func init() {
	tun.WintunTunnelType = "Tailscale"

	// Optionally set a static GUID from environment for single-adapter uses.
	// If you want multiple independent adapters, do NOT set this env var.
	if guidStr := os.Getenv("WINTUN_STATIC_GUID"); guidStr != "" {
		guid, err := windows.GUIDFromString(guidStr)
		if err != nil {
			panic(err)
		}
		tun.WintunStaticRequestedGUID = &guid
	} else {
		// Leave tun.WintunStaticRequestedGUID nil so each process gets its own adapter.
		tun.WintunStaticRequestedGUID = nil
	}
}

func interfaceName(dev tun.Device) (string, error) {
	guid, err := winipcfg.LUID(dev.(*tun.NativeTun).LUID()).GUID()
	if err != nil {
		return "", err
	}
	return guid.String(), nil
}
