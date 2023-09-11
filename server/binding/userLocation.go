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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cityId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"UserLocationRecord\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"countyId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chengShiId\",\"type\":\"bytes32\"}],\"name\":\"SetCityMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"intoCityAddress_\",\"type\":\"address\"}],\"name\":\"adminSetCityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pledgeStakeAddress_\",\"type\":\"address\"}],\"name\":\"adminSetPledgeStakeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint56\",\"name\":\"secondsPerDay_\",\"type\":\"uint56\"}],\"name\":\"adminSetSecondsPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chengShiIDCityIdSet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityIdChengShiID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityIdExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityIdExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cityIdNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityIdToChengShiIDExits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cityIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cityIdsNoRepeat\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cityInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getChengShiIdByAddress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"countyId\",\"type\":\"bytes32\"}],\"name\":\"getChengShiIdByCountyId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCityNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getCountyId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chengShiId\",\"type\":\"bytes32\"}],\"name\":\"getCountyIdsByChengShiId\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intoCityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"noRepeatCityIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgeStakeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondsPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cityId_\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"location_\",\"type\":\"string\"}],\"name\":\"setUserLocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cityId_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"location_\",\"type\":\"string\"}],\"name\":\"setUserLocationTest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userCityId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userLocationInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"userNumberOfCity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// ChengShiIDCityIdSet is a free data retrieval call binding the contract method 0xbec15656.
//
// Solidity: function chengShiIDCityIdSet(bytes32 , uint256 ) view returns(bytes32)
func (_UserLocation *UserLocationCaller) ChengShiIDCityIdSet(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "chengShiIDCityIdSet", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChengShiIDCityIdSet is a free data retrieval call binding the contract method 0xbec15656.
//
// Solidity: function chengShiIDCityIdSet(bytes32 , uint256 ) view returns(bytes32)
func (_UserLocation *UserLocationSession) ChengShiIDCityIdSet(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _UserLocation.Contract.ChengShiIDCityIdSet(&_UserLocation.CallOpts, arg0, arg1)
}

// ChengShiIDCityIdSet is a free data retrieval call binding the contract method 0xbec15656.
//
// Solidity: function chengShiIDCityIdSet(bytes32 , uint256 ) view returns(bytes32)
func (_UserLocation *UserLocationCallerSession) ChengShiIDCityIdSet(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _UserLocation.Contract.ChengShiIDCityIdSet(&_UserLocation.CallOpts, arg0, arg1)
}

// CityIdChengShiID is a free data retrieval call binding the contract method 0x95af990f.
//
// Solidity: function cityIdChengShiID(bytes32 ) view returns(bytes32)
func (_UserLocation *UserLocationCaller) CityIdChengShiID(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "cityIdChengShiID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CityIdChengShiID is a free data retrieval call binding the contract method 0x95af990f.
//
// Solidity: function cityIdChengShiID(bytes32 ) view returns(bytes32)
func (_UserLocation *UserLocationSession) CityIdChengShiID(arg0 [32]byte) ([32]byte, error) {
	return _UserLocation.Contract.CityIdChengShiID(&_UserLocation.CallOpts, arg0)
}

// CityIdChengShiID is a free data retrieval call binding the contract method 0x95af990f.
//
// Solidity: function cityIdChengShiID(bytes32 ) view returns(bytes32)
func (_UserLocation *UserLocationCallerSession) CityIdChengShiID(arg0 [32]byte) ([32]byte, error) {
	return _UserLocation.Contract.CityIdChengShiID(&_UserLocation.CallOpts, arg0)
}

// CityIdExist is a free data retrieval call binding the contract method 0x1f6cd476.
//
// Solidity: function cityIdExist(bytes32 ) view returns(bool)
func (_UserLocation *UserLocationCaller) CityIdExist(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "cityIdExist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CityIdExist is a free data retrieval call binding the contract method 0x1f6cd476.
//
// Solidity: function cityIdExist(bytes32 ) view returns(bool)
func (_UserLocation *UserLocationSession) CityIdExist(arg0 [32]byte) (bool, error) {
	return _UserLocation.Contract.CityIdExist(&_UserLocation.CallOpts, arg0)
}

// CityIdExist is a free data retrieval call binding the contract method 0x1f6cd476.
//
// Solidity: function cityIdExist(bytes32 ) view returns(bool)
func (_UserLocation *UserLocationCallerSession) CityIdExist(arg0 [32]byte) (bool, error) {
	return _UserLocation.Contract.CityIdExist(&_UserLocation.CallOpts, arg0)
}

// CityIdExists is a free data retrieval call binding the contract method 0x94ee0a50.
//
// Solidity: function cityIdExists(bytes32 ) view returns(bool)
func (_UserLocation *UserLocationCaller) CityIdExists(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "cityIdExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CityIdExists is a free data retrieval call binding the contract method 0x94ee0a50.
//
// Solidity: function cityIdExists(bytes32 ) view returns(bool)
func (_UserLocation *UserLocationSession) CityIdExists(arg0 [32]byte) (bool, error) {
	return _UserLocation.Contract.CityIdExists(&_UserLocation.CallOpts, arg0)
}

// CityIdExists is a free data retrieval call binding the contract method 0x94ee0a50.
//
// Solidity: function cityIdExists(bytes32 ) view returns(bool)
func (_UserLocation *UserLocationCallerSession) CityIdExists(arg0 [32]byte) (bool, error) {
	return _UserLocation.Contract.CityIdExists(&_UserLocation.CallOpts, arg0)
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

// CityIdToChengShiIDExits is a free data retrieval call binding the contract method 0x744395f3.
//
// Solidity: function cityIdToChengShiIDExits(bytes32 , bytes32 ) view returns(bool)
func (_UserLocation *UserLocationCaller) CityIdToChengShiIDExits(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "cityIdToChengShiIDExits", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CityIdToChengShiIDExits is a free data retrieval call binding the contract method 0x744395f3.
//
// Solidity: function cityIdToChengShiIDExits(bytes32 , bytes32 ) view returns(bool)
func (_UserLocation *UserLocationSession) CityIdToChengShiIDExits(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _UserLocation.Contract.CityIdToChengShiIDExits(&_UserLocation.CallOpts, arg0, arg1)
}

// CityIdToChengShiIDExits is a free data retrieval call binding the contract method 0x744395f3.
//
// Solidity: function cityIdToChengShiIDExits(bytes32 , bytes32 ) view returns(bool)
func (_UserLocation *UserLocationCallerSession) CityIdToChengShiIDExits(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _UserLocation.Contract.CityIdToChengShiIDExits(&_UserLocation.CallOpts, arg0, arg1)
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

// CityIdsNoRepeat is a free data retrieval call binding the contract method 0x5de48e18.
//
// Solidity: function cityIdsNoRepeat(uint256 ) view returns(bytes32)
func (_UserLocation *UserLocationCaller) CityIdsNoRepeat(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "cityIdsNoRepeat", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CityIdsNoRepeat is a free data retrieval call binding the contract method 0x5de48e18.
//
// Solidity: function cityIdsNoRepeat(uint256 ) view returns(bytes32)
func (_UserLocation *UserLocationSession) CityIdsNoRepeat(arg0 *big.Int) ([32]byte, error) {
	return _UserLocation.Contract.CityIdsNoRepeat(&_UserLocation.CallOpts, arg0)
}

// CityIdsNoRepeat is a free data retrieval call binding the contract method 0x5de48e18.
//
// Solidity: function cityIdsNoRepeat(uint256 ) view returns(bytes32)
func (_UserLocation *UserLocationCallerSession) CityIdsNoRepeat(arg0 *big.Int) ([32]byte, error) {
	return _UserLocation.Contract.CityIdsNoRepeat(&_UserLocation.CallOpts, arg0)
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

// GetChengShiIdByAddress is a free data retrieval call binding the contract method 0x981d5563.
//
// Solidity: function getChengShiIdByAddress(address user) view returns(bytes32)
func (_UserLocation *UserLocationCaller) GetChengShiIdByAddress(opts *bind.CallOpts, user common.Address) ([32]byte, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "getChengShiIdByAddress", user)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetChengShiIdByAddress is a free data retrieval call binding the contract method 0x981d5563.
//
// Solidity: function getChengShiIdByAddress(address user) view returns(bytes32)
func (_UserLocation *UserLocationSession) GetChengShiIdByAddress(user common.Address) ([32]byte, error) {
	return _UserLocation.Contract.GetChengShiIdByAddress(&_UserLocation.CallOpts, user)
}

// GetChengShiIdByAddress is a free data retrieval call binding the contract method 0x981d5563.
//
// Solidity: function getChengShiIdByAddress(address user) view returns(bytes32)
func (_UserLocation *UserLocationCallerSession) GetChengShiIdByAddress(user common.Address) ([32]byte, error) {
	return _UserLocation.Contract.GetChengShiIdByAddress(&_UserLocation.CallOpts, user)
}

// GetChengShiIdByCountyId is a free data retrieval call binding the contract method 0x0337fe26.
//
// Solidity: function getChengShiIdByCountyId(bytes32 countyId) view returns(bytes32)
func (_UserLocation *UserLocationCaller) GetChengShiIdByCountyId(opts *bind.CallOpts, countyId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "getChengShiIdByCountyId", countyId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetChengShiIdByCountyId is a free data retrieval call binding the contract method 0x0337fe26.
//
// Solidity: function getChengShiIdByCountyId(bytes32 countyId) view returns(bytes32)
func (_UserLocation *UserLocationSession) GetChengShiIdByCountyId(countyId [32]byte) ([32]byte, error) {
	return _UserLocation.Contract.GetChengShiIdByCountyId(&_UserLocation.CallOpts, countyId)
}

// GetChengShiIdByCountyId is a free data retrieval call binding the contract method 0x0337fe26.
//
// Solidity: function getChengShiIdByCountyId(bytes32 countyId) view returns(bytes32)
func (_UserLocation *UserLocationCallerSession) GetChengShiIdByCountyId(countyId [32]byte) ([32]byte, error) {
	return _UserLocation.Contract.GetChengShiIdByCountyId(&_UserLocation.CallOpts, countyId)
}

// GetCityNumber is a free data retrieval call binding the contract method 0x383a9096.
//
// Solidity: function getCityNumber() view returns(uint256)
func (_UserLocation *UserLocationCaller) GetCityNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "getCityNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCityNumber is a free data retrieval call binding the contract method 0x383a9096.
//
// Solidity: function getCityNumber() view returns(uint256)
func (_UserLocation *UserLocationSession) GetCityNumber() (*big.Int, error) {
	return _UserLocation.Contract.GetCityNumber(&_UserLocation.CallOpts)
}

// GetCityNumber is a free data retrieval call binding the contract method 0x383a9096.
//
// Solidity: function getCityNumber() view returns(uint256)
func (_UserLocation *UserLocationCallerSession) GetCityNumber() (*big.Int, error) {
	return _UserLocation.Contract.GetCityNumber(&_UserLocation.CallOpts)
}

// GetCountyId is a free data retrieval call binding the contract method 0x74cc7c80.
//
// Solidity: function getCountyId(address user) view returns(bytes32)
func (_UserLocation *UserLocationCaller) GetCountyId(opts *bind.CallOpts, user common.Address) ([32]byte, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "getCountyId", user)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCountyId is a free data retrieval call binding the contract method 0x74cc7c80.
//
// Solidity: function getCountyId(address user) view returns(bytes32)
func (_UserLocation *UserLocationSession) GetCountyId(user common.Address) ([32]byte, error) {
	return _UserLocation.Contract.GetCountyId(&_UserLocation.CallOpts, user)
}

// GetCountyId is a free data retrieval call binding the contract method 0x74cc7c80.
//
// Solidity: function getCountyId(address user) view returns(bytes32)
func (_UserLocation *UserLocationCallerSession) GetCountyId(user common.Address) ([32]byte, error) {
	return _UserLocation.Contract.GetCountyId(&_UserLocation.CallOpts, user)
}

// GetCountyIdsByChengShiId is a free data retrieval call binding the contract method 0x294643be.
//
// Solidity: function getCountyIdsByChengShiId(bytes32 chengShiId) view returns(bytes32[])
func (_UserLocation *UserLocationCaller) GetCountyIdsByChengShiId(opts *bind.CallOpts, chengShiId [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "getCountyIdsByChengShiId", chengShiId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetCountyIdsByChengShiId is a free data retrieval call binding the contract method 0x294643be.
//
// Solidity: function getCountyIdsByChengShiId(bytes32 chengShiId) view returns(bytes32[])
func (_UserLocation *UserLocationSession) GetCountyIdsByChengShiId(chengShiId [32]byte) ([][32]byte, error) {
	return _UserLocation.Contract.GetCountyIdsByChengShiId(&_UserLocation.CallOpts, chengShiId)
}

// GetCountyIdsByChengShiId is a free data retrieval call binding the contract method 0x294643be.
//
// Solidity: function getCountyIdsByChengShiId(bytes32 chengShiId) view returns(bytes32[])
func (_UserLocation *UserLocationCallerSession) GetCountyIdsByChengShiId(chengShiId [32]byte) ([][32]byte, error) {
	return _UserLocation.Contract.GetCountyIdsByChengShiId(&_UserLocation.CallOpts, chengShiId)
}

// GetDay is a free data retrieval call binding the contract method 0x14ba5c09.
//
// Solidity: function getDay() view returns(uint256)
func (_UserLocation *UserLocationCaller) GetDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "getDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDay is a free data retrieval call binding the contract method 0x14ba5c09.
//
// Solidity: function getDay() view returns(uint256)
func (_UserLocation *UserLocationSession) GetDay() (*big.Int, error) {
	return _UserLocation.Contract.GetDay(&_UserLocation.CallOpts)
}

// GetDay is a free data retrieval call binding the contract method 0x14ba5c09.
//
// Solidity: function getDay() view returns(uint256)
func (_UserLocation *UserLocationCallerSession) GetDay() (*big.Int, error) {
	return _UserLocation.Contract.GetDay(&_UserLocation.CallOpts)
}

// GetUserNumber is a free data retrieval call binding the contract method 0x20c9ad2c.
//
// Solidity: function getUserNumber() view returns(uint256)
func (_UserLocation *UserLocationCaller) GetUserNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "getUserNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNumber is a free data retrieval call binding the contract method 0x20c9ad2c.
//
// Solidity: function getUserNumber() view returns(uint256)
func (_UserLocation *UserLocationSession) GetUserNumber() (*big.Int, error) {
	return _UserLocation.Contract.GetUserNumber(&_UserLocation.CallOpts)
}

// GetUserNumber is a free data retrieval call binding the contract method 0x20c9ad2c.
//
// Solidity: function getUserNumber() view returns(uint256)
func (_UserLocation *UserLocationCallerSession) GetUserNumber() (*big.Int, error) {
	return _UserLocation.Contract.GetUserNumber(&_UserLocation.CallOpts)
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

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() view returns(uint256)
func (_UserLocation *UserLocationCaller) SecondsPerDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserLocation.contract.Call(opts, &out, "secondsPerDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() view returns(uint256)
func (_UserLocation *UserLocationSession) SecondsPerDay() (*big.Int, error) {
	return _UserLocation.Contract.SecondsPerDay(&_UserLocation.CallOpts)
}

// SecondsPerDay is a free data retrieval call binding the contract method 0x63809953.
//
// Solidity: function secondsPerDay() view returns(uint256)
func (_UserLocation *UserLocationCallerSession) SecondsPerDay() (*big.Int, error) {
	return _UserLocation.Contract.SecondsPerDay(&_UserLocation.CallOpts)
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

// SetCityMapping is a paid mutator transaction binding the contract method 0x03174d5a.
//
// Solidity: function SetCityMapping(bytes32 countyId, bytes32 chengShiId) returns()
func (_UserLocation *UserLocationTransactor) SetCityMapping(opts *bind.TransactOpts, countyId [32]byte, chengShiId [32]byte) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "SetCityMapping", countyId, chengShiId)
}

// SetCityMapping is a paid mutator transaction binding the contract method 0x03174d5a.
//
// Solidity: function SetCityMapping(bytes32 countyId, bytes32 chengShiId) returns()
func (_UserLocation *UserLocationSession) SetCityMapping(countyId [32]byte, chengShiId [32]byte) (*types.Transaction, error) {
	return _UserLocation.Contract.SetCityMapping(&_UserLocation.TransactOpts, countyId, chengShiId)
}

// SetCityMapping is a paid mutator transaction binding the contract method 0x03174d5a.
//
// Solidity: function SetCityMapping(bytes32 countyId, bytes32 chengShiId) returns()
func (_UserLocation *UserLocationTransactorSession) SetCityMapping(countyId [32]byte, chengShiId [32]byte) (*types.Transaction, error) {
	return _UserLocation.Contract.SetCityMapping(&_UserLocation.TransactOpts, countyId, chengShiId)
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

// AdminSetSecondsPerDay is a paid mutator transaction binding the contract method 0xace00fce.
//
// Solidity: function adminSetSecondsPerDay(uint56 secondsPerDay_) returns()
func (_UserLocation *UserLocationTransactor) AdminSetSecondsPerDay(opts *bind.TransactOpts, secondsPerDay_ *big.Int) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "adminSetSecondsPerDay", secondsPerDay_)
}

// AdminSetSecondsPerDay is a paid mutator transaction binding the contract method 0xace00fce.
//
// Solidity: function adminSetSecondsPerDay(uint56 secondsPerDay_) returns()
func (_UserLocation *UserLocationSession) AdminSetSecondsPerDay(secondsPerDay_ *big.Int) (*types.Transaction, error) {
	return _UserLocation.Contract.AdminSetSecondsPerDay(&_UserLocation.TransactOpts, secondsPerDay_)
}

// AdminSetSecondsPerDay is a paid mutator transaction binding the contract method 0xace00fce.
//
// Solidity: function adminSetSecondsPerDay(uint56 secondsPerDay_) returns()
func (_UserLocation *UserLocationTransactorSession) AdminSetSecondsPerDay(secondsPerDay_ *big.Int) (*types.Transaction, error) {
	return _UserLocation.Contract.AdminSetSecondsPerDay(&_UserLocation.TransactOpts, secondsPerDay_)
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

// NoRepeatCityIds is a paid mutator transaction binding the contract method 0xc653a2f0.
//
// Solidity: function noRepeatCityIds(uint256 start, uint256 end) returns()
func (_UserLocation *UserLocationTransactor) NoRepeatCityIds(opts *bind.TransactOpts, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "noRepeatCityIds", start, end)
}

// NoRepeatCityIds is a paid mutator transaction binding the contract method 0xc653a2f0.
//
// Solidity: function noRepeatCityIds(uint256 start, uint256 end) returns()
func (_UserLocation *UserLocationSession) NoRepeatCityIds(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _UserLocation.Contract.NoRepeatCityIds(&_UserLocation.TransactOpts, start, end)
}

// NoRepeatCityIds is a paid mutator transaction binding the contract method 0xc653a2f0.
//
// Solidity: function noRepeatCityIds(uint256 start, uint256 end) returns()
func (_UserLocation *UserLocationTransactorSession) NoRepeatCityIds(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _UserLocation.Contract.NoRepeatCityIds(&_UserLocation.TransactOpts, start, end)
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

// SetUserLocationTest is a paid mutator transaction binding the contract method 0x4f5879d8.
//
// Solidity: function setUserLocationTest(bytes32 cityId_, address user, string location_) returns()
func (_UserLocation *UserLocationTransactor) SetUserLocationTest(opts *bind.TransactOpts, cityId_ [32]byte, user common.Address, location_ string) (*types.Transaction, error) {
	return _UserLocation.contract.Transact(opts, "setUserLocationTest", cityId_, user, location_)
}

// SetUserLocationTest is a paid mutator transaction binding the contract method 0x4f5879d8.
//
// Solidity: function setUserLocationTest(bytes32 cityId_, address user, string location_) returns()
func (_UserLocation *UserLocationSession) SetUserLocationTest(cityId_ [32]byte, user common.Address, location_ string) (*types.Transaction, error) {
	return _UserLocation.Contract.SetUserLocationTest(&_UserLocation.TransactOpts, cityId_, user, location_)
}

// SetUserLocationTest is a paid mutator transaction binding the contract method 0x4f5879d8.
//
// Solidity: function setUserLocationTest(bytes32 cityId_, address user, string location_) returns()
func (_UserLocation *UserLocationTransactorSession) SetUserLocationTest(cityId_ [32]byte, user common.Address, location_ string) (*types.Transaction, error) {
	return _UserLocation.Contract.SetUserLocationTest(&_UserLocation.TransactOpts, cityId_, user, location_)
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
