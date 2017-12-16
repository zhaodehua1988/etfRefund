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
	"enode://97bb9dfad70ce50925c133c1f50c2cc5e1c3e51c028843897a31840f8984466cb4abebd3c368b24243d04face17d5ed0a62f631bc8a431b095be611545900a68@39.107.61.127:61910", // china
	"enode://63e56f57fa0925f40b83bdefe131e90b0f220b05b664bb17e37717555c99fa22b3285552d4d7727445fed3f826d9b91d2ffc7914b38d53475a44fffc328318d6@39.107.61.127:61911", // china2
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://97bb9dfad70ce50925c133c1f50c2cc5e1c3e51c028843897a31840f8984466cb4abebd3c368b24243d04face17d5ed0a62f631bc8a431b095be611545900a68@39.107.61.127:61910", // china
	"enode://63e56f57fa0925f40b83bdefe131e90b0f220b05b664bb17e37717555c99fa22b3285552d4d7727445fed3f826d9b91d2ffc7914b38d53475a44fffc328318d6@39.107.61.127:61911", // china2
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
