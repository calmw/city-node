// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package intoCityNode

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

// BgtTokenMetaData contains all meta data concerning the BgtToken contract.
var BgtTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addProhibitFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addProhibitTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"inProhibitFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"inProhibitTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeProhibitFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeProhibitTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMinter_\",\"type\":\"bool\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BgtTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BgtTokenMetaData.ABI instead.
var BgtTokenABI = BgtTokenMetaData.ABI

// BgtToken is an auto generated Go binding around an Ethereum contract.
type BgtToken struct {
	BgtTokenCaller     // Read-only binding to the contract
	BgtTokenTransactor // Write-only binding to the contract
	BgtTokenFilterer   // Log filterer for contract events
}

// BgtTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BgtTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BgtTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BgtTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BgtTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BgtTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BgtTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BgtTokenSession struct {
	Contract     *BgtToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BgtTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BgtTokenCallerSession struct {
	Contract *BgtTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BgtTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BgtTokenTransactorSession struct {
	Contract     *BgtTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BgtTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BgtTokenRaw struct {
	Contract *BgtToken // Generic contract binding to access the raw methods on
}

// BgtTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BgtTokenCallerRaw struct {
	Contract *BgtTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BgtTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BgtTokenTransactorRaw struct {
	Contract *BgtTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBgtToken creates a new instance of BgtToken, bound to a specific deployed contract.
func NewBgtToken(address common.Address, backend bind.ContractBackend) (*BgtToken, error) {
	contract, err := bindBgtToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BgtToken{BgtTokenCaller: BgtTokenCaller{contract: contract}, BgtTokenTransactor: BgtTokenTransactor{contract: contract}, BgtTokenFilterer: BgtTokenFilterer{contract: contract}}, nil
}

// NewBgtTokenCaller creates a new read-only instance of BgtToken, bound to a specific deployed contract.
func NewBgtTokenCaller(address common.Address, caller bind.ContractCaller) (*BgtTokenCaller, error) {
	contract, err := bindBgtToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BgtTokenCaller{contract: contract}, nil
}

// NewBgtTokenTransactor creates a new write-only instance of BgtToken, bound to a specific deployed contract.
func NewBgtTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BgtTokenTransactor, error) {
	contract, err := bindBgtToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BgtTokenTransactor{contract: contract}, nil
}

// NewBgtTokenFilterer creates a new log filterer instance of BgtToken, bound to a specific deployed contract.
func NewBgtTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BgtTokenFilterer, error) {
	contract, err := bindBgtToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BgtTokenFilterer{contract: contract}, nil
}

// bindBgtToken binds a generic wrapper to an already deployed contract.
func bindBgtToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BgtTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BgtToken *BgtTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BgtToken.Contract.BgtTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BgtToken *BgtTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BgtToken.Contract.BgtTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BgtToken *BgtTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BgtToken.Contract.BgtTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BgtToken *BgtTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BgtToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BgtToken *BgtTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BgtToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BgtToken *BgtTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BgtToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BgtToken *BgtTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BgtToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BgtToken *BgtTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BgtToken.Contract.Allowance(&_BgtToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BgtToken *BgtTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BgtToken.Contract.Allowance(&_BgtToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BgtToken *BgtTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BgtToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BgtToken *BgtTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BgtToken.Contract.BalanceOf(&_BgtToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BgtToken *BgtTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BgtToken.Contract.BalanceOf(&_BgtToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BgtToken *BgtTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BgtToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BgtToken *BgtTokenSession) Decimals() (uint8, error) {
	return _BgtToken.Contract.Decimals(&_BgtToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BgtToken *BgtTokenCallerSession) Decimals() (uint8, error) {
	return _BgtToken.Contract.Decimals(&_BgtToken.CallOpts)
}

// InProhibitFrom is a free data retrieval call binding the contract method 0x64258251.
//
// Solidity: function inProhibitFrom(address account) view returns(bool)
func (_BgtToken *BgtTokenCaller) InProhibitFrom(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BgtToken.contract.Call(opts, &out, "inProhibitFrom", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InProhibitFrom is a free data retrieval call binding the contract method 0x64258251.
//
// Solidity: function inProhibitFrom(address account) view returns(bool)
func (_BgtToken *BgtTokenSession) InProhibitFrom(account common.Address) (bool, error) {
	return _BgtToken.Contract.InProhibitFrom(&_BgtToken.CallOpts, account)
}

// InProhibitFrom is a free data retrieval call binding the contract method 0x64258251.
//
// Solidity: function inProhibitFrom(address account) view returns(bool)
func (_BgtToken *BgtTokenCallerSession) InProhibitFrom(account common.Address) (bool, error) {
	return _BgtToken.Contract.InProhibitFrom(&_BgtToken.CallOpts, account)
}

// InProhibitTo is a free data retrieval call binding the contract method 0x2f119a36.
//
// Solidity: function inProhibitTo(address account) view returns(bool)
func (_BgtToken *BgtTokenCaller) InProhibitTo(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BgtToken.contract.Call(opts, &out, "inProhibitTo", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InProhibitTo is a free data retrieval call binding the contract method 0x2f119a36.
//
// Solidity: function inProhibitTo(address account) view returns(bool)
func (_BgtToken *BgtTokenSession) InProhibitTo(account common.Address) (bool, error) {
	return _BgtToken.Contract.InProhibitTo(&_BgtToken.CallOpts, account)
}

// InProhibitTo is a free data retrieval call binding the contract method 0x2f119a36.
//
// Solidity: function inProhibitTo(address account) view returns(bool)
func (_BgtToken *BgtTokenCallerSession) InProhibitTo(account common.Address) (bool, error) {
	return _BgtToken.Contract.InProhibitTo(&_BgtToken.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_BgtToken *BgtTokenCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BgtToken.contract.Call(opts, &out, "isMinter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_BgtToken *BgtTokenSession) IsMinter(account common.Address) (bool, error) {
	return _BgtToken.Contract.IsMinter(&_BgtToken.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_BgtToken *BgtTokenCallerSession) IsMinter(account common.Address) (bool, error) {
	return _BgtToken.Contract.IsMinter(&_BgtToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BgtToken *BgtTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BgtToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BgtToken *BgtTokenSession) Name() (string, error) {
	return _BgtToken.Contract.Name(&_BgtToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BgtToken *BgtTokenCallerSession) Name() (string, error) {
	return _BgtToken.Contract.Name(&_BgtToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BgtToken *BgtTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BgtToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BgtToken *BgtTokenSession) Owner() (common.Address, error) {
	return _BgtToken.Contract.Owner(&_BgtToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BgtToken *BgtTokenCallerSession) Owner() (common.Address, error) {
	return _BgtToken.Contract.Owner(&_BgtToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BgtToken *BgtTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BgtToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BgtToken *BgtTokenSession) Symbol() (string, error) {
	return _BgtToken.Contract.Symbol(&_BgtToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BgtToken *BgtTokenCallerSession) Symbol() (string, error) {
	return _BgtToken.Contract.Symbol(&_BgtToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BgtToken *BgtTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BgtToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BgtToken *BgtTokenSession) TotalSupply() (*big.Int, error) {
	return _BgtToken.Contract.TotalSupply(&_BgtToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BgtToken *BgtTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BgtToken.Contract.TotalSupply(&_BgtToken.CallOpts)
}

// AddProhibitFrom is a paid mutator transaction binding the contract method 0x0d15202b.
//
// Solidity: function addProhibitFrom(address account) returns()
func (_BgtToken *BgtTokenTransactor) AddProhibitFrom(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "addProhibitFrom", account)
}

// AddProhibitFrom is a paid mutator transaction binding the contract method 0x0d15202b.
//
// Solidity: function addProhibitFrom(address account) returns()
func (_BgtToken *BgtTokenSession) AddProhibitFrom(account common.Address) (*types.Transaction, error) {
	return _BgtToken.Contract.AddProhibitFrom(&_BgtToken.TransactOpts, account)
}

// AddProhibitFrom is a paid mutator transaction binding the contract method 0x0d15202b.
//
// Solidity: function addProhibitFrom(address account) returns()
func (_BgtToken *BgtTokenTransactorSession) AddProhibitFrom(account common.Address) (*types.Transaction, error) {
	return _BgtToken.Contract.AddProhibitFrom(&_BgtToken.TransactOpts, account)
}

// AddProhibitTo is a paid mutator transaction binding the contract method 0x2cd63548.
//
// Solidity: function addProhibitTo(address account) returns()
func (_BgtToken *BgtTokenTransactor) AddProhibitTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "addProhibitTo", account)
}

// AddProhibitTo is a paid mutator transaction binding the contract method 0x2cd63548.
//
// Solidity: function addProhibitTo(address account) returns()
func (_BgtToken *BgtTokenSession) AddProhibitTo(account common.Address) (*types.Transaction, error) {
	return _BgtToken.Contract.AddProhibitTo(&_BgtToken.TransactOpts, account)
}

// AddProhibitTo is a paid mutator transaction binding the contract method 0x2cd63548.
//
// Solidity: function addProhibitTo(address account) returns()
func (_BgtToken *BgtTokenTransactorSession) AddProhibitTo(account common.Address) (*types.Transaction, error) {
	return _BgtToken.Contract.AddProhibitTo(&_BgtToken.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BgtToken *BgtTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BgtToken *BgtTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.Approve(&_BgtToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BgtToken *BgtTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.Approve(&_BgtToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_BgtToken *BgtTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_BgtToken *BgtTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.Burn(&_BgtToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_BgtToken *BgtTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.Burn(&_BgtToken.TransactOpts, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BgtToken *BgtTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BgtToken *BgtTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.DecreaseAllowance(&_BgtToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BgtToken *BgtTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.DecreaseAllowance(&_BgtToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BgtToken *BgtTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BgtToken *BgtTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.IncreaseAllowance(&_BgtToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BgtToken *BgtTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.IncreaseAllowance(&_BgtToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BgtToken *BgtTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BgtToken *BgtTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.Mint(&_BgtToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_BgtToken *BgtTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.Mint(&_BgtToken.TransactOpts, to, amount)
}

// RemoveProhibitFrom is a paid mutator transaction binding the contract method 0x9d69c13c.
//
// Solidity: function removeProhibitFrom(address account) returns()
func (_BgtToken *BgtTokenTransactor) RemoveProhibitFrom(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "removeProhibitFrom", account)
}

// RemoveProhibitFrom is a paid mutator transaction binding the contract method 0x9d69c13c.
//
// Solidity: function removeProhibitFrom(address account) returns()
func (_BgtToken *BgtTokenSession) RemoveProhibitFrom(account common.Address) (*types.Transaction, error) {
	return _BgtToken.Contract.RemoveProhibitFrom(&_BgtToken.TransactOpts, account)
}

// RemoveProhibitFrom is a paid mutator transaction binding the contract method 0x9d69c13c.
//
// Solidity: function removeProhibitFrom(address account) returns()
func (_BgtToken *BgtTokenTransactorSession) RemoveProhibitFrom(account common.Address) (*types.Transaction, error) {
	return _BgtToken.Contract.RemoveProhibitFrom(&_BgtToken.TransactOpts, account)
}

// RemoveProhibitTo is a paid mutator transaction binding the contract method 0xc249865c.
//
// Solidity: function removeProhibitTo(address account) returns()
func (_BgtToken *BgtTokenTransactor) RemoveProhibitTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "removeProhibitTo", account)
}

// RemoveProhibitTo is a paid mutator transaction binding the contract method 0xc249865c.
//
// Solidity: function removeProhibitTo(address account) returns()
func (_BgtToken *BgtTokenSession) RemoveProhibitTo(account common.Address) (*types.Transaction, error) {
	return _BgtToken.Contract.RemoveProhibitTo(&_BgtToken.TransactOpts, account)
}

// RemoveProhibitTo is a paid mutator transaction binding the contract method 0xc249865c.
//
// Solidity: function removeProhibitTo(address account) returns()
func (_BgtToken *BgtTokenTransactorSession) RemoveProhibitTo(account common.Address) (*types.Transaction, error) {
	return _BgtToken.Contract.RemoveProhibitTo(&_BgtToken.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BgtToken *BgtTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BgtToken *BgtTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _BgtToken.Contract.RenounceOwnership(&_BgtToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BgtToken *BgtTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BgtToken.Contract.RenounceOwnership(&_BgtToken.TransactOpts)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address account, bool isMinter_) returns()
func (_BgtToken *BgtTokenTransactor) SetMinter(opts *bind.TransactOpts, account common.Address, isMinter_ bool) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "setMinter", account, isMinter_)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address account, bool isMinter_) returns()
func (_BgtToken *BgtTokenSession) SetMinter(account common.Address, isMinter_ bool) (*types.Transaction, error) {
	return _BgtToken.Contract.SetMinter(&_BgtToken.TransactOpts, account, isMinter_)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address account, bool isMinter_) returns()
func (_BgtToken *BgtTokenTransactorSession) SetMinter(account common.Address, isMinter_ bool) (*types.Transaction, error) {
	return _BgtToken.Contract.SetMinter(&_BgtToken.TransactOpts, account, isMinter_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BgtToken *BgtTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BgtToken *BgtTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.Transfer(&_BgtToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BgtToken *BgtTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.Transfer(&_BgtToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_BgtToken *BgtTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_BgtToken *BgtTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.TransferFrom(&_BgtToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_BgtToken *BgtTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BgtToken.Contract.TransferFrom(&_BgtToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BgtToken *BgtTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BgtToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BgtToken *BgtTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BgtToken.Contract.TransferOwnership(&_BgtToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BgtToken *BgtTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BgtToken.Contract.TransferOwnership(&_BgtToken.TransactOpts, newOwner)
}

// BgtTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BgtToken contract.
type BgtTokenApprovalIterator struct {
	Event *BgtTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BgtTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BgtTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BgtTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BgtTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BgtTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BgtTokenApproval represents a Approval event raised by the BgtToken contract.
type BgtTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BgtToken *BgtTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BgtTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BgtToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BgtTokenApprovalIterator{contract: _BgtToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BgtToken *BgtTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BgtTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BgtToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BgtTokenApproval)
				if err := _BgtToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BgtToken *BgtTokenFilterer) ParseApproval(log types.Log) (*BgtTokenApproval, error) {
	event := new(BgtTokenApproval)
	if err := _BgtToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BgtTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BgtToken contract.
type BgtTokenOwnershipTransferredIterator struct {
	Event *BgtTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BgtTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BgtTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BgtTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BgtTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BgtTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BgtTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BgtToken contract.
type BgtTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BgtToken *BgtTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BgtTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BgtToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BgtTokenOwnershipTransferredIterator{contract: _BgtToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BgtToken *BgtTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BgtTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BgtToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BgtTokenOwnershipTransferred)
				if err := _BgtToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BgtToken *BgtTokenFilterer) ParseOwnershipTransferred(log types.Log) (*BgtTokenOwnershipTransferred, error) {
	event := new(BgtTokenOwnershipTransferred)
	if err := _BgtToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BgtTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BgtToken contract.
type BgtTokenTransferIterator struct {
	Event *BgtTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BgtTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BgtTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BgtTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BgtTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BgtTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BgtTokenTransfer represents a Transfer event raised by the BgtToken contract.
type BgtTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BgtToken *BgtTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BgtTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BgtToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BgtTokenTransferIterator{contract: _BgtToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BgtToken *BgtTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BgtTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BgtToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BgtTokenTransfer)
				if err := _BgtToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BgtToken *BgtTokenFilterer) ParseTransfer(log types.Log) (*BgtTokenTransfer, error) {
	event := new(BgtTokenTransfer)
	if err := _BgtToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
