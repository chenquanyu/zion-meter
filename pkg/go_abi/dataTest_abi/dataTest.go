// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dataTest

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

var (
	MethodCostManyGas = "costManyGas"

	MethodHasHugeCallData = "hasHugeCallData"

	MethodReset = "reset"

	MethodStartTime = "startTime"

	MethodTxNum = "txNum"
)

// DataTestABI is the input ABI used to generate the binding from.
const DataTestABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"complexity\",\"type\":\"uint256\"}],\"name\":\"costManyGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"hasHugeCallData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_startTime\",\"type\":\"uint64\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"txNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataTestBin is the compiled bytecode used for deploying new contracts.
var DataTestBin = "0x608060405234801561001057600080fd5b50610919806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633a4ef5441461005c57806378e979251461007a5780638a229c6514610098578063e365981c146100b4578063ef22dc29146100d0575b600080fd5b6100646100ec565b6040516100719190610333565b60405180910390f35b6100826100f2565b60405161008f9190610371565b60405180910390f35b6100b260048036038101906100ad91906105cc565b61010c565b005b6100ce60048036038101906100c99190610641565b6101c5565b005b6100ea60048036038101906100e5919061069a565b6101f9565b005b60015481565b600260009054906101000a900467ffffffffffffffff1681565b60005b81518110156101c157600382828151811061012d5761012c6106f6565b5b6020026020010151604051610142919061079f565b908152602001604051809103902060009054906101000a900460ff16156003838381518110610174576101736106f6565b5b6020026020010151604051610189919061079f565b908152602001604051809103902060006101000a81548160ff02191690831515021790555080806101b9906107e5565b91505061010f565b5050565b80600260006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600060018190555050565b60005b81811015610259576001600080828254610216919061082d565b925050819055508260046000805481526020019081526020016000209080519060200190610245929190610277565b508080610251906107e5565b9150506101fc565b50600180600082825461026c919061082d565b925050819055505050565b828054610283906108b2565b90600052602060002090601f0160209004810192826102a557600085556102ec565b82601f106102be57805160ff19168380011785556102ec565b828001600101855582156102ec579182015b828111156102eb5782518255916020019190600101906102d0565b5b5090506102f991906102fd565b5090565b5b808211156103165760008160009055506001016102fe565b5090565b6000819050919050565b61032d8161031a565b82525050565b60006020820190506103486000830184610324565b92915050565b600067ffffffffffffffff82169050919050565b61036b8161034e565b82525050565b60006020820190506103866000830184610362565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103ee826103a5565b810181811067ffffffffffffffff8211171561040d5761040c6103b6565b5b80604052505050565b600061042061038c565b905061042c82826103e5565b919050565b600067ffffffffffffffff82111561044c5761044b6103b6565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff821115610482576104816103b6565b5b61048b826103a5565b9050602081019050919050565b82818337600083830152505050565b60006104ba6104b584610467565b610416565b9050828152602081018484840111156104d6576104d5610462565b5b6104e1848285610498565b509392505050565b600082601f8301126104fe576104fd6103a0565b5b813561050e8482602086016104a7565b91505092915050565b600061052a61052584610431565b610416565b9050808382526020820190506020840283018581111561054d5761054c61045d565b5b835b8181101561059457803567ffffffffffffffff811115610572576105716103a0565b5b80860161057f89826104e9565b8552602085019450505060208101905061054f565b5050509392505050565b600082601f8301126105b3576105b26103a0565b5b81356105c3848260208601610517565b91505092915050565b6000602082840312156105e2576105e1610396565b5b600082013567ffffffffffffffff811115610600576105ff61039b565b5b61060c8482850161059e565b91505092915050565b61061e8161034e565b811461062957600080fd5b50565b60008135905061063b81610615565b92915050565b60006020828403121561065757610656610396565b5b60006106658482850161062c565b91505092915050565b6106778161031a565b811461068257600080fd5b50565b6000813590506106948161066e565b92915050565b600080604083850312156106b1576106b0610396565b5b600083013567ffffffffffffffff8111156106cf576106ce61039b565b5b6106db858286016104e9565b92505060206106ec85828601610685565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081519050919050565b600081905092915050565b60005b8381101561075957808201518184015260208101905061073e565b83811115610768576000848401525b50505050565b600061077982610725565b6107838185610730565b935061079381856020860161073b565b80840191505092915050565b60006107ab828461076e565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006107f08261031a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610822576108216107b6565b5b600182019050919050565b60006108388261031a565b91506108438361031a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610878576108776107b6565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806108ca57607f821691505b6020821081036108dd576108dc610883565b5b5091905056fea2646970667358221220d1fe5b9bf52fdb63213f0ae288e5438c0f6f9fda0337c5336e4e0ab0274af84064736f6c634300080d0033"

// DeployDataTest deploys a new Ethereum contract, binding an instance of DataTest to it.
func DeployDataTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataTest, error) {
	parsed, err := abi.JSON(strings.NewReader(DataTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataTest{DataTestCaller: DataTestCaller{contract: contract}, DataTestTransactor: DataTestTransactor{contract: contract}, DataTestFilterer: DataTestFilterer{contract: contract}}, nil
}

// DataTest is an auto generated Go binding around an Ethereum contract.
type DataTest struct {
	DataTestCaller     // Read-only binding to the contract
	DataTestTransactor // Write-only binding to the contract
	DataTestFilterer   // Log filterer for contract events
}

// DataTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataTestSession struct {
	Contract     *DataTest         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataTestCallerSession struct {
	Contract *DataTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DataTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataTestTransactorSession struct {
	Contract     *DataTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DataTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataTestRaw struct {
	Contract *DataTest // Generic contract binding to access the raw methods on
}

// DataTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataTestCallerRaw struct {
	Contract *DataTestCaller // Generic read-only contract binding to access the raw methods on
}

// DataTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataTestTransactorRaw struct {
	Contract *DataTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataTest creates a new instance of DataTest, bound to a specific deployed contract.
func NewDataTest(address common.Address, backend bind.ContractBackend) (*DataTest, error) {
	contract, err := bindDataTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataTest{DataTestCaller: DataTestCaller{contract: contract}, DataTestTransactor: DataTestTransactor{contract: contract}, DataTestFilterer: DataTestFilterer{contract: contract}}, nil
}

// NewDataTestCaller creates a new read-only instance of DataTest, bound to a specific deployed contract.
func NewDataTestCaller(address common.Address, caller bind.ContractCaller) (*DataTestCaller, error) {
	contract, err := bindDataTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataTestCaller{contract: contract}, nil
}

// NewDataTestTransactor creates a new write-only instance of DataTest, bound to a specific deployed contract.
func NewDataTestTransactor(address common.Address, transactor bind.ContractTransactor) (*DataTestTransactor, error) {
	contract, err := bindDataTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataTestTransactor{contract: contract}, nil
}

// NewDataTestFilterer creates a new log filterer instance of DataTest, bound to a specific deployed contract.
func NewDataTestFilterer(address common.Address, filterer bind.ContractFilterer) (*DataTestFilterer, error) {
	contract, err := bindDataTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataTestFilterer{contract: contract}, nil
}

// bindDataTest binds a generic wrapper to an already deployed contract.
func bindDataTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataTest *DataTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataTest.Contract.DataTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataTest *DataTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataTest.Contract.DataTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataTest *DataTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataTest.Contract.DataTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataTest *DataTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataTest *DataTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataTest *DataTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataTest.Contract.contract.Transact(opts, method, params...)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_DataTest *DataTestCaller) StartTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DataTest.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_DataTest *DataTestSession) StartTime() (uint64, error) {
	return _DataTest.Contract.StartTime(&_DataTest.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint64)
func (_DataTest *DataTestCallerSession) StartTime() (uint64, error) {
	return _DataTest.Contract.StartTime(&_DataTest.CallOpts)
}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_DataTest *DataTestCaller) TxNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataTest.contract.Call(opts, &out, "txNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_DataTest *DataTestSession) TxNum() (*big.Int, error) {
	return _DataTest.Contract.TxNum(&_DataTest.CallOpts)
}

// TxNum is a free data retrieval call binding the contract method 0x3a4ef544.
//
// Solidity: function txNum() view returns(uint256)
func (_DataTest *DataTestCallerSession) TxNum() (*big.Int, error) {
	return _DataTest.Contract.TxNum(&_DataTest.CallOpts)
}

// CostManyGas is a paid mutator transaction binding the contract method 0xef22dc29.
//
// Solidity: function costManyGas(bytes input, uint256 complexity) returns()
func (_DataTest *DataTestTransactor) CostManyGas(opts *bind.TransactOpts, input []byte, complexity *big.Int) (*types.Transaction, error) {
	return _DataTest.contract.Transact(opts, "costManyGas", input, complexity)
}

// CostManyGas is a paid mutator transaction binding the contract method 0xef22dc29.
//
// Solidity: function costManyGas(bytes input, uint256 complexity) returns()
func (_DataTest *DataTestSession) CostManyGas(input []byte, complexity *big.Int) (*types.Transaction, error) {
	return _DataTest.Contract.CostManyGas(&_DataTest.TransactOpts, input, complexity)
}

// CostManyGas is a paid mutator transaction binding the contract method 0xef22dc29.
//
// Solidity: function costManyGas(bytes input, uint256 complexity) returns()
func (_DataTest *DataTestTransactorSession) CostManyGas(input []byte, complexity *big.Int) (*types.Transaction, error) {
	return _DataTest.Contract.CostManyGas(&_DataTest.TransactOpts, input, complexity)
}

// HasHugeCallData is a paid mutator transaction binding the contract method 0x8a229c65.
//
// Solidity: function hasHugeCallData(bytes[] inputs) returns()
func (_DataTest *DataTestTransactor) HasHugeCallData(opts *bind.TransactOpts, inputs [][]byte) (*types.Transaction, error) {
	return _DataTest.contract.Transact(opts, "hasHugeCallData", inputs)
}

// HasHugeCallData is a paid mutator transaction binding the contract method 0x8a229c65.
//
// Solidity: function hasHugeCallData(bytes[] inputs) returns()
func (_DataTest *DataTestSession) HasHugeCallData(inputs [][]byte) (*types.Transaction, error) {
	return _DataTest.Contract.HasHugeCallData(&_DataTest.TransactOpts, inputs)
}

// HasHugeCallData is a paid mutator transaction binding the contract method 0x8a229c65.
//
// Solidity: function hasHugeCallData(bytes[] inputs) returns()
func (_DataTest *DataTestTransactorSession) HasHugeCallData(inputs [][]byte) (*types.Transaction, error) {
	return _DataTest.Contract.HasHugeCallData(&_DataTest.TransactOpts, inputs)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_DataTest *DataTestTransactor) Reset(opts *bind.TransactOpts, _startTime uint64) (*types.Transaction, error) {
	return _DataTest.contract.Transact(opts, "reset", _startTime)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_DataTest *DataTestSession) Reset(_startTime uint64) (*types.Transaction, error) {
	return _DataTest.Contract.Reset(&_DataTest.TransactOpts, _startTime)
}

// Reset is a paid mutator transaction binding the contract method 0xe365981c.
//
// Solidity: function reset(uint64 _startTime) returns()
func (_DataTest *DataTestTransactorSession) Reset(_startTime uint64) (*types.Transaction, error) {
	return _DataTest.Contract.Reset(&_DataTest.TransactOpts, _startTime)
}
