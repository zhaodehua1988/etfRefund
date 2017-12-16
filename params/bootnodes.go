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
	"enode://9a91981f4a9762de87099213613a303449d56dcd867890f736024e3307058efb01317661fac89ccf664cfc96b949d9685ba4bf9a465a16782fdc0201cf2db828@59.110.26.199:61910", // china
	"enode://b0872f07be344a9351569b0c3fe9d7c28df02711969235eb152b4702c5fcb1c475b3e9956d358e82d17ade55917b8118fa19e319b7d16dcc71557ee4709edd6e@59.110.26.199:61911", // china
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://9a91981f4a9762de87099213613a303449d56dcd867890f736024e3307058efb01317661fac89ccf664cfc96b949d9685ba4bf9a465a16782fdc0201cf2db828@59.110.26.199:61910", // china
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://9a91981f4a9762de87099213613a303449d56dcd867890f736024e3307058efb01317661fac89ccf664cfc96b949d9685ba4bf9a465a16782fdc0201cf2db828@59.110.26.199:61910", // china
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	"enode://9a91981f4a9762de87099213613a303449d56dcd867890f736024e3307058efb01317661fac89ccf664cfc96b949d9685ba4bf9a465a16782fdc0201cf2db828@59.110.26.199:61910", // china
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://9a91981f4a9762de87099213613a303449d56dcd867890f736024e3307058efb01317661fac89ccf664cfc96b949d9685ba4bf9a465a16782fdc0201cf2db828@59.110.26.199:61910", // china
}
