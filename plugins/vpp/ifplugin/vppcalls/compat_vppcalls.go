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

package vppcalls

import (
	govppapi "git.fd.io/govpp.git/api"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/af_packet"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/bfd"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/interfaces"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/ip"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/memif"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/nat"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/stn"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/tap"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/tapv2"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/vxlan"
)

func (handler *ifVppHandler) CheckMsgCompatibilityForInterface() error {
	msgs := []govppapi.Message{
		&memif.MemifCreate{},
		&memif.MemifCreateReply{},
		&memif.MemifDelete{},
		&memif.MemifDeleteReply{},
		&memif.MemifDump{},
		&memif.MemifDetails{},
		&memif.MemifSocketFilenameDump{},
		&memif.MemifSocketFilenameDetails{},

		&interfaces.CreateLoopback{},
		&interfaces.CreateLoopbackReply{},

		&vxlan.VxlanAddDelTunnel{},
		&vxlan.VxlanAddDelTunnelReply{},
		&vxlan.VxlanTunnelDump{},
		&vxlan.VxlanTunnelDetails{},

		&af_packet.AfPacketCreate{},
		&af_packet.AfPacketCreateReply{},
		&af_packet.AfPacketDelete{},
		&af_packet.AfPacketDeleteReply{},

		&tap.TapConnect{},
		&tap.TapConnectReply{},
		&tap.TapDelete{},
		&tap.TapDeleteReply{},
		&tap.SwInterfaceTapDump{},
		&tap.SwInterfaceTapDetails{},

		&tapv2.TapCreateV2{},
		&tapv2.TapCreateV2Reply{},
		&tapv2.TapDeleteV2{},
		&tapv2.TapDeleteV2Reply{},

		&interfaces.SwInterfaceDump{},
		&interfaces.SwInterfaceDetails{},
		&interfaces.SwInterfaceEvent{},
		&interfaces.SwInterfaceSetFlags{},
		&interfaces.SwInterfaceSetFlagsReply{},
		&interfaces.SwInterfaceAddDelAddress{},
		&interfaces.SwInterfaceAddDelAddressReply{},
		&interfaces.SwInterfaceSetMacAddress{},
		&interfaces.SwInterfaceSetMacAddressReply{},
		&interfaces.SwInterfaceSetTable{},
		&interfaces.SwInterfaceSetTableReply{},
		&interfaces.SwInterfaceGetTable{},
		&interfaces.SwInterfaceGetTableReply{},
		&interfaces.SwInterfaceSetUnnumbered{},
		&interfaces.SwInterfaceSetUnnumberedReply{},
		&interfaces.SwInterfaceTagAddDel{},
		&interfaces.SwInterfaceTagAddDelReply{},
		&interfaces.SwInterfaceSetMtu{},
		&interfaces.SwInterfaceSetMtuReply{},
		&interfaces.HwInterfaceSetMtu{},
		&interfaces.HwInterfaceSetMtuReply{},

		&ip.IPAddressDump{},
		&ip.IPAddressDetails{},
		&ip.IPFibDump{},
		&ip.IPFibDetails{},
		&ip.IPTableAddDel{},
		&ip.IPTableAddDelReply{},
		&ip.IPContainerProxyAddDel{},
		&ip.IPContainerProxyAddDelReply{},
	}
	return handler.callsChannel.CheckMessageCompatibility(msgs...)
}

func (handler *bfdVppHandler) CheckMsgCompatibilityForBfd() error {
	msgs := []govppapi.Message{
		&bfd.BfdUDPAdd{},
		&bfd.BfdUDPAddReply{},
		&bfd.BfdUDPMod{},
		&bfd.BfdUDPModReply{},
		&bfd.BfdUDPDel{},
		&bfd.BfdUDPDelReply{},
		&bfd.BfdAuthSetKey{},
		&bfd.BfdAuthSetKeyReply{},
		&bfd.BfdAuthDelKey{},
		&bfd.BfdAuthDelKeyReply{},
	}
	return handler.callsChannel.CheckMessageCompatibility(msgs...)
}

func (handler *natVppHandler) CheckMsgCompatibilityForNat() error {
	msgs := []govppapi.Message{
		&nat.Nat44AddDelAddressRange{},
		&nat.Nat44AddDelAddressRangeReply{},
		&nat.Nat44ForwardingEnableDisable{},
		&nat.Nat44ForwardingEnableDisableReply{},
		&nat.Nat44InterfaceAddDelFeature{},
		&nat.Nat44InterfaceAddDelFeatureReply{},
		&nat.Nat44AddDelStaticMapping{},
		&nat.Nat44AddDelStaticMappingReply{},
		&nat.Nat44AddDelLbStaticMapping{},
		&nat.Nat44AddDelLbStaticMappingReply{},
	}
	return handler.callsChannel.CheckMessageCompatibility(msgs...)
}

func (handler *stnVppHandler) CheckMsgCompatibilityForStn() error {
	msgs := []govppapi.Message{
		&stn.StnAddDelRule{},
		&stn.StnAddDelRuleReply{},
	}
	return handler.callsChannel.CheckMessageCompatibility(msgs...)
}
