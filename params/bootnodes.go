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
	"enode://5b402786639ad4744886dff3829599fba920552fa0d99a9fab4be6f2d6e5364a2f51fe92d24d33c2f630467abdd9131c0fa73e16ada7e2fe130ec691dfac3b0f@3.77.213.104:13931",
	"enode://7bd9530f3c65cc2c4fcbede569b07d226ff2c22bfdd8b4dcb22fa9ece9ba35d16435aeb17ed96460debd27c4a31921056a8613aa016f2a6f0cb5ab5fe413e69e@3.124.171.228:13931",
	"enode://914200237e58de0c4b830d42466819ed4434f7a3b1de9f97b62ed03a2420ae7b9a79497c20f87b010dd093c6c1cd9a0f82ce509fbfad6bc7f1b2081723322cb2@3.77.178.42:13931",
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
