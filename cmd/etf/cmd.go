package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/core/state"
	"os"
	"math/big"
	"github.com/ethereum/go-ethereum/rlp"
	"bytes"
)



var etfRefundFileHead = `// Copyright 2018 The go-etf Authors
// This file is part of the go-etf library.
//
// The go-etf library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-etf library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-etf library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)
// ETFReclaimBalanceBlockExtra is the block header extra-data field to set for the etfRefund fork
// point and a number of consecutive blocks to allow fast/light syncers to correctly
// pick the side they want  ("etfRefund-hard-fork").
//var ETFRefundContractBlockExtra = common.FromHex("0x64616f2d686172642d666f726b")
var ETFRefundContractBlockExtra = common.FromHex("0x0102030405060708090a0b0c0d")
// ETFReclaimBalanceExtraRange is the number of consecutive blocks from the etfRefund fork point
// to override the extra-data in to prevent no-fork attacks.
var ETFRefundContractExtraRange = big.NewInt(10)

// ETFRefundContract is the address of the refund contract to send etfRefundContract balances to.
var ETFRefundContractAddress = common.HexToAddress("0x37215c4bcfeb37badb9b466c6c0ff2ef6d7238f9")//接收ETF分叉前智能合约余额的账户

//ETFForkBlock[4730660]'s stadetrie head hash
//var ETFForkBlockStateTrieHeadHash = common.HexToHash("0x800cb7aabf718fb2db4e6c910e053d9930f64c6ceb4877010b923eb937edf6b2") //ETF分叉节点[4730660]的statetrie 哈希

func ETFRefundDrainList() []common.Address {
	return []common.Address{
`
var (
		version = "etf cmd version 2.0"
		emptyCodeHash = crypto.Keccak256(nil)
		codeFile = "./contractAddr.txt"
		logFIle = "log.csv"
		logTitle = "合约地址,余额\r\n"
		logFockFile = "FockedContract.csv"
		key = common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9")
	)

func GetContractAddr(stateDBPath string ,stateRoot string) error {

	fmt.Println(version)
	fmt.Println("get contract Addr", stateDBPath,"stateRootHex",stateRoot)

	db, err := ethdb.NewLDBDatabase(stateDBPath, 1024, 512)
	if err != nil {
		utils.Fatalf("error opening db: %s", err.Error())
	}
	cacheDb := state.NewDatabase(db)
	//head := common.HexToHash("0x800cb7aabf718fb2db4e6c910e053d9930f64c6ceb4877010b923eb937edf6b2")
	head := common.HexToHash(stateRoot)
	root, err := cacheDb.OpenTrie(head)
	if err != nil {
		utils.Fatalf("error opening trie: %s", err.Error())
	}
	//把获取到的合约地址保存到本地./contractAddr.txt

	f, err := os.OpenFile(codeFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(etfRefundFileHead)
	if err != nil {
		panic(err)
	}
	f.Sync()
	//把分叉之前合约地址和余额保存为csv格式文件
	fLog ,err:=os.OpenFile(logFIle,os.O_RDWR | os.O_CREATE,0666)
	if err != nil{
		panic(err)
	}
	defer fLog.Close()

	_,err = fLog.WriteString(logTitle)
	if err != nil{
		panic(err)
	}

	//把分叉之后合约地址和余额保存为csv格式文件
	fFockLog ,err:=os.OpenFile(logFockFile,os.O_RDWR | os.O_CREATE,0666)
	if err != nil{
		panic(err)
	}
	defer fFockLog.Close()
	_,err = fFockLog.WriteString(logTitle)
	if err != nil{
		panic(err)
	}

	totalWei:=big.NewInt(0)
	i := root.NodeIterator(nil)

	for i.Next(true) {

		if i.Leaf() {

			hAddr := i.LeafKey()
			addr := root.GetKey(hAddr)

			var data state.Account
			if err := rlp.DecodeBytes(i.LeafBlob(), &data); err != nil {
				//t.Fatal("Failed to decode state object", "addrHash", addr, "err", err)
				return err
			}
			if (data.Balance.Cmp(big.NewInt(0)) > 0) && !bytes.Equal(data.CodeHash, emptyCodeHash) {
				totalWei.Add(totalWei,data.Balance)

				//判断是否为分叉之前的合约
				storageTrie,err := cacheDb.OpenStorageTrie(common.BytesToHash(hAddr), data.Root)
				if err != nil{
					fmt.Println("openstorageTry err !")
					continue
				}
				value:= common.Hash{}
				enc,err := storageTrie.TryGet(key[:])
				if err == nil { //err != nil 说明可以查到这个key，所以是分叉之后的，分叉之后的智能合约账户余额保留
					if len(enc) > 0 {
						_, content, _, err := rlp.Split(enc)
						if err != nil {

						}
						value.SetBytes(content)
					}
				}

				 if value == key {//把分叉后的合约地址和余额
					 addrWriteToLog := fmt.Sprintf("\"%s\",%s\r\n",common.BytesToAddress(addr).String(),data.Balance.String())
					 _,err =fFockLog.WriteString(addrWriteToLog)
					 if err != nil {
						 fmt.Println("write to log.csv err",error(err))
						 panic(err)
					 }
				 	continue
				 }

				//fmt.Println(n,":","address: ", common.BytesToAddress(addr).Hex(), "balance: ", data.Balance,"totalBalance",totalWei,"(ether",new(big.Int).Div(totalWei,big.NewInt(1000000000000000000)),")")
				addrWriteToFile := fmt.Sprintf("\n        common.HexToAddress(\"%s\"),", common.BytesToAddress(addr).String())
				_, err = f.WriteString(addrWriteToFile)
				if err != nil {
					fmt.Println("write to file err",error(err))
					panic(err)
				}
				addrWriteToLog := fmt.Sprintf("\"%s\",%s\r\n",common.BytesToAddress(addr).String(),data.Balance.String())
				_,err =fLog.WriteString(addrWriteToLog)
				if err != nil {
					fmt.Println("write to log.csv err",error(err))
					panic(err)
				}
			}
		}
	}
	_, err = f.WriteString("\n    }\n}\n")
	if err != nil {
		panic(err)
	}

	fmt.Println("----------------end of file --------------------")

	return nil
}