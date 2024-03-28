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

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"contains\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_substr\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"endsWith\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_substr\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"equal\",\"inputs\":[{\"name\":\"_a\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_b\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"equalCaseFold\",\"inputs\":[{\"name\":\"_a\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_b\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexOf\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_substr\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"padEnd\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_len\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_pad\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"padStart\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_len\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_pad\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"repeat\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_count\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"replace\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_old\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_new\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_n\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"replaceAll\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_old\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_new\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"split\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startsWith\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_substr\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toLowerCase\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toUpperCase\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"trim\",\"inputs\":[{\"name\":\"_str\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
}

// StringsABI is the input ABI used to generate the binding from.
// Deprecated: Use StringsMetaData.ABI instead.
var StringsABI = StringsMetaData.ABI

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}

// Contains is a free data retrieval call binding the contract method 0x3fb18aec.
//
// Solidity: function contains(string _str, string _substr) view returns(bool)
func (_Strings *StringsCaller) Contains(opts *bind.CallOpts, _str string, _substr string) (bool, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "contains", _str, _substr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x3fb18aec.
//
// Solidity: function contains(string _str, string _substr) view returns(bool)
func (_Strings *StringsSession) Contains(_str string, _substr string) (bool, error) {
	return _Strings.Contract.Contains(&_Strings.CallOpts, _str, _substr)
}

// Contains is a free data retrieval call binding the contract method 0x3fb18aec.
//
// Solidity: function contains(string _str, string _substr) view returns(bool)
func (_Strings *StringsCallerSession) Contains(_str string, _substr string) (bool, error) {
	return _Strings.Contract.Contains(&_Strings.CallOpts, _str, _substr)
}

// EndsWith is a free data retrieval call binding the contract method 0xded03449.
//
// Solidity: function endsWith(string _str, string _substr) view returns(bool)
func (_Strings *StringsCaller) EndsWith(opts *bind.CallOpts, _str string, _substr string) (bool, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "endsWith", _str, _substr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EndsWith is a free data retrieval call binding the contract method 0xded03449.
//
// Solidity: function endsWith(string _str, string _substr) view returns(bool)
func (_Strings *StringsSession) EndsWith(_str string, _substr string) (bool, error) {
	return _Strings.Contract.EndsWith(&_Strings.CallOpts, _str, _substr)
}

// EndsWith is a free data retrieval call binding the contract method 0xded03449.
//
// Solidity: function endsWith(string _str, string _substr) view returns(bool)
func (_Strings *StringsCallerSession) EndsWith(_str string, _substr string) (bool, error) {
	return _Strings.Contract.EndsWith(&_Strings.CallOpts, _str, _substr)
}

// Equal is a free data retrieval call binding the contract method 0x46bdca9a.
//
// Solidity: function equal(string _a, string _b) view returns(bool)
func (_Strings *StringsCaller) Equal(opts *bind.CallOpts, _a string, _b string) (bool, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "equal", _a, _b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Equal is a free data retrieval call binding the contract method 0x46bdca9a.
//
// Solidity: function equal(string _a, string _b) view returns(bool)
func (_Strings *StringsSession) Equal(_a string, _b string) (bool, error) {
	return _Strings.Contract.Equal(&_Strings.CallOpts, _a, _b)
}

// Equal is a free data retrieval call binding the contract method 0x46bdca9a.
//
// Solidity: function equal(string _a, string _b) view returns(bool)
func (_Strings *StringsCallerSession) Equal(_a string, _b string) (bool, error) {
	return _Strings.Contract.Equal(&_Strings.CallOpts, _a, _b)
}

// EqualCaseFold is a free data retrieval call binding the contract method 0x4a012253.
//
// Solidity: function equalCaseFold(string _a, string _b) view returns(bool)
func (_Strings *StringsCaller) EqualCaseFold(opts *bind.CallOpts, _a string, _b string) (bool, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "equalCaseFold", _a, _b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EqualCaseFold is a free data retrieval call binding the contract method 0x4a012253.
//
// Solidity: function equalCaseFold(string _a, string _b) view returns(bool)
func (_Strings *StringsSession) EqualCaseFold(_a string, _b string) (bool, error) {
	return _Strings.Contract.EqualCaseFold(&_Strings.CallOpts, _a, _b)
}

// EqualCaseFold is a free data retrieval call binding the contract method 0x4a012253.
//
// Solidity: function equalCaseFold(string _a, string _b) view returns(bool)
func (_Strings *StringsCallerSession) EqualCaseFold(_a string, _b string) (bool, error) {
	return _Strings.Contract.EqualCaseFold(&_Strings.CallOpts, _a, _b)
}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string _str, string _substr) view returns(uint256)
func (_Strings *StringsCaller) IndexOf(opts *bind.CallOpts, _str string, _substr string) (*big.Int, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "indexOf", _str, _substr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string _str, string _substr) view returns(uint256)
func (_Strings *StringsSession) IndexOf(_str string, _substr string) (*big.Int, error) {
	return _Strings.Contract.IndexOf(&_Strings.CallOpts, _str, _substr)
}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string _str, string _substr) view returns(uint256)
func (_Strings *StringsCallerSession) IndexOf(_str string, _substr string) (*big.Int, error) {
	return _Strings.Contract.IndexOf(&_Strings.CallOpts, _str, _substr)
}

// PadEnd is a free data retrieval call binding the contract method 0x30c84451.
//
// Solidity: function padEnd(string _str, uint16 _len, string _pad) view returns(string)
func (_Strings *StringsCaller) PadEnd(opts *bind.CallOpts, _str string, _len uint16, _pad string) (string, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "padEnd", _str, _len, _pad)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PadEnd is a free data retrieval call binding the contract method 0x30c84451.
//
// Solidity: function padEnd(string _str, uint16 _len, string _pad) view returns(string)
func (_Strings *StringsSession) PadEnd(_str string, _len uint16, _pad string) (string, error) {
	return _Strings.Contract.PadEnd(&_Strings.CallOpts, _str, _len, _pad)
}

// PadEnd is a free data retrieval call binding the contract method 0x30c84451.
//
// Solidity: function padEnd(string _str, uint16 _len, string _pad) view returns(string)
func (_Strings *StringsCallerSession) PadEnd(_str string, _len uint16, _pad string) (string, error) {
	return _Strings.Contract.PadEnd(&_Strings.CallOpts, _str, _len, _pad)
}

// PadStart is a free data retrieval call binding the contract method 0xcca60dfa.
//
// Solidity: function padStart(string _str, uint16 _len, string _pad) view returns(string)
func (_Strings *StringsCaller) PadStart(opts *bind.CallOpts, _str string, _len uint16, _pad string) (string, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "padStart", _str, _len, _pad)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PadStart is a free data retrieval call binding the contract method 0xcca60dfa.
//
// Solidity: function padStart(string _str, uint16 _len, string _pad) view returns(string)
func (_Strings *StringsSession) PadStart(_str string, _len uint16, _pad string) (string, error) {
	return _Strings.Contract.PadStart(&_Strings.CallOpts, _str, _len, _pad)
}

// PadStart is a free data retrieval call binding the contract method 0xcca60dfa.
//
// Solidity: function padStart(string _str, uint16 _len, string _pad) view returns(string)
func (_Strings *StringsCallerSession) PadStart(_str string, _len uint16, _pad string) (string, error) {
	return _Strings.Contract.PadStart(&_Strings.CallOpts, _str, _len, _pad)
}

// Repeat is a free data retrieval call binding the contract method 0x408182c4.
//
// Solidity: function repeat(string _str, uint16 _count) view returns(string)
func (_Strings *StringsCaller) Repeat(opts *bind.CallOpts, _str string, _count uint16) (string, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "repeat", _str, _count)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Repeat is a free data retrieval call binding the contract method 0x408182c4.
//
// Solidity: function repeat(string _str, uint16 _count) view returns(string)
func (_Strings *StringsSession) Repeat(_str string, _count uint16) (string, error) {
	return _Strings.Contract.Repeat(&_Strings.CallOpts, _str, _count)
}

// Repeat is a free data retrieval call binding the contract method 0x408182c4.
//
// Solidity: function repeat(string _str, uint16 _count) view returns(string)
func (_Strings *StringsCallerSession) Repeat(_str string, _count uint16) (string, error) {
	return _Strings.Contract.Repeat(&_Strings.CallOpts, _str, _count)
}

// Replace is a free data retrieval call binding the contract method 0xcaeac0a9.
//
// Solidity: function replace(string _str, string _old, string _new, uint16 _n) view returns(string)
func (_Strings *StringsCaller) Replace(opts *bind.CallOpts, _str string, _old string, _new string, _n uint16) (string, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "replace", _str, _old, _new, _n)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Replace is a free data retrieval call binding the contract method 0xcaeac0a9.
//
// Solidity: function replace(string _str, string _old, string _new, uint16 _n) view returns(string)
func (_Strings *StringsSession) Replace(_str string, _old string, _new string, _n uint16) (string, error) {
	return _Strings.Contract.Replace(&_Strings.CallOpts, _str, _old, _new, _n)
}

// Replace is a free data retrieval call binding the contract method 0xcaeac0a9.
//
// Solidity: function replace(string _str, string _old, string _new, uint16 _n) view returns(string)
func (_Strings *StringsCallerSession) Replace(_str string, _old string, _new string, _n uint16) (string, error) {
	return _Strings.Contract.Replace(&_Strings.CallOpts, _str, _old, _new, _n)
}

// ReplaceAll is a free data retrieval call binding the contract method 0xca02701f.
//
// Solidity: function replaceAll(string _str, string _old, string _new) view returns(string)
func (_Strings *StringsCaller) ReplaceAll(opts *bind.CallOpts, _str string, _old string, _new string) (string, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "replaceAll", _str, _old, _new)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReplaceAll is a free data retrieval call binding the contract method 0xca02701f.
//
// Solidity: function replaceAll(string _str, string _old, string _new) view returns(string)
func (_Strings *StringsSession) ReplaceAll(_str string, _old string, _new string) (string, error) {
	return _Strings.Contract.ReplaceAll(&_Strings.CallOpts, _str, _old, _new)
}

// ReplaceAll is a free data retrieval call binding the contract method 0xca02701f.
//
// Solidity: function replaceAll(string _str, string _old, string _new) view returns(string)
func (_Strings *StringsCallerSession) ReplaceAll(_str string, _old string, _new string) (string, error) {
	return _Strings.Contract.ReplaceAll(&_Strings.CallOpts, _str, _old, _new)
}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string _str, string _delim) view returns(string[])
func (_Strings *StringsCaller) Split(opts *bind.CallOpts, _str string, _delim string) ([]string, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "split", _str, _delim)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string _str, string _delim) view returns(string[])
func (_Strings *StringsSession) Split(_str string, _delim string) ([]string, error) {
	return _Strings.Contract.Split(&_Strings.CallOpts, _str, _delim)
}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string _str, string _delim) view returns(string[])
func (_Strings *StringsCallerSession) Split(_str string, _delim string) ([]string, error) {
	return _Strings.Contract.Split(&_Strings.CallOpts, _str, _delim)
}

// StartsWith is a free data retrieval call binding the contract method 0xadf069ea.
//
// Solidity: function startsWith(string _str, string _substr) view returns(bool)
func (_Strings *StringsCaller) StartsWith(opts *bind.CallOpts, _str string, _substr string) (bool, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "startsWith", _str, _substr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StartsWith is a free data retrieval call binding the contract method 0xadf069ea.
//
// Solidity: function startsWith(string _str, string _substr) view returns(bool)
func (_Strings *StringsSession) StartsWith(_str string, _substr string) (bool, error) {
	return _Strings.Contract.StartsWith(&_Strings.CallOpts, _str, _substr)
}

// StartsWith is a free data retrieval call binding the contract method 0xadf069ea.
//
// Solidity: function startsWith(string _str, string _substr) view returns(bool)
func (_Strings *StringsCallerSession) StartsWith(_str string, _substr string) (bool, error) {
	return _Strings.Contract.StartsWith(&_Strings.CallOpts, _str, _substr)
}

// ToLowerCase is a free data retrieval call binding the contract method 0x97b21440.
//
// Solidity: function toLowerCase(string _str) view returns(string)
func (_Strings *StringsCaller) ToLowerCase(opts *bind.CallOpts, _str string) (string, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "toLowerCase", _str)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToLowerCase is a free data retrieval call binding the contract method 0x97b21440.
//
// Solidity: function toLowerCase(string _str) view returns(string)
func (_Strings *StringsSession) ToLowerCase(_str string) (string, error) {
	return _Strings.Contract.ToLowerCase(&_Strings.CallOpts, _str)
}

// ToLowerCase is a free data retrieval call binding the contract method 0x97b21440.
//
// Solidity: function toLowerCase(string _str) view returns(string)
func (_Strings *StringsCallerSession) ToLowerCase(_str string) (string, error) {
	return _Strings.Contract.ToLowerCase(&_Strings.CallOpts, _str)
}

// ToUpperCase is a free data retrieval call binding the contract method 0x6299a85b.
//
// Solidity: function toUpperCase(string _str) view returns(string)
func (_Strings *StringsCaller) ToUpperCase(opts *bind.CallOpts, _str string) (string, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "toUpperCase", _str)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToUpperCase is a free data retrieval call binding the contract method 0x6299a85b.
//
// Solidity: function toUpperCase(string _str) view returns(string)
func (_Strings *StringsSession) ToUpperCase(_str string) (string, error) {
	return _Strings.Contract.ToUpperCase(&_Strings.CallOpts, _str)
}

// ToUpperCase is a free data retrieval call binding the contract method 0x6299a85b.
//
// Solidity: function toUpperCase(string _str) view returns(string)
func (_Strings *StringsCallerSession) ToUpperCase(_str string) (string, error) {
	return _Strings.Contract.ToUpperCase(&_Strings.CallOpts, _str)
}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string _str) view returns(string)
func (_Strings *StringsCaller) Trim(opts *bind.CallOpts, _str string) (string, error) {
	var out []interface{}
	err := _Strings.contract.Call(opts, &out, "trim", _str)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string _str) view returns(string)
func (_Strings *StringsSession) Trim(_str string) (string, error) {
	return _Strings.Contract.Trim(&_Strings.CallOpts, _str)
}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string _str) view returns(string)
func (_Strings *StringsCallerSession) Trim(_str string) (string, error) {
	return _Strings.Contract.Trim(&_Strings.CallOpts, _str)
}
