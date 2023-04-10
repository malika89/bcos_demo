// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// LotteryRecordABI is the input ABI used to generate the binding from.
const LotteryRecordABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_orderHash\",\"type\":\"string\"},{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"saveTransactionData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_txStr\",\"type\":\"string\"}],\"name\":\"getTransactionData\",\"outputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"orderHash\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"isTxExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LotteryRecordBin is the compiled bytecode used for deploying new contracts.
var LotteryRecordBin = "0x608060405234801561001057600080fd5b50610b3e806100206000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806356b90c951461006757806383197ef014610172578063beddf82c14610189578063e89648a11461033d575b600080fd5b34801561007357600080fd5b50610154600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506103be565b60405180826000191660001916815260200191505060405180910390f35b34801561017e57600080fd5b50610187610744565b005b34801561019557600080fd5b506101f0600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061075d565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561029857808201518184015260208101905061027d565b50505050905090810190601f1680156102c55780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156102fe5780820151818401526020810190506102e3565b50505050905090810190601f16801561032b5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34801561034957600080fd5b506103a4600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506109e9565b604051808215151515815260200191505060405180910390f35b60008060008686868642604051602001808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140184805190602001908083835b602083101515610482578051825260208201915060208101905060208303925061045d565b6001836020036101000a03801982511681845116808217855250505050505090500183805190602001908083835b6020831015156104d557805182526020820191506020810190506020830392506104b0565b6001836020036101000a038019825116818451168082178552505050505050905001828152602001955050505050506040516020818303038152906040526040518082805190602001908083835b6020831015156105485780518252602082019150602081019050602083039250610523565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209150610580856109e9565b1515156105f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f54524332304461746120616c726561647920657869737473000000000000000081525060200191505060405180910390fd5b6000856040518082805190602001908083835b60208310151561062d5780518252602082019150602081019050602083039250610608565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390209050868160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550858160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084816002019080519060200190610700929190610a6d565b5083816003019080519060200190610719929190610a6d565b5060018160040160006101000a81548160ff0219169083151502179055508192505050949350505050565b3373ffffffffffffffffffffffffffffffffffffffff16ff5b600080606080600061076e866109e9565b15156107e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f5472616e73616374696f6e20646f6573206e6f7420657869737400000000000081525060200191505060405180910390fd5b6000866040518082805190602001908083835b60208310151561081a57805182526020820191506020810190506020830392506107f5565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260020183600301818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109365780601f1061090b57610100808354040283529160200191610936565b820191906000526020600020905b81548152906001019060200180831161091957829003601f168201915b50505050509150808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109d25780601f106109a7576101008083540402835291602001916109d2565b820191906000526020600020905b8154815290600101906020018083116109b557829003601f168201915b505050505090509450945094509450509193509193565b600080826040518082805190602001908083835b602083101515610a2257805182526020820191506020810190506020830392506109fd565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060040160009054906101000a900460ff169050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610aae57805160ff1916838001178555610adc565b82800160010185558215610adc579182015b82811115610adb578251825591602001919060010190610ac0565b5b509050610ae99190610aed565b5090565b610b0f91905b80821115610b0b576000816000905550600101610af3565b5090565b905600a165627a7a72305820b2bd031c559ba70d163ed7d52e29e94ab50ea7f0a26d66df60d91caf64eb7ebb0029"

// DeployLotteryRecord deploys a new contract, binding an instance of LotteryRecord to it.
func DeployLotteryRecord(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LotteryRecord, error) {
	parsed, err := abi.JSON(strings.NewReader(LotteryRecordABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LotteryRecordBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LotteryRecord{LotteryRecordCaller: LotteryRecordCaller{contract: contract}, LotteryRecordTransactor: LotteryRecordTransactor{contract: contract}, LotteryRecordFilterer: LotteryRecordFilterer{contract: contract}}, nil
}

func AsyncDeployLotteryRecord(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(LotteryRecordABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(LotteryRecordBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// LotteryRecord is an auto generated Go binding around a Solidity contract.
type LotteryRecord struct {
	LotteryRecordCaller     // Read-only binding to the contract
	LotteryRecordTransactor // Write-only binding to the contract
	LotteryRecordFilterer   // Log filterer for contract events
}

// LotteryRecordCaller is an auto generated read-only Go binding around a Solidity contract.
type LotteryRecordCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryRecordTransactor is an auto generated write-only Go binding around a Solidity contract.
type LotteryRecordTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryRecordFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type LotteryRecordFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryRecordSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type LotteryRecordSession struct {
	Contract     *LotteryRecord    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LotteryRecordCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type LotteryRecordCallerSession struct {
	Contract *LotteryRecordCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LotteryRecordTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type LotteryRecordTransactorSession struct {
	Contract     *LotteryRecordTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LotteryRecordRaw is an auto generated low-level Go binding around a Solidity contract.
type LotteryRecordRaw struct {
	Contract *LotteryRecord // Generic contract binding to access the raw methods on
}

// LotteryRecordCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type LotteryRecordCallerRaw struct {
	Contract *LotteryRecordCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryRecordTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type LotteryRecordTransactorRaw struct {
	Contract *LotteryRecordTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLotteryRecord creates a new instance of LotteryRecord, bound to a specific deployed contract.
func NewLotteryRecord(address common.Address, backend bind.ContractBackend) (*LotteryRecord, error) {
	contract, err := bindLotteryRecord(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LotteryRecord{LotteryRecordCaller: LotteryRecordCaller{contract: contract}, LotteryRecordTransactor: LotteryRecordTransactor{contract: contract}, LotteryRecordFilterer: LotteryRecordFilterer{contract: contract}}, nil
}

// NewLotteryRecordCaller creates a new read-only instance of LotteryRecord, bound to a specific deployed contract.
func NewLotteryRecordCaller(address common.Address, caller bind.ContractCaller) (*LotteryRecordCaller, error) {
	contract, err := bindLotteryRecord(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryRecordCaller{contract: contract}, nil
}

// NewLotteryRecordTransactor creates a new write-only instance of LotteryRecord, bound to a specific deployed contract.
func NewLotteryRecordTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryRecordTransactor, error) {
	contract, err := bindLotteryRecord(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryRecordTransactor{contract: contract}, nil
}

// NewLotteryRecordFilterer creates a new log filterer instance of LotteryRecord, bound to a specific deployed contract.
func NewLotteryRecordFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryRecordFilterer, error) {
	contract, err := bindLotteryRecord(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryRecordFilterer{contract: contract}, nil
}

// bindLotteryRecord binds a generic wrapper to an already deployed contract.
func bindLotteryRecord(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LotteryRecordABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LotteryRecord *LotteryRecordRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LotteryRecord.Contract.LotteryRecordCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LotteryRecord *LotteryRecordRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _LotteryRecord.Contract.LotteryRecordTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LotteryRecord *LotteryRecordRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _LotteryRecord.Contract.LotteryRecordTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LotteryRecord *LotteryRecordCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LotteryRecord.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LotteryRecord *LotteryRecordTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _LotteryRecord.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LotteryRecord *LotteryRecordTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _LotteryRecord.Contract.contract.Transact(opts, method, params...)
}

// GetTransactionData is a free data retrieval call binding the contract method 0xbeddf82c.
//
// Solidity: function getTransactionData(string _txStr) constant returns(address sender, address receiver, string orderHash, string data)
func (_LotteryRecord *LotteryRecordCaller) GetTransactionData(opts *bind.CallOpts, _txStr string) (struct {
	Sender    common.Address
	Receiver  common.Address
	OrderHash string
	Data      string
}, error) {
	ret := new(struct {
		Sender    common.Address
		Receiver  common.Address
		OrderHash string
		Data      string
	})
	out := ret
	err := _LotteryRecord.contract.Call(opts, out, "getTransactionData", _txStr)
	return *ret, err
}

// GetTransactionData is a free data retrieval call binding the contract method 0xbeddf82c.
//
// Solidity: function getTransactionData(string _txStr) constant returns(address sender, address receiver, string orderHash, string data)
func (_LotteryRecord *LotteryRecordSession) GetTransactionData(_txStr string) (struct {
	Sender    common.Address
	Receiver  common.Address
	OrderHash string
	Data      string
}, error) {
	return _LotteryRecord.Contract.GetTransactionData(&_LotteryRecord.CallOpts, _txStr)
}

// GetTransactionData is a free data retrieval call binding the contract method 0xbeddf82c.
//
// Solidity: function getTransactionData(string _txStr) constant returns(address sender, address receiver, string orderHash, string data)
func (_LotteryRecord *LotteryRecordCallerSession) GetTransactionData(_txStr string) (struct {
	Sender    common.Address
	Receiver  common.Address
	OrderHash string
	Data      string
}, error) {
	return _LotteryRecord.Contract.GetTransactionData(&_LotteryRecord.CallOpts, _txStr)
}

// IsTxExists is a free data retrieval call binding the contract method 0xe89648a1.
//
// Solidity: function isTxExists(string key) constant returns(bool)
func (_LotteryRecord *LotteryRecordCaller) IsTxExists(opts *bind.CallOpts, key string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LotteryRecord.contract.Call(opts, out, "isTxExists", key)
	return *ret0, err
}

// IsTxExists is a free data retrieval call binding the contract method 0xe89648a1.
//
// Solidity: function isTxExists(string key) constant returns(bool)
func (_LotteryRecord *LotteryRecordSession) IsTxExists(key string) (bool, error) {
	return _LotteryRecord.Contract.IsTxExists(&_LotteryRecord.CallOpts, key)
}

// IsTxExists is a free data retrieval call binding the contract method 0xe89648a1.
//
// Solidity: function isTxExists(string key) constant returns(bool)
func (_LotteryRecord *LotteryRecordCallerSession) IsTxExists(key string) (bool, error) {
	return _LotteryRecord.Contract.IsTxExists(&_LotteryRecord.CallOpts, key)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_LotteryRecord *LotteryRecordTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _LotteryRecord.contract.Transact(opts, "destroy")
}

func (_LotteryRecord *LotteryRecordTransactor) AsyncDestroy(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryRecord.contract.AsyncTransact(opts, handler, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_LotteryRecord *LotteryRecordSession) Destroy() (*types.Transaction, *types.Receipt, error) {
	return _LotteryRecord.Contract.Destroy(&_LotteryRecord.TransactOpts)
}

func (_LotteryRecord *LotteryRecordSession) AsyncDestroy(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _LotteryRecord.Contract.AsyncDestroy(handler, &_LotteryRecord.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_LotteryRecord *LotteryRecordTransactorSession) Destroy() (*types.Transaction, *types.Receipt, error) {
	return _LotteryRecord.Contract.Destroy(&_LotteryRecord.TransactOpts)
}

func (_LotteryRecord *LotteryRecordTransactorSession) AsyncDestroy(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _LotteryRecord.Contract.AsyncDestroy(handler, &_LotteryRecord.TransactOpts)
}

// SaveTransactionData is a paid mutator transaction binding the contract method 0x56b90c95.
//
// Solidity: function saveTransactionData(address _sender, address _receiver, string _orderHash, string _data) returns(bytes32)
func (_LotteryRecord *LotteryRecordTransactor) SaveTransactionData(opts *bind.TransactOpts, _sender common.Address, _receiver common.Address, _orderHash string, _data string) (*types.Transaction, *types.Receipt, error) {
	return _LotteryRecord.contract.Transact(opts, "saveTransactionData", _sender, _receiver, _orderHash, _data)
}

func (_LotteryRecord *LotteryRecordTransactor) AsyncSaveTransactionData(handler func(*types.Receipt, error), opts *bind.TransactOpts, _sender common.Address, _receiver common.Address, _orderHash string, _data string) (*types.Transaction, error) {
	return _LotteryRecord.contract.AsyncTransact(opts, handler, "saveTransactionData", _sender, _receiver, _orderHash, _data)
}

// SaveTransactionData is a paid mutator transaction binding the contract method 0x56b90c95.
//
// Solidity: function saveTransactionData(address _sender, address _receiver, string _orderHash, string _data) returns(bytes32)
func (_LotteryRecord *LotteryRecordSession) SaveTransactionData(_sender common.Address, _receiver common.Address, _orderHash string, _data string) (*types.Transaction, *types.Receipt, error) {
	return _LotteryRecord.Contract.SaveTransactionData(&_LotteryRecord.TransactOpts, _sender, _receiver, _orderHash, _data)
}

func (_LotteryRecord *LotteryRecordSession) AsyncSaveTransactionData(handler func(*types.Receipt, error), _sender common.Address, _receiver common.Address, _orderHash string, _data string) (*types.Transaction, error) {
	return _LotteryRecord.Contract.AsyncSaveTransactionData(handler, &_LotteryRecord.TransactOpts, _sender, _receiver, _orderHash, _data)
}

// SaveTransactionData is a paid mutator transaction binding the contract method 0x56b90c95.
//
// Solidity: function saveTransactionData(address _sender, address _receiver, string _orderHash, string _data) returns(bytes32)
func (_LotteryRecord *LotteryRecordTransactorSession) SaveTransactionData(_sender common.Address, _receiver common.Address, _orderHash string, _data string) (*types.Transaction, *types.Receipt, error) {
	return _LotteryRecord.Contract.SaveTransactionData(&_LotteryRecord.TransactOpts, _sender, _receiver, _orderHash, _data)
}

func (_LotteryRecord *LotteryRecordTransactorSession) AsyncSaveTransactionData(handler func(*types.Receipt, error), _sender common.Address, _receiver common.Address, _orderHash string, _data string) (*types.Transaction, error) {
	return _LotteryRecord.Contract.AsyncSaveTransactionData(handler, &_LotteryRecord.TransactOpts, _sender, _receiver, _orderHash, _data)
}
