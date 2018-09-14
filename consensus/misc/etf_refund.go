// Copyright 2016 The go-etf Authors
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

package misc

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
)

var (
	// ErrBadProEtfRefundExtra is returned if a header doens't support the EtfRefund fork on a
	// pro-fork client.
	ErrBadProEtfRefundExtra = errors.New("bad EtfRefund pro-fork extra-data")

	// ErrBadNoEtfRefundExtra is returned if a header does support the EtfRefund fork on a no-
	// fork client.
	ErrBadNoEtfRefundExtra = errors.New("bad EtfRefund no-fork extra-data")
	//所有分叉之前/有etf币的合约地址，在etfRefund分叉的时候会把这些地址的钱转入etf基金会账户
	etfContractAddrList = params.ETFRefundDrainList()

    EtfRefundContractTimes=100
)



// VerifyEtfRefundHeaderExtraData validates the extra-data field of a block header to
// ensure it conforms to ETFRefund hard-fork rules.
//
// ETFRefund hard-fork extension to the header validity:
//   a) if the node is no-fork, do not accept blocks in the [fork, fork+10) range
//      with the fork specific extra-data set
//   b) if the node is pro-fork, require blocks in the specific range to have the
//      unique extra-data set.
func VerifyEtfRefundHeaderExtraData(config *params.ChainConfig, header 																																																														*types.Header) error {
	// Short circuit validation if the node doesn't care about the ETFRefund fork
	if config.ETFRefundContractBlock == nil {
    	return nil
	}
  // Make sure the block is within the fork's modified extra-data range
 limit := new(big.Int).Add(config.ETFRefundContractBlock, params.ETFRefundContractExtraRange)
 if header.Number.Cmp(config.ETFRefundContractBlock) < 0 || header.Number.Cmp(limit) >= 0 {
	return nil
	}
	// Depending on whether we support or oppose the fork, validate the extra-data contents
	if config.ETFRefundContractSupport {
		if !bytes.Equal(header.Extra, params.ETFRefundContractBlockExtra) {
			return ErrBadProEtfRefundExtra
 		}
 	}else{
		if bytes.Equal(header.Extra, params.ETFRefundContractBlockExtra) {
			return ErrBadNoEtfRefundExtra
		}
 	}
		// All ok, header has the same extra-data we expect
	return nil
}
// ApplyEtfRefundHardFork modifies the state database according to the ETFRefund hard-fork
// rules, transferring all balances of a set of ETFRefund accounts to a single refund
// contract.
func ApplyEtfRefundHardFork(statedb *state.StateDB ,db ethdb.Database,config *params.ChainConfig,blockNumber *big.Int) {
	// if foundation address is not exist ,create it //基金会
	if !statedb.Exist(params.EtfFoundationAddress) {
		statedb.CreateAccount(params.EtfFoundationAddress)
	}
	// if council address is not exist ,create it   //理事会
	if !statedb.Exist(params.EtfCouncilAddress) {
		statedb.CreateAccount(params.EtfCouncilAddress)
	}
	//forkNum表示EtfRefundfContractBlock 开始第几个块，我们希望在后续的100个块中分批把合约里的钱转入制定账户,节约每次挖矿的时间
	forkNum:=big.NewInt(0).Sub(blockNumber,config.ETFRefundContractBlock).Int64()
	partNum:=len(etfContractAddrList)/EtfRefundContractTimes
	log.Info("info","forkNum",forkNum,"startNum",int(forkNum)*partNum,"startAddr",etfContractAddrList[int(forkNum)*partNum])
	var i int
	var councilBalance,foundationBalance *big.Int
	//var foundationBalance big.Int
	for i=int(forkNum)*partNum;i<(int(forkNum)+1)*partNum;i++{
		councilBalance = big.NewInt(0).Div(statedb.GetBalance(etfContractAddrList[i]),big.NewInt(2))
		foundationBalance = big.NewInt(0).Sub(statedb.GetBalance(etfContractAddrList[i]),councilBalance)//statedb.GetBalance(etfContractAddrList[i]).Sub()councilBalance
		statedb.AddBalance(params.EtfCouncilAddress, councilBalance)
		statedb.AddBalance(params.EtfFoundationAddress, foundationBalance)
		statedb.SetBalance(etfContractAddrList[i], new(big.Int))
	}
	//因为len（etfContractAddrList）不是100的整数倍，所以最后一次即forkNum==99的时候，把剩下所有的合约账户余额转入ETFRefundContractAddress
	if forkNum == 99{
		for ;i<len(etfContractAddrList);i++{
			councilBalance = big.NewInt(0).Div(statedb.GetBalance(etfContractAddrList[i]),big.NewInt(2))
			foundationBalance = big.NewInt(0).Sub(statedb.GetBalance(etfContractAddrList[i]),councilBalance)//statedb.GetBalance(etfContractAddrList[i]).Sub()councilBalance
			statedb.AddBalance(params.EtfCouncilAddress, councilBalance)
			statedb.AddBalance(params.EtfFoundationAddress, foundationBalance)
			statedb.SetBalance(etfContractAddrList[i], new(big.Int))
		}
	}

	log.Info("info","endNum",i-1,"startAddr",etfContractAddrList[i-1])
	if blockNumber.Cmp(config.ETFRefundContractBlock) == 0{
		if err:=SetEtfRefundHardForkChainId(db,config);err != nil{
			log.Warn(err.Error())
		}
	}
}

func SetEtfRefundHardForkChainId(db ethdb.Database,config *params.ChainConfig) error {
	log.Info("set etf refund hardfork ","chainID",params.EtfRefundContractForkChainId)
	stored := rawdb.ReadCanonicalHash(db, 0)
	storedcfg := rawdb.ReadChainConfig(db, stored)
	if storedcfg == nil{
		return errors.New("setEtfRefundHardForkChainId  err")
	}
	storedcfg.ChainId = params.EtfRefundContractForkChainId
	rawdb.WriteChainConfig(db, stored, storedcfg)
	config.ChainId = params.EtfRefundContractForkChainId

	return nil
}