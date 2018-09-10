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
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61911", // china
	"enode://2f10c57a897d9c417ffd5d69ac790acfacd3f214500dc7f24ecd5fe343f4f3e1e3e75bd8818fecc8e64759af27faf14089e0a2bec4e3901941ab6623e5c2557b@39.107.26.73:61910",  // china
	"enode://74d02a84ba4abf3281f4bfa0bb86bd8447155b7a91eda16c76a955dd8ce35df64f96762e5b85567e58d383c6f7042d614d626b36f5bfac16726f1975ba28253f@39.107.47.78:61910",  // china
	"enode://64e4dd3f53e5c01e7e5ef92aa79746cdc49d310d3bca18017caa81aca8fc68cad696e1b2f95f8332e0dca10cd9d3a065a5362b079182c62ca801c81a6a5db633@39.107.26.140:61910", // china
	"enode://7d7c0d58893c1b523b38f1907f0f40f872faa1ae079fbcfcf0867a56132f4facd6d21eccbe19bf7fa3f0c692f00d2ded958fd041ceb1c399336de2b4175f946b@47.52.251.187:61910", // HK
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://c6255fd229b7837bc6bc1ee3e99d0fa00a6fef33e7da377357d1fa738262328b2740c343b07a10f7a9be0ca0d293f8d053951c91fa287c90ac321e5e25bc84b8@47.52.114.83:30303", // HK
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61910", // china
	"enode://7d7c0d58893c1b523b38f1907f0f40f872faa1ae079fbcfcf0867a56132f4facd6d21eccbe19bf7fa3f0c692f00d2ded958fd041ceb1c399336de2b4175f946b@47.52.251.187:61910", // HK
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
//var DiscoveryV5Bootnodes = []string{
//	"enode://91115bff3a0391a45b2003f7f8ab3c767ce24ec6ff5a2c9a51264b3421dbd9762e78b0b99d7ee03a65bad1ecbb1f14e37fefa3e3a7448bf7505ac8dec4c809ad@47.93.251.181:61911", // china
//	"enode://7d7c0d58893c1b523b38f1907f0f40f872faa1ae079fbcfcf0867a56132f4facd6d21eccbe19bf7fa3f0c692f00d2ded958fd041ceb1c399336de2b4175f946b@47.52.251.187:61910", // HK
//}
var DiscoveryV5Bootnodes = MainnetBootnodes
