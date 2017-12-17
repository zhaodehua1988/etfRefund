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
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61910", // china
	"enode://2c6573c62628233bbe500e3c8f5b2f91b80e1cdfb6a7b3e18732a59f5a560d44c1ebef0a48b67bfd99021b0af23e21dffd5c6da7d98678a7958ec077164a715b@47.52.251.187:61910", // HK
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61910", // china
	"enode://2c6573c62628233bbe500e3c8f5b2f91b80e1cdfb6a7b3e18732a59f5a560d44c1ebef0a48b67bfd99021b0af23e21dffd5c6da7d98678a7958ec077164a715b@47.52.251.187:61910", // HK
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61910", // china
	"enode://2c6573c62628233bbe500e3c8f5b2f91b80e1cdfb6a7b3e18732a59f5a560d44c1ebef0a48b67bfd99021b0af23e21dffd5c6da7d98678a7958ec077164a715b@47.52.251.187:61910", // HK
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61910", // china
	"enode://2c6573c62628233bbe500e3c8f5b2f91b80e1cdfb6a7b3e18732a59f5a560d44c1ebef0a48b67bfd99021b0af23e21dffd5c6da7d98678a7958ec077164a715b@47.52.251.187:61910", // HK
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61910", // china
	"enode://2c6573c62628233bbe500e3c8f5b2f91b80e1cdfb6a7b3e18732a59f5a560d44c1ebef0a48b67bfd99021b0af23e21dffd5c6da7d98678a7958ec077164a715b@47.52.251.187:61910", // HK
}
