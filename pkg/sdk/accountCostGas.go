package sdk

import (
	"encoding/hex"
	dataTest "github.com/dylenfu/zion-meter/pkg/go_abi/dataTest_abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"math/big"
)


func (c *Account) DataDeploy() (common.Address, error) {
	if c.sender == nil {
		return common.EmptyAddress, ErrNoSender
	}

	auth := c.makeDeployAuth()
	addr, tx, _, err := dataTest.DeployDataTest(auth, c.sender.client)
	if err != nil {
		return common.EmptyAddress, err
	}
	if err := c.waitTransaction(tx.Hash()); err != nil {
		return common.EmptyAddress, err
	}
	return addr, nil
}

func (c *Account) DataReset(contract common.Address, startTime uint64) (common.Hash, error) {
	if c.sender == nil {
		return common.EmptyHash, ErrNoSender
	}
	st, err := dataTest.NewDataTest(contract, c.sender.client)
	if err != nil {
		return common.EmptyHash, err
	}
	auth := c.makeAuth()
	auth.GasLimit = 500000
	tx, err := st.Reset(auth, startTime)
	if err != nil {
		return common.EmptyHash, err
	}
	if err := c.waitTransaction(tx.Hash()); err != nil {
		return common.EmptyHash, err
	}
	return tx.Hash(), nil
}


var (
	input, _  = abi.NewType("bytes", "", nil)
	complexity, _ = abi.NewType("uint256", "", nil)
	arguments  = abi.Arguments{
		{Type: input, Name: "input"},
		{Type: complexity, Name: "complexity"},
	}
)

func (c *Account) CostManyGas(contract common.Address) (common.Hash, uint64, error) {
	if c.sender == nil {
		return common.EmptyHash, c.nonce, ErrNoSender
	}

	//auth := c.makeAuth()
	//nonce, err := c.sender.client.NonceAt(context.Background(), c.address, nil)
	//if err != nil {
	//	return common.EmptyHash, nonce, err
	//}
	//auth.Nonce = new(big.Int).SetUint64(nonce)
	//
	//st, err := stat.NewStat(contract, c.sender.client)
	//if err != nil {
	//	return common.EmptyHash, auth.Nonce.Uint64(), err
	//}
	//if tx, err := st.Add(auth); err != nil {
	//	return common.EmptyHash, auth.Nonce.Uint64(), err
	//} else {
	//	return tx.Hash(), auth.Nonce.Uint64(), nil
	//}

	//nonce, err := c.sender.client.NonceAt(context.Background(), c.address, nil)
	//if err != nil {
	//	return common.EmptyHash, nonce, err
	//}
	//
	//c.nonce = nonce

	inputarg,err:=hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	complexityarg:=big.NewInt(8)
	//args,err:=arguments.Pack(inputarg,complexityarg)
	if err != nil {
		return common.EmptyHash, c.nonce, err
	}
	payload, err := utils.PackMethod(ABI, "costManyGas",inputarg,complexityarg)
	if err != nil {
		return common.EmptyHash, c.nonce, err
	}

	originNonce := c.nonce

	tx, err := c.newSignedTx(contract, big.NewInt(0), payload)
	if err != nil {
		return common.EmptyHash, originNonce, err
	}

	if err := c.SendTx(tx); err != nil {
		return common.EmptyHash, originNonce, err
	}

	return tx.Hash(), originNonce, nil
}

func (c *Account) DataTxNum(contract common.Address) (uint64, error) {
	if c.sender == nil {
		return 0, ErrNoSender
	}

	st, err := dataTest.NewDataTest(contract, c.sender.client)
	if err != nil {
		return 0, err
	}
	if num, err := st.TxNum(nil); err != nil {
		return 0, err
	} else {
		return num.Uint64(), nil
	}
}

func (c *Account) DataExist(contract common.Address) bool {
	if c.sender == nil {
		return false
	}

	st, err := dataTest.NewDataTest(contract, c.sender.client)
	if err != nil {
		return false
	}

	if startTime, _ := st.StartTime(nil); startTime > 0 {
		return true
	} else {
		return false
	}
}

