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

// AppraiseMetaData contains all meta data concerning the Appraise contract.
var AppraiseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cityAddress_\",\"type\":\"address\"}],\"name\":\"adminSetCity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cityPioneerAddress_\",\"type\":\"address\"}],\"name\":\"adminSetPioneer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"starAddress_\",\"type\":\"address\"}],\"name\":\"adminSetStar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pioneerAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"chengShiId_\",\"type\":\"bytes32\"}],\"name\":\"appraise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"city\",\"outputs\":[{\"internalType\":\"contractICity\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pioneer\",\"outputs\":[{\"internalType\":\"contractICityPioneer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pioneerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pioneerMonthWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"star\",\"outputs\":[{\"internalType\":\"contractIStar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AppraiseABI is the input ABI used to generate the binding from.
// Deprecated: Use AppraiseMetaData.ABI instead.
var AppraiseABI = AppraiseMetaData.ABI

// Appraise is an auto generated Go binding around an Ethereum contract.
type Appraise struct {
	AppraiseCaller     // Read-only binding to the contract
	AppraiseTransactor // Write-only binding to the contract
	AppraiseFilterer   // Log filterer for contract events
}

// AppraiseCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppraiseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppraiseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppraiseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppraiseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppraiseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppraiseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppraiseSession struct {
	Contract     *Appraise         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppraiseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppraiseCallerSession struct {
	Contract *AppraiseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AppraiseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppraiseTransactorSession struct {
	Contract     *AppraiseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AppraiseRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppraiseRaw struct {
	Contract *Appraise // Generic contract binding to access the raw methods on
}

// AppraiseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppraiseCallerRaw struct {
	Contract *AppraiseCaller // Generic read-only contract binding to access the raw methods on
}

// AppraiseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppraiseTransactorRaw struct {
	Contract *AppraiseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppraise creates a new instance of Appraise, bound to a specific deployed contract.
func NewAppraise(address common.Address, backend bind.ContractBackend) (*Appraise, error) {
	contract, err := bindAppraise(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Appraise{AppraiseCaller: AppraiseCaller{contract: contract}, AppraiseTransactor: AppraiseTransactor{contract: contract}, AppraiseFilterer: AppraiseFilterer{contract: contract}}, nil
}

// NewAppraiseCaller creates a new read-only instance of Appraise, bound to a specific deployed contract.
func NewAppraiseCaller(address common.Address, caller bind.ContractCaller) (*AppraiseCaller, error) {
	contract, err := bindAppraise(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppraiseCaller{contract: contract}, nil
}

// NewAppraiseTransactor creates a new write-only instance of Appraise, bound to a specific deployed contract.
func NewAppraiseTransactor(address common.Address, transactor bind.ContractTransactor) (*AppraiseTransactor, error) {
	contract, err := bindAppraise(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppraiseTransactor{contract: contract}, nil
}

// NewAppraiseFilterer creates a new log filterer instance of Appraise, bound to a specific deployed contract.
func NewAppraiseFilterer(address common.Address, filterer bind.ContractFilterer) (*AppraiseFilterer, error) {
	contract, err := bindAppraise(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppraiseFilterer{contract: contract}, nil
}

// bindAppraise binds a generic wrapper to an already deployed contract.
func bindAppraise(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AppraiseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Appraise *AppraiseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Appraise.Contract.AppraiseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Appraise *AppraiseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Appraise.Contract.AppraiseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Appraise *AppraiseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Appraise.Contract.AppraiseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Appraise *AppraiseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Appraise.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Appraise *AppraiseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Appraise.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Appraise *AppraiseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Appraise.Contract.contract.Transact(opts, method, params...)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Appraise *AppraiseCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Appraise.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Appraise *AppraiseSession) AllAdmins() ([]common.Address, error) {
	return _Appraise.Contract.AllAdmins(&_Appraise.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Appraise *AppraiseCallerSession) AllAdmins() ([]common.Address, error) {
	return _Appraise.Contract.AllAdmins(&_Appraise.CallOpts)
}

// City is a free data retrieval call binding the contract method 0xfd0160fe.
//
// Solidity: function city() view returns(address)
func (_Appraise *AppraiseCaller) City(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Appraise.contract.Call(opts, &out, "city")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// City is a free data retrieval call binding the contract method 0xfd0160fe.
//
// Solidity: function city() view returns(address)
func (_Appraise *AppraiseSession) City() (common.Address, error) {
	return _Appraise.Contract.City(&_Appraise.CallOpts)
}

// City is a free data retrieval call binding the contract method 0xfd0160fe.
//
// Solidity: function city() view returns(address)
func (_Appraise *AppraiseCallerSession) City() (common.Address, error) {
	return _Appraise.Contract.City(&_Appraise.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Appraise *AppraiseCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Appraise.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Appraise *AppraiseSession) IsAdmin(account common.Address) (bool, error) {
	return _Appraise.Contract.IsAdmin(&_Appraise.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Appraise *AppraiseCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _Appraise.Contract.IsAdmin(&_Appraise.CallOpts, account)
}

// Pioneer is a free data retrieval call binding the contract method 0x7f8e1229.
//
// Solidity: function pioneer() view returns(address)
func (_Appraise *AppraiseCaller) Pioneer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Appraise.contract.Call(opts, &out, "pioneer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pioneer is a free data retrieval call binding the contract method 0x7f8e1229.
//
// Solidity: function pioneer() view returns(address)
func (_Appraise *AppraiseSession) Pioneer() (common.Address, error) {
	return _Appraise.Contract.Pioneer(&_Appraise.CallOpts)
}

// Pioneer is a free data retrieval call binding the contract method 0x7f8e1229.
//
// Solidity: function pioneer() view returns(address)
func (_Appraise *AppraiseCallerSession) Pioneer() (common.Address, error) {
	return _Appraise.Contract.Pioneer(&_Appraise.CallOpts)
}

// PioneerBatch is a free data retrieval call binding the contract method 0x3cbcbb1c.
//
// Solidity: function pioneerBatch(address ) view returns(uint256)
func (_Appraise *AppraiseCaller) PioneerBatch(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Appraise.contract.Call(opts, &out, "pioneerBatch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PioneerBatch is a free data retrieval call binding the contract method 0x3cbcbb1c.
//
// Solidity: function pioneerBatch(address ) view returns(uint256)
func (_Appraise *AppraiseSession) PioneerBatch(arg0 common.Address) (*big.Int, error) {
	return _Appraise.Contract.PioneerBatch(&_Appraise.CallOpts, arg0)
}

// PioneerBatch is a free data retrieval call binding the contract method 0x3cbcbb1c.
//
// Solidity: function pioneerBatch(address ) view returns(uint256)
func (_Appraise *AppraiseCallerSession) PioneerBatch(arg0 common.Address) (*big.Int, error) {
	return _Appraise.Contract.PioneerBatch(&_Appraise.CallOpts, arg0)
}

// PioneerMonthWeight is a free data retrieval call binding the contract method 0x4fc80c63.
//
// Solidity: function pioneerMonthWeight(address , uint256 ) view returns(uint256)
func (_Appraise *AppraiseCaller) PioneerMonthWeight(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Appraise.contract.Call(opts, &out, "pioneerMonthWeight", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PioneerMonthWeight is a free data retrieval call binding the contract method 0x4fc80c63.
//
// Solidity: function pioneerMonthWeight(address , uint256 ) view returns(uint256)
func (_Appraise *AppraiseSession) PioneerMonthWeight(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Appraise.Contract.PioneerMonthWeight(&_Appraise.CallOpts, arg0, arg1)
}

// PioneerMonthWeight is a free data retrieval call binding the contract method 0x4fc80c63.
//
// Solidity: function pioneerMonthWeight(address , uint256 ) view returns(uint256)
func (_Appraise *AppraiseCallerSession) PioneerMonthWeight(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Appraise.Contract.PioneerMonthWeight(&_Appraise.CallOpts, arg0, arg1)
}

// Star is a free data retrieval call binding the contract method 0x2c6dbb1c.
//
// Solidity: function star() view returns(address)
func (_Appraise *AppraiseCaller) Star(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Appraise.contract.Call(opts, &out, "star")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Star is a free data retrieval call binding the contract method 0x2c6dbb1c.
//
// Solidity: function star() view returns(address)
func (_Appraise *AppraiseSession) Star() (common.Address, error) {
	return _Appraise.Contract.Star(&_Appraise.CallOpts)
}

// Star is a free data retrieval call binding the contract method 0x2c6dbb1c.
//
// Solidity: function star() view returns(address)
func (_Appraise *AppraiseCallerSession) Star() (common.Address, error) {
	return _Appraise.Contract.Star(&_Appraise.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Appraise *AppraiseTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Appraise.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Appraise *AppraiseSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.AddAdmin(&_Appraise.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Appraise *AppraiseTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.AddAdmin(&_Appraise.TransactOpts, account)
}

// AdminSetCity is a paid mutator transaction binding the contract method 0x70d01861.
//
// Solidity: function adminSetCity(address cityAddress_) returns()
func (_Appraise *AppraiseTransactor) AdminSetCity(opts *bind.TransactOpts, cityAddress_ common.Address) (*types.Transaction, error) {
	return _Appraise.contract.Transact(opts, "adminSetCity", cityAddress_)
}

// AdminSetCity is a paid mutator transaction binding the contract method 0x70d01861.
//
// Solidity: function adminSetCity(address cityAddress_) returns()
func (_Appraise *AppraiseSession) AdminSetCity(cityAddress_ common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.AdminSetCity(&_Appraise.TransactOpts, cityAddress_)
}

// AdminSetCity is a paid mutator transaction binding the contract method 0x70d01861.
//
// Solidity: function adminSetCity(address cityAddress_) returns()
func (_Appraise *AppraiseTransactorSession) AdminSetCity(cityAddress_ common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.AdminSetCity(&_Appraise.TransactOpts, cityAddress_)
}

// AdminSetPioneer is a paid mutator transaction binding the contract method 0x6f61df99.
//
// Solidity: function adminSetPioneer(address cityPioneerAddress_) returns()
func (_Appraise *AppraiseTransactor) AdminSetPioneer(opts *bind.TransactOpts, cityPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _Appraise.contract.Transact(opts, "adminSetPioneer", cityPioneerAddress_)
}

// AdminSetPioneer is a paid mutator transaction binding the contract method 0x6f61df99.
//
// Solidity: function adminSetPioneer(address cityPioneerAddress_) returns()
func (_Appraise *AppraiseSession) AdminSetPioneer(cityPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.AdminSetPioneer(&_Appraise.TransactOpts, cityPioneerAddress_)
}

// AdminSetPioneer is a paid mutator transaction binding the contract method 0x6f61df99.
//
// Solidity: function adminSetPioneer(address cityPioneerAddress_) returns()
func (_Appraise *AppraiseTransactorSession) AdminSetPioneer(cityPioneerAddress_ common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.AdminSetPioneer(&_Appraise.TransactOpts, cityPioneerAddress_)
}

// AdminSetStar is a paid mutator transaction binding the contract method 0x66804f9b.
//
// Solidity: function adminSetStar(address starAddress_) returns()
func (_Appraise *AppraiseTransactor) AdminSetStar(opts *bind.TransactOpts, starAddress_ common.Address) (*types.Transaction, error) {
	return _Appraise.contract.Transact(opts, "adminSetStar", starAddress_)
}

// AdminSetStar is a paid mutator transaction binding the contract method 0x66804f9b.
//
// Solidity: function adminSetStar(address starAddress_) returns()
func (_Appraise *AppraiseSession) AdminSetStar(starAddress_ common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.AdminSetStar(&_Appraise.TransactOpts, starAddress_)
}

// AdminSetStar is a paid mutator transaction binding the contract method 0x66804f9b.
//
// Solidity: function adminSetStar(address starAddress_) returns()
func (_Appraise *AppraiseTransactorSession) AdminSetStar(starAddress_ common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.AdminSetStar(&_Appraise.TransactOpts, starAddress_)
}

// Appraise is a paid mutator transaction binding the contract method 0x68bdb3ed.
//
// Solidity: function appraise(address pioneerAddress_, bytes32 chengShiId_) returns()
func (_Appraise *AppraiseTransactor) Appraise(opts *bind.TransactOpts, pioneerAddress_ common.Address, chengShiId_ [32]byte) (*types.Transaction, error) {
	return _Appraise.contract.Transact(opts, "appraise", pioneerAddress_, chengShiId_)
}

// Appraise is a paid mutator transaction binding the contract method 0x68bdb3ed.
//
// Solidity: function appraise(address pioneerAddress_, bytes32 chengShiId_) returns()
func (_Appraise *AppraiseSession) Appraise(pioneerAddress_ common.Address, chengShiId_ [32]byte) (*types.Transaction, error) {
	return _Appraise.Contract.Appraise(&_Appraise.TransactOpts, pioneerAddress_, chengShiId_)
}

// Appraise is a paid mutator transaction binding the contract method 0x68bdb3ed.
//
// Solidity: function appraise(address pioneerAddress_, bytes32 chengShiId_) returns()
func (_Appraise *AppraiseTransactorSession) Appraise(pioneerAddress_ common.Address, chengShiId_ [32]byte) (*types.Transaction, error) {
	return _Appraise.Contract.Appraise(&_Appraise.TransactOpts, pioneerAddress_, chengShiId_)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Appraise *AppraiseTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _Appraise.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Appraise *AppraiseSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.BatchAddAdmin(&_Appraise.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Appraise *AppraiseTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.BatchAddAdmin(&_Appraise.TransactOpts, amounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Appraise *AppraiseTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Appraise.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Appraise *AppraiseSession) Initialize() (*types.Transaction, error) {
	return _Appraise.Contract.Initialize(&_Appraise.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Appraise *AppraiseTransactorSession) Initialize() (*types.Transaction, error) {
	return _Appraise.Contract.Initialize(&_Appraise.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Appraise *AppraiseTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Appraise.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Appraise *AppraiseSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.RemoveAdmin(&_Appraise.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Appraise *AppraiseTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _Appraise.Contract.RemoveAdmin(&_Appraise.TransactOpts, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Appraise *AppraiseTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Appraise.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Appraise *AppraiseSession) RenounceAdmin() (*types.Transaction, error) {
	return _Appraise.Contract.RenounceAdmin(&_Appraise.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Appraise *AppraiseTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _Appraise.Contract.RenounceAdmin(&_Appraise.TransactOpts)
}

// AppraiseAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the Appraise contract.
type AppraiseAdminAddedIterator struct {
	Event *AppraiseAdminAdded // Event containing the contract specifics and raw log

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
func (it *AppraiseAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppraiseAdminAdded)
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
		it.Event = new(AppraiseAdminAdded)
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
func (it *AppraiseAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppraiseAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppraiseAdminAdded represents a AdminAdded event raised by the Appraise contract.
type AppraiseAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Appraise *AppraiseFilterer) FilterAdminAdded(opts *bind.FilterOpts, account []common.Address) (*AppraiseAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Appraise.contract.FilterLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &AppraiseAdminAddedIterator{contract: _Appraise.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Appraise *AppraiseFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *AppraiseAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Appraise.contract.WatchLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppraiseAdminAdded)
				if err := _Appraise.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// ParseAdminAdded is a log parse operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Appraise *AppraiseFilterer) ParseAdminAdded(log types.Log) (*AppraiseAdminAdded, error) {
	event := new(AppraiseAdminAdded)
	if err := _Appraise.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppraiseAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Appraise contract.
type AppraiseAdminRemovedIterator struct {
	Event *AppraiseAdminRemoved // Event containing the contract specifics and raw log

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
func (it *AppraiseAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppraiseAdminRemoved)
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
		it.Event = new(AppraiseAdminRemoved)
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
func (it *AppraiseAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppraiseAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppraiseAdminRemoved represents a AdminRemoved event raised by the Appraise contract.
type AppraiseAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Appraise *AppraiseFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*AppraiseAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Appraise.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AppraiseAdminRemovedIterator{contract: _Appraise.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Appraise *AppraiseFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *AppraiseAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Appraise.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppraiseAdminRemoved)
				if err := _Appraise.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Appraise *AppraiseFilterer) ParseAdminRemoved(log types.Log) (*AppraiseAdminRemoved, error) {
	event := new(AppraiseAdminRemoved)
	if err := _Appraise.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppraiseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Appraise contract.
type AppraiseInitializedIterator struct {
	Event *AppraiseInitialized // Event containing the contract specifics and raw log

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
func (it *AppraiseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppraiseInitialized)
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
		it.Event = new(AppraiseInitialized)
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
func (it *AppraiseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppraiseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppraiseInitialized represents a Initialized event raised by the Appraise contract.
type AppraiseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Appraise *AppraiseFilterer) FilterInitialized(opts *bind.FilterOpts) (*AppraiseInitializedIterator, error) {

	logs, sub, err := _Appraise.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AppraiseInitializedIterator{contract: _Appraise.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Appraise *AppraiseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AppraiseInitialized) (event.Subscription, error) {

	logs, sub, err := _Appraise.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppraiseInitialized)
				if err := _Appraise.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Appraise *AppraiseFilterer) ParseInitialized(log types.Log) (*AppraiseInitialized, error) {
	event := new(AppraiseInitialized)
	if err := _Appraise.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}