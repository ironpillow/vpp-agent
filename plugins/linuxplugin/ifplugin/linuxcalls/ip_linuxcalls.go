// +build !windows,!darwin

// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linuxcalls

import (
	"bytes"
	"net"

	"time"

	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/cn-infra/logging/measure"
	"github.com/vishvananda/netlink"
)

// AddInterfaceIP calls AddrAdd Netlink API.
func AddInterfaceIP(log logging.Logger, ifName string, addr *net.IPNet, timeLog measure.StopWatchEntry) error {
	start := time.Now()
	defer func() {
		if timeLog != nil {
			timeLog.LogTimeEntry(time.Since(start))
		}
	}()

	link, err := netlink.LinkByName(ifName)
	if err != nil {
		return err
	}

	exAddrList, err := netlink.AddrList(link, netlink.FAMILY_ALL)
	if err != nil {
		return err
	}

	// The check is because of link local addresses which sometimes cannot be reassigned
	var ipExists bool
	for _, exAddr := range exAddrList {
		if bytes.Compare(exAddr.IP, addr.IP) == 0 {
			ipExists = true
			break
		}
	}
	if ipExists {
		log.Debugf("Cannot assign %v to interface %v, IP already exists", addr.IP.String(), ifName)
		return nil
	} else {
		return netlink.AddrAdd(link, &netlink.Addr{IPNet: addr})
	}
}

// DelInterfaceIP calls AddrDel Netlink API.
func DelInterfaceIP(ifName string, addr *net.IPNet, timeLog measure.StopWatchEntry) error {
	start := time.Now()
	defer func() {
		if timeLog != nil {
			timeLog.LogTimeEntry(time.Since(start))
		}
	}()

	link, err := netlink.LinkByName(ifName)
	if err != nil {
		return err
	}
	address := &netlink.Addr{IPNet: addr}
	return netlink.AddrDel(link, address)
}

// SetInterfaceMTU calls LinkSetMTU Netlink API.
func SetInterfaceMTU(ifName string, mtu int, timeLog measure.StopWatchEntry) error {
	start := time.Now()
	defer func() {
		if timeLog != nil {
			timeLog.LogTimeEntry(time.Since(start))
		}
	}()

	link, err := netlink.LinkByName(ifName)
	if err != nil {
		return err
	}
	return netlink.LinkSetMTU(link, mtu)
}
