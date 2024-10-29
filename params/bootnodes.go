// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://914200237e58de0c4b830d42466819ed4434f7a3b1de9f97b62ed03a2420ae7b9a79497c20f87b010dd093c6c1cd9a0f82ce509fbfad6bc7f1b2081723322cb2@3.77.178.42:13931",
	"enode://77ead0d548d61cfedafd049d877c52a229948f6bc250e6f08f1a9ca2de2bea2e7a34528f4a21d80e6117c4295c5db0a3fc9c6aac20e7c23ced39e9c4203ae7c7@3.74.57.216:13931",
	"enode://47a0c54cd2441d0a2028e621d87558687bc1990bbeb5cc236e99d2ac54a01f84edd0c400add25991ba614cd08ee711fb46d192a22b52e96f214b94e0b13b31fe@54.238.213.203:13931",
	"enode://5283e84c58d52187cc2b49642766f98d9467a707412d0cf6a46390a3e9d3fef009bbae9e591edf4ae036f4da0914ddb5abd6a7c902fa0f4465a4a60e09645f56@44.201.236.113:13931",
	"enode://a41b3f51704e49bd2fc94db866a5e77b8c90154e5ea7dfd74d6c12794eb65ed145462b5b4b6b12fda715aaba121b53889adbd8bb0be4cb154697b221fa650c69@18.163.3.21:13931",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://d46374654d198bfff22161c96d8fce12bfbbb7c47dc5708751b55d1d12ffa0d4ebcc0b54019567f0ef51c07e727441b6aeb7d2552c56b2f967202cccf37e42fd@175.41.153.124:13931",
	"enode://1fa645a25b554a2f6bfe9c894a43d04679f0a64bf38e7d77328b87c37a1b76a682ce31378d3212e60f682d307eddee886c91c7861f1471f21d018210087a43ae@54.255.14.223:13931",
}

var V5Bootnodes = []string{}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
