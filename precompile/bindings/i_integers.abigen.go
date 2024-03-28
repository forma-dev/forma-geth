// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IntegersMetaData contains all meta data concerning the Integers contract.
var IntegersMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"fromHexString\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toHexString\",\"inputs\":[{\"name\":\"_i\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"_i\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"_i\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
}

// IntegersABI is the input ABI used to generate the binding from.
// Deprecated: Use IntegersMetaData.ABI instead.
var IntegersABI = IntegersMetaData.ABI

// Integers is an auto generated Go binding around an Ethereum contract.
type Integers struct {
	IntegersCaller     // Read-only binding to the contract
	IntegersTransactor // Write-only binding to the contract
	IntegersFilterer   // Log filterer for contract events
}

// IntegersCaller is an auto generated read-only Go binding around an Ethereum contract.
type IntegersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntegersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IntegersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntegersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IntegersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntegersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IntegersSession struct {
	Contract     *Integers         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IntegersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IntegersCallerSession struct {
	Contract *IntegersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IntegersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IntegersTransactorSession struct {
	Contract     *IntegersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IntegersRaw is an auto generated low-level Go binding around an Ethereum contract.
type IntegersRaw struct {
	Contract *Integers // Generic contract binding to access the raw methods on
}

// IntegersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IntegersCallerRaw struct {
	Contract *IntegersCaller // Generic read-only contract binding to access the raw methods on
}

// IntegersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IntegersTransactorRaw struct {
	Contract *IntegersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIntegers creates a new instance of Integers, bound to a specific deployed contract.
func NewIntegers(address common.Address, backend bind.ContractBackend) (*Integers, error) {
	contract, err := bindIntegers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Integers{IntegersCaller: IntegersCaller{contract: contract}, IntegersTransactor: IntegersTransactor{contract: contract}, IntegersFilterer: IntegersFilterer{contract: contract}}, nil
}

// NewIntegersCaller creates a new read-only instance of Integers, bound to a specific deployed contract.
func NewIntegersCaller(address common.Address, caller bind.ContractCaller) (*IntegersCaller, error) {
	contract, err := bindIntegers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IntegersCaller{contract: contract}, nil
}

// NewIntegersTransactor creates a new write-only instance of Integers, bound to a specific deployed contract.
func NewIntegersTransactor(address common.Address, transactor bind.ContractTransactor) (*IntegersTransactor, error) {
	contract, err := bindIntegers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IntegersTransactor{contract: contract}, nil
}

// NewIntegersFilterer creates a new log filterer instance of Integers, bound to a specific deployed contract.
func NewIntegersFilterer(address common.Address, filterer bind.ContractFilterer) (*IntegersFilterer, error) {
	contract, err := bindIntegers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IntegersFilterer{contract: contract}, nil
}

// bindIntegers binds a generic wrapper to an already deployed contract.
func bindIntegers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IntegersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Integers *IntegersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Integers.Contract.IntegersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Integers *IntegersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Integers.Contract.IntegersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Integers *IntegersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Integers.Contract.IntegersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Integers *IntegersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Integers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Integers *IntegersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Integers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Integers *IntegersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Integers.Contract.contract.Transact(opts, method, params...)
}

// FromHexString is a free data retrieval call binding the contract method 0x9978443d.
//
// Solidity: function fromHexString(string _str) view returns(uint256)
func (_Integers *IntegersCaller) FromHexString(opts *bind.CallOpts, _str string) (*big.Int, error) {
	var out []interface{}
	err := _Integers.contract.Call(opts, &out, "fromHexString", _str)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FromHexString is a free data retrieval call binding the contract method 0x9978443d.
//
// Solidity: function fromHexString(string _str) view returns(uint256)
func (_Integers *IntegersSession) FromHexString(_str string) (*big.Int, error) {
	return _Integers.Contract.FromHexString(&_Integers.CallOpts, _str)
}

// FromHexString is a free data retrieval call binding the contract method 0x9978443d.
//
// Solidity: function fromHexString(string _str) view returns(uint256)
func (_Integers *IntegersCallerSession) FromHexString(_str string) (*big.Int, error) {
	return _Integers.Contract.FromHexString(&_Integers.CallOpts, _str)
}

// ToHexString is a free data retrieval call binding the contract method 0x8fba8d5c.
//
// Solidity: function toHexString(uint256 _i) view returns(string)
func (_Integers *IntegersCaller) ToHexString(opts *bind.CallOpts, _i *big.Int) (string, error) {
	var out []interface{}
	err := _Integers.contract.Call(opts, &out, "toHexString", _i)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToHexString is a free data retrieval call binding the contract method 0x8fba8d5c.
//
// Solidity: function toHexString(uint256 _i) view returns(string)
func (_Integers *IntegersSession) ToHexString(_i *big.Int) (string, error) {
	return _Integers.Contract.ToHexString(&_Integers.CallOpts, _i)
}

// ToHexString is a free data retrieval call binding the contract method 0x8fba8d5c.
//
// Solidity: function toHexString(uint256 _i) view returns(string)
func (_Integers *IntegersCallerSession) ToHexString(_i *big.Int) (string, error) {
	return _Integers.Contract.ToHexString(&_Integers.CallOpts, _i)
}

// ToString is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 _i) view returns(string)
func (_Integers *IntegersCaller) ToString(opts *bind.CallOpts, _i *big.Int) (string, error) {
	var out []interface{}
	err := _Integers.contract.Call(opts, &out, "toString", _i)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 _i) view returns(string)
func (_Integers *IntegersSession) ToString(_i *big.Int) (string, error) {
	return _Integers.Contract.ToString(&_Integers.CallOpts, _i)
}

// ToString is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 _i) view returns(string)
func (_Integers *IntegersCallerSession) ToString(_i *big.Int) (string, error) {
	return _Integers.Contract.ToString(&_Integers.CallOpts, _i)
}

// ToString0 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 _i) view returns(string)
func (_Integers *IntegersCaller) ToString0(opts *bind.CallOpts, _i *big.Int) (string, error) {
	var out []interface{}
	err := _Integers.contract.Call(opts, &out, "toString0", _i)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString0 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 _i) view returns(string)
func (_Integers *IntegersSession) ToString0(_i *big.Int) (string, error) {
	return _Integers.Contract.ToString0(&_Integers.CallOpts, _i)
}

// ToString0 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 _i) view returns(string)
func (_Integers *IntegersCallerSession) ToString0(_i *big.Int) (string, error) {
	return _Integers.Contract.ToString0(&_Integers.CallOpts, _i)
}
