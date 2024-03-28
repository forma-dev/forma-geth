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

// Base64MetaData contains all meta data concerning the Base64 contract.
var Base64MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"encode\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"encodeURL\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
}

// Base64ABI is the input ABI used to generate the binding from.
// Deprecated: Use Base64MetaData.ABI instead.
var Base64ABI = Base64MetaData.ABI

// Base64 is an auto generated Go binding around an Ethereum contract.
type Base64 struct {
	Base64Caller     // Read-only binding to the contract
	Base64Transactor // Write-only binding to the contract
	Base64Filterer   // Log filterer for contract events
}

// Base64Caller is an auto generated read-only Go binding around an Ethereum contract.
type Base64Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Base64Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Base64Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base64Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Base64Session struct {
	Contract     *Base64           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Base64CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Base64CallerSession struct {
	Contract *Base64Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Base64TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Base64TransactorSession struct {
	Contract     *Base64Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Base64Raw is an auto generated low-level Go binding around an Ethereum contract.
type Base64Raw struct {
	Contract *Base64 // Generic contract binding to access the raw methods on
}

// Base64CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Base64CallerRaw struct {
	Contract *Base64Caller // Generic read-only contract binding to access the raw methods on
}

// Base64TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Base64TransactorRaw struct {
	Contract *Base64Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBase64 creates a new instance of Base64, bound to a specific deployed contract.
func NewBase64(address common.Address, backend bind.ContractBackend) (*Base64, error) {
	contract, err := bindBase64(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Base64{Base64Caller: Base64Caller{contract: contract}, Base64Transactor: Base64Transactor{contract: contract}, Base64Filterer: Base64Filterer{contract: contract}}, nil
}

// NewBase64Caller creates a new read-only instance of Base64, bound to a specific deployed contract.
func NewBase64Caller(address common.Address, caller bind.ContractCaller) (*Base64Caller, error) {
	contract, err := bindBase64(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Base64Caller{contract: contract}, nil
}

// NewBase64Transactor creates a new write-only instance of Base64, bound to a specific deployed contract.
func NewBase64Transactor(address common.Address, transactor bind.ContractTransactor) (*Base64Transactor, error) {
	contract, err := bindBase64(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Base64Transactor{contract: contract}, nil
}

// NewBase64Filterer creates a new log filterer instance of Base64, bound to a specific deployed contract.
func NewBase64Filterer(address common.Address, filterer bind.ContractFilterer) (*Base64Filterer, error) {
	contract, err := bindBase64(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Base64Filterer{contract: contract}, nil
}

// bindBase64 binds a generic wrapper to an already deployed contract.
func bindBase64(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Base64MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base64 *Base64Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base64.Contract.Base64Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base64 *Base64Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base64.Contract.Base64Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base64 *Base64Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base64.Contract.Base64Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base64 *Base64CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base64.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base64 *Base64TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base64.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base64 *Base64TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base64.Contract.contract.Transact(opts, method, params...)
}

// Encode is a free data retrieval call binding the contract method 0x12496a1b.
//
// Solidity: function encode(bytes data) view returns(string)
func (_Base64 *Base64Caller) Encode(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _Base64.contract.Call(opts, &out, "encode", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Encode is a free data retrieval call binding the contract method 0x12496a1b.
//
// Solidity: function encode(bytes data) view returns(string)
func (_Base64 *Base64Session) Encode(data []byte) (string, error) {
	return _Base64.Contract.Encode(&_Base64.CallOpts, data)
}

// Encode is a free data retrieval call binding the contract method 0x12496a1b.
//
// Solidity: function encode(bytes data) view returns(string)
func (_Base64 *Base64CallerSession) Encode(data []byte) (string, error) {
	return _Base64.Contract.Encode(&_Base64.CallOpts, data)
}

// EncodeURL is a free data retrieval call binding the contract method 0xbff52db3.
//
// Solidity: function encodeURL(bytes data) view returns(string)
func (_Base64 *Base64Caller) EncodeURL(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _Base64.contract.Call(opts, &out, "encodeURL", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EncodeURL is a free data retrieval call binding the contract method 0xbff52db3.
//
// Solidity: function encodeURL(bytes data) view returns(string)
func (_Base64 *Base64Session) EncodeURL(data []byte) (string, error) {
	return _Base64.Contract.EncodeURL(&_Base64.CallOpts, data)
}

// EncodeURL is a free data retrieval call binding the contract method 0xbff52db3.
//
// Solidity: function encodeURL(bytes data) view returns(string)
func (_Base64 *Base64CallerSession) EncodeURL(data []byte) (string, error) {
	return _Base64.Contract.EncodeURL(&_Base64.CallOpts, data)
}
