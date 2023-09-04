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

// UserLocationMetaData contains all meta data concerning the UserLocation contract.
var UserLocationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cityId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"UserLocationRecord\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"intoCityAddress_\",\"type\":\"address\"}],\"name\":\"adminSetCityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pledgeStakeAddress_\",\"type\":\"address\"}],\"name\":\"adminSetPledgeStakeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cityIdNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cityIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllCityNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intoCityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgeStakeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cityId_\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"location_\",\"type\":\"string\"}],\"name\":\"setUserLocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userCityId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userLocationInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"userNumberOfCity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UserLocationABI is the input ABI used to generate the binding from.
// Deprecated: Use UserLocationMetaData.ABI instead.
var UserLocationABI = UserLocationMetaData.ABI

// UserLocation is an auto generated Go binding around an Ethereum contract.
type UserLocation struct {
	UserLocationCaller     // Read-only binding to the contract
	UserLocationTransactor // Write-only binding to the contract
	UserLocationFilterer   // Log filterer for contract events
}

// UserLocationCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserLocationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserLocationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserLocationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserLocationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserLocationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserLocationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserLocationSession struct {
	Contract     *UserLocation     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserLocationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserLocationCallerSession struct {
	Contract *UserLocationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UserLocationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserLocationTransactorSession struct {
	Contract     *UserLocationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UserLocationRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserLocationRaw struct {
	Contract *UserLocation // Generic contract binding to access the raw methods on
}

// UserLocationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserLocationCallerRaw struct {
	Contract *UserLocationCaller // Generic read-only contract binding to access the raw methods on
}

// UserLocationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserLocationTransactorRaw struct {
	Contract *UserLocationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserLocation creates a new instance of UserLocation, bound to a specific deployed contract.
func NewUserLocation(address common.Address, backend bind.ContractBackend) (*UserLocation, error) {
	contract, err := bindUserLocation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserLocation{UserLocationCaller: UserLocationCaller{contract: contract}, UserLocationTransactor: UserLocationTransactor{contract: contract}, UserLocationFilterer: UserLocationFilterer{contract: contract}}, nil
}

// NewUserLocationCaller creates a new read-only instance of UserLocation, bound to a specific deployed contract.
func NewUserLocationCaller(address common.Address, caller bind.ContractCaller) (*UserLocationCaller, error) {
	contract, err := bindUserLocation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserLocationCaller{contract: contract}, nil
}

// NewUserLocationTransactor creates a new write-only instance of UserLocation, bound to a specific deployed contract.
func NewUserLocationTransactor(address common.Address, transactor bind.ContractTransactor) (*UserLocationTransactor, error) {
	contract, err := bindUserLocation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserLocationTransactor{contract: contract}, nil
}

// NewUserLocationFilterer creates a new log filterer instance of UserLocation, bound to a specific deployed contract.
func NewUserLocationFilterer(address common.Address, filterer bind.ContractFilterer) (*UserLocationFilterer, error) {
	contract, err := bindUserLocation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserLocationFilterer{contract: contract}, nil
}

// bindUserLocation binds a generic wrapper to an already deployed contract.
func bindUserLocation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserLocationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserLocation *UserLocationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserLocation.Contract.UserLocationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserLocation *UserLocationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserLocation.Contract.UserLocationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserLocation *UserLocationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserLocation.Contract.UserLocationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserLocation *UserLocationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserLocation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserLocation *UserLocationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserLocation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserLocation *UserLocationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserLocation.Contract.contract.Transact(opts, method, params...)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_UserLocation *UserLocationCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_UserLocation *UserLocationSession) AllAdmins() ([]common.Address, error) {
	return _UserLocation.Contract.AllAdmins(&_UserLocation.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_UserLocation *UserLocationCallerSession) AllAdmins() ([]common.Address, error) {
	return _UserLocation.Contract.AllAdmins(&_UserLocation.CallOpts)
}

// CityIdNum is a free data retrieval call binding the contract method 0xf896a228.
//
// Solidity: function cityIdNum() view returns(uint256)
func (_UserLocation *UserLocationCaller) CityIdNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "cityIdNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CityIdNum is a free data retrieval call binding the contract method 0xf896a228.
//
// Solidity: function cityIdNum() view returns(uint256)
func (_UserLocation *UserLocationSession) CityIdNum() (*big.Int, error) {
	return _UserLocation.Contract.CityIdNum(&_UserLocation.CallOpts)
}

// CityIdNum is a free data retrieval call binding the contract method 0xf896a228.
//
// Solidity: function cityIdNum() view returns(uint256)
func (_UserLocation *UserLocationCallerSession) CityIdNum() (*big.Int, error) {
	return _UserLocation.Contract.CityIdNum(&_UserLocation.CallOpts)
}

// CityIds is a free data retrieval call binding the contract method 0x75e6a6d4.
//
// Solidity: function cityIds(uint256 ) view returns(bytes32)
func (_UserLocation *UserLocationCaller) CityIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "cityIds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CityIds is a free data retrieval call binding the contract method 0x75e6a6d4.
//
// Solidity: function cityIds(uint256 ) view returns(bytes32)
func (_UserLocation *UserLocationSession) CityIds(arg0 *big.Int) ([32]byte, error) {
	return _UserLocation.Contract.CityIds(&_UserLocation.CallOpts, arg0)
}

// CityIds is a free data retrieval call binding the contract method 0x75e6a6d4.
//
// Solidity: function cityIds(uint256 ) view returns(bytes32)
func (_UserLocation *UserLocationCallerSession) CityIds(arg0 *big.Int) ([32]byte, error) {
	return _UserLocation.Contract.CityIds(&_UserLocation.CallOpts, arg0)
}

// CityInfo is a free data retrieval call binding the contract method 0x72e53ecd.
//
// Solidity: function cityInfo(bytes32 ) view returns(string)
func (_UserLocation *UserLocationCaller) CityInfo(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "cityInfo", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CityInfo is a free data retrieval call binding the contract method 0x72e53ecd.
//
// Solidity: function cityInfo(bytes32 ) view returns(string)
func (_UserLocation *UserLocationSession) CityInfo(arg0 [32]byte) (string, error) {
	return _UserLocation.Contract.CityInfo(&_UserLocation.CallOpts, arg0)
}

// CityInfo is a free data retrieval call binding the contract method 0x72e53ecd.
//
// Solidity: function cityInfo(bytes32 ) view returns(string)
func (_UserLocation *UserLocationCallerSession) CityInfo(arg0 [32]byte) (string, error) {
	return _UserLocation.Contract.CityInfo(&_UserLocation.CallOpts, arg0)
}

// GetAllCityNumber is a free data retrieval call binding the contract method 0x957466c7.
//
// Solidity: function getAllCityNumber() view returns(uint256)
func (_UserLocation *UserLocationCaller) GetAllCityNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "getAllCityNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllCityNumber is a free data retrieval call binding the contract method 0x957466c7.
//
// Solidity: function getAllCityNumber() view returns(uint256)
func (_UserLocation *UserLocationSession) GetAllCityNumber() (*big.Int, error) {
	return _UserLocation.Contract.GetAllCityNumber(&_UserLocation.CallOpts)
}

// GetAllCityNumber is a free data retrieval call binding the contract method 0x957466c7.
//
// Solidity: function getAllCityNumber() view returns(uint256)
func (_UserLocation *UserLocationCallerSession) GetAllCityNumber() (*big.Int, error) {
	return _UserLocation.Contract.GetAllCityNumber(&_UserLocation.CallOpts)
}

// IntoCityAddress is a free data retrieval call binding the contract method 0xf9c454de.
//
// Solidity: function intoCityAddress() view returns(address)
func (_UserLocation *UserLocationCaller) IntoCityAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "intoCityAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IntoCityAddress is a free data retrieval call binding the contract method 0xf9c454de.
//
// Solidity: function intoCityAddress() view returns(address)
func (_UserLocation *UserLocationSession) IntoCityAddress() (common.Address, error) {
	return _UserLocation.Contract.IntoCityAddress(&_UserLocation.CallOpts)
}

// IntoCityAddress is a free data retrieval call binding the contract method 0xf9c454de.
//
// Solidity: function intoCityAddress() view returns(address)
func (_UserLocation *UserLocationCallerSession) IntoCityAddress() (common.Address, error) {
	return _UserLocation.Contract.IntoCityAddress(&_UserLocation.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_UserLocation *UserLocationCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_UserLocation *UserLocationSession) IsAdmin(account common.Address) (bool, error) {
	return _UserLocation.Contract.IsAdmin(&_UserLocation.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_UserLocation *UserLocationCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _UserLocation.Contract.IsAdmin(&_UserLocation.CallOpts, account)
}

// PledgeStakeAddress is a free data retrieval call binding the contract method 0x7a4e39f7.
//
// Solidity: function pledgeStakeAddress() view returns(address)
func (_UserLocation *UserLocationCaller) PledgeStakeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "pledgeStakeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PledgeStakeAddress is a free data retrieval call binding the contract method 0x7a4e39f7.
//
// Solidity: function pledgeStakeAddress() view returns(address)
func (_UserLocation *UserLocationSession) PledgeStakeAddress() (common.Address, error) {
	return _UserLocation.Contract.PledgeStakeAddress(&_UserLocation.CallOpts)
}

// PledgeStakeAddress is a free data retrieval call binding the contract method 0x7a4e39f7.
//
// Solidity: function pledgeStakeAddress() view returns(address)
func (_UserLocation *UserLocationCallerSession) PledgeStakeAddress() (common.Address, error) {
	return _UserLocation.Contract.PledgeStakeAddress(&_UserLocation.CallOpts)
}

// UserCityId is a free data retrieval call binding the contract method 0xd4c22976.
//
// Solidity: function userCityId(address ) view returns(bytes32)
func (_UserLocation *UserLocationCaller) UserCityId(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "userCityId", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UserCityId is a free data retrieval call binding the contract method 0xd4c22976.
//
// Solidity: function userCityId(address ) view returns(bytes32)
func (_UserLocation *UserLocationSession) UserCityId(arg0 common.Address) ([32]byte, error) {
	return _UserLocation.Contract.UserCityId(&_UserLocation.CallOpts, arg0)
}

// UserCityId is a free data retrieval call binding the contract method 0xd4c22976.
//
// Solidity: function userCityId(address ) view returns(bytes32)
func (_UserLocation *UserLocationCallerSession) UserCityId(arg0 common.Address) ([32]byte, error) {
	return _UserLocation.Contract.UserCityId(&_UserLocation.CallOpts, arg0)
}

// UserLocationInfo is a free data retrieval call binding the contract method 0x2f25dcf7.
//
// Solidity: function userLocationInfo(address ) view returns(string)
func (_UserLocation *UserLocationCaller) UserLocationInfo(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "userLocationInfo", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UserLocationInfo is a free data retrieval call binding the contract method 0x2f25dcf7.
//
// Solidity: function userLocationInfo(address ) view returns(string)
func (_UserLocation *UserLocationSession) UserLocationInfo(arg0 common.Address) (string, error) {
	return _UserLocation.Contract.UserLocationInfo(&_UserLocation.CallOpts, arg0)
}

// UserLocationInfo is a free data retrieval call binding the contract method 0x2f25dcf7.
//
// Solidity: function userLocationInfo(address ) view returns(string)
func (_UserLocation *UserLocationCallerSession) UserLocationInfo(arg0 common.Address) (string, error) {
	return _UserLocation.Contract.UserLocationInfo(&_UserLocation.CallOpts, arg0)
}

// UserNumberOfCity is a free data retrieval call binding the contract method 0xd03834b5.
//
// Solidity: function userNumberOfCity(bytes32 ) view returns(uint256)
func (_UserLocation *UserLocationCaller) UserNumberOfCity(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "userNumberOfCity", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserNumberOfCity is a free data retrieval call binding the contract method 0xd03834b5.
//
// Solidity: function userNumberOfCity(bytes32 ) view returns(uint256)
func (_UserLocation *UserLocationSession) UserNumberOfCity(arg0 [32]byte) (*big.Int, error) {
	return _UserLocation.Contract.UserNumberOfCity(&_UserLocation.CallOpts, arg0)
}

// UserNumberOfCity is a free data retrieval call binding the contract method 0xd03834b5.
//
// Solidity: function userNumberOfCity(bytes32 ) view returns(uint256)
func (_UserLocation *UserLocationCallerSession) UserNumberOfCity(arg0 [32]byte) (*big.Int, error) {
	return _UserLocation.Contract.UserNumberOfCity(&_UserLocation.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_UserLocation *UserLocationTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_UserLocation *UserLocationSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _UserLocation.Contract.AddAdmin(&_UserLocation.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_UserLocation *UserLocationTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _UserLocation.Contract.AddAdmin(&_UserLocation.TransactOpts, account)
}

// AdminSetCityAddress is a paid mutator transaction binding the contract method 0xa94c843d.
//
// Solidity: function adminSetCityAddress(address intoCityAddress_) returns()
func (_UserLocation *UserLocationTransactor) AdminSetCityAddress(opts *bind.TransactOpts, intoCityAddress_ common.Address) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "adminSetCityAddress", intoCityAddress_)
}

// AdminSetCityAddress is a paid mutator transaction binding the contract method 0xa94c843d.
//
// Solidity: function adminSetCityAddress(address intoCityAddress_) returns()
func (_UserLocation *UserLocationSession) AdminSetCityAddress(intoCityAddress_ common.Address) (*types.Transaction, error) {
	return _UserLocation.Contract.AdminSetCityAddress(&_UserLocation.TransactOpts, intoCityAddress_)
}

// AdminSetCityAddress is a paid mutator transaction binding the contract method 0xa94c843d.
//
// Solidity: function adminSetCityAddress(address intoCityAddress_) returns()
func (_UserLocation *UserLocationTransactorSession) AdminSetCityAddress(intoCityAddress_ common.Address) (*types.Transaction, error) {
	return _UserLocation.Contract.AdminSetCityAddress(&_UserLocation.TransactOpts, intoCityAddress_)
}

// AdminSetPledgeStakeAddress is a paid mutator transaction binding the contract method 0xefddfbeb.
//
// Solidity: function adminSetPledgeStakeAddress(address pledgeStakeAddress_) returns()
func (_UserLocation *UserLocationTransactor) AdminSetPledgeStakeAddress(opts *bind.TransactOpts, pledgeStakeAddress_ common.Address) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "adminSetPledgeStakeAddress", pledgeStakeAddress_)
}

// AdminSetPledgeStakeAddress is a paid mutator transaction binding the contract method 0xefddfbeb.
//
// Solidity: function adminSetPledgeStakeAddress(address pledgeStakeAddress_) returns()
func (_UserLocation *UserLocationSession) AdminSetPledgeStakeAddress(pledgeStakeAddress_ common.Address) (*types.Transaction, error) {
	return _UserLocation.Contract.AdminSetPledgeStakeAddress(&_UserLocation.TransactOpts, pledgeStakeAddress_)
}

// AdminSetPledgeStakeAddress is a paid mutator transaction binding the contract method 0xefddfbeb.
//
// Solidity: function adminSetPledgeStakeAddress(address pledgeStakeAddress_) returns()
func (_UserLocation *UserLocationTransactorSession) AdminSetPledgeStakeAddress(pledgeStakeAddress_ common.Address) (*types.Transaction, error) {
	return _UserLocation.Contract.AdminSetPledgeStakeAddress(&_UserLocation.TransactOpts, pledgeStakeAddress_)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_UserLocation *UserLocationTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_UserLocation *UserLocationSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _UserLocation.Contract.BatchAddAdmin(&_UserLocation.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_UserLocation *UserLocationTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _UserLocation.Contract.BatchAddAdmin(&_UserLocation.TransactOpts, amounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_UserLocation *UserLocationTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_UserLocation *UserLocationSession) Initialize() (*types.Transaction, error) {
	return _UserLocation.Contract.Initialize(&_UserLocation.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_UserLocation *UserLocationTransactorSession) Initialize() (*types.Transaction, error) {
	return _UserLocation.Contract.Initialize(&_UserLocation.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_UserLocation *UserLocationTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_UserLocation *UserLocationSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _UserLocation.Contract.RemoveAdmin(&_UserLocation.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_UserLocation *UserLocationTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _UserLocation.Contract.RemoveAdmin(&_UserLocation.TransactOpts, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_UserLocation *UserLocationTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_UserLocation *UserLocationSession) RenounceAdmin() (*types.Transaction, error) {
	return _UserLocation.Contract.RenounceAdmin(&_UserLocation.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_UserLocation *UserLocationTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _UserLocation.Contract.RenounceAdmin(&_UserLocation.TransactOpts)
}

// SetUserLocation is a paid mutator transaction binding the contract method 0x56a07295.
//
// Solidity: function setUserLocation(bytes32 cityId_, string location_) returns()
func (_UserLocation *UserLocationTransactor) SetUserLocation(opts *bind.TransactOpts, cityId_ [32]byte, location_ string) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "setUserLocation", cityId_, location_)
}

// SetUserLocation is a paid mutator transaction binding the contract method 0x56a07295.
//
// Solidity: function setUserLocation(bytes32 cityId_, string location_) returns()
func (_UserLocation *UserLocationSession) SetUserLocation(cityId_ [32]byte, location_ string) (*types.Transaction, error) {
	return _UserLocation.Contract.SetUserLocation(&_UserLocation.TransactOpts, cityId_, location_)
}

// SetUserLocation is a paid mutator transaction binding the contract method 0x56a07295.
//
// Solidity: function setUserLocation(bytes32 cityId_, string location_) returns()
func (_UserLocation *UserLocationTransactorSession) SetUserLocation(cityId_ [32]byte, location_ string) (*types.Transaction, error) {
	return _UserLocation.Contract.SetUserLocation(&_UserLocation.TransactOpts, cityId_, location_)
}

// UserLocationAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the UserLocation contract.
type UserLocationAdminAddedIterator struct {
	Event *UserLocationAdminAdded // Event containing the contract specifics and raw log

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
func (it *UserLocationAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserLocationAdminAdded)
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
		it.Event = new(UserLocationAdminAdded)
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
func (it *UserLocationAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserLocationAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserLocationAdminAdded represents a AdminAdded event raised by the UserLocation contract.
type UserLocationAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_UserLocation *UserLocationFilterer) FilterAdminAdded(opts *bind.FilterOpts, account []common.Address) (*UserLocationAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _UserLocation.contract.FilterLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &UserLocationAdminAddedIterator{contract: _UserLocation.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_UserLocation *UserLocationFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *UserLocationAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _UserLocation.contract.WatchLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserLocationAdminAdded)
				if err := _UserLocation.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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
func (_UserLocation *UserLocationFilterer) ParseAdminAdded(log types.Log) (*UserLocationAdminAdded, error) {
	event := new(UserLocationAdminAdded)
	if err := _UserLocation.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserLocationAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the UserLocation contract.
type UserLocationAdminRemovedIterator struct {
	Event *UserLocationAdminRemoved // Event containing the contract specifics and raw log

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
func (it *UserLocationAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserLocationAdminRemoved)
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
		it.Event = new(UserLocationAdminRemoved)
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
func (it *UserLocationAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserLocationAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserLocationAdminRemoved represents a AdminRemoved event raised by the UserLocation contract.
type UserLocationAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_UserLocation *UserLocationFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*UserLocationAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _UserLocation.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &UserLocationAdminRemovedIterator{contract: _UserLocation.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_UserLocation *UserLocationFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *UserLocationAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _UserLocation.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserLocationAdminRemoved)
				if err := _UserLocation.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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
func (_UserLocation *UserLocationFilterer) ParseAdminRemoved(log types.Log) (*UserLocationAdminRemoved, error) {
	event := new(UserLocationAdminRemoved)
	if err := _UserLocation.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserLocationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UserLocation contract.
type UserLocationInitializedIterator struct {
	Event *UserLocationInitialized // Event containing the contract specifics and raw log

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
func (it *UserLocationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserLocationInitialized)
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
		it.Event = new(UserLocationInitialized)
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
func (it *UserLocationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserLocationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserLocationInitialized represents a Initialized event raised by the UserLocation contract.
type UserLocationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UserLocation *UserLocationFilterer) FilterInitialized(opts *bind.FilterOpts) (*UserLocationInitializedIterator, error) {

	logs, sub, err := _UserLocation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UserLocationInitializedIterator{contract: _UserLocation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UserLocation *UserLocationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UserLocationInitialized) (event.Subscription, error) {

	logs, sub, err := _UserLocation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserLocationInitialized)
				if err := _UserLocation.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_UserLocation *UserLocationFilterer) ParseInitialized(log types.Log) (*UserLocationInitialized, error) {
	event := new(UserLocationInitialized)
	if err := _UserLocation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserLocationUserLocationRecordIterator is returned from FilterUserLocationRecord and is used to iterate over the raw logs and unpacked data for UserLocationRecord events raised by the UserLocation contract.
type UserLocationUserLocationRecordIterator struct {
	Event *UserLocationUserLocationRecord // Event containing the contract specifics and raw log

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
func (it *UserLocationUserLocationRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserLocationUserLocationRecord)
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
		it.Event = new(UserLocationUserLocationRecord)
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
func (it *UserLocationUserLocationRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserLocationUserLocationRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserLocationUserLocationRecord represents a UserLocationRecord event raised by the UserLocation contract.
type UserLocationUserLocationRecord struct {
	User     common.Address
	CityId   [32]byte
	Location string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUserLocationRecord is a free log retrieval operation binding the contract event 0xcac51e2b22a0c7477d48ef20bc67251d74aa161293e1fa5cdf916a8066a6be74.
//
// Solidity: event UserLocationRecord(address user, bytes32 cityId, string location)
func (_UserLocation *UserLocationFilterer) FilterUserLocationRecord(opts *bind.FilterOpts) (*UserLocationUserLocationRecordIterator, error) {

	logs, sub, err := _UserLocation.contract.FilterLogs(opts, "UserLocationRecord")
	if err != nil {
		return nil, err
	}
	return &UserLocationUserLocationRecordIterator{contract: _UserLocation.contract, event: "UserLocationRecord", logs: logs, sub: sub}, nil
}

// WatchUserLocationRecord is a free log subscription operation binding the contract event 0xcac51e2b22a0c7477d48ef20bc67251d74aa161293e1fa5cdf916a8066a6be74.
//
// Solidity: event UserLocationRecord(address user, bytes32 cityId, string location)
func (_UserLocation *UserLocationFilterer) WatchUserLocationRecord(opts *bind.WatchOpts, sink chan<- *UserLocationUserLocationRecord) (event.Subscription, error) {

	logs, sub, err := _UserLocation.contract.WatchLogs(opts, "UserLocationRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserLocationUserLocationRecord)
				if err := _UserLocation.contract.UnpackLog(event, "UserLocationRecord", log); err != nil {
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

// ParseUserLocationRecord is a log parse operation binding the contract event 0xcac51e2b22a0c7477d48ef20bc67251d74aa161293e1fa5cdf916a8066a6be74.
//
// Solidity: event UserLocationRecord(address user, bytes32 cityId, string location)
func (_UserLocation *UserLocationFilterer) ParseUserLocationRecord(log types.Log) (*UserLocationUserLocationRecord, error) {
	event := new(UserLocationUserLocationRecord)
	if err := _UserLocation.contract.UnpackLog(event, "UserLocationRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
