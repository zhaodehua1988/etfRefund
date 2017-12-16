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

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://77ee57b66ae3a512d59437743ce410229d2aca1266156780e7a94af5c4aec419efc5541a80de75f00a89e02051914cfff60e5034648916ebeb87113120cc7338@59.110.26.199:61910", // china
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:31910", // china2
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://77ee57b66ae3a512d59437743ce410229d2aca1266156780e7a94af5c4aec419efc5541a80de75f00a89e02051914cfff60e5034648916ebeb87113120cc7338@59.110.26.199:61910", // china
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:31910", // china2
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://97bb9dfad70ce50925c133c1f50c2cc5e1c3e51c028843897a31840f8984466cb4abebd3c368b24243d04face17d5ed0a62f631bc8a431b095be611545900a68@39.107.61.127:61910", // china
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	"enode://97bb9dfad70ce50925c133c1f50c2cc5e1c3e51c028843897a31840f8984466cb4abebd3c368b24243d04face17d5ed0a62f631bc8a431b095be611545900a68@39.107.61.127:61910", // china
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://97bb9dfad70ce50925c133c1f50c2cc5e1c3e51c028843897a31840f8984466cb4abebd3c368b24243d04face17d5ed0a62f631bc8a431b095be611545900a68@39.107.61.127:61910", // china
}
