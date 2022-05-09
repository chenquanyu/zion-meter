/*
 * Copyright (C) 2021 The Zion Authors
 * This file is part of The Zion library.
 *
 * The Zion is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The Zion is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The Zion.  If not, see <http://www.gnu.org/licenses/>.
 */

package sdk

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

var (
	testMaster      *Account
	testAmount      = big.NewInt(1000000000000000) //0.001 eth
	testChainID     = uint64(60801)
	testUrl         = "http://127.0.0.1:22000"
	testMainNodeKey = "4b0c9b9d685db17ac9f295cb12f9d7d2369f5bf524b3ce52ce424031cafda1ae"
)

func TestMain(m *testing.M) {
	sender, _ := NewSender(testUrl, testChainID)
	testMaster, _ = MasterAccount(sender, testMainNodeKey)
	os.Exit(m.Run())
}

// go test -v github.com/dylenfu/zion-meter/pkg/sdk -run TestTransfer
func TestTransfer(t *testing.T) {
	n := 100
	to := common.HexToAddress("0x67CDE763bD045B14898d8B044F8afC8695ae8608")

	balance, _ := testMaster.BalanceOf(to, nil)
	t.Logf("balance before transfer %s, nonce before transfer %d", balance.String(), testMaster.nonce)

	for i := 0; i < n; i++ {
		if _, err := testMaster.Transfer(to, testAmount); err != nil {
			t.Fatal(err)
		}
	}

	time.Sleep(15 * time.Second)
	balance, _ = testMaster.BalanceOf(to, nil)
	t.Logf("balance after transfer %s, nonce after transfer %d", balance.String(), testMaster.nonce)
}

// go test -v github.com/dylenfu/zion-meter/pkg/sdk -run TestStat
func TestData(t *testing.T) {
	n := 10000
	startTime := uint64(time.Now().Unix())

	acc, err := NewAccount()
	if err != nil {
		t.Fatal(err)
	}

	balance, _ := new(big.Int).SetString("1000000000000000000", 10)
	if _, err := testMaster.TransferWithConfirm(acc.Address(), balance); err != nil {
		t.Fatal(err)
	}

	contract, err := testMaster.DataDeploy()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("contract address %s, start time %d, nonce before testing %d", contract.Hex(), startTime, acc.nonce)

	sender, err := NewSender(testUrl, testChainID)
	acc.SetSender(sender)
	for i := 0; i < n; i++ {
		if _, _, err := acc.CostManyGas(contract); err != nil {
			t.Error(err)
		}
	}

	time.Sleep(15 * time.Second)

	total, err := testMaster.DataTxNum(contract)
	if err != nil {
		t.Fatal(err)
	}
	endTime := uint64(time.Now().Unix())
	t.Logf("end time %d, spent %d, nonce after testing %d, total tx number %d", endTime, endTime-startTime, acc.nonce, total)
}
func TestDataPack(t *testing.T){
	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		t.Fatal(err)
	}
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/9bca539684b6408d9dbcbb179e593eab")
	if err != nil {
		t.Fatal(err)
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	authAddress := crypto.PubkeyToAddress(*privateKey.Public().(*ecdsa.PublicKey))
	nonce, err := client.PendingNonceAt(context.Background(), authAddress)
	if err != nil {
		t.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		t.Fatal(err)
	}
	auth.From = authAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(0)) // in wei
	var gasLimit uint64 = 0
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice



	inputarg:= []byte("0xffffff")
	complexityarg:=big.NewInt(6)
	args,err:=arguments.Pack(inputarg,complexityarg)
	if err != nil {
		t.Fatal(err)
	}else{
		print(args)
	}

	payload, err := utils.PackMethod(ABI, "costManyGas",inputarg,complexityarg)

	if err != nil {
		t.Fatal(err)
	}else{
		print(payload)
	}

	p,err:=ABI.Pack("costManyGas",inputarg,complexityarg)
	if err != nil {
		t.Fatal(err)
	}else {
		print(p)
	}

	print(args,"--",payload,"--",p)
	datacontract:=common.HexToAddress("")
	tx:=types.NewTx(&types.LegacyTx{
		Nonce:auth.Nonce.Uint64(),
		GasPrice:auth.GasPrice,
		Gas:auth.GasLimit,
		To: &datacontract,
		Value:auth.Value,
		Data:payload,
	})
	signer := types.LatestSignerForChainID(chainId)
	tx, err = types.SignTx(tx, signer, privateKey)
	if err != nil {
		t.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Fatal(err)
	}

	println(tx.Hash().Hex())


}
func TestDataPackMethod(t *testing.T){
	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		t.Fatal(err)
	}
	authAddress := crypto.PubkeyToAddress(*privateKey.Public().(*ecdsa.PublicKey))
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/9bca539684b6408d9dbcbb179e593eab")
	if err != nil {
		t.Fatal(err)
	}
	url:="https://ropsten.infura.io/v3/9bca539684b6408d9dbcbb179e593eab"
	chainId, err := client.ChainID(context.Background())
	ss:=chainId.String()
	chainid,err:=strconv.Atoi(ss)

	sender,err:=NewSender(url,uint64(chainid))
	if err != nil {
		t.Fatal(err)
	}
	nonce, err := client.PendingNonceAt(context.Background(), authAddress)
	if err != nil {
		t.Fatal(err)
	}
	c:=&Account{
		pk:privateKey,
		address: authAddress,
		nonce:nonce,
		sender:sender,
	}
	datacontract:=common.HexToAddress("0x3021edb975013e5BF6Efa4fBC39b70B1453d9082")
	inputarg,err:=hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	if err != nil {
		t.Fatal(err)
	}
	complexityarg:=big.NewInt(8)
	payload, err :=utils.PackMethod(ABI, "costManyGas",inputarg,complexityarg)
	if err != nil {
		t.Fatal(err)
	}
	print(common.Bytes2Hex(payload))
	//p,err:=hex.DecodeString("0xef22dc29000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000003ffffff0000000000000000000000000000000000000000000000000000000000")
	txx, err := c.newSignedTx(datacontract, big.NewInt(0), payload)
	if err != nil {
		t.Fatal(err)
	}

	if err := c.SendTx(txx); err != nil {
		t.Fatal(err)
	}
	println(txx.Hash().Hex())
}
func Test(t *testing.T){
	b:=[]byte("0xfffff")
	//s:=string(b)
	bb:=common.Bytes2Hex(b)
	println(bb)

	s:="ffffff"
	d,err:=hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	bbb:=common.Bytes2Hex(d)
	println(d,"-----",bbb)

	inputarg:=d
	complexityarg:=big.NewInt(8)
	payload, err :=utils.PackMethod(ABI, "costManyGas",inputarg,complexityarg)
	println(common.Bytes2Hex(payload))
}

