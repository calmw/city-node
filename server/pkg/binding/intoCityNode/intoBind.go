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

// IntoBindUIDAddress is an auto generated low-level Go binding around an user-defined struct.
type IntoBindUIDAddress struct {
	Account common.Address
	Uid     string
}

// IntoBindMetaData contains all meta data concerning the IntoBind contract.
var IntoBindMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"}],\"name\":\"Bind\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"BindMain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"}],\"name\":\"BindRelation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"UnBind\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"UnBindMain\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"}],\"name\":\"bind\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bindAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"inv\",\"type\":\"address\"}],\"name\":\"bindAndRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"bindMain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"bindMainAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inv\",\"type\":\"address\"}],\"name\":\"bindRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inv\",\"type\":\"address\"}],\"name\":\"bindUIDAndRelationWithAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bindUid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"bindWithAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_str1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_str2\",\"type\":\"string\"}],\"name\":\"checkBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"getBindAddrs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"getBindStatusWithAddress\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"uid\",\"type\":\"string[]\"}],\"name\":\"getMainBindAddressMany\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"internalType\":\"structIntoBind.UIDAddress[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"uid\",\"type\":\"string[]\"}],\"name\":\"getMainBindAddressManyReturnAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"getUidsWithAddress\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBind\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"mainAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRelationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"unBind\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"name\":\"unBindMainAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"unBindMainAddressWithAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"unBindWithAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IntoBindABI is the input ABI used to generate the binding from.
// Deprecated: Use IntoBindMetaData.ABI instead.
var IntoBindABI = IntoBindMetaData.ABI

// IntoBind is an auto generated Go binding around an Ethereum contract.
type IntoBind struct {
	IntoBindCaller     // Read-only binding to the contract
	IntoBindTransactor // Write-only binding to the contract
	IntoBindFilterer   // Log filterer for contract events
}

// IntoBindCaller is an auto generated read-only Go binding around an Ethereum contract.
type IntoBindCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntoBindTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IntoBindTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntoBindFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IntoBindFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntoBindSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IntoBindSession struct {
	Contract     *IntoBind         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IntoBindCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IntoBindCallerSession struct {
	Contract *IntoBindCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IntoBindTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IntoBindTransactorSession struct {
	Contract     *IntoBindTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IntoBindRaw is an auto generated low-level Go binding around an Ethereum contract.
type IntoBindRaw struct {
	Contract *IntoBind // Generic contract binding to access the raw methods on
}

// IntoBindCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IntoBindCallerRaw struct {
	Contract *IntoBindCaller // Generic read-only contract binding to access the raw methods on
}

// IntoBindTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IntoBindTransactorRaw struct {
	Contract *IntoBindTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIntoBind creates a new instance of IntoBind, bound to a specific deployed contract.
func NewIntoBind(address common.Address, backend bind.ContractBackend) (*IntoBind, error) {
	contract, err := bindIntoBind(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IntoBind{IntoBindCaller: IntoBindCaller{contract: contract}, IntoBindTransactor: IntoBindTransactor{contract: contract}, IntoBindFilterer: IntoBindFilterer{contract: contract}}, nil
}

// NewIntoBindCaller creates a new read-only instance of IntoBind, bound to a specific deployed contract.
func NewIntoBindCaller(address common.Address, caller bind.ContractCaller) (*IntoBindCaller, error) {
	contract, err := bindIntoBind(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IntoBindCaller{contract: contract}, nil
}

// NewIntoBindTransactor creates a new write-only instance of IntoBind, bound to a specific deployed contract.
func NewIntoBindTransactor(address common.Address, transactor bind.ContractTransactor) (*IntoBindTransactor, error) {
	contract, err := bindIntoBind(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IntoBindTransactor{contract: contract}, nil
}

// NewIntoBindFilterer creates a new log filterer instance of IntoBind, bound to a specific deployed contract.
func NewIntoBindFilterer(address common.Address, filterer bind.ContractFilterer) (*IntoBindFilterer, error) {
	contract, err := bindIntoBind(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IntoBindFilterer{contract: contract}, nil
}

// bindIntoBind binds a generic wrapper to an already deployed contract.
func bindIntoBind(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IntoBindMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IntoBind *IntoBindRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IntoBind.Contract.IntoBindCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IntoBind *IntoBindRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntoBind.Contract.IntoBindTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IntoBind *IntoBindRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IntoBind.Contract.IntoBindTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IntoBind *IntoBindCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IntoBind.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IntoBind *IntoBindTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntoBind.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IntoBind *IntoBindTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IntoBind.Contract.contract.Transact(opts, method, params...)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_IntoBind *IntoBindCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_IntoBind *IntoBindSession) AllAdmins() ([]common.Address, error) {
	return _IntoBind.Contract.AllAdmins(&_IntoBind.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_IntoBind *IntoBindCallerSession) AllAdmins() ([]common.Address, error) {
	return _IntoBind.Contract.AllAdmins(&_IntoBind.CallOpts)
}

// BindAddrs is a free data retrieval call binding the contract method 0x53b333a8.
//
// Solidity: function bindAddrs(string , uint256 ) view returns(address)
func (_IntoBind *IntoBindCaller) BindAddrs(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "bindAddrs", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BindAddrs is a free data retrieval call binding the contract method 0x53b333a8.
//
// Solidity: function bindAddrs(string , uint256 ) view returns(address)
func (_IntoBind *IntoBindSession) BindAddrs(arg0 string, arg1 *big.Int) (common.Address, error) {
	return _IntoBind.Contract.BindAddrs(&_IntoBind.CallOpts, arg0, arg1)
}

// BindAddrs is a free data retrieval call binding the contract method 0x53b333a8.
//
// Solidity: function bindAddrs(string , uint256 ) view returns(address)
func (_IntoBind *IntoBindCallerSession) BindAddrs(arg0 string, arg1 *big.Int) (common.Address, error) {
	return _IntoBind.Contract.BindAddrs(&_IntoBind.CallOpts, arg0, arg1)
}

// BindMain is a free data retrieval call binding the contract method 0xa6abf450.
//
// Solidity: function bindMain(string ) view returns(bool)
func (_IntoBind *IntoBindCaller) BindMain(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "bindMain", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BindMain is a free data retrieval call binding the contract method 0xa6abf450.
//
// Solidity: function bindMain(string ) view returns(bool)
func (_IntoBind *IntoBindSession) BindMain(arg0 string) (bool, error) {
	return _IntoBind.Contract.BindMain(&_IntoBind.CallOpts, arg0)
}

// BindMain is a free data retrieval call binding the contract method 0xa6abf450.
//
// Solidity: function bindMain(string ) view returns(bool)
func (_IntoBind *IntoBindCallerSession) BindMain(arg0 string) (bool, error) {
	return _IntoBind.Contract.BindMain(&_IntoBind.CallOpts, arg0)
}

// BindUid is a free data retrieval call binding the contract method 0xdf4431d0.
//
// Solidity: function bindUid(address ) view returns(string)
func (_IntoBind *IntoBindCaller) BindUid(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "bindUid", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BindUid is a free data retrieval call binding the contract method 0xdf4431d0.
//
// Solidity: function bindUid(address ) view returns(string)
func (_IntoBind *IntoBindSession) BindUid(arg0 common.Address) (string, error) {
	return _IntoBind.Contract.BindUid(&_IntoBind.CallOpts, arg0)
}

// BindUid is a free data retrieval call binding the contract method 0xdf4431d0.
//
// Solidity: function bindUid(address ) view returns(string)
func (_IntoBind *IntoBindCallerSession) BindUid(arg0 common.Address) (string, error) {
	return _IntoBind.Contract.BindUid(&_IntoBind.CallOpts, arg0)
}

// CheckBool is a free data retrieval call binding the contract method 0x565f21ce.
//
// Solidity: function checkBool(string _str1, string _str2) pure returns(bool)
func (_IntoBind *IntoBindCaller) CheckBool(opts *bind.CallOpts, _str1 string, _str2 string) (bool, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "checkBool", _str1, _str2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckBool is a free data retrieval call binding the contract method 0x565f21ce.
//
// Solidity: function checkBool(string _str1, string _str2) pure returns(bool)
func (_IntoBind *IntoBindSession) CheckBool(_str1 string, _str2 string) (bool, error) {
	return _IntoBind.Contract.CheckBool(&_IntoBind.CallOpts, _str1, _str2)
}

// CheckBool is a free data retrieval call binding the contract method 0x565f21ce.
//
// Solidity: function checkBool(string _str1, string _str2) pure returns(bool)
func (_IntoBind *IntoBindCallerSession) CheckBool(_str1 string, _str2 string) (bool, error) {
	return _IntoBind.Contract.CheckBool(&_IntoBind.CallOpts, _str1, _str2)
}

// GetBindAddrs is a free data retrieval call binding the contract method 0xfe39cb15.
//
// Solidity: function getBindAddrs(string uid) view returns(address[])
func (_IntoBind *IntoBindCaller) GetBindAddrs(opts *bind.CallOpts, uid string) ([]common.Address, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "getBindAddrs", uid)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBindAddrs is a free data retrieval call binding the contract method 0xfe39cb15.
//
// Solidity: function getBindAddrs(string uid) view returns(address[])
func (_IntoBind *IntoBindSession) GetBindAddrs(uid string) ([]common.Address, error) {
	return _IntoBind.Contract.GetBindAddrs(&_IntoBind.CallOpts, uid)
}

// GetBindAddrs is a free data retrieval call binding the contract method 0xfe39cb15.
//
// Solidity: function getBindAddrs(string uid) view returns(address[])
func (_IntoBind *IntoBindCallerSession) GetBindAddrs(uid string) ([]common.Address, error) {
	return _IntoBind.Contract.GetBindAddrs(&_IntoBind.CallOpts, uid)
}

// GetBindStatusWithAddress is a free data retrieval call binding the contract method 0xfba53d94.
//
// Solidity: function getBindStatusWithAddress(address[] addrs) view returns(bool[])
func (_IntoBind *IntoBindCaller) GetBindStatusWithAddress(opts *bind.CallOpts, addrs []common.Address) ([]bool, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "getBindStatusWithAddress", addrs)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// GetBindStatusWithAddress is a free data retrieval call binding the contract method 0xfba53d94.
//
// Solidity: function getBindStatusWithAddress(address[] addrs) view returns(bool[])
func (_IntoBind *IntoBindSession) GetBindStatusWithAddress(addrs []common.Address) ([]bool, error) {
	return _IntoBind.Contract.GetBindStatusWithAddress(&_IntoBind.CallOpts, addrs)
}

// GetBindStatusWithAddress is a free data retrieval call binding the contract method 0xfba53d94.
//
// Solidity: function getBindStatusWithAddress(address[] addrs) view returns(bool[])
func (_IntoBind *IntoBindCallerSession) GetBindStatusWithAddress(addrs []common.Address) ([]bool, error) {
	return _IntoBind.Contract.GetBindStatusWithAddress(&_IntoBind.CallOpts, addrs)
}

// GetMainBindAddressMany is a free data retrieval call binding the contract method 0x0e183d2d.
//
// Solidity: function getMainBindAddressMany(string[] uid) view returns((address,string)[])
func (_IntoBind *IntoBindCaller) GetMainBindAddressMany(opts *bind.CallOpts, uid []string) ([]IntoBindUIDAddress, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "getMainBindAddressMany", uid)

	if err != nil {
		return *new([]IntoBindUIDAddress), err
	}

	out0 := *abi.ConvertType(out[0], new([]IntoBindUIDAddress)).(*[]IntoBindUIDAddress)

	return out0, err

}

// GetMainBindAddressMany is a free data retrieval call binding the contract method 0x0e183d2d.
//
// Solidity: function getMainBindAddressMany(string[] uid) view returns((address,string)[])
func (_IntoBind *IntoBindSession) GetMainBindAddressMany(uid []string) ([]IntoBindUIDAddress, error) {
	return _IntoBind.Contract.GetMainBindAddressMany(&_IntoBind.CallOpts, uid)
}

// GetMainBindAddressMany is a free data retrieval call binding the contract method 0x0e183d2d.
//
// Solidity: function getMainBindAddressMany(string[] uid) view returns((address,string)[])
func (_IntoBind *IntoBindCallerSession) GetMainBindAddressMany(uid []string) ([]IntoBindUIDAddress, error) {
	return _IntoBind.Contract.GetMainBindAddressMany(&_IntoBind.CallOpts, uid)
}

// GetMainBindAddressManyReturnAddress is a free data retrieval call binding the contract method 0x9173f213.
//
// Solidity: function getMainBindAddressManyReturnAddress(string[] uid) view returns(address[])
func (_IntoBind *IntoBindCaller) GetMainBindAddressManyReturnAddress(opts *bind.CallOpts, uid []string) ([]common.Address, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "getMainBindAddressManyReturnAddress", uid)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMainBindAddressManyReturnAddress is a free data retrieval call binding the contract method 0x9173f213.
//
// Solidity: function getMainBindAddressManyReturnAddress(string[] uid) view returns(address[])
func (_IntoBind *IntoBindSession) GetMainBindAddressManyReturnAddress(uid []string) ([]common.Address, error) {
	return _IntoBind.Contract.GetMainBindAddressManyReturnAddress(&_IntoBind.CallOpts, uid)
}

// GetMainBindAddressManyReturnAddress is a free data retrieval call binding the contract method 0x9173f213.
//
// Solidity: function getMainBindAddressManyReturnAddress(string[] uid) view returns(address[])
func (_IntoBind *IntoBindCallerSession) GetMainBindAddressManyReturnAddress(uid []string) ([]common.Address, error) {
	return _IntoBind.Contract.GetMainBindAddressManyReturnAddress(&_IntoBind.CallOpts, uid)
}

// GetUidsWithAddress is a free data retrieval call binding the contract method 0x663b928c.
//
// Solidity: function getUidsWithAddress(address[] addrs) view returns(string[])
func (_IntoBind *IntoBindCaller) GetUidsWithAddress(opts *bind.CallOpts, addrs []common.Address) ([]string, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "getUidsWithAddress", addrs)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetUidsWithAddress is a free data retrieval call binding the contract method 0x663b928c.
//
// Solidity: function getUidsWithAddress(address[] addrs) view returns(string[])
func (_IntoBind *IntoBindSession) GetUidsWithAddress(addrs []common.Address) ([]string, error) {
	return _IntoBind.Contract.GetUidsWithAddress(&_IntoBind.CallOpts, addrs)
}

// GetUidsWithAddress is a free data retrieval call binding the contract method 0x663b928c.
//
// Solidity: function getUidsWithAddress(address[] addrs) view returns(string[])
func (_IntoBind *IntoBindCallerSession) GetUidsWithAddress(addrs []common.Address) ([]string, error) {
	return _IntoBind.Contract.GetUidsWithAddress(&_IntoBind.CallOpts, addrs)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_IntoBind *IntoBindCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_IntoBind *IntoBindSession) IsAdmin(account common.Address) (bool, error) {
	return _IntoBind.Contract.IsAdmin(&_IntoBind.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_IntoBind *IntoBindCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _IntoBind.Contract.IsAdmin(&_IntoBind.CallOpts, account)
}

// IsBind is a free data retrieval call binding the contract method 0xf4b35538.
//
// Solidity: function isBind(address ) view returns(bool)
func (_IntoBind *IntoBindCaller) IsBind(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "isBind", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBind is a free data retrieval call binding the contract method 0xf4b35538.
//
// Solidity: function isBind(address ) view returns(bool)
func (_IntoBind *IntoBindSession) IsBind(arg0 common.Address) (bool, error) {
	return _IntoBind.Contract.IsBind(&_IntoBind.CallOpts, arg0)
}

// IsBind is a free data retrieval call binding the contract method 0xf4b35538.
//
// Solidity: function isBind(address ) view returns(bool)
func (_IntoBind *IntoBindCallerSession) IsBind(arg0 common.Address) (bool, error) {
	return _IntoBind.Contract.IsBind(&_IntoBind.CallOpts, arg0)
}

// MainAddr is a free data retrieval call binding the contract method 0x40d817f9.
//
// Solidity: function mainAddr(string ) view returns(address)
func (_IntoBind *IntoBindCaller) MainAddr(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _IntoBind.contract.Call(opts, &out, "mainAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainAddr is a free data retrieval call binding the contract method 0x40d817f9.
//
// Solidity: function mainAddr(string ) view returns(address)
func (_IntoBind *IntoBindSession) MainAddr(arg0 string) (common.Address, error) {
	return _IntoBind.Contract.MainAddr(&_IntoBind.CallOpts, arg0)
}

// MainAddr is a free data retrieval call binding the contract method 0x40d817f9.
//
// Solidity: function mainAddr(string ) view returns(address)
func (_IntoBind *IntoBindCallerSession) MainAddr(arg0 string) (common.Address, error) {
	return _IntoBind.Contract.MainAddr(&_IntoBind.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_IntoBind *IntoBindTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_IntoBind *IntoBindSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.AddAdmin(&_IntoBind.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_IntoBind *IntoBindTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.AddAdmin(&_IntoBind.TransactOpts, account)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_IntoBind *IntoBindTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_IntoBind *IntoBindSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.BatchAddAdmin(&_IntoBind.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_IntoBind *IntoBindTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.BatchAddAdmin(&_IntoBind.TransactOpts, amounts)
}

// Bind is a paid mutator transaction binding the contract method 0xc2797e7f.
//
// Solidity: function bind(string uid, bool isMain) returns()
func (_IntoBind *IntoBindTransactor) Bind(opts *bind.TransactOpts, uid string, isMain bool) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "bind", uid, isMain)
}

// Bind is a paid mutator transaction binding the contract method 0xc2797e7f.
//
// Solidity: function bind(string uid, bool isMain) returns()
func (_IntoBind *IntoBindSession) Bind(uid string, isMain bool) (*types.Transaction, error) {
	return _IntoBind.Contract.Bind(&_IntoBind.TransactOpts, uid, isMain)
}

// Bind is a paid mutator transaction binding the contract method 0xc2797e7f.
//
// Solidity: function bind(string uid, bool isMain) returns()
func (_IntoBind *IntoBindTransactorSession) Bind(uid string, isMain bool) (*types.Transaction, error) {
	return _IntoBind.Contract.Bind(&_IntoBind.TransactOpts, uid, isMain)
}

// BindAndRelation is a paid mutator transaction binding the contract method 0x310f1da7.
//
// Solidity: function bindAndRelation(string uid, bool isMain, address inv) returns()
func (_IntoBind *IntoBindTransactor) BindAndRelation(opts *bind.TransactOpts, uid string, isMain bool, inv common.Address) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "bindAndRelation", uid, isMain, inv)
}

// BindAndRelation is a paid mutator transaction binding the contract method 0x310f1da7.
//
// Solidity: function bindAndRelation(string uid, bool isMain, address inv) returns()
func (_IntoBind *IntoBindSession) BindAndRelation(uid string, isMain bool, inv common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.BindAndRelation(&_IntoBind.TransactOpts, uid, isMain, inv)
}

// BindAndRelation is a paid mutator transaction binding the contract method 0x310f1da7.
//
// Solidity: function bindAndRelation(string uid, bool isMain, address inv) returns()
func (_IntoBind *IntoBindTransactorSession) BindAndRelation(uid string, isMain bool, inv common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.BindAndRelation(&_IntoBind.TransactOpts, uid, isMain, inv)
}

// BindMainAddress is a paid mutator transaction binding the contract method 0x0e7eac0a.
//
// Solidity: function bindMainAddress(string uid) returns()
func (_IntoBind *IntoBindTransactor) BindMainAddress(opts *bind.TransactOpts, uid string) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "bindMainAddress", uid)
}

// BindMainAddress is a paid mutator transaction binding the contract method 0x0e7eac0a.
//
// Solidity: function bindMainAddress(string uid) returns()
func (_IntoBind *IntoBindSession) BindMainAddress(uid string) (*types.Transaction, error) {
	return _IntoBind.Contract.BindMainAddress(&_IntoBind.TransactOpts, uid)
}

// BindMainAddress is a paid mutator transaction binding the contract method 0x0e7eac0a.
//
// Solidity: function bindMainAddress(string uid) returns()
func (_IntoBind *IntoBindTransactorSession) BindMainAddress(uid string) (*types.Transaction, error) {
	return _IntoBind.Contract.BindMainAddress(&_IntoBind.TransactOpts, uid)
}

// BindRelation is a paid mutator transaction binding the contract method 0x9d56b80a.
//
// Solidity: function bindRelation(address inv) returns()
func (_IntoBind *IntoBindTransactor) BindRelation(opts *bind.TransactOpts, inv common.Address) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "bindRelation", inv)
}

// BindRelation is a paid mutator transaction binding the contract method 0x9d56b80a.
//
// Solidity: function bindRelation(address inv) returns()
func (_IntoBind *IntoBindSession) BindRelation(inv common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.BindRelation(&_IntoBind.TransactOpts, inv)
}

// BindRelation is a paid mutator transaction binding the contract method 0x9d56b80a.
//
// Solidity: function bindRelation(address inv) returns()
func (_IntoBind *IntoBindTransactorSession) BindRelation(inv common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.BindRelation(&_IntoBind.TransactOpts, inv)
}

// BindUIDAndRelationWithAdmin is a paid mutator transaction binding the contract method 0xf42266e3.
//
// Solidity: function bindUIDAndRelationWithAdmin(string uid, bool isMain, address sender, address inv) returns()
func (_IntoBind *IntoBindTransactor) BindUIDAndRelationWithAdmin(opts *bind.TransactOpts, uid string, isMain bool, sender common.Address, inv common.Address) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "bindUIDAndRelationWithAdmin", uid, isMain, sender, inv)
}

// BindUIDAndRelationWithAdmin is a paid mutator transaction binding the contract method 0xf42266e3.
//
// Solidity: function bindUIDAndRelationWithAdmin(string uid, bool isMain, address sender, address inv) returns()
func (_IntoBind *IntoBindSession) BindUIDAndRelationWithAdmin(uid string, isMain bool, sender common.Address, inv common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.BindUIDAndRelationWithAdmin(&_IntoBind.TransactOpts, uid, isMain, sender, inv)
}

// BindUIDAndRelationWithAdmin is a paid mutator transaction binding the contract method 0xf42266e3.
//
// Solidity: function bindUIDAndRelationWithAdmin(string uid, bool isMain, address sender, address inv) returns()
func (_IntoBind *IntoBindTransactorSession) BindUIDAndRelationWithAdmin(uid string, isMain bool, sender common.Address, inv common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.BindUIDAndRelationWithAdmin(&_IntoBind.TransactOpts, uid, isMain, sender, inv)
}

// BindWithAdmin is a paid mutator transaction binding the contract method 0x48743984.
//
// Solidity: function bindWithAdmin(string uid, bool isMain, address sender) returns()
func (_IntoBind *IntoBindTransactor) BindWithAdmin(opts *bind.TransactOpts, uid string, isMain bool, sender common.Address) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "bindWithAdmin", uid, isMain, sender)
}

// BindWithAdmin is a paid mutator transaction binding the contract method 0x48743984.
//
// Solidity: function bindWithAdmin(string uid, bool isMain, address sender) returns()
func (_IntoBind *IntoBindSession) BindWithAdmin(uid string, isMain bool, sender common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.BindWithAdmin(&_IntoBind.TransactOpts, uid, isMain, sender)
}

// BindWithAdmin is a paid mutator transaction binding the contract method 0x48743984.
//
// Solidity: function bindWithAdmin(string uid, bool isMain, address sender) returns()
func (_IntoBind *IntoBindTransactorSession) BindWithAdmin(uid string, isMain bool, sender common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.BindWithAdmin(&_IntoBind.TransactOpts, uid, isMain, sender)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IntoBind *IntoBindTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IntoBind *IntoBindSession) Initialize() (*types.Transaction, error) {
	return _IntoBind.Contract.Initialize(&_IntoBind.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IntoBind *IntoBindTransactorSession) Initialize() (*types.Transaction, error) {
	return _IntoBind.Contract.Initialize(&_IntoBind.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_IntoBind *IntoBindTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_IntoBind *IntoBindSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.RemoveAdmin(&_IntoBind.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_IntoBind *IntoBindTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.RemoveAdmin(&_IntoBind.TransactOpts, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_IntoBind *IntoBindTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_IntoBind *IntoBindSession) RenounceAdmin() (*types.Transaction, error) {
	return _IntoBind.Contract.RenounceAdmin(&_IntoBind.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_IntoBind *IntoBindTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _IntoBind.Contract.RenounceAdmin(&_IntoBind.TransactOpts)
}

// SetRelationAddress is a paid mutator transaction binding the contract method 0xa3fca939.
//
// Solidity: function setRelationAddress(address _address) returns()
func (_IntoBind *IntoBindTransactor) SetRelationAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "setRelationAddress", _address)
}

// SetRelationAddress is a paid mutator transaction binding the contract method 0xa3fca939.
//
// Solidity: function setRelationAddress(address _address) returns()
func (_IntoBind *IntoBindSession) SetRelationAddress(_address common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.SetRelationAddress(&_IntoBind.TransactOpts, _address)
}

// SetRelationAddress is a paid mutator transaction binding the contract method 0xa3fca939.
//
// Solidity: function setRelationAddress(address _address) returns()
func (_IntoBind *IntoBindTransactorSession) SetRelationAddress(_address common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.SetRelationAddress(&_IntoBind.TransactOpts, _address)
}

// UnBind is a paid mutator transaction binding the contract method 0x907326ac.
//
// Solidity: function unBind(string uid) returns()
func (_IntoBind *IntoBindTransactor) UnBind(opts *bind.TransactOpts, uid string) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "unBind", uid)
}

// UnBind is a paid mutator transaction binding the contract method 0x907326ac.
//
// Solidity: function unBind(string uid) returns()
func (_IntoBind *IntoBindSession) UnBind(uid string) (*types.Transaction, error) {
	return _IntoBind.Contract.UnBind(&_IntoBind.TransactOpts, uid)
}

// UnBind is a paid mutator transaction binding the contract method 0x907326ac.
//
// Solidity: function unBind(string uid) returns()
func (_IntoBind *IntoBindTransactorSession) UnBind(uid string) (*types.Transaction, error) {
	return _IntoBind.Contract.UnBind(&_IntoBind.TransactOpts, uid)
}

// UnBindMainAddress is a paid mutator transaction binding the contract method 0xac1607b0.
//
// Solidity: function unBindMainAddress(string uid) returns()
func (_IntoBind *IntoBindTransactor) UnBindMainAddress(opts *bind.TransactOpts, uid string) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "unBindMainAddress", uid)
}

// UnBindMainAddress is a paid mutator transaction binding the contract method 0xac1607b0.
//
// Solidity: function unBindMainAddress(string uid) returns()
func (_IntoBind *IntoBindSession) UnBindMainAddress(uid string) (*types.Transaction, error) {
	return _IntoBind.Contract.UnBindMainAddress(&_IntoBind.TransactOpts, uid)
}

// UnBindMainAddress is a paid mutator transaction binding the contract method 0xac1607b0.
//
// Solidity: function unBindMainAddress(string uid) returns()
func (_IntoBind *IntoBindTransactorSession) UnBindMainAddress(uid string) (*types.Transaction, error) {
	return _IntoBind.Contract.UnBindMainAddress(&_IntoBind.TransactOpts, uid)
}

// UnBindMainAddressWithAdmin is a paid mutator transaction binding the contract method 0x970f26aa.
//
// Solidity: function unBindMainAddressWithAdmin(string uid, address sender) returns()
func (_IntoBind *IntoBindTransactor) UnBindMainAddressWithAdmin(opts *bind.TransactOpts, uid string, sender common.Address) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "unBindMainAddressWithAdmin", uid, sender)
}

// UnBindMainAddressWithAdmin is a paid mutator transaction binding the contract method 0x970f26aa.
//
// Solidity: function unBindMainAddressWithAdmin(string uid, address sender) returns()
func (_IntoBind *IntoBindSession) UnBindMainAddressWithAdmin(uid string, sender common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.UnBindMainAddressWithAdmin(&_IntoBind.TransactOpts, uid, sender)
}

// UnBindMainAddressWithAdmin is a paid mutator transaction binding the contract method 0x970f26aa.
//
// Solidity: function unBindMainAddressWithAdmin(string uid, address sender) returns()
func (_IntoBind *IntoBindTransactorSession) UnBindMainAddressWithAdmin(uid string, sender common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.UnBindMainAddressWithAdmin(&_IntoBind.TransactOpts, uid, sender)
}

// UnBindWithAdmin is a paid mutator transaction binding the contract method 0xafd43ba1.
//
// Solidity: function unBindWithAdmin(string uid, address sender) returns()
func (_IntoBind *IntoBindTransactor) UnBindWithAdmin(opts *bind.TransactOpts, uid string, sender common.Address) (*types.Transaction, error) {
	return _IntoBind.contract.Transact(opts, "unBindWithAdmin", uid, sender)
}

// UnBindWithAdmin is a paid mutator transaction binding the contract method 0xafd43ba1.
//
// Solidity: function unBindWithAdmin(string uid, address sender) returns()
func (_IntoBind *IntoBindSession) UnBindWithAdmin(uid string, sender common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.UnBindWithAdmin(&_IntoBind.TransactOpts, uid, sender)
}

// UnBindWithAdmin is a paid mutator transaction binding the contract method 0xafd43ba1.
//
// Solidity: function unBindWithAdmin(string uid, address sender) returns()
func (_IntoBind *IntoBindTransactorSession) UnBindWithAdmin(uid string, sender common.Address) (*types.Transaction, error) {
	return _IntoBind.Contract.UnBindWithAdmin(&_IntoBind.TransactOpts, uid, sender)
}

// IntoBindAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the IntoBind contract.
type IntoBindAdminAddedIterator struct {
	Event *IntoBindAdminAdded // Event containing the contract specifics and raw log

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
func (it *IntoBindAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntoBindAdminAdded)
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
		it.Event = new(IntoBindAdminAdded)
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
func (it *IntoBindAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntoBindAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntoBindAdminAdded represents a AdminAdded event raised by the IntoBind contract.
type IntoBindAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_IntoBind *IntoBindFilterer) FilterAdminAdded(opts *bind.FilterOpts, account []common.Address) (*IntoBindAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IntoBind.contract.FilterLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &IntoBindAdminAddedIterator{contract: _IntoBind.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_IntoBind *IntoBindFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *IntoBindAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IntoBind.contract.WatchLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntoBindAdminAdded)
				if err := _IntoBind.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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
func (_IntoBind *IntoBindFilterer) ParseAdminAdded(log types.Log) (*IntoBindAdminAdded, error) {
	event := new(IntoBindAdminAdded)
	if err := _IntoBind.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntoBindAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the IntoBind contract.
type IntoBindAdminRemovedIterator struct {
	Event *IntoBindAdminRemoved // Event containing the contract specifics and raw log

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
func (it *IntoBindAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntoBindAdminRemoved)
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
		it.Event = new(IntoBindAdminRemoved)
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
func (it *IntoBindAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntoBindAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntoBindAdminRemoved represents a AdminRemoved event raised by the IntoBind contract.
type IntoBindAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_IntoBind *IntoBindFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*IntoBindAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IntoBind.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &IntoBindAdminRemovedIterator{contract: _IntoBind.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_IntoBind *IntoBindFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *IntoBindAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IntoBind.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntoBindAdminRemoved)
				if err := _IntoBind.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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
func (_IntoBind *IntoBindFilterer) ParseAdminRemoved(log types.Log) (*IntoBindAdminRemoved, error) {
	event := new(IntoBindAdminRemoved)
	if err := _IntoBind.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntoBindBindIterator is returned from FilterBind and is used to iterate over the raw logs and unpacked data for Bind events raised by the IntoBind contract.
type IntoBindBindIterator struct {
	Event *IntoBindBind // Event containing the contract specifics and raw log

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
func (it *IntoBindBindIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntoBindBind)
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
		it.Event = new(IntoBindBind)
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
func (it *IntoBindBindIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntoBindBindIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntoBindBind represents a Bind event raised by the IntoBind contract.
type IntoBindBind struct {
	From   common.Address
	Uid    string
	IsMain bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBind is a free log retrieval operation binding the contract event 0xef2d9da059041adb88becbf698f278eabeccb23e046f0ae367dd57be0cce7b22.
//
// Solidity: event Bind(address from, string uid, bool isMain)
func (_IntoBind *IntoBindFilterer) FilterBind(opts *bind.FilterOpts) (*IntoBindBindIterator, error) {

	logs, sub, err := _IntoBind.contract.FilterLogs(opts, "Bind")
	if err != nil {
		return nil, err
	}
	return &IntoBindBindIterator{contract: _IntoBind.contract, event: "Bind", logs: logs, sub: sub}, nil
}

// WatchBind is a free log subscription operation binding the contract event 0xef2d9da059041adb88becbf698f278eabeccb23e046f0ae367dd57be0cce7b22.
//
// Solidity: event Bind(address from, string uid, bool isMain)
func (_IntoBind *IntoBindFilterer) WatchBind(opts *bind.WatchOpts, sink chan<- *IntoBindBind) (event.Subscription, error) {

	logs, sub, err := _IntoBind.contract.WatchLogs(opts, "Bind")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntoBindBind)
				if err := _IntoBind.contract.UnpackLog(event, "Bind", log); err != nil {
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

// ParseBind is a log parse operation binding the contract event 0xef2d9da059041adb88becbf698f278eabeccb23e046f0ae367dd57be0cce7b22.
//
// Solidity: event Bind(address from, string uid, bool isMain)
func (_IntoBind *IntoBindFilterer) ParseBind(log types.Log) (*IntoBindBind, error) {
	event := new(IntoBindBind)
	if err := _IntoBind.contract.UnpackLog(event, "Bind", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntoBindBindMainIterator is returned from FilterBindMain and is used to iterate over the raw logs and unpacked data for BindMain events raised by the IntoBind contract.
type IntoBindBindMainIterator struct {
	Event *IntoBindBindMain // Event containing the contract specifics and raw log

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
func (it *IntoBindBindMainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntoBindBindMain)
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
		it.Event = new(IntoBindBindMain)
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
func (it *IntoBindBindMainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntoBindBindMainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntoBindBindMain represents a BindMain event raised by the IntoBind contract.
type IntoBindBindMain struct {
	From common.Address
	Uid  string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBindMain is a free log retrieval operation binding the contract event 0x6b09508f5dc2290916e4c2e6c1866177a1bb2024b4c8a9a06dcf929d350f9153.
//
// Solidity: event BindMain(address from, string uid)
func (_IntoBind *IntoBindFilterer) FilterBindMain(opts *bind.FilterOpts) (*IntoBindBindMainIterator, error) {

	logs, sub, err := _IntoBind.contract.FilterLogs(opts, "BindMain")
	if err != nil {
		return nil, err
	}
	return &IntoBindBindMainIterator{contract: _IntoBind.contract, event: "BindMain", logs: logs, sub: sub}, nil
}

// WatchBindMain is a free log subscription operation binding the contract event 0x6b09508f5dc2290916e4c2e6c1866177a1bb2024b4c8a9a06dcf929d350f9153.
//
// Solidity: event BindMain(address from, string uid)
func (_IntoBind *IntoBindFilterer) WatchBindMain(opts *bind.WatchOpts, sink chan<- *IntoBindBindMain) (event.Subscription, error) {

	logs, sub, err := _IntoBind.contract.WatchLogs(opts, "BindMain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntoBindBindMain)
				if err := _IntoBind.contract.UnpackLog(event, "BindMain", log); err != nil {
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

// ParseBindMain is a log parse operation binding the contract event 0x6b09508f5dc2290916e4c2e6c1866177a1bb2024b4c8a9a06dcf929d350f9153.
//
// Solidity: event BindMain(address from, string uid)
func (_IntoBind *IntoBindFilterer) ParseBindMain(log types.Log) (*IntoBindBindMain, error) {
	event := new(IntoBindBindMain)
	if err := _IntoBind.contract.UnpackLog(event, "BindMain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntoBindBindRelationIterator is returned from FilterBindRelation and is used to iterate over the raw logs and unpacked data for BindRelation events raised by the IntoBind contract.
type IntoBindBindRelationIterator struct {
	Event *IntoBindBindRelation // Event containing the contract specifics and raw log

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
func (it *IntoBindBindRelationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntoBindBindRelation)
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
		it.Event = new(IntoBindBindRelation)
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
func (it *IntoBindBindRelationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntoBindBindRelationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntoBindBindRelation represents a BindRelation event raised by the IntoBind contract.
type IntoBindBindRelation struct {
	Sender common.Address
	Parent common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBindRelation is a free log retrieval operation binding the contract event 0x9e09303a92bee5c6c1e2c0dedf14a376bb5dc36774caddd851052830f1a01f1d.
//
// Solidity: event BindRelation(address sender, address parent)
func (_IntoBind *IntoBindFilterer) FilterBindRelation(opts *bind.FilterOpts) (*IntoBindBindRelationIterator, error) {

	logs, sub, err := _IntoBind.contract.FilterLogs(opts, "BindRelation")
	if err != nil {
		return nil, err
	}
	return &IntoBindBindRelationIterator{contract: _IntoBind.contract, event: "BindRelation", logs: logs, sub: sub}, nil
}

// WatchBindRelation is a free log subscription operation binding the contract event 0x9e09303a92bee5c6c1e2c0dedf14a376bb5dc36774caddd851052830f1a01f1d.
//
// Solidity: event BindRelation(address sender, address parent)
func (_IntoBind *IntoBindFilterer) WatchBindRelation(opts *bind.WatchOpts, sink chan<- *IntoBindBindRelation) (event.Subscription, error) {

	logs, sub, err := _IntoBind.contract.WatchLogs(opts, "BindRelation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntoBindBindRelation)
				if err := _IntoBind.contract.UnpackLog(event, "BindRelation", log); err != nil {
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

// ParseBindRelation is a log parse operation binding the contract event 0x9e09303a92bee5c6c1e2c0dedf14a376bb5dc36774caddd851052830f1a01f1d.
//
// Solidity: event BindRelation(address sender, address parent)
func (_IntoBind *IntoBindFilterer) ParseBindRelation(log types.Log) (*IntoBindBindRelation, error) {
	event := new(IntoBindBindRelation)
	if err := _IntoBind.contract.UnpackLog(event, "BindRelation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntoBindInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IntoBind contract.
type IntoBindInitializedIterator struct {
	Event *IntoBindInitialized // Event containing the contract specifics and raw log

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
func (it *IntoBindInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntoBindInitialized)
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
		it.Event = new(IntoBindInitialized)
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
func (it *IntoBindInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntoBindInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntoBindInitialized represents a Initialized event raised by the IntoBind contract.
type IntoBindInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IntoBind *IntoBindFilterer) FilterInitialized(opts *bind.FilterOpts) (*IntoBindInitializedIterator, error) {

	logs, sub, err := _IntoBind.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IntoBindInitializedIterator{contract: _IntoBind.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IntoBind *IntoBindFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IntoBindInitialized) (event.Subscription, error) {

	logs, sub, err := _IntoBind.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntoBindInitialized)
				if err := _IntoBind.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IntoBind *IntoBindFilterer) ParseInitialized(log types.Log) (*IntoBindInitialized, error) {
	event := new(IntoBindInitialized)
	if err := _IntoBind.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntoBindUnBindIterator is returned from FilterUnBind and is used to iterate over the raw logs and unpacked data for UnBind events raised by the IntoBind contract.
type IntoBindUnBindIterator struct {
	Event *IntoBindUnBind // Event containing the contract specifics and raw log

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
func (it *IntoBindUnBindIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntoBindUnBind)
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
		it.Event = new(IntoBindUnBind)
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
func (it *IntoBindUnBindIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntoBindUnBindIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntoBindUnBind represents a UnBind event raised by the IntoBind contract.
type IntoBindUnBind struct {
	From common.Address
	Uid  string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUnBind is a free log retrieval operation binding the contract event 0xea97762ed95338290a631dc63c304d1d1e204d3c29f42d7c1a896761639c837c.
//
// Solidity: event UnBind(address from, string uid)
func (_IntoBind *IntoBindFilterer) FilterUnBind(opts *bind.FilterOpts) (*IntoBindUnBindIterator, error) {

	logs, sub, err := _IntoBind.contract.FilterLogs(opts, "UnBind")
	if err != nil {
		return nil, err
	}
	return &IntoBindUnBindIterator{contract: _IntoBind.contract, event: "UnBind", logs: logs, sub: sub}, nil
}

// WatchUnBind is a free log subscription operation binding the contract event 0xea97762ed95338290a631dc63c304d1d1e204d3c29f42d7c1a896761639c837c.
//
// Solidity: event UnBind(address from, string uid)
func (_IntoBind *IntoBindFilterer) WatchUnBind(opts *bind.WatchOpts, sink chan<- *IntoBindUnBind) (event.Subscription, error) {

	logs, sub, err := _IntoBind.contract.WatchLogs(opts, "UnBind")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntoBindUnBind)
				if err := _IntoBind.contract.UnpackLog(event, "UnBind", log); err != nil {
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

// ParseUnBind is a log parse operation binding the contract event 0xea97762ed95338290a631dc63c304d1d1e204d3c29f42d7c1a896761639c837c.
//
// Solidity: event UnBind(address from, string uid)
func (_IntoBind *IntoBindFilterer) ParseUnBind(log types.Log) (*IntoBindUnBind, error) {
	event := new(IntoBindUnBind)
	if err := _IntoBind.contract.UnpackLog(event, "UnBind", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntoBindUnBindMainIterator is returned from FilterUnBindMain and is used to iterate over the raw logs and unpacked data for UnBindMain events raised by the IntoBind contract.
type IntoBindUnBindMainIterator struct {
	Event *IntoBindUnBindMain // Event containing the contract specifics and raw log

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
func (it *IntoBindUnBindMainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntoBindUnBindMain)
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
		it.Event = new(IntoBindUnBindMain)
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
func (it *IntoBindUnBindMainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntoBindUnBindMainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntoBindUnBindMain represents a UnBindMain event raised by the IntoBind contract.
type IntoBindUnBindMain struct {
	From common.Address
	Uid  string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUnBindMain is a free log retrieval operation binding the contract event 0xdf4046136a3962d371b0f533904dbca3418a7e29657f66b6f1653a309fa0b13a.
//
// Solidity: event UnBindMain(address from, string uid)
func (_IntoBind *IntoBindFilterer) FilterUnBindMain(opts *bind.FilterOpts) (*IntoBindUnBindMainIterator, error) {

	logs, sub, err := _IntoBind.contract.FilterLogs(opts, "UnBindMain")
	if err != nil {
		return nil, err
	}
	return &IntoBindUnBindMainIterator{contract: _IntoBind.contract, event: "UnBindMain", logs: logs, sub: sub}, nil
}

// WatchUnBindMain is a free log subscription operation binding the contract event 0xdf4046136a3962d371b0f533904dbca3418a7e29657f66b6f1653a309fa0b13a.
//
// Solidity: event UnBindMain(address from, string uid)
func (_IntoBind *IntoBindFilterer) WatchUnBindMain(opts *bind.WatchOpts, sink chan<- *IntoBindUnBindMain) (event.Subscription, error) {

	logs, sub, err := _IntoBind.contract.WatchLogs(opts, "UnBindMain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntoBindUnBindMain)
				if err := _IntoBind.contract.UnpackLog(event, "UnBindMain", log); err != nil {
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

// ParseUnBindMain is a log parse operation binding the contract event 0xdf4046136a3962d371b0f533904dbca3418a7e29657f66b6f1653a309fa0b13a.
//
// Solidity: event UnBindMain(address from, string uid)
func (_IntoBind *IntoBindFilterer) ParseUnBindMain(log types.Log) (*IntoBindUnBindMain, error) {
	event := new(IntoBindUnBindMain)
	if err := _IntoBind.contract.UnpackLog(event, "UnBindMain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
