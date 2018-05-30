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
	"enode://7d7c0d58893c1b523b38f1907f0f40f872faa1ae079fbcfcf0867a56132f4facd6d21eccbe19bf7fa3f0c692f00d2ded958fd041ceb1c399336de2b4175f946b@47.52.251.187:61910", // HK
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61910", // china
	"enode://7d7c0d58893c1b523b38f1907f0f40f872faa1ae079fbcfcf0867a56132f4facd6d21eccbe19bf7fa3f0c692f00d2ded958fd041ceb1c399336de2b4175f946b@47.52.251.187:61910", // HK
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61910", // china
	"enode://7d7c0d58893c1b523b38f1907f0f40f872faa1ae079fbcfcf0867a56132f4facd6d21eccbe19bf7fa3f0c692f00d2ded958fd041ceb1c399336de2b4175f946b@47.52.251.187:61910", // HK
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61910", // china
	"enode://7d7c0d58893c1b523b38f1907f0f40f872faa1ae079fbcfcf0867a56132f4facd6d21eccbe19bf7fa3f0c692f00d2ded958fd041ceb1c399336de2b4175f946b@47.52.251.187:61910", // HK
}
